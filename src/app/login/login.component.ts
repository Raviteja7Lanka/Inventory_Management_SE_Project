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
  public test: any;

  
  ngOnInit(): void {
    let header = new HttpHeaders();
    header.set('Access-Control-Allow-Origin', '*');

    console.log('in init');
    console.log('above', this.test);

    // this.http.get<any>('http://localhost:8085/staff/smadipadi', {headers: header}).subscribe(data => {
    //     console.log(data);
    //     this.test = data.total;
    //     console.log('in', this.test);
    // });
    console.log('out', this.test);
    this.loginForm = this.formBuilder.group({
      email:[''],
      password:['']
    })
  }

  login(){
<<<<<<< Updated upstream
    this.http.get<any>("http://localhost:3000")
    .subscribe(res=>{
      const user = res.find((a:any)=>{
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
=======
    console.log(this.loginForm.value.email)
    let uname= this.loginForm.value.email
    let pass= this.loginForm.value.password
    let user= uname.split('@');
    console.log('user name is', user);
    const httpOptions={
      headers: new HttpHeaders(
        {
          'Content-Type': 'application/json',
          'Access-Control-Allow-Origin' : '*',
          'Access-Control-Allow-Methods': 'GET'
        }
      )
    } ;



    this.http.get<any>(`http://localhost:8085/staff/${user[0]}`).subscribe(res => {
        console.log(res);
        // this.test = res.total;
        console.log('in', this.test);
        const user = res.find((a:any)=>{
              console.log(this.loginForm.value.email)
              if (a.username === user[0] && a.password === this.loginForm.value.password)
              {
                alert("Login Success");
                this.loginForm.reset();
                this.router.navigate(['/home']) 
              }
            })

    });

    // this.http.get<any>("http://localhost:8085/staff/"+user[0],httpOptions)
    // .subscribe(res=>{
    //   const user = res.find((a:any)=>{
    //     console.log(this.loginForm.value.email)
    //     return a.username === user[0] && a.password === this.loginForm.value.password
    //   });

    // if(user[0]==="SMADIPADI" && pass==="SMADIPADI")
    // {
    //     alert("Login Success");
    //     this.loginForm.reset();
    //     this.router.navigate(['/home'])
    //   }else{
    //     alert("user not found");

    // }
>>>>>>> Stashed changes
  }
}