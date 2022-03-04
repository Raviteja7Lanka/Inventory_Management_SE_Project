import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { TabsComponent } from './tabs/tabs.component';
import { MatTabsModule } from '@angular/material/tabs';
import { MatTableModule } from '@angular/material/table';
import { SupplierTableComponent } from './supplier-table/supplier-table.component';
import { CustomerTableComponent } from './customer-table/customer-table.component';
import { MatButtonModule } from '@angular/material/button';
import { ThemePalette } from '@angular/material/core';
import { MyNavbarComponent } from './my-navbar/my-navbar.component';
import { LayoutModule } from '@angular/cdk/layout';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatIconModule } from '@angular/material/icon';
import { FlexLayoutModule } from '@angular/flex-layout';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatCardModule } from '@angular/material/card';
import { MatListModule } from '@angular/material/list';
import { HomeComponent } from './home/home.component';
import { HttpClientModule } from '@angular/common/http';
import { SignupComponent } from './signup/signup.component';
import { LoginComponent } from './login/login.component';
import { ForgotPasswordComponent } from './login/forgot-password/forgot-password.component';
import { MatDialogModule } from '@angular/material/dialog';
import { AddCustomerComponent } from './customer-table/add-customer/add-customer.component';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { AddSupplierComponent } from './supplier-table/add-supplier/add-supplier.component';
import { ProductsComponent } from './products/products.component';

@NgModule({
  declarations: [
    AppComponent,
    TabsComponent,
    CustomerTableComponent,
    MyNavbarComponent,
    SupplierTableComponent,
    HomeComponent,
    SignupComponent,
    LoginComponent,
    ForgotPasswordComponent,
    AddCustomerComponent,
    AddSupplierComponent,
    ProductsComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatTabsModule,
    MatTableModule,
    MatButtonModule,
    ReactiveFormsModule,
    LayoutModule,
    HttpClientModule,
    MatToolbarModule,
    MatSidenavModule,
    FlexLayoutModule,
    MatIconModule,
    MatCardModule,
    MatGridListModule,
    MatListModule,
    MatDialogModule,
    MatFormFieldModule,
    MatInputModule,
    ReactiveFormsModule,
    FormsModule,
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
