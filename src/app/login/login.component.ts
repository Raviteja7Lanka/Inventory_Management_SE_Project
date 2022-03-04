import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormGroup} from '@angular/forms'
import { Router } from '@angular/router';
@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  public loginForm!: FormGroup;
  constructor(private formBuilder:FormBuilder,private http:HttpClient, private router:Router) { }

  ngOnInit(): void {
    this.loginForm = this.formBuilder.group({
      email:[''],
      password:['']
    })
  }
  login(){
    console.log(this.loginForm.value.email)
    this.http.get<any>("http://localhost:8085/")
    .subscribe(res=>{
      const user = res.find((a:any)=>{
        console.log(this.loginForm.value.email)
        return a.email === this.loginForm.value.email && a.password === this.loginForm.value.password
      });
      if(user){
        alert("Login Success");
        this.loginForm.reset();
        this.router.navigate(['my-navbar'])
      }else{
        alert("user not found");

      }
    },err=>{
      alert("something went wrong")
    })
  }
}
