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
      ]),
      password: new FormControl("", [
        Validators.required,
        Validators.minLength(3),
      ]),
      passwordConfirm: new FormControl("", [
        Validators.required,
        Validators.minLength(3),
      ])
    })
  }

  onSignUp(): void {
    if (this.registerFormGroup.invalid) {
      this.message = "Please enter your email"
      return;
    }

    const password = this.registerFormGroup.get('password')?.value;
    const passwordConfirm = this.registerFormGroup.get('passwordConfirm')?.value;

    if (password !== passwordConfirm) {
      this.message = "Password not match"
      return
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
