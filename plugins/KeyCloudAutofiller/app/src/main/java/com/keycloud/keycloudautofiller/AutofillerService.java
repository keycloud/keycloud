package com.keycloud.keycloudautofiller;

import android.app.PendingIntent;
import android.app.assist.AssistStructure;
import android.content.Intent;
import android.content.IntentSender;
import android.os.CancellationSignal;
import android.service.autofill.AutofillService;
import android.service.autofill.FillCallback;
import android.service.autofill.FillContext;
import android.service.autofill.FillRequest;
import android.service.autofill.FillResponse;
import android.service.autofill.SaveCallback;
import android.service.autofill.SaveRequest;
import android.support.annotation.NonNull;
import android.util.Log;
import android.view.autofill.AutofillId;
import android.widget.RemoteViews;

import java.util.LinkedList;
import java.util.List;

public class AutofillerService extends AutofillService {

    @Override
    public void onFillRequest(@NonNull FillRequest request, @NonNull CancellationSignal cancellationSignal, @NonNull FillCallback callback) {
        // Get the structure from the request
        List<FillContext> context = request.getFillContexts();
        AssistStructure structure = context.get(context.size() - 1).getStructure();
        Log.i("FILLREQUEST", "FROM COMPONENT:" + structure.getActivityComponent().toString());

        List<AutofillId> autofillIds = new LinkedList<>();
        for(int i = 0; i < structure.getWindowNodeCount(); i++)
            traverseView(structure.getWindowNodeAt(i).getRootViewNode(), autofillIds);

        if(autofillIds.size() == 0){
            callback.onFailure("No ID detected");
            return;
        }
        RemoteViews authPresentation = new RemoteViews(getPackageName(), R.layout.multidataset_list_item);
        authPresentation.setTextViewText(R.id.text, "WebAuthn");
        authPresentation.setImageViewResource(R.id.icon, R.drawable.ic_lock_black_24dp);

        Intent authIntent = new Intent(this, KeyCloudAuthenticationActivity.class);

        // Send any additional data required to complete the request.
        IntentSender intentSender = PendingIntent.getActivity(
                this,
                0,
                authIntent,
                PendingIntent.FLAG_CANCEL_CURRENT
        ).getIntentSender();

        // Build a FillResponse object that requires authentication.
        FillResponse fillResponse = new FillResponse.Builder()
                .setAuthentication(autofillIds.toArray(new AutofillId[autofillIds.size()]), intentSender, authPresentation)
                .build();

        callback.onSuccess(fillResponse);
    }

    @Override
    public void onSaveRequest(@NonNull SaveRequest request, @NonNull SaveCallback callback) {
        //Not implemented saving of new passwords yet
    }

    public static void traverseView(AssistStructure.ViewNode node, List<AutofillId> autofillIds){
        String[] hints = node.getAutofillHints();
        if (hints != null){
            autofillIds.add(node.getAutofillId());
        }
        for(int i = 0; i < node.getChildCount(); i++){
            traverseView(node.getChildAt(i), autofillIds);
        }

    }
}
