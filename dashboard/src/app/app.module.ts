import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { appRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import {RouterModule, Routes} from '@angular/router';
import { SettingsComponent } from './settings/settings.component';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';

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
  ],
  imports: [
    BrowserModule,
    appRoutingModule,
    FormsModule,
    ReactiveFormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
