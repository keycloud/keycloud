import {Component, OnInit} from '@angular/core';
import * as passwordGenerator from '../util/pwgen.js';
import {FormBuilder, FormGroup} from '@angular/forms';
import {MatDialog, MatDialogConfig} from '@angular/material/dialog';
import {DialogComponent} from '../dialog/dialog.component';

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
    private formBuilder: FormBuilder,
    private dialog: MatDialog,
  ) {
  }

  ngOnInit() {
    this.newEntryForm = this.formBuilder.group({
      usernameInput: [],
      urlInput: []
    });
  }

  get f() { return this.newEntryForm.controls; }

  openDialog() {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = true;
    dialogConfig.autoFocus = true;
    dialogConfig.data = {
      id: 1,
      title: 'Angular For Beginners'
    };

    this.dialog.open(DialogComponent, dialogConfig);

    const dialogRef = this.dialog.open(DialogComponent, dialogConfig);

    dialogRef.afterClosed().subscribe(
      data => console.log(`Dialog output: ${data}`)
    );
  }

  removeEntry(id) {
    console.log(`remove pressed with id ${id}`);
    this.exampleEntries.splice(id, 1);
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
