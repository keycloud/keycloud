import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders, HttpResponse} from '@angular/common/http';
import {Observable} from 'rxjs';
import {UsernameEmail} from '../models/username-email';
import {UsernamePassword} from '../models/username-password';
import {mapTo} from 'rxjs/operators';
import {UsernameMasterPassword} from '../models/username-master-password';

@Injectable({ providedIn: 'root' })
export class UserService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json'}),
    withCredentials: true,
    observe: 'response' as 'response',
    responseType: 'text' as 'json'
  };

  constructor(
    private httpClient: HttpClient
  ) { }

  register(body: UsernameEmail): Observable<any> {
    return this.httpClient.post<UsernameEmail>(`http://127.0.0.1:8080/standard/register`, JSON.stringify(body),
      this.httpOptions);
  }

  login(body: UsernamePassword): Observable<any> {
    return this.httpClient.post<UsernamePassword>(`http://127.0.0.1:8080/standard/login`, JSON.stringify(body),
      this.httpOptions);
  }

  getUser(): Observable<any> {
    return this.httpClient.get<UsernameMasterPassword>(`http://127.0.0.1:8080/user`, this.httpOptions);
  }
}
