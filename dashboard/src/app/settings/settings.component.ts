import {Component, OnInit} from '@angular/core';
import {Router} from '@angular/router';
import {UserService} from '../services/user.service';
import {MatSnackBar} from '@angular/material/snack-bar';
import {Decoder} from '../util/decoder';
import {UserRegister} from '../models/user-register';
import {User} from '../models/user';

@Component({
  selector: 'app-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.css']
})
export class SettingsComponent implements OnInit {

  hide = true;
  user: User;
  masterpassword: string;

  constructor(
    private router: Router,
    private userService: UserService,
    private popOver: MatSnackBar,
    private decoder: Decoder,
  ) {
    this.getUser();
  }

  ngOnInit() {
  }

  add2FA() {
    const body = new UserRegister(this.user.username, '');
    this.userService.webauthnRegistrationStart(body).subscribe(
      resp => {
        const respBody = JSON.parse(resp.body);
        respBody.publicKey.challenge = this.decoder._decodeBuffer(respBody.publicKey.challenge);
        respBody.publicKey.user.id = this.decoder._decodeBuffer(respBody.publicKey.user.id);
        if (respBody.publicKey.allowCredentials) {
          // tslint:disable-next-line:prefer-for-of
          for (let i = 0; i < respBody.publicKey.excludeCredentials.length; i++) {
            respBody.publicKey.excludeCredentials[i].id = this.decoder._decodeBuffer(respBody.publicKey.excludeCredentials[i].id);
          }
        }
        navigator.credentials.create(respBody)
          .then(credential => {
            const requestBody = {
              username: this.user.username,
              mail: '',
              id: credential.id,
              // @ts-ignore
              rawId: this.decoder._encodeBuffer(credential.rawId),
              response: {
                // @ts-ignore
                attestationObject: this.decoder._encodeBuffer(credential.response.attestationObject),
                // @ts-ignore
                clientDataJSON: this.decoder._encodeBuffer(credential.response.clientDataJSON),
              },
              type: credential.type,
            };
            this.userService.webauthnRegistrationFinish(requestBody).subscribe(
              // tslint:disable-next-line:no-shadowed-variable
              resp => {
                if (resp.status === 201) {
                  this.popOver.open('Success! Your account is now secured via a second factor.', '', {duration: 5000});
                }
              }
            );
          })
          .catch(error => {
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
        this.user = new User(resp.username, resp.masterpassword);
        this.masterpassword = resp.masterpassword;
      },
      error => {
        if (error.status === 401) {
          this.popOver.open(`Please sign in to retrieve your user information.`,
            '', {duration: 5000});
          this.router.navigate(['/login']);
        }
      }
    );
  }
}
