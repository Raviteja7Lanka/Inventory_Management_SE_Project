import { HttpClient,HttpHeaders } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import {FormGroup,FormBuilder} from "@angular/forms";
import {Subscription} from "rxjs";
import { Router } from '@angular/router';
@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css']
})
export class SignupComponent implements OnInit {
  public  signupForm !: FormGroup;
  constructor(private formBuilder :FormBuilder, private http :HttpClient, private router:Router) { }

  ngOnInit(): void {
    this.signupForm = this.formBuilder.group(
      {
        firstname:[''],
        lastname:[''],
        email:[''],
        password:[''],
        mobile:['']
      }
    )
  }
  // signup(){
  //   console.log(this.signupForm.value);
  //   this.http.post<any>("http://localhost:8085/staff/all",this.signupForm.value)
  //   .subscribe({next : res=>{
  //     alert("signup Successfull");
  //     this.signupForm.reset();
  //     this .router.navigate(['login']);
  //   }},err=>{
  //     alert("something went wrong")
  //   }
  //   )
  // }
  //   // if(this.signupForm.value!=null)
  //   // {
  //   //     alert("Signup is SucessFull!! You Can Log in");
  //   //     this.signupForm.reset();
  //   //     this.router.navigate(['/login'])
  //   //   }else{
  //   //     alert("please enter valid data");

  //   // }

  // }
  signup(){
<<<<<<< Updated upstream
    this.http.post<any>("http://localhost:3000",this.signupForm.value)
    .subscribe(res=>{
      alert("signup Successfull");
      this.signupForm.reset();
      this .router.navigate(['login']);
    },err=>{
      alert("something went wrong")
    }
    )

=======
    const httpOptions={
      headers: new HttpHeaders(
        {
          'Content-Type': 'application/json',
          'Access-Control-Allow-Origin' : '*',
          'Access-Control-Allow-Methods': 'POST'
        }
      )
    } ;
    if(this.signupForm.value!=null)
    {
    console.log(this.signupForm.value);
    this.http.post<any>("http://localhost:8085/staff/all",this.signupForm.value)
    .subscribe ({
      next :(res)=> {
        alert("signup Successfull");
        this.signupForm.reset();
        this .router.navigate(['login']);
      },
      error :()=>{
        alert("something went wrong") 
      }
    })}
>>>>>>> Stashed changes
  }

}
