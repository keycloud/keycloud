import {Component, OnInit} from '@angular/core';
import {UserService} from '../services/user.service';
import {Router} from '@angular/router';
import {MatSnackBar} from '@angular/material/snack-bar';
import {Decoder} from '../util/decoder';
import {User} from '../models/user';
import {UserRegister} from '../models/user-register';

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
  use2FA = false;
  loginLoading = false;
  registerLoading = false;
  body: any;

  constructor(
    private userService: UserService,
    private router: Router,
    private popOver: MatSnackBar,
    private decoder: Decoder,
  ) {
  }

  ngOnInit() {
  }

  login() {
    if (this.use2FA) {
      this.body = new UserRegister(this.loginUsername, '');
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
                username: this.loginUsername,
                mail: '',
                id: credential.id,
                // @ts-ignore
                rawId: this._encodeBuffer(credential.rawId),
                response: {
                  // @ts-ignore
                  clientDataJSON: this.decoder._encodeBuffer(credential.response.clientDataJSON),
                  // @ts-ignore
                  authenticatorData: this.decoder._encodeBuffer(credential.response.authenticatorData),
                  // @ts-ignore
                  signature: this.decoder._encodeBuffer(credential.response.signature),
                  // @ts-ignore
                  userHandle: this.decoder._encodeBuffer(credential.response.userHandle),
                },
                type: credential.type,
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
              this.openErrorPopOver(error);
            });
        }
      );
    } else {
      this.loginLoading = true;
      this.body = new User(this.loginUsername, this.password);
      this.userService.login(this.body)
        .subscribe(data => {
          if (data.status === 200) {
            this.router.navigate(['/dashboard']);
          } else {
            console.log(data);
          }
        }, error => {
          this.openErrorPopOver(error);
          this.loginLoading = false;
        });
    }
  }

  register() {
    this.registerLoading = true;
    this.body = new UserRegister(this.registerUsername, this.email);
    this.userService.register(this.body)
      .subscribe(resp => {
        if (resp.status === 200) {
          this.router.navigate(['/settings'], {queryParams: {firstVisit: true}});
        }
      }, error => {
        this.openErrorPopOver(error);
        this.registerLoading = false;
      });
  }

  openErrorPopOver(error) {
    this.popOver.open(`Something went wrong! If this error persists, please contact us with the following error: ${error}`,
      '', {duration: 5000});
  }
}
