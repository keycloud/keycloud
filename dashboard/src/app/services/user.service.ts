import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import {UsernameEmail} from '../models/username-email';
import {UsernamePassword} from '../models/username-password';

@Injectable({ providedIn: 'root' })
export class UserService {

  constructor(
    private httpClient: HttpClient
  ) { }

  register(body: UsernameEmail): Observable<any> {
    return this.httpClient.post<UsernameEmail>(`http://127.0.0.1:8080/standard/register`, JSON.stringify(body),
      { observe: 'response', responseType: 'text' as 'json'});
  }

  login(body: UsernamePassword): Observable<any> {
    return this.httpClient.post<UsernamePassword>(`http://127.0.0.1:8080/standard/login`, JSON.stringify(body),
      { observe: 'response', responseType: 'text' as 'json'});
  }
}
