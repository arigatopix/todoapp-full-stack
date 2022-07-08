import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
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
    Authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjYsImV4cCI6MTY1NzI4NDQ3MywiaXNzIjoidG9kb2FwcCJ9.qSNoTsEU_MbVybIg5TYdLxd5pXE1c8jCRDIZL9qc8-0`,
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

  delete(td: Todo): Observable<Todo[]> {
    return of((this.todos = this.todos.filter((todo) => todo.id != td.id)));
  }

  toggle(td: Todo): Observable<TodoResponse> {
    return this.http.put<TodoResponse>(`/api/todos/${td.id}`, td, httpOptions);
  }
}
