import {Component, OnInit} from '@angular/core';
import {MatDialog, MatDialogConfig, MatDialogRef} from '@angular/material/dialog';
import {DialogComponent} from '../dialog/dialog.component';
import {MatSnackBar} from '@angular/material/snack-bar';
import {UserService} from '../services/user.service';
import {CrudService} from '../services/crud.service';
import {PasswordEntry} from '../models/password-entry';
import {Router} from '@angular/router';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {

  displayedColumns: string[] = ['No.', 'Username', 'Password', 'Url', 'Delete'];
  dataSource: PasswordEntry[] = [];

  constructor(
    private dialog: MatDialog,
    private popOver: MatSnackBar,
    private userService: UserService,
    private crudService: CrudService,
    private router: Router
  ) {
    this.initDashboard();
  }

  ngOnInit() {
  }

  initDashboard() {
    this.crudService.getListOfPasswords().subscribe(
      resp => {
        if (resp.status === 200) {
          const body = JSON.parse(resp.body);
          body.forEach(item => {
            const newEntry = new PasswordEntry(item.id, item.username, item.password, item.url, '');
            this.dataSource.push(newEntry);
          });
        } else {
          this.popOver.open(`${resp.status}`, '', {duration: 2000});
        }
      }, error => {
        if (error.status === 401) {
          this.popOver.open(`Please sign in to retrieve your user information.`,
            '', {duration: 5000});
          this.router.navigate(['/login']);
        } else {
          this.popOver.open(`Something went wrong! If this error persists, please contact us with the following error: ${error.error}`,
            '', {duration: 5000});
        }
      }
    );
  }

  openDialog() {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = true;
    dialogConfig.autoFocus = true;
    const dialogRef = this.dialog.open(DialogComponent, dialogConfig);
    dialogRef.afterClosed().subscribe(
      data => this.saveNewEntry(data)
    );
  }

  removeEntry(item) {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = true;
    dialogConfig.autoFocus = true;
    const dialogRef = this.dialog.open(ConfirmationDialogComponent, dialogConfig);
    dialogRef.afterClosed().subscribe(
      bool => {
        if (bool) {
          const index = this.dataSource.indexOf(item);
          const body = new PasswordEntry('', item.username, '', item.url, '');
          this.crudService.deletePassword(body).subscribe(
            resp => {
              this.dataSource.splice(index, 1);
              this.popOver.open('Deleted!', '', {duration: 2000});
            }, error => {
              this.popOver.open(`Something went wrong! If this error persists, please contact us with the following error: ${error.error}`,
                '', {duration: 5000});
            }
          );
        }
      }
    );
  }

  copyToClipboard(item) {
    const index = this.dataSource.indexOf(item);
    const password = this.dataSource[index].password;
    const selBox = document.createElement('textarea');
    selBox.style.position = 'fixed';
    selBox.style.left = '0';
    selBox.style.top = '0';
    selBox.style.opacity = '0';
    selBox.value = password;
    document.body.appendChild(selBox);
    selBox.focus();
    selBox.select();
    document.execCommand('copy');
    document.body.removeChild(selBox);
    this.popOver.open('Copied!', '', {duration: 2000});
  }

  saveNewEntry(data) {
    const newEntry = new PasswordEntry(
      'Reload to view id',
      data.username,
      data.password,
      data.url,
      '',
    );
    this.crudService.addPassword(newEntry).subscribe(
      resp => {
        if (resp.status === 200) {
          this.dataSource.push(newEntry);
          this.popOver.open('Saved', '', {duration: 2000});
        } else {
          this.popOver.open(`${resp.status}`, '', {duration: 2000});
        }
      }, error => {
        this.popOver.open(`Something went wrong! If this error persists, please contact us with the following error: ${error.error}`,
          '', {duration: 5000});
      }
    );
  }
}

@Component({
  selector: 'app-confirmation-dialog',
  templateUrl: './confirmation.component.html',
})
export class ConfirmationDialogComponent {

  constructor(
    public dialogRef: MatDialogRef<ConfirmationDialogComponent>) {
  }

  onNoClick(): void {
    this.dialogRef.close();
  }

}
