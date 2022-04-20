import { HttpClient } from '@angular/common/http';
import { Component, Inject, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';
import { Router, ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-addorder',
  templateUrl: './addorder.component.html',
  styleUrls: ['./addorder.component.css'],
})
export class AddorderComponent implements OnInit {

  constructor(@Inject(MAT_DIALOG_DATA) public data: any,private formBuilder:FormBuilder,private http:HttpClient, private router:Router,private route: ActivatedRoute) { }
  public  addOrderForm !: FormGroup
  type_order=""
  suppliers:any=[]
  customers:any=[]
  ngOnInit(): void {

    this.addOrderForm = this.formBuilder.group(
      {
        order_id: parseInt((Math.random()*10000).toString()).toString(),
        order_status: [''],
        order_date: [''],
        other_details:[ ''],
        order_type: [''],
        unit_price: [''],
        quantity: [''],
        discount: [''],
        total: [''],
        date_of_order: [''],
        product_id: [''],
        billnumber: parseInt((Math.random()*10000).toString()).toString(),
        payment_type: [''],
        order_by:['']
      }
    );

    
  }
  addOrder()
  {
    if(this.addOrderForm.value.order_date!=null)
      {
       let order_date= this.addOrderForm.value.order_date
       this.addOrderForm.value.order_date= order_date.toString()
       this.addOrderForm.value.date_of_order=order_date.toString()
      }
    console.log(this.addOrderForm.value)
    if(this.addOrderForm.value.order_type=="Supplier")
      this.addOrderForm.value.order_status="outstanding"
    else
      this.addOrderForm.value.order_status="in_progress"

    this.http.post<any>("http://localhost:8085/orders/add",this.addOrderForm.value)
    .subscribe(res=>{
      // alert("Added New Order");
      
      // this .router.navigate(['warehouse']);
    },err=>{
      alert("something went wrong")
    });
    console.log(this.addOrderForm.value)
    this.http.post<any>("http://localhost:8085/orders/details/add",this.addOrderForm.value)
    .subscribe(res=>{
      // alert("Added Order Details");
      
      // this .router.navigate(['warehouse']);
    },err=>{
      alert("something went wrong")
    });
    this.http.post<any>("http://localhost:8085/payment/add",this.addOrderForm.value)
    .subscribe(res=>{
      alert("Added new order");
      this.addOrderForm.reset();
      location.reload();
    },err=>{
      alert("something went wrong")
    });
  }
  closeAddOrder()
  {

  }

  onOrderTypeSelected(val:any)
  {
    this.type_order=val
    
    if(val=="Supplier")
    {
    this.http.get<any>('http://localhost:8085/supplier/all').subscribe({
      next: (res) => {
        this.suppliers = res;
        console.log(this.suppliers);
        location.reload();
      },
      error: () => {
        console.log('There was an error loading Suppliers Information');
      },
    });
  }
  if(val="Customer")
  {
    this.http.get<any>('http://localhost:8085/customer/all').subscribe({
      next: (res) => {
        this.customers = res;
        console.log(this.suppliers);
      },
      error: () => {
        console.log('There was an error loading Customers Information');
      },
    });
  }

  }
}
