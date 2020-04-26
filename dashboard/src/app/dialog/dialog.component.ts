import {Component, Inject, OnInit} from '@angular/core';
import {FormBuilder, FormGroup} from '@angular/forms';
import {MAT_DIALOG_DATA, MatDialog, MatDialogRef} from '@angular/material/dialog';
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
  private generatedPassword: string;
  newEntryForm: any;


  constructor(
    public dialog: MatDialog,
    private fb: FormBuilder,
    private dialogRef: MatDialogRef<DialogComponent>,
    @Inject(MAT_DIALOG_DATA) data,
  ) {
    this.description = data.description;
  }

  ngOnInit() {
    this.form = this.fb.group({
      description: [this.description, []],
    });
  }

  openDialog() {
    const dialogRef = this.dialog.open(DialogComponent, {
      width: '250px',
      data: {
        username: this.username,
        url: this.url,
        password: this.password,
      }
    });

    dialogRef.afterClosed().subscribe(
      data => console.log(`Dialog data: ${data}`)
    );
  }

  save() {
    this.dialogRef.close(this.form.value);
  }

  close() {
    this.dialogRef.close();
  }

  generatePassword() {
    this.generatedPassword = passwordGenerator.genPW();
  }


}
