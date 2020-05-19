package com.keycloud.keycloudautofiller;

import android.app.assist.AssistStructure;
import android.content.Intent;
import android.os.Bundle;
import android.support.v7.app.AppCompatActivity;
import android.view.View;
import android.widget.Button;
import android.widget.CheckBox;
import android.widget.CompoundButton;
import android.widget.TextView;

import static android.view.autofill.AutofillManager.EXTRA_ASSIST_STRUCTURE;

public class KeyCloudAuthenticationActivity extends AppCompatActivity {

    private TextView mUsername;
    private TextView mPassword;
    private CheckBox mUseWebauthn;

    private Intent mReplyIntent = null;

    private KeyCloudAPI mAPI;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_key_cloud_authentication);

        mUsername = findViewById(R.id.username_inpt);
        mPassword = findViewById(R.id.password_inpt);
        mUseWebauthn = findViewById(R.id.webauthn_box);
        Button mLogin = findViewById(R.id.login_btn);

        mUseWebauthn.setOnCheckedChangeListener(new CompoundButton.OnCheckedChangeListener() {
            @Override
            public void onCheckedChanged(CompoundButton buttonView, boolean isChecked) {
                mPassword.setEnabled(!isChecked);
            }
        });

        mLogin.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                String username = mUsername.getText().toString();
                if (mUseWebauthn.isChecked()){
                    mAPI.fillOutFromKeyCloud(username);
                }else {
                    String password = mPassword.getText().toString();
                    mAPI.fillOutFromKeyCloud(username, password);
                }
            }
        });

        AssistStructure mAutofillStructure = getIntent().getParcelableExtra(EXTRA_ASSIST_STRUCTURE);
        if (mAutofillStructure != null)
            mAPI = new KeyCloudAPI(this, mAutofillStructure);
    }

    @Override
    public void finish() {
        if (mReplyIntent != null) {
            setResult(RESULT_OK, mReplyIntent);
        } else {
            setResult(RESULT_CANCELED);
        }
        super.finish();
    }

    public void setReplyIntent(Intent replyIntent){
        mReplyIntent = replyIntent;
    }
}
