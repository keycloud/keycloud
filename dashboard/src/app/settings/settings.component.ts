import { Component, OnInit } from '@angular/core';
import {Router} from '@angular/router';
import {UserService} from '../services/user.service';
import {CrudService} from '../services/crud.service';

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
    private crudService: CrudService,
  ) {
    this.getPassword();
  }

  ngOnInit() {
  }

  add2FA() {

  }

  private getPassword() {
    this.crudService.getUser().subscribe(
      resp => {
        resp = JSON.parse(resp.body);
        this.password = resp.MasterPassword;
      }
    );
  }
}
