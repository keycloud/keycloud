package com.keycloud.keycloudautofiller;

import android.app.assist.AssistStructure;
import android.content.Intent;
import android.content.IntentSender;
import android.service.autofill.Dataset;
import android.service.autofill.FillResponse;
import android.util.Base64;
import android.util.Log;
import android.view.autofill.AutofillId;
import android.view.autofill.AutofillValue;
import android.widget.RemoteViews;
import android.widget.Toast;

import com.google.android.gms.fido.Fido;
import com.google.android.gms.fido.common.Transport;
import com.google.android.gms.fido.fido2.Fido2ApiClient;
import com.google.android.gms.fido.fido2.Fido2PendingIntent;
import com.google.android.gms.fido.fido2.api.common.AuthenticatorAssertionResponse;
import com.google.android.gms.fido.fido2.api.common.PublicKeyCredentialDescriptor;
import com.google.android.gms.fido.fido2.api.common.PublicKeyCredentialRequestOptions;
import com.google.android.gms.fido.fido2.api.common.PublicKeyCredentialType;
import com.google.android.gms.tasks.OnSuccessListener;
import com.google.android.gms.tasks.Task;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.net.CookieHandler;
import java.net.CookieManager;
import java.net.URL;
import java.nio.charset.Charset;
import java.security.SecureRandom;
import java.security.cert.X509Certificate;
import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import java.util.stream.StreamSupport;

import javax.net.ssl.HostnameVerifier;
import javax.net.ssl.HttpsURLConnection;
import javax.net.ssl.SSLContext;
import javax.net.ssl.SSLSession;
import javax.net.ssl.TrustManager;
import javax.net.ssl.X509TrustManager;

import static android.view.autofill.AutofillManager.EXTRA_AUTHENTICATION_RESULT;

class KeyCloudAPI {
    static final int SIGN_REQUEST_CODE = 2356;
    private KeyCloudAuthenticationActivity mActivity;

    private static final String WEB_API_URL = "https://10.0.2.2/";
    private static final String WEB_API_STANDARD_LOGIN_URL = WEB_API_URL + "standard/login";
    private static final String WEB_API_WEBAUTHN_LOGIN_URL = WEB_API_URL + "webauthn/";

    private HttpsURLConnection mConnection;

    private List<AutofillId> mAutofillIds;
    private AssistStructure mFillStructure;
    private String mQuery;

    private List<FillInfoSet> mKeyCloudFillInfo;

    KeyCloudAPI(KeyCloudAuthenticationActivity activity, AssistStructure structure){
        this.mActivity = activity;
        this.mFillStructure = structure;
        this.mAutofillIds  = new LinkedList<>();
        if (structure != null){
            this.mQuery = mFillStructure.getActivityComponent().toString();
            for(int i = 0; i < structure.getWindowNodeCount(); i++)
                AutofillerService.traverseView(structure.getWindowNodeAt(i).getRootViewNode(),
                        mAutofillIds);
        }

        // Preparation for login: Create a cookie manager
        CookieManager manager = new CookieManager();
        CookieHandler.setDefault(manager);
    }

    void fillOutFromKeyCloud(final String username){
        final KeyCloudAPI self = this;
        new Thread(new Runnable() {
            @Override
            public void run() {
                checkError(self ,self.loginWithWebAuthn(username), "Authentication failed");
            }
        }).start();
    }

    void fillOutFromKeyCloud(final String username, final String password){
        final KeyCloudAPI self = this;
        new Thread(new Runnable() {
            @Override
            public void run() {
                checkError(self ,self.loginWithPassword(username, password), "Wrong Credentials");
                checkError(self, self.gatherFillingInfo(), "Error while looking for information");
                self.replyFillRequest();
            }
        }).start();
    }

