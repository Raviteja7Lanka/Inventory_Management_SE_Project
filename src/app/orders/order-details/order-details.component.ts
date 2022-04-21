import { HttpClient } from '@angular/common/http';
import { Component, Inject, OnInit } from '@angular/core';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';

@Component({
  selector: 'app-order-details',
  templateUrl: './order-details.component.html',
  styleUrls: ['./order-details.component.css']
})
export class OrderDetailsComponent implements OnInit {

  order=null
  order_details:any= []
  constructor(@Inject(MAT_DIALOG_DATA) public data: any,private http:HttpClient) {
    this.order=this.data
   }

  ngOnInit(): void {
    
    if(this.order!=null)
     { 
    this.http.get<any>(`http://localhost:8085/order/details/${this.order['order_id']}`).subscribe(res => {
      this.order_details=[res];
      console.log(res)
    });  

  }
  }

}
