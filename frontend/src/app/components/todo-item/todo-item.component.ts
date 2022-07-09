import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { faTimes } from '@fortawesome/free-solid-svg-icons';
import { Todo } from '../../interfaces/Todo';

@Component({
  selector: 'app-todo-item',
  templateUrl: './todo-item.component.html',
  styleUrls: ['./todo-item.component.css'],
})
export class TodoItemComponent implements OnInit {
  @Output() onDeleteTodo: EventEmitter<Todo> = new EventEmitter();
  @Output() onToggleHandler: EventEmitter<Todo> = new EventEmitter();
  @Input() todo!: Todo;

  faTimes = faTimes;
  constructor() {}

  ngOnInit(): void {}

  onDeleteHandler(todo: Todo) {
    this.onDeleteTodo.emit(todo);
  }

  onCheckboxChange(todo: Todo) {
    this.onToggleHandler.emit(todo);
  }
}
