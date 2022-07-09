import { Component, Input, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {
  registerFormGroup! : FormGroup;

  @Input() message: string = ''

  constructor(private authService: AuthService, private router: Router) { }

  ngOnInit(): void {
    this.registerFormGroup = new FormGroup({
      email: new FormControl('', [
        Validators.required,
        Validators.pattern(/^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$/)
      ])
    })
  }

  onSignUp(): void {
    if (this.registerFormGroup.invalid) {
      this.message = "Please enter your email"
      return;
    }

    this.authService.register(this.registerFormGroup.value).subscribe({
      next: ()=> {
        this.router.navigateByUrl("/todo")
      },
      error: (err) => {
        this.message = err.error.message
      }
    })

    this.registerFormGroup.reset()
    this.message = ""
  }

}
