import { Component, OnInit, Inject } from '@angular/core';
import { MatDialog, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { AddCustomerComponent } from './add-customer/add-customer.component';
import { DeleteCustomerComponent } from './delete-customer/delete-customer.component';
import { EditCustomerComponent } from './edit-customer/edit-customer.component';

@Component({
  selector: 'app-customer-table',
  templateUrl: './customer-table.component.html',
  styleUrls: ['./customer-table.component.css'],
})
export class CustomerTableComponent implements OnInit {
  constructor(private dialog: MatDialog) {}

  ngOnInit(): void {}
  displayedColumns = [
    'customerId',
    'name',
    'address',
    'phone',
    'email',
    'fax',
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

  datasource = customerData;
}

export interface customers {
  customerId: number;
  name: string;
  address: string;
  phone: number;
  email: string;
  fax: number;
  otherDetails: string;
}

const customerData: customers[] = [
  {
    customerId: 1,
    name: 'ABC',
    address: 'abcaddress',
    phone: 12345,
    email: 'abc@gmail.com',
    fax: 123456,
    otherDetails: 'nothing',
  },
  {
    customerId: 2,
    name: 'ABC',
    address: 'abcaddress',
    phone: 12345,
    email: 'abc@gmail.com',
    fax: 123456,
    otherDetails: 'nothing',
  },
  {
    customerId: 3,
    name: 'ABC',
    address: 'abcaddress',
    phone: 12345,
    email: 'abc@gmail.com',
    fax: 123456,
    otherDetails: 'nothing',
  },
  {
    customerId: 4,
    name: 'ABC',
    address: 'abcaddress',
    phone: 12345,
    email: 'abc@gmail.com',
    fax: 123456,
    otherDetails: 'nothing',
  },
];
