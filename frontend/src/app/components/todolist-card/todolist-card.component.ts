import { Component, OnInit } from '@angular/core';
import { TodoService } from 'src/app/services/todo.service';
import { Todo } from '../../interfaces/Todo';

@Component({
  selector: 'app-todolist-card',
  templateUrl: './todolist-card.component.html',
  styleUrls: ['./todolist-card.component.css'],
})
export class TodolistCardComponent implements OnInit {
  todos: Todo[] = [];

  constructor(private todoService: TodoService) {}

  ngOnInit(): void {
    this.todoService.get().subscribe((todos) => {
      this.todos = todos;
    });
  }

  createTodo(todo: Todo) {
    return this.todoService.create(todo).subscribe((todo) => {
      this.todos.push(todo);
    });
  }

  deleteTodo(todo: Todo) {
    return this.todoService
      .delete(todo)
      .subscribe((todos) => (this.todos = todos));
  }

  toggle(todo: Todo) {
    todo.compleated = !todo.compleated;

    this.todoService.toggle(todo).subscribe();
  }
}
