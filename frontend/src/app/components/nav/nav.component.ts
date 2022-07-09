import { Component, Input, OnInit } from '@angular/core';
import { AuthService } from 'src/app/services/auth.service';
import { BehaviorSubject } from 'rxjs';


@Component({
  selector: 'app-nav',
  templateUrl: './nav.component.html',
  styleUrls: ['./nav.component.css']
})
export class NavComponent implements OnInit {
  @Input() email!: string;
  @Input() signedin!: BehaviorSubject<boolean>

  
  constructor() { 
   
  }

  ngOnInit(): void {}
}
