import { HttpClient,HttpHeaders } from '@angular/common/http';
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
    let uname= this.loginForm.value.email
    let pass= this.loginForm.value.password
    let user= uname.split('@')
    const httpOptions={
      headers: new HttpHeaders(
        {
          'Content-Type': 'application/json',
          'Access-Control-Allow-Origin' : '*',
          'Access-Control-Allow-Methods': 'GET'
        }
      )
    } ;
    // this.http.get<any>("http://localhost:8085/staff/"+user[0],httpOptions)
    // .subscribe(res=>{
    //   const user = res.find((a:any)=>{
    //     console.log(this.loginForm.value.email)
    //     return a.username === user[0] && a.password === this.loginForm.value.password
    //   });

    if(user[0]==="SMADIPADI" && pass==="SMADIPADI")
    {
        alert("Login Success");
        this.loginForm.reset();
        this.router.navigate(['/home'])
      }else{
        alert("user not found");

    }
  }
}

