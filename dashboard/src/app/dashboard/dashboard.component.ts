import {Component, OnInit} from '@angular/core';
import {FormBuilder, FormGroup} from '@angular/forms';
import {MatDialog, MatDialogConfig} from '@angular/material/dialog';
import {DialogComponent} from '../dialog/dialog.component';
import {MatSnackBar} from '@angular/material/snack-bar';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {

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
    private dialog: MatDialog,
    private popOver: MatSnackBar,
  ) {
  }

  ngOnInit() { }

  openDialog() {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = true;
    dialogConfig.autoFocus = true;

    const dialogRef = this.dialog.open(DialogComponent, dialogConfig);

    dialogRef.afterClosed().subscribe(
      data => this.saveNewEntry(data)
    );
  }

  removeEntry(id) {
    console.log(`remove pressed with id ${id}`);
    this.exampleEntries.splice(id, 1);
    this.popOver.open('Deleted!', '', {duration: 2000});
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
    this.popOver.open('Copied!', '', {duration: 2000});
  }

  saveNewEntry(data) {
    const newEntry = {
      i : this.exampleEntries.length,
      Username : data.username,
      Url : data.url,
      Password : data.password,
    };
    this.exampleEntries.push(newEntry);
  }

}
