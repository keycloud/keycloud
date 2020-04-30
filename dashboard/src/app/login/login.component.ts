import { Component, OnInit } from '@angular/core';
import {UserService} from '../services/user.service';
import {Router} from '@angular/router';
import {MatSnackBar} from '@angular/material/snack-bar';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  username: string;
  password: string;
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
    this.userService.login()
      .subscribe( data => {
        this.router.navigate(['/dashboard']);
      }, error => {
        this.popOver.open(`Something went wrong! If this error persists, please contact us with the following error: ${error}`,
          '', {duration: 5000});
        this.loginLoading = false;
      });
  }

  register() {
    this.registerLoading = true;
    this.userService.register(this.body)
      .subscribe( data => {
        this.router.navigate(['/settings'], { queryParams: { firstVisit: true}}); // Upon first visit navigate to register
        // and show the generated pw TODO add handler for generated pw
      }, error => {
        this.popOver.open(`Something went wrong! If this error persists, please contact us with the following error: ${error}`,
          '', {duration: 5000});
        this.registerLoading = false;
      });
  }
}
