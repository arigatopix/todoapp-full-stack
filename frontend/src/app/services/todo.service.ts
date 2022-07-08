import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { Todo } from '../interfaces/Todo';

@Injectable({
  providedIn: 'root',
})
export class TodoService {
  todos: Todo[] = [];
  constructor() {}

  get(): Observable<Todo[]> {
    return of(this.todos);
  }

  create(td: Todo): Observable<Todo> {
    const newId = this.todos.length + 1;
    const todo: Todo = {
      id: newId,
      title: td.title,
      compleated: td.compleated,
    };
    return of(todo);
  }

  delete(td: Todo): Observable<Todo[]> {
    return of((this.todos = this.todos.filter((todo) => todo.id != td.id)));
  }

  toggle(td: Todo): Observable<Todo> {
    return of(td);
  }
}
