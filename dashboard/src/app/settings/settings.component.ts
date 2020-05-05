import { Component, OnInit } from '@angular/core';
import {Router} from '@angular/router';
import {UserService} from '../services/user.service';

@Component({
  selector: 'app-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.css']
})
export class SettingsComponent implements OnInit {

  password: string;
  hide = true;

  constructor(
    private router: Router,
    private userService: UserService,
  ) {
    this.getPassword();
  }

  ngOnInit() {
  }

  add2FA() {

  }

  private getPassword() {
    this.userService.getUser().subscribe(
      resp => {
        resp = JSON.parse(resp.body);
        this.password = resp.MasterPassword;
      }
    );
  }
}
