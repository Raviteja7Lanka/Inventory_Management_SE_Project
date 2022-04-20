<<<<<<< HEAD
import { HttpClient } from '@angular/common/http';
import { Component, Inject, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';
import { Router, ActivatedRoute } from '@angular/router';
=======
import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators, Form } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { ApiService } from 'src/app/services/api.service';
>>>>>>> e5e0ea2d36763bee6cb148035414e4eb72e47e8f

@Component({
  selector: 'app-addorder',
  templateUrl: './addorder.component.html',
  styleUrls: ['./addorder.component.css'],
})
export class AddorderComponent implements OnInit {
<<<<<<< HEAD

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
=======
  constructor(
    private formbuilder: FormBuilder,
    private dialogRef: MatDialogRef<AddorderComponent>,
    private api: ApiService
  ) {}
  addOrderForm!: FormGroup;
  ngOnInit(): void {}
  orderAdd() {}
>>>>>>> e5e0ea2d36763bee6cb148035414e4eb72e47e8f
}
