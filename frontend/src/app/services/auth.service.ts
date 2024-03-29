import { Injectable } from '@angular/core';
import { tap } from 'rxjs/operators';
import { BehaviorSubject } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';

interface LoginCredentials {
  email: string;
  password: string;
}

interface RegisterCredentials {
  email: string;
  password: string;
  passwordConfirm: string;
}

interface GetMeResponse {
  code: number;
  message: string;
  data: {
    id: number;
    email: string;
  };
}

interface LoginResponse {
  code: number;
  message: string;
  data: {
    email: string;
    token: string;
  };
}

interface RegisterResponse {
  code: number;
  message: string;
  data: {
    email: string;
    token: string;
  };
}

const httpOptions = {
  withCredentials: true,
  headers: new HttpHeaders({
    'Content-Type': 'application/json',
  }),
};

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  isAuth$ = new BehaviorSubject<boolean>(null || false);
  emailAuth$ = new BehaviorSubject('');

  constructor(private http: HttpClient) {}

  getMe() {
    return this.http.get<GetMeResponse>('/api/auth/me', httpOptions).pipe(
      tap((res) => {
        if (res.code === 200) {
          this.isAuth$.next(true);
          this.emailAuth$.next(res.data.email);
        }
      })
    );
  }

  login(credentials: LoginCredentials) {
    return this.http
      .post<LoginResponse>('/api/auth/login', credentials, {
        headers: new HttpHeaders({
          'Content-Type': 'application/json',
        }),
      })
      .pipe(
        tap((res) => {
          if (res.code === 200) {
            this.isAuth$.next(true);
            this.emailAuth$.next(res.data.email);
          }
        })
      );
  }

  register(credentials: RegisterCredentials) {
    return this.http
      .post<RegisterResponse>('/api/auth/register', credentials, {
        headers: new HttpHeaders({
          'Content-Type': 'application/json',
        }),
      })
      .pipe(
        tap((res) => {
          if (res.code === 200) {
            this.isAuth$.next(true);
            this.emailAuth$.next(res.data.email);
          }
        })
      );
  }

  logout() {
    return this.http
      .post(
        '/api/auth/logout',
        {},
        {
          headers: new HttpHeaders({
            'Content-Type': 'application/json',
          }),
        }
      )
      .pipe(
        tap(() => {
          this.isAuth$.next(false);
          this.emailAuth$.next('');
        })
      );
  }
}
