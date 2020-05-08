import { Component, OnInit } from '@angular/core';
import {UserService} from '../services/user.service';
import {Router} from '@angular/router';
import {MatSnackBar} from '@angular/material/snack-bar';
import {UsernameEmail} from '../models/username-email';
import {UsernamePassword} from '../models/username-password';
import {Decoder} from '../util/decoder';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  loginUsername: string;
  registerUsername: string;
  password: string;
  email: string;
  use2FA: boolean;
  loginLoading = false;
  registerLoading = false;
  body: any;

  constructor(
    private userService: UserService,
    private router: Router,
    private popOver: MatSnackBar,
    private decoder: Decoder,
  ) { }

  ngOnInit() {
  }

  login() {
    if (this.use2FA) {
      this.body = new UsernameEmail(this.loginUsername, '');
      this.userService.webauthnLoginStart(this.body).subscribe(
        resp => {
          const respBody = JSON.parse(resp.body);
          respBody.publicKey.challenge = this.decoder._decodeBuffer(respBody.publicKey.challenge);
          if (respBody.publicKey.allowCredentials) {
            // tslint:disable-next-line:prefer-for-of
            for (let i = 0; i < respBody.publicKey.allowCredentials.length; i++) {
              respBody.publicKey.allowCredentials[i].id = this.decoder._decodeBuffer(respBody.publicKey.allowCredentials[i].id);
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
                username: this.loginUsername,
              };
              this.userService.webauthnLoginFinish(requestBody).subscribe(
                // tslint:disable-next-line:no-shadowed-variable
                resp => {
                  if (resp.status === 200) {
                    this.router.navigate(['/dashboard']);
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
    } else {
      this.loginLoading = true;
      this.body = new UsernamePassword(this.loginUsername, this.password);
      this.userService.login(this.body)
        .subscribe( data => {
          if (data.status === 200) {
            this.router.navigate(['/dashboard']);
          } else {
            console.log(data);
          }
        }, error => {
          this.popOver.open(`Something went wrong! If this error persists, please contact us with the following error: ${error}`,
            '', {duration: 5000});
          this.loginLoading = false;
        });
    }
  }

  register() {
    this.registerLoading = true;
    this.body = new UsernameEmail(this.registerUsername, this.email);
    this.userService.register(this.body)
      .subscribe( resp => {
        if (resp.status === 200) {
          this.router.navigate(['/settings'], { queryParams: { firstVisit: true}});
        }
      }, error => {
        console.log(error);
        this.popOver.open(`Something went wrong! If this error persists, please contact us with the following error: ${error.error}`,
          '', {duration: 5000});
        this.registerLoading = false;
      });
  }
}
