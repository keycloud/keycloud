import { Component, OnInit } from '@angular/core';
import {UserService} from '../services/user.service';
import {Router} from '@angular/router';
import {MatSnackBar} from '@angular/material/snack-bar';
import {UsernameEmail} from '../models/username-email';
import {UsernamePassword} from '../models/username-password';

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

  webAuthnConfig = {
    timeout: 30000,
    username: undefined,
    pw: undefined
  };

  constructor(
    private userService: UserService,
    private router: Router,
    private popOver: MatSnackBar,
  ) { }

  ngOnInit() {
  }

  login() {
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

  async register() {
    this.registerLoading = true;
    this.body = new UsernameEmail(this.registerUsername, this.email);
    this.userService.register(this.body)
      .subscribe( resp => {
        if (resp.status === 200) {
          this.router.navigate(['/settings'], { queryParams: { firstVisit: true}});
        }
         // Upon first visit navigate to register
        // and show the generated pw TODO add handler for generated pw
      }, error => {
        console.log(error);
        this.popOver.open(`Something went wrong! If this error persists, please contact us with the following error: ${error.error}`,
          '', {duration: 5000});
        this.registerLoading = false;
      });
  }
}
