import { Component, Input, OnInit } from '@angular/core';
import { faTimes } from '@fortawesome/free-solid-svg-icons';
import  { Todo } from '../../interfaces/Todo'

@Component({
  selector: 'app-todo-item',
  templateUrl: './todo-item.component.html',
  styleUrls: ['./todo-item.component.css'],
})
export class TodoItemComponent implements OnInit {

  
  @Input() todo!: Todo;

  faTimes = faTimes;
  constructor() {}

  ngOnInit(): void {}
}
