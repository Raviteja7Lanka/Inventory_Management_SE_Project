import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CustomerTableComponent } from './customer-table/customer-table.component';
import { HomeComponent } from './home/home.component';
import { ForgotPasswordComponent } from './login/forgot-password/forgot-password.component';
import { LoginComponent } from './login/login.component';
import { MyNavbarComponent } from './my-navbar/my-navbar.component';
import { SignupComponent } from './signup/signup.component';
import { SupplierTableComponent } from './supplier-table/supplier-table.component';
import { TabsComponent } from './tabs/tabs.component';

const routes: Routes = [
  { path: '',redirectTo: '/login',pathMatch: 'full'},
  { path:'login', component:LoginComponent},
  { path:'signup',component:SignupComponent},
  { path:'forgot-password',component:ForgotPasswordComponent},
  {
    path:'',
    component:MyNavbarComponent,
    children:[
       { path: 'tabs', component: TabsComponent },
       { path: 'my-navbar', component:MyNavbarComponent},
       { path: 'supplier-table',component:SupplierTableComponent},
       { path: 'customer-table',component:CustomerTableComponent},
       { path: 'home',component:HomeComponent}
    ]
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
