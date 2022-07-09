import { Component, OnInit, EventEmitter, Output } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { Todo } from '../../interfaces/Todo';

@Component({
  selector: 'app-add-todo',
  templateUrl: './add-todo.component.html',
  styleUrls: ['./add-todo.component.css'],
})
export class AddTodoComponent implements OnInit {
  @Output() addTodoHandler: EventEmitter<Todo> = new EventEmitter();

  todoFormGroup = new FormGroup({
    id: new FormControl(),
    title: new FormControl(),
    completed: new FormControl(),
  });
  constructor() {}

  ngOnInit(): void {}

  onClick(): void {
    if (this.todoFormGroup.get('title')?.value) {
      this.addTodoHandler.emit({
        title: this.todoFormGroup.get('title')?.value,
        completed: false,
      });

      this.todoFormGroup.reset();
    }
  }
}
