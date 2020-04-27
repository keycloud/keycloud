import {Component, Inject, OnInit} from '@angular/core';
import {FormBuilder, FormGroup} from '@angular/forms';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import * as passwordGenerator from '../util/pwgen';

export interface DialogData {
  username: string;
  url: string;
  password: string;
}

@Component({
  selector: 'app-dialog',
  templateUrl: './dialog.component.html',
  styleUrls: ['./dialog.component.css']
})
export class DialogComponent implements OnInit {

  username: string;
  url: string;
  password: string;

  form: FormGroup;
  description: string;
  generatedPassword: string;


  constructor(
    private fb: FormBuilder,
    private dialogRef: MatDialogRef<DialogComponent>,
    @Inject(MAT_DIALOG_DATA) data,
  ) {}

  ngOnInit() {}

  close() {
    this.dialogRef.close();
  }

  generatePassword() {
    this.generatedPassword = passwordGenerator.genPW();
  }


}
