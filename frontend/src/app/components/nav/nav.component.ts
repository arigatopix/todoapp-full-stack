import { Component, OnInit } from '@angular/core';
import { AuthService } from 'src/app/services/auth.service';
import { BehaviorSubject } from 'rxjs';

@Component({
  selector: 'app-nav',
  templateUrl: './nav.component.html',
  styleUrls: ['./nav.component.css'],
})
export class NavComponent implements OnInit {
  emailAuth$: BehaviorSubject<string>;
  signedin$: BehaviorSubject<boolean>;

  constructor(private authService: AuthService) {
    this.signedin$ = this.authService.isAuth$;
    this.emailAuth$ = this.authService.emailAuth$;
  }

  ngOnInit(): void {
    this.authService.getMe().subscribe(() => {});
  }
}
