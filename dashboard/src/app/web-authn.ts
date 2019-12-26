export class WebAuthn {
  // Decode a base64 string into a Uint8Array.
  _decodeBuffer(value: string) {
    return Uint8Array.from(atob(value), c => c.charCodeAt(0));
  }

  // Encode an ArrayBuffer into a base64 string.
  _encodeBuffer(value: Iterable<number>) {
    return btoa(new Uint8Array(value).reduce((s, byte) => s + String.fromCharCode(byte), ''));
  }

  // Checks whether the status returned matches the status given.
  _checkStatus(status: any) {
    return res => {
      if (res.status === status) {
        return res;
      }
      throw new Error(res.statusText);
    };
  }

  register(config: any) {
    return fetch('/standard/register', {
      method: 'POST',
      body: JSON.stringify({username: config.username}),
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json'
      }
    }).then(response => {
      if(response.status !== 200) {
        throw new Error(response.statusText);
      } else {
        // redirect to dashboard
      }
    });
  }

  login(config: any, use2FA: boolean) {
    if (use2FA === false) {
      return fetch('/standard/login', {
        method: 'POST',
        credentials: 'same-origin',
        body: JSON.stringify({username: config.username, password: config.pw}),
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json'
        }
      }).then(response => {
        if(response.status !== 200) {
          throw new Error(response.statusText);
        } else {
          // redirect to dashboard
        }
      });
    }

    // @ts-ignore
    return fetch('/webauthn/login/start', {
      method: 'POST',
      body: JSON.stringify({username: config.username}),
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json'
      }
    })
      .then(this._checkStatus(200))
      .then(res => res.json())
      .then(res => {
        res.publicKey.challenge = this._decodeBuffer(res.publicKey.challenge);
        if (res.publicKey.allowCredentials) {
          // tslint:disable-next-line:prefer-for-of
          for (let i = 0; i < res.publicKey.allowCredentials.length; i++) {
            res.publicKey.allowCredentials[i].id = this._decodeBuffer(res.publicKey.allowCredentials[i].id);
          }
        }
        return res;
      })
      .then(res => navigator.credentials.get(res))
      .then(credential => {
        return fetch('/webauthn/login/finish', {
          method: 'POST',
          headers: {
            Accept: 'application/json',
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            id: credential.id,
            rawId: this._encodeBuffer(credential.rawId),
            response: {
              clientDataJSON: this._encodeBuffer(credential.response.clientDataJSON),
              authenticatorData: this._encodeBuffer(credential.response.authenticatorData),
              signature: this._encodeBuffer(credential.response.signature),
              userHandle: this._encodeBuffer(credential.response.userHandle),
            },
            type: credential.type,
            username: config.username
          }),
        });
      })
      .then(this._checkStatus(200));
  }
}
