import { HttpClient } from '@angular/common/http';
import { Component, Inject, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-addproduct',
  templateUrl: './addproduct.component.html',
  styleUrls: ['./addproduct.component.css'],
})
export class AddproductComponent implements OnInit {
  category=null
  suppliers=null
  constructor(@Inject(MAT_DIALOG_DATA) public data: any,private formBuilder:FormBuilder,private http:HttpClient, private router:Router,private route: ActivatedRoute) {
    this.category=data
  }
  public  addProductForm !: FormGroup
  ngOnInit(): void {
    this.http.get<any>(`http://localhost:8085/supplier/all`).subscribe(res => {
      console.log(res);
      this.suppliers=res;
    });
    this.addProductForm = this.formBuilder.group(
      {
        product_name:[''],
        product_descr_ip_tion: [''],
        warehouse_id: [''],
        product_id: parseInt((Math.random()*10000).toString()).toString(),
        category_id:[''],
        product_status: [''],
        product_quantity: [''],
        product_price: [''],
        product_unit: [''],
        supplier_id:[''],
        other_details:['']
      }
      
    )
  }

  closeAddProduct()
  {
    
  }
  addProduct()
  {
    if(this.category!=null)
    {
    this.addProductForm.value.category_id= this.category['category_id']==null?'0': this.category['category_id']
    this.addProductForm.value.warehouse_id= this.category['warehouse_id']==null?'0': this.category['warehouse_id']
    }
    console.log(this.addProductForm.value)
    this.http.post<any>("http://localhost:8085/product/add",this.addProductForm.value)
    .subscribe(res=>{
      alert("Added New Product");
      this.addProductForm.reset();
      this .router.navigate(['warehouse']);
    },err=>{
      alert("something went wrong")
    });
  }

}
