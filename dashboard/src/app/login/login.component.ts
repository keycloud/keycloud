import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  username: string;
  password: string;
  use2FA: boolean;

  webAuthnConfig = {
    timeout: 30000,
    username: undefined,
    pw: undefined
  };

  constructor() { }

  ngOnInit() {
  }

  login() {
    console.log(`${this.username} ${this.password} ${this.use2FA}`);
  }

  register() {

  }
}
