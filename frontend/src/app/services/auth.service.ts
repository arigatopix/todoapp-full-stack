import { Injectable } from '@angular/core';
import { tap } from 'rxjs/operators';
import { BehaviorSubject } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Response } from '../interfaces/Response';

interface LoginCredentails {
  email: string
}

interface GetMeResponse {
  code: number
  message: string
  data: {
    id: number
    email: string
  }
}

interface LoginResponse {
  code: number
  message: string
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

  isAuth$ = new BehaviorSubject(false)

  constructor(private http: HttpClient) { }

  getMe() {
    return this.http.get<GetMeResponse>('/api/auth/me', httpOptions).pipe(
      tap((res) => {
        if (res.message === 'ok') {
          this.isAuth$.next(true)
        }
      })
    )
  }

  login(credentials : LoginCredentails) {
    return this.http.post<LoginResponse>('/api/auth/login', credentials, httpOptions).pipe(
      tap((res) => {
        if (res.message === 'ok') {
          this.isAuth$.next(true)
        }
      })
    )
  }

  logout() {
    return this.http.post('/api/auth/logout', {}, httpOptions).pipe(
      tap(() => {
        this.isAuth$.next(false)
      })
    )
  }
}
