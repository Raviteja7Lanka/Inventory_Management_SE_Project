import { Component, OnInit } from '@angular/core';
import { RouterLink, Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder } from '@angular/forms';
import { MatDialog } from '@angular/material/dialog';

@Component({
  selector: 'app-products',
  templateUrl: './products.component.html',
  styleUrls: ['./products.component.css'],
})
export class ProductsComponent implements OnInit {
  constructor(private dialog: MatDialog,private formBuilder:FormBuilder,private http:HttpClient, private router:Router) {}
  categories:any=[]
  products:any=[]
  productsExist=true;
  ngOnInit(): void {
    this.http.get<any>(`http://localhost:8085/category/all`).subscribe(res => {
      console.log(res);
      this.categories=res;
    });
  }
  getProducts(category:any)
  {
    this.http.get<any>(`http://localhost:8085/product/cat/${category.category_id}`).subscribe(res => {
      console.log(res);
      this.products=res;
      if(this.products.length==0)
        this.productsExist=false
    });
    
  }

}
