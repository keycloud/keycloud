import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {Observable} from 'rxjs';
import {environment} from '../../environments/environment.prod';
import {PasswordEntry} from '../models/password-entry';

@Injectable({
  providedIn: 'root'
})
export class CrudService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json'}),
    withCredentials: true,
    observe: 'response' as 'response',
    responseType: 'text' as 'json'
  };

  apiUrl = environment.apiUrl;

  constructor(
    private httpClient: HttpClient
  ) { }

  getListOfPasswords(): Observable<any> {
    return this.httpClient.get(`${this.apiUrl}/passwords`, this.httpOptions);
  }

  addPassword(body: PasswordEntry): Observable<any> {
    return this.httpClient.post<PasswordEntry>(`${this.apiUrl}/password`, JSON.stringify(body), this.httpOptions);
  }

  deletePassword(body: PasswordEntry): Observable<any> {
    return this.httpClient.request<PasswordEntry>('delete', `${this.apiUrl}/password`, {body: JSON.stringify(body)});
  }
}
