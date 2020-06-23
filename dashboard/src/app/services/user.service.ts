import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {Observable} from 'rxjs';
import {User} from '../models/user';
import {UserRegister} from '../models/user-register';

@Injectable({providedIn: 'root'})
export class UserService {

  httpOptions = {
    headers: new HttpHeaders({'Content-Type': 'application/json'}),
    withCredentials: true,
    observe: 'response' as 'response',
    responseType: 'text' as 'json'
  };

  constructor(
    private httpClient: HttpClient
  ) {
  }

  getUser(): Observable<any> {
    return this.httpClient.get<User>(`/user`, this.httpOptions);
  }

  register(body: UserRegister): Observable<any> {
    return this.httpClient.post<UserRegister>(`/standard/register`, JSON.stringify(body),
      this.httpOptions);
  }

  login(body: User): Observable<any> {
    return this.httpClient.post<User>(`/standard/login`, JSON.stringify(body),
      this.httpOptions);
  }

  logout(): Observable<any> {
    return this.httpClient.post(`/logout`, '', this.httpOptions);
  }

  webauthnRegistrationStart(body: UserRegister): Observable<any> {
    return this.httpClient.post<UserRegister>(`/webauthn/registration/start`, JSON.stringify(body), this.httpOptions);
  }

  webauthnRegistrationFinish(body: any): Observable<any> {
    return this.httpClient.post(`/webauthn/registration/finish`, JSON.stringify(body), this.httpOptions);
  }

  webauthnLoginStart(body: UserRegister): Observable<any> {
    return this.httpClient.post<UserRegister>(`/webauthn/login/start`, JSON.stringify(body), this.httpOptions);
  }

  webauthnLoginFinish(body: any): Observable<any> {
    return this.httpClient.post(`/webauthn/login/finish`, JSON.stringify(body), this.httpOptions);
  }
}
