import { Component, OnInit } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css'],
})
export class HomeComponent implements OnInit {
  signedin$: BehaviorSubject<boolean>;

  constructor(private authService: AuthService) {
    this.signedin$ = this.authService.isAuth$;
  }

  ngOnInit(): void {}
}
