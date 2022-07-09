import { Component } from '@angular/core';
import { AuthService } from './services/auth.service';
import { BehaviorSubject } from 'rxjs';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent {
  email! : string;

  signedid$: BehaviorSubject<boolean>;

  title = 'TodoApp';

  constructor(private authService: AuthService) {
    this.signedid$ = this.authService.isAuth$;
  }

  ngOnInit(): void {
    this.authService.getMe().subscribe((res) =>{
      this.email = res.data.email
    })
  }
}
