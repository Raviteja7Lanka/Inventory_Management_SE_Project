import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class ApiService {
  constructor(private http: HttpClient) {}

  postSupplier(data: any) {
    return this.http.post<any>('http://localhost:8085/supplier/add', data);
  }
  getAllSuppliers(){
    return this.http.get<any>('http://localhost:8085/supplier/all')
  }
  getAllCustomers(){
    return this.http.get<any>('http://localhost:8085/customer/all') 
  }
  getSupplier() {
    return this.http.get<any>('SupplierAPI');
  }
  postCustomer(data: any) {
    return this.http.post<any>('http://localhost:8085/customer/add', data);
  }
  getCustomer() {
    return this.http.get<any>('CustomerAPI');
  }
}
