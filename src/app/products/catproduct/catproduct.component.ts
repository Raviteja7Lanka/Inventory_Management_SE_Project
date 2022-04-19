import { Component, OnInit } from '@angular/core';
import { ProductsComponent } from '../products.component';
import { MatDialog } from '@angular/material/dialog';
import { ApiService } from 'src/app/services/api.service';
import { AddproductComponent } from './addproduct/addproduct.component';
import { DeleteproductComponent } from './deleteproduct/deleteproduct.component';
import { EditproductComponent } from './editproduct/editproduct.component';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-catproduct',
  templateUrl: './catproduct.component.html',
  styleUrls: ['./catproduct.component.css'],
})
export class CatproductComponent implements OnInit {
  constructor(private dialog: MatDialog,private router: Router,private http: HttpClient ) {
    let p=this.router.getCurrentNavigation()
    if (p!=null && p.extras.state) {
      this.category = p.extras.state['category'];
    }
  }
  category=null;
  products=null;
  ngOnInit(): void {
    console.log("From Categories");
    console.log(this.category)
    if(this.category!= null)
    {
      let category_id= this.category['category_id']==null?'0': this.category['category_id'];
      let warehouse_id= this.category['warehouse_id']==null?'0': this.category['warehouse_id'];
      console.log(category_id)
      console.log(warehouse_id)
      this.http.get<any>(`http://localhost:8085/category-products/${warehouse_id}/${category_id}`).subscribe(res => {
      console.log(res);
      this.products=res;
    });
    }
  }

  add_product() {
    this.dialog.open(AddproductComponent, {
      width: '60%',
      height: '95%',
      data: this.category
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

  editProduct(product:any)
  {

  }

}

