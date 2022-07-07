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

}
