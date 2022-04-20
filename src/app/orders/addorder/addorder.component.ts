import { HttpClient } from '@angular/common/http';
import { Component, Inject, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';
import { Router, ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-addorder',
  templateUrl: './addorder.component.html',
  styleUrls: ['./addorder.component.css']
})
export class AddorderComponent implements OnInit {

  constructor(@Inject(MAT_DIALOG_DATA) public data: any,private formBuilder:FormBuilder,private http:HttpClient, private router:Router,private route: ActivatedRoute) { }
  public  addOrderForm !: FormGroup
  ngOnInit(): void {

    this.addOrderForm = this.formBuilder.group(
      {
        order_name:[''],
        order_descr_ip_tion: [''],
        warehouse_id: [''],
        order_id: parseInt((Math.random()*10000).toString()).toString(),
        category_id:[''],
        order_status: [''],
        order_quantity: [''],
        order_price: [''],
        order_unit: [''],
        supplier_id:[''],
        other_details:['']
      }
    );
  }
  addOrder()
  {

  }
  closeAddOrder()
  {

  }
}
