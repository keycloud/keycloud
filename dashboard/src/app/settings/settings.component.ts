import { Component, OnInit } from '@angular/core';
import {Router} from '@angular/router';
import {UserService} from '../services/user.service';
import {CrudService} from '../services/crud.service';
import {UsernameMasterPassword} from '../models/username-master-password';
import {UsernameEmail} from '../models/username-email';
import {MatSnackBar} from '@angular/material/snack-bar';

@Component({
  selector: 'app-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.css']
})
export class SettingsComponent implements OnInit {

  password: string;
  hide = true;
  user: UsernameMasterPassword;

  constructor(
    private router: Router,
    private userService: UserService,
    private popOver: MatSnackBar,
  ) {
    this.getUser();
  }

  ngOnInit() {
  }

  add2FA() {
    const body = new UsernameEmail(this.user.Name, '');
    this.userService.webauthnRegistrationStart(body).subscribe(
      resp => {
        const respBody = JSON.parse(resp.body);
        console.log(respBody);
        respBody.publicKey.challenge = this._decodeBuffer(respBody.publicKey.challenge);
        if (respBody.publicKey.allowCredentials) {
          // tslint:disable-next-line:prefer-for-of
          for (let i = 0; i < respBody.publicKey.allowCredentials.length; i++) {
            respBody.publicKey.allowCredentials[i].id = this._decodeBuffer(respBody.publicKey.allowCredentials[i].id);
          }
        }
        navigator.credentials.get(respBody)
          .then(credential => {
            const requestBody = {
              id: credential.id,
              // @ts-ignore
              rawId: this._encodeBuffer(credential.rawId),
              response: {
                // @ts-ignore
                clientDataJSON: this._encodeBuffer(credential.response.clientDataJSON),
                // @ts-ignore
                authenticatorData: this._encodeBuffer(credential.response.authenticatorData),
                // @ts-ignore
                signature: this._encodeBuffer(credential.response.signature),
                // @ts-ignore
                userHandle: this._encodeBuffer(credential.response.userHandle),
              },
              type: credential.type,
              username: this.user.Name,
            };
            this.userService.webauthnRegistrationFinish(requestBody).subscribe(
              // tslint:disable-next-line:no-shadowed-variable
              resp => {
                if (resp.status === 200) {
                  this.popOver.open('Success! Your account is now secured via a second factor.', '', {duration: 5000});
              }
              }
            );
        })
          .catch(error => {
            console.log(error);
            this.popOver.open(`Something went wrong! If this error persists, please contact us with the following error: ${error.error}`,
              '', {duration: 5000});
        });
      }
    );
  }

  private getUser() {
    this.userService.getUser().subscribe(
      resp => {
        resp = JSON.parse(resp.body);
        console.log(resp);
        this.user = new UsernameMasterPassword(resp.Name, resp.MasterPassword);
        this.password = this.user.MasterPassword;
      }
    );
  }

  private _decodeBuffer(value: string) {
    return Uint8Array.from(atob(value), c => c.charCodeAt(0));
  }

  // Encode an ArrayBuffer into a base64 string.
  private _encodeBuffer(value: Iterable<number>) {
    return btoa(new Uint8Array(value).reduce((s, byte) => s + String.fromCharCode(byte), ''));
  }
}
