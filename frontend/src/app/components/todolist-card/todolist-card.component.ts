import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TodoService } from 'src/app/services/todo.service';
import { Todo } from '../../interfaces/Todo';

@Component({
  selector: 'app-todolist-card',
  templateUrl: './todolist-card.component.html',
  styleUrls: ['./todolist-card.component.css'],
})
export class TodolistCardComponent implements OnInit {
  todos: Todo[] = [];


  message: string = '';

  constructor(private todoService: TodoService, private router: Router) {}

  ngOnInit(): void {
    this.todoService.get().subscribe({
      next: (res) => {
        this.todos = res.data;
      },
      error: (err) => {
        this.router.navigateByUrl("login")
        return this.message = err.error.message
      }
    });
  }

  createTodo(todo: Todo) {
    this.todoService.create(todo).subscribe({
      next: (res) => this.todos.push(res.data),
      error: (err) => (this.message = err.message),
    });
  }

  deleteTodo(todo: Todo) {
    return this.todoService.delete(todo).subscribe({
      next: () =>
        (this.todos = this.todos.filter((td) => td.id !== todo.id)),
      error: (err) =>{
        return this.message = err.message
      },
    });
  }

  toggle(todo: Todo) {
    todo.completed = !todo.completed;
    this.todoService.toggle(todo).subscribe();
  }
}
