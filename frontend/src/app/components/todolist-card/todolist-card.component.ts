import { Component, OnInit } from '@angular/core';
import  { Todo } from '../../interfaces/Todo'
import { todos } from '../../mock-todo'

@Component({
  selector: 'app-todolist-card',
  templateUrl: './todolist-card.component.html',
  styleUrls: ['./todolist-card.component.css']
})
export class TodolistCardComponent implements OnInit {

  todos: Todo[] = todos;
  
  constructor() { }

  ngOnInit(): void {
  }

  createTodo(todo: Todo) {
    todos.push(todo)
  }

  deleteTodo(todo: Todo) {
    return this.todos = this.todos.filter(td => td.id !== todo.id)
  }
}
