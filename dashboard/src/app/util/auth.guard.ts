import {ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot} from '@angular/router';
import {Injectable} from '@angular/core';
import {CookieService} from 'ngx-cookie-service';

@Injectable({ providedIn: 'root' })
export class AuthGuard implements CanActivate {

  constructor(
    private router: Router,
    private cookieService: CookieService
  ) {}

  canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot) {
    return true;
    if (this.cookieService.check('keycloud-main')) {  // TODO: cookie is not set in browser but send with the request headers
      return true;
    } else {
      this.router.navigate(['/login'], { queryParams: { returnUrl: state.url}});
      return false;
    }
  }
}