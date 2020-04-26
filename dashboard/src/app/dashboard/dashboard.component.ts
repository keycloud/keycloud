import { Component, OnInit } from '@angular/core';
import {$} from 'protractor';
import * as passwordGenerator from '../util/pwgen.js';
import {FormBuilder, FormGroup} from '@angular/forms';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {

  generatedPassword: string;
  newEntryForm: FormGroup;

  header = ['i', 'Username', 'Url', 'Password', 'Delete'];

  exampleEntries = [
    {
      i : 0,
      Username : 'Mark',
      Url : 'https://www.google.com',
      Password : 'example',
    }, {
      i : 1,
      Username : 'Jacob',
      Url : 'https://www.google.com',
      Password : 'example',
    }, {
      i : 2,
      Username : 'Larry',
      Url : 'https://www.google.com',
      Password : 'example',
    }
  ];

  constructor(
    private formBuilder: FormBuilder
  ) {
  }

  ngOnInit() {
    this.newEntryForm = this.formBuilder.group({
      usernameInput: [],
      urlInput: []
    });
  }

  get f() { return this.newEntryForm.controls; }

  addCustomField() {
    // TODO: This JQeury function cannot be used, causes error
    /*$(`<div class="form-row custom-field-row-added" style="margin-bottom: 15px">
                                    <div class="col">
                                        <input type="text" class="form-control" placeholder="Custom Field Name">
                                    </div>
                                    <div class="col">
                                        <input type="text" class="form-control" placeholder="Custom Field content">
                                    </div>
                                    <div class="col">
                                        <input class="form-check-input big-checkbox" type="checkbox">
                                        <label class="form-check-label" style="font-size: x-large;margin-left: 15px;">
                                            Encrypt
                                        </label>
                                    </div>
                                </div>`).insertBefore('#btn-add-field-group');*/
  }

  updateModal() {
    // TODO: This JQeury function cannot be used, causes error
    // $('.custom-field-row-added').remove();
  }

  removeEntry(id) {
    console.log(`remove pressed with id ${id}`);
    this.exampleEntries.splice(id, 1);
    // TODO: This JQeury function cannot be used, causes error
      // $('.toast').toast('show');
  }

  generatePassword() {
    this.generatedPassword = passwordGenerator.genPW();
  }

  copyToClipboard(id) {
    console.log(`copy pressed with id ${id}`);
    const pw = this.exampleEntries[id].Password;
    const selBox = document.createElement('textarea');
    selBox.style.position = 'fixed';
    selBox.style.left = '0';
    selBox.style.top = '0';
    selBox.style.opacity = '0';
    selBox.value = pw;
    document.body.appendChild(selBox);
    selBox.focus();
    selBox.select();
    document.execCommand('copy');
    document.body.removeChild(selBox);
    window.alert('Copied!');
  }

  saveNewEntry() {
    const newEntry = {
      i : this.exampleEntries.length,
      Username : '',
      Url : '',
      Password : '',
    };
    newEntry.Username = this.f.usernameInput.value;
    newEntry.Url = this.f.urlInput.value;
    newEntry.Password = this.f.pwInput.value;
    this.exampleEntries.push(newEntry);
  }

}