    private int loginWithWebAuthn(String username){
        try {
            URL url = new URL(WEB_API_WEBAUTHN_LOGIN_URL + "login/start");

            trustAllCertificates();
            mConnection = (HttpsURLConnection) url.openConnection();
            mConnection.setRequestMethod("POST");
            mConnection.setDoOutput(true);

            OutputStream out = mConnection.getOutputStream();

            out.write(("{\"username\": \"" + username.trim() + "\", " +
                        "\"mail\": \"\"}").getBytes());
            out.flush();
            out.close();

            int responseCode = mConnection.getResponseCode();
            if(responseCode != 200){
                return responseCode;
            }

            InputStream inputStream = mConnection.getInputStream();
            BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(inputStream));

            StringBuilder jsonObj = new StringBuilder();
            String line;
            while((line = bufferedReader.readLine()) != null){
                jsonObj.append(line);
            }

            JSONObject publicKey = new JSONObject(jsonObj.toString()).getJSONObject("publicKey");
            JSONArray publicKeyDescriptors = publicKey.getJSONArray("allowCredentials");

            List<PublicKeyCredentialDescriptor> publicKeyCredentialDescriptors = new LinkedList<>();
            for(int i = 0; i < publicKeyDescriptors.length(); i++){
                //public Key Descriptor
                JSONArray descriptorArray = publicKeyDescriptors.getJSONObject(i).getJSONArray("id");
                byte[] descriptor = new byte[descriptorArray.length()];
                for (int j = 0; j < descriptorArray.length(); j++){
                    descriptor[j] = (byte) descriptorArray.getInt(j);
                }
                //Authenticator transports
                JSONArray authTransport = publicKeyDescriptors.getJSONObject(i).getJSONArray("transports");
                List<Transport> transports = new LinkedList<>();
                for(int j = 0; j < authTransport.length(); j++){
                    transports.add(Transport.fromString(authTransport.getString(j)));
                }

                publicKeyCredentialDescriptors.add(
                        new PublicKeyCredentialDescriptor(
                                PublicKeyCredentialType.PUBLIC_KEY.toString(),
                                descriptor,
                                transports));
            }

            PublicKeyCredentialRequestOptions options =
                    new PublicKeyCredentialRequestOptions.Builder()
                            .setRpId("10.0.2.2")
                            .setAllowList(publicKeyCredentialDescriptors)
                            .setChallenge(Base64.decode(publicKey.getString("challenge"), Base64.NO_WRAP))
                            .build();
            Fido2ApiClient fido2ApiClient = Fido.getFido2ApiClient(mActivity);
            Task<Fido2PendingIntent> fido2PendingIntentTask = fido2ApiClient.getSignIntent(options);
            fido2PendingIntentTask.addOnSuccessListener(
                    new OnSuccessListener<Fido2PendingIntent>() {
                        @Override
                        public void onSuccess(Fido2PendingIntent fido2PendingIntent) {
                            if (fido2PendingIntent.hasPendingIntent()) {
                                try {
                                    fido2PendingIntent.launchPendingIntent(
                                            mActivity,
                                            SIGN_REQUEST_CODE);
                                } catch (IntentSender.SendIntentException e) {
                                    e.printStackTrace();
                                }
                            }
                        }
                    });

            Log.i("WEBAUTHN", "");
            return 200;

        } catch (IOException | JSONException | Transport.UnsupportedTransportException e) {
            e.printStackTrace();
        }

