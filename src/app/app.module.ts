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
import { DeleteSupplierComponent } from './supplier-table/delete-supplier/delete-supplier.component';
import { EditSupplierComponent } from './supplier-table/edit-supplier/edit-supplier.component';
import { DeleteCustomerComponent } from './customer-table/delete-customer/delete-customer.component';
import { EditCustomerComponent } from './customer-table/edit-customer/edit-customer.component';
import { WarehouseComponent } from './services/warehouse/warehouse.component';
import { WarehouseUpdateComponent } from './services/warehouse-update/warehouse-update.component';
import { WarehouseCreateComponent } from './services/warehouse-create/warehouse-create.component';
import { OrdersComponent } from './orders/orders.component';
import { MatSelect, MatSelectModule } from '@angular/material/select';
import { MatMenuModule } from '@angular/material/menu';
import { ReportsComponent } from './reports/reports.component';
import { AboutComponent } from './about/about.component';
import { CatproductComponent } from './products/catproduct/catproduct.component';
import { AddproductComponent } from './products/catproduct/addproduct/addproduct.component';
import { DeleteproductComponent } from './products/catproduct/deleteproduct/deleteproduct.component';
import { EditproductComponent } from './products/catproduct/editproduct/editproduct.component';
import { CategoriesComponent } from './services/categories/categories.component';
import { PendingordersComponent } from './orders/pendingorders/pendingorders.component';
import { ProgressordersComponent } from './orders/progressorders/progressorders.component';
import { OutstandingordersComponent } from './orders/outstandingorders/outstandingorders.component';
import { AddorderComponent } from './orders/addorder/addorder.component';
import { EditorderComponent } from './orders/editorder/editorder.component';
import { DeleteorderComponent } from './orders/deleteorder/deleteorder.component';
import { CategoryCreateComponent } from './services/category-create/category-create.component';
import { ProfileComponent } from './profile/profile.component';
import { OrderDetailsComponent } from './orders/order-details/order-details.component';
import { ViewOrdersComponent } from './supplier-table/view-orders/view-orders.component';

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
    DeleteSupplierComponent,
    EditSupplierComponent,
    DeleteCustomerComponent,
    EditCustomerComponent,
    WarehouseComponent,
    WarehouseUpdateComponent,
    WarehouseCreateComponent,
    OrdersComponent,
    ReportsComponent,
    AboutComponent,
    CatproductComponent,
    AddproductComponent,
    DeleteproductComponent,
    EditproductComponent,
    CategoriesComponent,
    PendingordersComponent,
    ProgressordersComponent,
    OutstandingordersComponent,
    AddorderComponent,
    EditorderComponent,
    DeleteorderComponent,
    CategoryCreateComponent,
    ProfileComponent,
    OrderDetailsComponent,
    ViewOrdersComponent,
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
    MatMenuModule,
    MatIconModule,
    MatCardModule,
    MatGridListModule,
    MatListModule,
    MatDialogModule,
    MatFormFieldModule,
    MatInputModule,
    ReactiveFormsModule,
    FormsModule,
    MatSelectModule,
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
