import { Routes, RouterModule } from '@angular/router';
import { LoginComponent} from './login/login.component';
import {DashboardComponent} from './dashboard/dashboard.component';
import {SettingsComponent} from './settings/settings.component';
import {AuthGuard} from './util/auth.guard';

const routes: Routes = [
  { path: '', component: LoginComponent },
  { path: 'login', component: LoginComponent },
  { path: 'dashboard', component: DashboardComponent, canActivate: [AuthGuard] },
  { path: 'settings', component: SettingsComponent, canActivate: [AuthGuard] },
  { path: '*', redirectTo: '' }
];

export const appRoutingModule = RouterModule.forRoot(routes);
