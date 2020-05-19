import {Component, Injector} from '@angular/core';
import {Router} from '@angular/router';
import {UserService} from './services/user.service';
import {MatSnackBar} from '@angular/material/snack-bar';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'KeyCloud';

  constructor(
    public router: Router,
    public userService: UserService,
    public popOver: MatSnackBar,
  ) {
  }

  logout() {
    this.userService.logout().subscribe(
      _ => this.router.navigate(['/login']),
      _ => this.popOver.open('Can\'t log out, as you are not logged in!', '', {duration: 5000})
    );
  }
}
