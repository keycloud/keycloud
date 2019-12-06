// override some of the default configuration options
// see the docs for a full list of configuration options
const webAuthnConfig = {
    timeout: 30000,
    username: undefined,
    pw: undefined
};
$(document).ready(function() {
    // when user clicks submit in the register form, start the registration process
    $("#register-form").submit(function (event) {
        // register -> get cookie; GET user(cookie) -> userinfo; route to main.html
        event.preventDefault();
        webAuthnConfig.username = $(event.target).children("input[name=username]")[0].value;
        new WebAuthn().register();
    });

    // when user clicks submit in the login form, start the log in process
    $("#login-form").submit(function (event) {
        event.preventDefault();
        webAuthnConfig.username = $(event.target).children("input[name=username]")[0].value;
        webAuthnConfig.pw = $(event.target).children("input[name=password]")[0].value;
		new WebAuthn().login();
    });

});

class WebAuthn {
    // Decode a base64 string into a Uint8Array.
    static _decodeBuffer(value) {
        return Uint8Array.from(atob(value), c => c.charCodeAt(0));
    }

    // Encode an ArrayBuffer into a base64 string.
    static _encodeBuffer(value) {
        return btoa(new Uint8Array(value).reduce((s, byte) => s + String.fromCharCode(byte), ''));
    }

    // Checks whether the status returned matches the status given.
    static _checkStatus(status) {
        return res => {
            if (res.status === status) {
                return res;
            }
            throw new Error(res.statusText);
        };
    }

    register() {
        return fetch('/standard/register', {
            method: 'POST',
            body: JSON.stringify({'username': webAuthnConfig.username}),
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            }
        }).then(function(response){
                if(response.status !== 200){
                    throw new Error(response.statusText);
                }else {
                    new WebAuthn().redirect();
                }
            });
    }

    login() {
        if ($("#use2FA").prop('checked') === false){
            return fetch('/standard/login', {
                method: 'POST',
                credentials: 'same-origin',
                body: JSON.stringify({'username': webAuthnConfig.username, 'password': webAuthnConfig.pw}),
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                }
            }).then(function (response) {
                if(response.status !== 200){
                    throw new Error(response.statusText);
                }else {
                    new WebAuthn().redirect();
                }
            });
        }

        return fetch('/webauthn/login/start', {
            method: 'POST',
            body: JSON.stringify({'username': webAuthnConfig.username}),
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            }
        })
            .then(WebAuthn._checkStatus(200))
            .then(res => res.json())
            .then(res => {
                res.publicKey.challenge = WebAuthn._decodeBuffer(res.publicKey.challenge);
                if (res.publicKey.allowCredentials) {
                    for (let i = 0; i < res.publicKey.allowCredentials.length; i++) {
                        res.publicKey.allowCredentials[i].id = WebAuthn._decodeBuffer(res.publicKey.allowCredentials[i].id);
                    }
                }
                return res;
            })
            .then(res => navigator.credentials.get(res))
            .then(credential => {
                return fetch('/webauthn/login/finish', {
                    method: 'POST',
                    headers: {
                        'Accept': 'application/json',
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        id: credential.id,
                        rawId: WebAuthn._encodeBuffer(credential.rawId),
                        response: {
                            clientDataJSON: WebAuthn._encodeBuffer(credential.response.clientDataJSON),
                            authenticatorData: WebAuthn._encodeBuffer(credential.response.authenticatorData),
                            signature: WebAuthn._encodeBuffer(credential.response.signature),
                            userHandle: WebAuthn._encodeBuffer(credential.response.userHandle),
                        },
                        type: credential.type,
                        username: webAuthnConfig.username
                    }),
                })
            })
            .then(WebAuthn._checkStatus(200));
    }

    redirect() {
        window.location.replace('main.html')
    }
}