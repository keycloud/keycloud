import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.css']
})
export class SettingsComponent implements OnInit {

  password: string;
  hide = true;

  constructor() { }

  ngOnInit() {
  }

  add2FA() {

  }
}
