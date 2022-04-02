import { Component, OnInit } from '@angular/core';
import { ProductsComponent } from '../products.component';
import { MatDialog } from '@angular/material/dialog';
import { ApiService } from 'src/app/services/api.service';
import { AddproductComponent } from './addproduct/addproduct.component';
import { DeleteproductComponent } from './deleteproduct/deleteproduct.component';
import { EditproductComponent } from './editproduct/editproduct.component';

@Component({
  selector: 'app-catproduct',
  templateUrl: './catproduct.component.html',
  styleUrls: ['./catproduct.component.css'],
})
export class CatproductComponent implements OnInit {
  constructor(private dialog: MatDialog, private api: ApiService) {}

  ngOnInit(): void {}

  add_product() {
    this.dialog.open(AddproductComponent, {
      width: '40%',
    });
  }
  delete_product() {
    this.dialog.open(DeleteproductComponent, {
      width: '40%',
    });
  }
  edit_product() {
    this.dialog.open(EditproductComponent, {
      width: '40%',
    });
  }

  datasource = products;

  displayedColumns = [
    'ID',
    'Name',
    'Description',
    'Unit',
    'Price',
    'Quantity',
    'Status',
  ];
}
const products: any = [];
