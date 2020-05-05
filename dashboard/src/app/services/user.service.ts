import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {Observable} from 'rxjs';
import {UsernameEmail} from '../models/username-email';
import {UsernamePassword} from '../models/username-password';
import {UsernameMasterPassword} from '../models/username-master-password';
import {UsernamePasswordUrl} from '../models/username-password-url';
import {UsernameUrl} from '../models/username-url';

@Injectable({ providedIn: 'root' })
export class UserService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json'}),
    withCredentials: true,
    observe: 'response' as 'response',
    responseType: 'text' as 'json'
  };

  apiUrl = 'http://127.0.0.1:8080'; // TODO: should be provided in environment variables

  constructor(
    private httpClient: HttpClient
  ) { }

  register(body: UsernameEmail): Observable<any> {
    return this.httpClient.post<UsernameEmail>(`${this.apiUrl}/standard/register`, JSON.stringify(body),
      this.httpOptions);
  }

  login(body: UsernamePassword): Observable<any> {
    return this.httpClient.post<UsernamePassword>(`${this.apiUrl}/standard/login`, JSON.stringify(body),
      this.httpOptions);
  }

  getUser(): Observable<any> {
    return this.httpClient.get<UsernameMasterPassword>(`${this.apiUrl}/user`, this.httpOptions);
  }

  getListOfPasswords(): Observable<any> {
    return this.httpClient.get(`${this.apiUrl}/passwords`, this.httpOptions);
  }

  addPassword(body: UsernamePasswordUrl): Observable<any> {
    return this.httpClient.post<UsernamePasswordUrl>(`${this.apiUrl}/password`, JSON.stringify(body), this.httpOptions);
  }

  deletePassword(body: UsernameUrl): Observable<any> {
    console.log(body);
    return this.httpClient.request<UsernameUrl>('delete', `${this.apiUrl}/password`, {body: JSON.stringify(body)});
  }
}
