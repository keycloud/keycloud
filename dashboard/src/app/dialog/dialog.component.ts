import {Component, OnInit} from '@angular/core';
import {MatDialogRef} from '@angular/material/dialog';
import * as passwordGenerator from '../util/pwgen';

@Component({
  selector: 'app-dialog',
  templateUrl: './dialog.component.html',
  styleUrls: ['./dialog.component.css']
})
export class DialogComponent implements OnInit {

  username: string;
  url: string;
  password: string;

  constructor(
    private dialogRef: MatDialogRef<DialogComponent>,
  ) {
  }

  ngOnInit() {
  }

  close() {
    this.dialogRef.close();
  }

  generatePassword() {
    this.password = passwordGenerator.genPW();
  }
}
