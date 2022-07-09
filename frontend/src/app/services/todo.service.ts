import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Todo } from '../interfaces/Todo';
import { Response } from '../interfaces/Response';
import { HttpClient, HttpHeaders } from '@angular/common/http';

interface TodosResponse {
  Response: Response;
  data: Todo[];
}

interface TodoResponse {
  Response: Response;
  data: Todo;
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
export class TodoService {
  todos: Todo[] = [];

  constructor(private http: HttpClient) {}

  get(): Observable<TodosResponse> {
    return this.http.get<TodosResponse>('/api/todos', httpOptions);
  }

  create(td: Todo): Observable<TodoResponse> {
    return this.http.post<TodoResponse>(`/api/todos`, td, httpOptions);
  }

  delete(td: Todo): Observable<TodoResponse> {
    return this.http.delete<TodoResponse>(`/api/todos/${td.id}`, httpOptions);
  }

  toggle(td: Todo): Observable<TodoResponse> {
    return this.http.put<TodoResponse>(`/api/todos/${td.id}`, td, httpOptions);
  }
}
