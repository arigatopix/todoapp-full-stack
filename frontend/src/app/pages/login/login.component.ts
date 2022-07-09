import { compileNgModule } from '@angular/compiler';
import { Component, Input, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators, FormBuilder } from '@angular/forms';
import { Location } from '@angular/common';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  loginFormGroup!: FormGroup;

  @Input() message: string = ''

  constructor(private authService: AuthService, private location: Location) { }

  ngOnInit(): void {
    this.loginFormGroup = new FormGroup({
      email: new FormControl("", [
        Validators.required,
        Validators.pattern(/^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$/)
      ])
    })
  }

  onLogin() {
    if (this.loginFormGroup.invalid) {
      this.message = "Please enter your email"
      return;
    }

    this.authService.login(this.loginFormGroup.value).subscribe({
      next: ()=> {
        this.location.back()
      },
      error: (err) => {
        this.message = err.error.message
      }
    })
      
    

    this.loginFormGroup.reset()
    this.message = ""
  }
}
