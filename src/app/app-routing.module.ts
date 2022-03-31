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
import { OrdersComponent } from './orders/orders.component';
import { ProductsComponent } from './products/products.component';
import { ReportsComponent } from './reports/reports.component';
import { AboutComponent } from './about/about.component';
import { CatproductComponent } from './products/catproduct/catproduct.component';

const routes: Routes = [
  { path: '', redirectTo: '/login', pathMatch: 'full' },
  { path: 'login', component: LoginComponent },
  { path: 'signup', component: SignupComponent },
  { path: 'forgot-password', component: ForgotPasswordComponent },

  {
    path: '',
    component: MyNavbarComponent,
    children: [
      { path: 'tabs', component: TabsComponent },
      { path: 'my-navbar', component: MyNavbarComponent },
      { path: 'supplier-table', component: SupplierTableComponent },
      { path: 'customer-table', component: CustomerTableComponent },
      { path: 'home', component: HomeComponent },
      { path: 'orders', component: OrdersComponent },
      { path: 'products', component: ProductsComponent },
      { path: 'reports', component: ReportsComponent },
      { path: 'about', component: AboutComponent },
      { path: 'catprod', component: CatproductComponent },
    ],
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
