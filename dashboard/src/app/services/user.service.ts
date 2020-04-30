import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {environment} from '../../environments/environment';
import {Observable} from 'rxjs';

@Injectable({ providedIn: 'root' })
export class UserService {

  constructor(
    private httpClient: HttpClient
  ) { }

  // tslint:disable-next-line:ban-types
  register(body: any): Observable<Object> {
    return this.httpClient.post(`${environment.apiUrl}`, body);
  }

  // tslint:disable-next-line:ban-types
  login(): Observable<Object> {
    return this.httpClient.get(`${environment.apiUrl}`);
  }
}
