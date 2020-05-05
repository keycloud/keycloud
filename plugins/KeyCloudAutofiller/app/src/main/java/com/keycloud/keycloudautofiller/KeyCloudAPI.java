package com.keycloud.keycloudautofiller;

import android.app.assist.AssistStructure;
import android.content.Intent;
import android.service.autofill.Dataset;
import android.service.autofill.FillResponse;
import android.view.autofill.AutofillId;
import android.view.autofill.AutofillValue;
import android.widget.RemoteViews;
import android.widget.Toast;

import java.io.IOException;
import java.io.OutputStream;
import java.net.CookieHandler;
import java.net.CookieManager;
import java.net.HttpURLConnection;
import java.net.URL;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

import static android.view.autofill.AutofillManager.EXTRA_AUTHENTICATION_RESULT;

class KeyCloudAPI {
    private KeyCloudAuthenticationActivity mActivity;

    private static final String WEB_API_URL = "http://10.0.2.2:8080/";
    private static final String WEB_API_STANDARD_LOGIN_URL = WEB_API_URL + "standard/login";
    private static final String WEB_API_WEBAUTHN_LOGIN_URL = WEB_API_URL + "webauthn/";

    private HttpURLConnection mConnection;

    private List<AutofillId> mAutofillIds  = new LinkedList<>();;
    private AssistStructure mFillStructure;
    private String mQuery;

    private List<FillInfoSet> mKeyCloudFillInfo;

    KeyCloudAPI(KeyCloudAuthenticationActivity activity, AssistStructure structure){
        this.mActivity = activity;
        this.mFillStructure = structure;
        if (structure != null)
            this.mQuery = mFillStructure.getActivityComponent().toString();

        // Preparation for login: Create a cookie manager
        CookieManager manager = new CookieManager();
        CookieHandler.setDefault(manager);
    }

    void fillOutFromKeyCloud(String username){
        this.loginWithWebAuthn(username);
        this.gatherFillingInfo();
        this.replyFillRequest();
    }

    void fillOutFromKeyCloud(final String username, final String password){
        final KeyCloudAPI self = this;
        new Thread(new Runnable() {
            @Override
            public void run() {
                if(self.loginWithPassword(username, password) != 200){
                    self.mActivity.runOnUiThread(new Runnable() {
                        @Override
                        public void run() {
                            Toast.makeText(self.mActivity, "Wrong Credentials", Toast.LENGTH_SHORT).show();
                        }
                    });
                }
                self.gatherFillingInfo();
                self.replyFillRequest();
            }
        }).start();
    }

    private void loginWithWebAuthn(String username){
        //Todo: Login to keycloud with WebAuthn and username
        Toast.makeText(
                mActivity.getApplicationContext(),
                "Not yet Implemented",
                Toast.LENGTH_LONG
        ).show();
    }

    private int loginWithPassword(String username, String password){
        try {

            URL loginUrl = new URL(WEB_API_STANDARD_LOGIN_URL);
            mConnection = (HttpURLConnection) loginUrl.openConnection();

            mConnection.setRequestMethod("POST");
            mConnection.setDoOutput(true);

            OutputStream out = mConnection.getOutputStream();

            out.write(("{\"username\": \"test\", \"password\": \"qnNQmCnZhHcJfDwJ\"}").getBytes());
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

    private void gatherFillingInfo() {
        mKeyCloudFillInfo = new LinkedList<>();
        try {
            URL url = new URL(WEB_API_URL + "password?query...");
            mConnection = (HttpURLConnection) url.openConnection();
            mConnection.setRequestMethod("GET");

            //Parse json response...
            //Todo: Interact with KeyCloud to get stored username, password, etc. for mQuery
            //      (Component name as url) and fill mKeyCloudFillInfo with Information

        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    private class FillInfoSet{
        Map<AutofillId, String> mInfoSet;
        String mDisplayName;

        public FillInfoSet(Map<AutofillId, String> mInfoSet, String mDisplayName) {
            this.mInfoSet = mInfoSet;
            this.mDisplayName = mDisplayName;
        }
    }
}
