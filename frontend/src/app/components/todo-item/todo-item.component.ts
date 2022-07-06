import { Component, OnInit } from '@angular/core';
import { faTimes } from '@fortawesome/free-solid-svg-icons';
import  { Todo } from '../../interfaces/Todo'
import { todos } from '../../mock-todo'

@Component({
  selector: 'app-todo-item',
  templateUrl: './todo-item.component.html',
  styleUrls: ['./todo-item.component.css'],
})
export class TodoItemComponent implements OnInit {

  todos: Todo[] = todos;

  faTimes = faTimes;
  constructor() {}

  ngOnInit(): void {}
}