        return -1;
    }

    static void checkError(final KeyCloudAPI self, int response, final String msg){
        if(response != 200){
            self.mActivity.runOnUiThread(new Runnable() {
                @Override
                public void run() {
                    Toast.makeText(self.mActivity, msg, Toast.LENGTH_SHORT).show();
                }
            });
        }
    }

    private int loginWithPassword(String username, String password){
        try {

            URL loginUrl = new URL(WEB_API_STANDARD_LOGIN_URL);
            trustAllCertificates();
            mConnection = (HttpsURLConnection) loginUrl.openConnection();

            mConnection.setRequestMethod("POST");
            mConnection.setDoOutput(true);

            OutputStream out = mConnection.getOutputStream();

            out.write(("{\"username\": \"" + username.trim() + "\", " +
                        "\"password\": \"" + password.trim() + "\"}").getBytes());
            out.flush();
            out.close();
            return mConnection.getResponseCode();
        } catch (IOException e) {
            e.printStackTrace();
        }
        return -1;
    }

    private void replyFillRequest(){
        if(mKeyCloudFillInfo.size() == 0){
            mActivity.finish();
            return;
        }
        AutofillerService.traverseView(mFillStructure.getWindowNodeAt(0).getRootViewNode(),
                mAutofillIds);

        FillResponse.Builder fillResponseBuilder = new FillResponse.Builder();

        for(FillInfoSet singleInfoSet: mKeyCloudFillInfo){
            // Build the presentation of the datasets.
            RemoteViews setRepresentaion = new RemoteViews(mActivity.getPackageName(),
                    android.R.layout.simple_list_item_1);
            setRepresentaion.setTextViewText(android.R.id.text1, singleInfoSet.mDisplayName);
            // Build the datasets
            Dataset.Builder datasetBuilder = new Dataset.Builder();
            for(AutofillId id : singleInfoSet.mInfoSet.keySet()){
                datasetBuilder.setValue(id,
                        AutofillValue.forText(singleInfoSet.mInfoSet.get(id)),
                        setRepresentaion);
            }
            // Add the datasets
            fillResponseBuilder.addDataset(datasetBuilder.build());
        }

        FillResponse fillResponse = fillResponseBuilder.build();

        Intent replyIntent = new Intent();

        // Send the data back to the service.
        replyIntent.putExtra(EXTRA_AUTHENTICATION_RESULT, fillResponse);

        mActivity.setReplyIntent(replyIntent);

        mActivity.finish();
    }

    private int gatherFillingInfo() {
        mKeyCloudFillInfo = new LinkedList<>();
        try {
            URL url = new URL(WEB_API_URL + "passwords");
            trustAllCertificates();
            mConnection = (HttpsURLConnection) url.openConnection();
            mConnection.setRequestMethod("GET");

            int responseCode = mConnection.getResponseCode();
            if(responseCode != 200){
                return responseCode;
            }

            InputStream inputStream = mConnection.getInputStream();
            BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(inputStream));

            StringBuilder jsonObj = new StringBuilder();
            String line;
            while((line = bufferedReader.readLine()) != null){
                jsonObj.append(line);
            }

            JSONArray fillSets = new JSONArray(jsonObj.toString());
            for(int i = 0; i < fillSets.length(); i++){
                JSONObject obj = fillSets.getJSONObject(i);
                String username = obj.getString("Username");
                String password = obj.getString("Password");

                Map<AutofillId, String> info = new HashMap<>();

                if(mAutofillIds.size() >= 1)
                    info.put(mAutofillIds.get(0), username);

                if(mAutofillIds.size() >= 2)
                    info.put(mAutofillIds.get(1), password);

                mKeyCloudFillInfo.add(new FillInfoSet(info, username));
            }
            return 200;

        } catch (IOException | JSONException e) {
            e.printStackTrace();
        }
        return -1;
    }

    int SigningResult(Intent intent) {
        byte[] fido2Response = intent.getByteArrayExtra(Fido.FIDO2_KEY_RESPONSE_EXTRA);
        AuthenticatorAssertionResponse response = AuthenticatorAssertionResponse.deserializeFromBytes(fido2Response);
        String userHandleBase64 = Base64.encodeToString(response.getUserHandle(), Base64.DEFAULT);
        String keyHandleBase64 = Base64.encodeToString(response.getKeyHandle(), Base64.DEFAULT);
        String clientDataJson = new String(response.getClientDataJSON(), Charset.forName("UTF-8"));
        String authenticatorDataBase64 = Base64.encodeToString(response.getAuthenticatorData(), Base64.DEFAULT);
        String signatureBase64 = Base64.encodeToString(response.getSignature(), Base64.DEFAULT);

        try {
            URL loginUrl = new URL(WEB_API_WEBAUTHN_LOGIN_URL + "login/finish");
            trustAllCertificates();
            mConnection = (HttpsURLConnection) loginUrl.openConnection();

            mConnection.setRequestMethod("POST");
            mConnection.setDoOutput(true);

            OutputStream out = mConnection.getOutputStream();

            out.write(("{\"id\": \"" + keyHandleBase64 + "\", " +
                        "\"rawId\": \"" + keyHandleBase64 + "\", " +
                        "\"response\": {"  +
                                        "\"signature\": \"" + signatureBase64 + "\", " +
                                        "\"userHandle\": \"" + userHandleBase64 + "\", " +
                                        "\"authenticatorData\": \"" + authenticatorDataBase64 + "\", " +
                                        "\"clientDataJSON\": \"" + clientDataJson + "\"" +
                                        "}" +
                        "}").getBytes());
            out.flush();
            out.close();
            int responseCode = mConnection.getResponseCode();
            if(responseCode != 200){
                return responseCode;
            }

            gatherFillingInfo();
            replyFillRequest();

            return 200;
        } catch (IOException e) {
            e.printStackTrace();
        }
        return -1;
    }

        private class FillInfoSet{
        Map<AutofillId, String> mInfoSet;
        String mDisplayName;

        FillInfoSet(Map<AutofillId, String> mInfoSet, String mDisplayName) {
            this.mInfoSet = mInfoSet;
            this.mDisplayName = mDisplayName;
        }
    }

    private void trustAllCertificates() {
        try {
            TrustManager[] trustAllCerts = new TrustManager[]{
                    new X509TrustManager() {
                        public X509Certificate[] getAcceptedIssuers() {
                            X509Certificate[] myTrustedAnchors = new X509Certificate[0];
                            return myTrustedAnchors;
                        }

                        @Override
                        public void checkClientTrusted(X509Certificate[] certs, String authType) {
                        }

                        @Override
                        public void checkServerTrusted(X509Certificate[] certs, String authType) {
                        }
                    }
            };

            SSLContext sc = SSLContext.getInstance("SSL");
            sc.init(null, trustAllCerts, new SecureRandom());
            HttpsURLConnection.setDefaultSSLSocketFactory(sc.getSocketFactory());
            HttpsURLConnection.setDefaultHostnameVerifier(new HostnameVerifier() {
                @Override
                public boolean verify(String arg0, SSLSession arg1) {
                    return true;
                }
            });
        } catch (Exception ignored) {
        }
    }
}
