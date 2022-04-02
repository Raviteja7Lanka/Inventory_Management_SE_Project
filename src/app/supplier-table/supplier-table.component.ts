import { Component, OnInit } from '@angular/core';
import { AddSupplierComponent } from './add-supplier/add-supplier.component';
import { MatDialog } from '@angular/material/dialog';
import { ApiService } from '../services/api.service';
import { HttpClient } from '@angular/common/http';
import { DeleteSupplierComponent } from './delete-supplier/delete-supplier.component';
import { EditSupplierComponent } from './edit-supplier/edit-supplier.component';

@Component({
  selector: 'app-supplier-table',
  templateUrl: './supplier-table.component.html',
  styleUrls: ['./supplier-table.component.css'],
})
export class SupplierTableComponent implements OnInit {
  constructor(private dialog: MatDialog, private api: ApiService) {}
  datasource: any = [];
  ngOnInit(): void {
    this.api.getAllSuppliers().subscribe({
      next: (res) => {
        console.log(this.datasource);
        this.datasource= res
      },
      error: () => {
        alert('There was an error loading Suppliers Information');

        console.log(this.datasource);
      },
    });
  }
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
  // datasource = supplierData;
}
export interface suppliers {
  supplierId: number;
  name: string;
  address: string;
  phone: number;
  email: string;
  fax: number;
  otherDetails: string;
}


const supplierData: suppliers[] = [
  {
    supplierId: 1,
    name: 'Ravi',
    address: '4000 SW 37TH BLVD',
    phone: 12345568,
    email: 'ravi@gmail.com',
    fax: 123456,
    otherDetails: 'nothing',
  },
  {
    supplierId: 1,
    name: 'Aditya',
    address: '3999 SW 40TH BLVD',
    phone: 19999454,
    email: 'aditya@gmail.com',
    fax: 9999,
    otherDetails: 'nothing',
  },
  {
    supplierId: 1,
    name: 'Aditya',
    address: '3999 SW 40TH BLVD',
    phone: 19999454,
    email: 'aditya@gmail.com',
    fax: 9999,
    otherDetails: 'nothing',
  },
  {
    supplierId: 1,
    name: 'Aditya',
    address: '3999 SW 40TH BLVD',
    phone: 19999454,
    email: 'aditya@gmail.com',
    fax: 9999,
    otherDetails: 'nothing',
  }
];
