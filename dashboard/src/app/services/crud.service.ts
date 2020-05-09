import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {Observable} from 'rxjs';
import {UsernameMasterPassword} from '../models/username-master-password';
import {UsernamePasswordUrl} from '../models/username-password-url';
import {UsernameUrl} from '../models/username-url';

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

  apiUrl = 'http://127.0.0.1:8080'; // TODO: should be provided in environment variables

  constructor(
    private httpClient: HttpClient
  ) { }

  getListOfPasswords(): Observable<any> {
    return this.httpClient.get(`${this.apiUrl}/passwords`, this.httpOptions);
  }

  addPassword(body: UsernamePasswordUrl): Observable<any> {
    return this.httpClient.post<UsernamePasswordUrl>(`${this.apiUrl}/password`, JSON.stringify(body), this.httpOptions);
  }

  deletePassword(body: UsernameUrl): Observable<any> {
    return this.httpClient.request<UsernameUrl>('delete', `${this.apiUrl}/password`, {body: JSON.stringify(body)});
  }
}
