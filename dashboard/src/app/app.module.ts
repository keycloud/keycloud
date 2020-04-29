import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import {MatDialogModule} from '@angular/material/dialog';

import { appRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { Routes} from '@angular/router';
import { SettingsComponent } from './settings/settings.component';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import { DialogComponent } from './dialog/dialog.component';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatInputModule} from '@angular/material/input';
import {MatButtonToggleModule} from '@angular/material/button-toggle';
import {MatButtonModule} from '@angular/material/button';
import {MatTooltipModule} from '@angular/material/tooltip';
import {MatSnackBarModule} from '@angular/material/snack-bar';
import {MatCheckboxModule} from '@angular/material/checkbox';
import {MatDividerModule} from '@angular/material/divider';
import {MatIconModule} from '@angular/material/icon';

const appRoutes: Routes = [
  { path : 'login', component : LoginComponent},
  { path : 'dashboard', component : DashboardComponent},
  ];

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    DashboardComponent,
    SettingsComponent,
    DialogComponent,
  ],
  imports: [
    BrowserModule,
    appRoutingModule,
    FormsModule,
    ReactiveFormsModule,
    MatDialogModule,
    BrowserAnimationsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonToggleModule,
    MatButtonModule,
    MatTooltipModule,
    MatSnackBarModule,
    MatCheckboxModule,
    MatDividerModule,
    MatIconModule,
  ],
  providers: [
  ],
  bootstrap: [
    AppComponent
  ],
  entryComponents: [
  ]
})
export class AppModule { }
