import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {Observable} from 'rxjs';
import {environment} from '../../environments/environment.prod';
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

  apiUrl = environment.apiUrl;

  constructor(
    private httpClient: HttpClient
  ) {
  }

  getUser(): Observable<any> {
    return this.httpClient.get<User>(`${this.apiUrl}/user`, this.httpOptions);
  }

  register(body: UserRegister): Observable<any> {
    return this.httpClient.post<UserRegister>(`${this.apiUrl}/standard/register`, JSON.stringify(body),
      this.httpOptions);
  }

  login(body: User): Observable<any> {
    return this.httpClient.post<User>(`${this.apiUrl}/standard/login`, JSON.stringify(body),
      this.httpOptions);
  }

  logout(): Observable<any> {
    return this.httpClient.post(`${this.apiUrl}/logout`, '', this.httpOptions);
  }

  webauthnRegistrationStart(body: UserRegister): Observable<any> {
    return this.httpClient.post<UserRegister>(`${this.apiUrl}/webauthn/registration/start`, JSON.stringify(body), this.httpOptions);
  }

  webauthnRegistrationFinish(body: any): Observable<any> {
    return this.httpClient.post(`${this.apiUrl}/webauthn/registration/finish`, JSON.stringify(body), this.httpOptions);
  }

  webauthnLoginStart(body: UserRegister): Observable<any> {
    return this.httpClient.post<UserRegister>(`${this.apiUrl}/webauthn/login/start`, JSON.stringify(body), this.httpOptions);
  }

  webauthnLoginFinish(body: any): Observable<any> {
    return this.httpClient.post(`${this.apiUrl}/webauthn/login/finish`, JSON.stringify(body), this.httpOptions);
  }
}
