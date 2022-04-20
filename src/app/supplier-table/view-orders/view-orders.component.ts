import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { OrderDetailsComponent } from 'src/app/orders/order-details/order-details.component';

@Component({
  selector: 'app-view-orders',
  templateUrl: './view-orders.component.html',
  styleUrls: ['./view-orders.component.css']
})
export class ViewOrdersComponent implements OnInit {

  supplier:any="0"
  constructor(private dialog: MatDialog,private router: Router,private http: HttpClient ) {
    let p=this.router.getCurrentNavigation()
    if (p!=null && p.extras.state) {
      this.supplier = p.extras.state['supplier'];
    }
  }
  outstanding_orders=[]
  completed_orders=[]
  ngOnInit(): void {

    this.http.get<any>(`http://localhost:8085/orders/by/${this.supplier.supplier_id}/outstanding`).subscribe(res => {
    // console.log(res);
    this.outstanding_orders=res;
  });
  this.http.get<any>(`http://localhost:8085/orders/by/${this.supplier.supplier_id}/completed`).subscribe(res => {
    // console.log(res);
    this.completed_orders=res;
  });
  }

  displayedColumns = [
    'order_id',
    'date_of_order',
    'order_type',
    'action',
  ];
  viewOrderDetails(order:any)
  {
    
    this.dialog.open(OrderDetailsComponent, {
      width: '80%',
      height:'60%',
      data: order
    });
  }
  markOrderComplete(order:any)
  {
    if(order!=null)
    {
      order.order_status="completed"
    }
    this.http.put<any>(`http://localhost:8085/orders/${order.order_id}`,order)
    .subscribe(res=>{
      alert("Order Marked Complete");
      
      // this .router.navigate(['warehouse']);
    },err=>{
      alert("something went wrong")
    }); 
    
  }


}
