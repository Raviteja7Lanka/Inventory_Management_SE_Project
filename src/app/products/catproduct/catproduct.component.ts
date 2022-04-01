import { Component, OnInit } from '@angular/core';
import { ProductsComponent } from '../products.component';

@Component({
  selector: 'app-catproduct',
  templateUrl: './catproduct.component.html',
  styleUrls: ['./catproduct.component.css'],
})
export class CatproductComponent implements OnInit {
  constructor() {}

  ngOnInit(): void {}

  add_product() {}
  delete_product() {}
  edit_product() {}

  datasource = products;

  displayedColumns = [];
}
const products: any = [];
