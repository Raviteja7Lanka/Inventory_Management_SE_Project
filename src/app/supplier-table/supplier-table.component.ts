import { Component, OnInit } from '@angular/core';
import { AddSupplierComponent } from './add-supplier/add-supplier.component';
import { MatDialog } from '@angular/material/dialog';
import { ApiService } from '../services/api.service';
import { DeleteSupplierComponent } from './delete-supplier/delete-supplier.component';
import { EditSupplierComponent } from './edit-supplier/edit-supplier.component';

@Component({
  selector: 'app-supplier-table',
  templateUrl: './supplier-table.component.html',
  styleUrls: ['./supplier-table.component.css'],
})
export class SupplierTableComponent implements OnInit {
  constructor(private dialog: MatDialog, private api: ApiService) {}

  ngOnInit(): void {}
  displayedColumns = [
    'supplierId',
    'name',
    'address',
    'phone',
    'email',
    'otherDetails',
  ];

  add_supplier() {
    this.dialog.open(AddSupplierComponent, {
      width: '40%',
    });
  }

  delete_supplier() {
    this.dialog.open(DeleteSupplierComponent, {
      width: '40%',
    });
  }

  edit_supplier() {
    this.dialog.open(EditSupplierComponent, {
      width: '40%',
    });
  }

  /* 
   <--------------- After Getting the API, Comment out the below code ----------------> 
   datasource=this.api.get
  
  */
  datasource = supplierData;
}

export interface suppliers {
  supplierId: number;
  name: string;
  address: string;
  phone: number;
  email: string;
  // fax: number;
  otherDetails: string;
}

const supplierData: suppliers[] = [
  {
    supplierId: 1,
    name: 'ABC',
    address: 'abcaddress',
    phone: 12345,
    email: 'abc@gmail.com',
    // fax: 123456,
    otherDetails: 'nothing',
  },
  {
    supplierId: 1,
    name: 'ABC',
    address: 'abcaddress',
    phone: 12345,
    email: 'abc@gmail.com',
    // fax: 123456,
    otherDetails: 'nothing',
  },
  {
    supplierId: 1,
    name: 'ABC',
    address: 'abcaddress',
    phone: 12345,
    email: 'abc@gmail.com',
    // fax: 123456,
    otherDetails: 'nothing',
  },
  {
    supplierId: 1,
    name: 'ABC',
    address: 'abcaddress',
    phone: 12345,
    email: 'abc@gmail.com',
    // fax: 123456,
    otherDetails: 'nothing',
  },
];
