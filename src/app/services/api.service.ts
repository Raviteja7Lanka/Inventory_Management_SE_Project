import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class ApiService {
  constructor(private http: HttpClient) {}

  postSupplier(data: any) {
    return this.http.post<any>('SupplierAPI', data);
  }
  getSupplier() {
    return this.http.get<any>('SupplierAPI');
  }
  postCustomer(data: any) {
    return this.http.post<any>('CustomerAPI', data);
  }
  getCustomer() {
    return this.http.get<any>('CustomerAPI');
  }
}
