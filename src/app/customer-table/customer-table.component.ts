import { Component, OnInit, Inject } from '@angular/core';
import { ApiService } from '../services/api.service';
import { MatDialog, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { AddCustomerComponent } from './add-customer/add-customer.component';
import { DeleteCustomerComponent } from './delete-customer/delete-customer.component';
import { EditCustomerComponent } from './edit-customer/edit-customer.component';
import { NavigationExtras, Route, Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-customer-table',
  templateUrl: './customer-table.component.html',
  styleUrls: ['./customer-table.component.css'],
})
export class CustomerTableComponent implements OnInit {
  constructor(private dialog: MatDialog,private api: ApiService,private router:Router,private http: HttpClient) {}
  datasource: any;
  ngOnInit(): void {
    this.api.getAllCustomers().subscribe(res => {
      console.log(res);
      this.datasource=res
  });
}
  displayedColumns = [
    'customerId',
    'name',
    'address',
    'phone',
    'email',
    // 'fax',
    'otherDetails',
  ];

  add_customer() {
    this.dialog.open(AddCustomerComponent, {
      width: '40%',
    });
  }

  delete_customer() {
    this.dialog.open(DeleteCustomerComponent, {
      width: '40%',
    });
  }

  edit_customer() {
    this.dialog.open(EditCustomerComponent, {
      width: '40%',
    });
  }
  viewOrders(customer:any)
  {
    let navigationExtras: NavigationExtras = {
      state: {
        customer: customer
      }
    };
    // this.router.navigate(['/user/raja'], navigationExtras);
    this.router.navigate(["/view-customer-orders"], navigationExtras);

  }
  deleteCustomer(customer:any)
  {
    let customer_id= customer['customer_id']
    this.http.delete<any>(`http://localhost:8085/customer/${customer_id}`).subscribe(res => {
      console.log(res);
      location.reload();
    });
  
  }


}

export interface customers {
  customerId: number;
  name: string;
  address: string;
  phone: number;
  email: string;
  // fax: number;
  otherDetails: string;
}

const customerData: customers[] = [
  {
    customerId: 1,
    name: 'Ravi',
    address: '4000 SW 37TH BLVD',
    phone: 12345568,
    email: 'ravi@gmail.com',
    // fax: 123456,
    otherDetails: 'nothing',
  },
  {
    customerId: 2,
    name: 'SURYA',
    address: '3800 SW 35TH BLVD',
    phone: 76543245,
    email: 'surya@gmail.com',
    // fax: 123456,
    otherDetails: 'nothing',
  },
  {
    customerId: 3,
    name: 'Aditya',
    address: '3999 SW 40TH BLVD',
    phone: 19999454,
    email: 'aditya@gmail.com',
    // fax: 9999,
    otherDetails: 'nothing',
  },
  {
    customerId: 4,
    name: 'saikiran',
    address: '4002 SW 49TH BLVD',
    phone: 10000345,
    email: 'saikiran@gmail.com',
    // fax: 123456,
    otherDetails: 'nothing',
  },
];
