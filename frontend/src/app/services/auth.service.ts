import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Response } from '../interfaces/Response';

interface LoginCredentails {
  email: string
}

interface LoginResponse {
  response : Response
  data: {
    email: string
    token: string
  }
}

const httpOptions = {
  withCredentials: true,
  headers: new HttpHeaders({
    'Content-Type': 'application/json',
  }),
};

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor(private http: HttpClient) { }

  login(credentials : LoginCredentails) {
    return this.http.post<LoginResponse>('/api/auth/login', credentials, httpOptions)
  }
}
