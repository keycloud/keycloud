import { Routes, RouterModule } from '@angular/router';
import { LoginComponent} from './login/login.component';
import {DashboardComponent} from './dashboard/dashboard.component';

const routes: Routes = [
  { path: '', component: LoginComponent }, // TODO might be good to have a landing page
  { path: 'login', component: LoginComponent },
  { path: 'dashboard', component: DashboardComponent },
  { path: '*', redirectTo: '' }
];

export const appRoutingModule = RouterModule.forRoot(routes);
