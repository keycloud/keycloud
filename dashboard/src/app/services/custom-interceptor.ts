import {Injectable} from '@angular/core';
import {HttpEvent, HttpHandler, HttpInterceptor, HttpRequest} from '@angular/common/http';
import {Observable} from 'rxjs';
import {environment} from '../../environments/environment.prod';

@Injectable({providedIn: 'root'})
export class CustomInterceptor implements HttpInterceptor {

  constructor() {
  }

  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    // console.log's for debugginggit add
    console.log(req);
    const apiReq = req.clone({
      url: `https://${environment.apiUrl}${req.url}`,
      withCredentials: true
    });
    console.log(apiReq);
    return next.handle(apiReq);
  }
}
