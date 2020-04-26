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
    this.renderTable();
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

  addTableRow(value) {
    // TODO: This JQeury function cannot be used, causes error
    /*$('#pwTable').prepend(`<tr class="entry">
    <th scope="row">${value.i}</th>
    <td>${value.Username}</td>
    <td><a target="_blank" rel="noopener noreferrer" href="${value.Url}">${value.Url}</a></td>
    <td><button type="button" class="btn btn-info" id="cp${value.i}" onclick="copyToClipboard(this.id)">
    <i class="fa fa-clipboard" ></i> Copy to Clipboard</button></td>
    <td><button type="button" class="btn btn-danger" id="rm${value.i}" onclick="removeEntry(this.id)">
    <i class="fa fa-remove"></i></button></td>
    </tr>`);*/
  }

  updateModal() {
    // TODO: This JQeury function cannot be used, causes error
    // $('.custom-field-row-added').remove();
  }

  removeEntry(id) {
      this.exampleEntries.splice(id.slice(2, ), 1);
      this.renderTable(); // works so far, needs some work done on the indices
    // TODO: This JQeury function cannot be used, causes error
      // $('.toast').toast('show');
  }

  generatePassword() {
    this.generatedPassword = passwordGenerator.genPW();
  }

  copyToClipboard(id) {
    const pw = this.exampleEntries[id.slice(2, )].Password;
    window.prompt('Copy to clipboard: Ctrl+C', pw); // Workaround for now, could use other, prettier techniques
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
    this.renderTable();
  }

  renderTable() {
    // TODO: This JQeury function cannot be used, causes error
    // $('.entry').remove(); // clear table
    this.exampleEntries.reverse();  // bc of callback
    this.exampleEntries.forEach(this.addTableRow);
    this.exampleEntries.reverse();
  }

}
