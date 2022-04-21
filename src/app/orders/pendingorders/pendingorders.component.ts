import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder } from '@angular/forms';
import { MatDialog } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { OrderDetailsComponent } from '../order-details/order-details.component';

@Component({
  selector: 'app-pendingorders',
  templateUrl: './pendingorders.component.html',
  styleUrls: ['./pendingorders.component.css'],
})
export class PendingordersComponent implements OnInit {
  constructor(private dialog: MatDialog,private formBuilder:FormBuilder,private http:HttpClient, private router:Router) {}
  datasource= [];
  ngOnInit(): void {

    this.http.get<any>(`http://localhost:8085/orders/status/completed`).subscribe(res => {
      this.datasource=res;
      console.log(res)
  });   
  }
  viewOrderDetails(order:any)
  {
    
    this.dialog.open(OrderDetailsComponent, {
      width: '80%',
      height:'30%',
      data: order
    });
  }
  // datasource: any = [
  //   {
  //     order_id: 1,
  //     date_of_order: '01/26/2022',
  //     customer_id: '1',
  //     supplier_id: '5',
  //     order_Details: 'nothing',
  //   },
  //   {
  //     order_id: 2,
  //     date_of_order: '01/26/2022',
  //     customer_id: '76',
  //     supplier_id: '2',
  //     order_Details: 'nothing',
  //   },
  //   {
  //     order_id: 3,
  //     date_of_order: '01/26/2022',
  //     customer_id: '46',
  //     supplier_id: '82',
  //     order_Details: 'nothing',
  //   },
  //   {
  //     order_id: 4,
  //     date_of_order: '01/26/2022',
  //     customer_id: '100',
  //     supplier_id: '33',
  //     order_Details: 'nothing',
  //   },
  //   {
  //     order_id: 5,
  //     date_of_order: '01/26/2022',
  //     customer_id: '22',
  //     supplier_id: '46',
  //     order_Details: 'nothing',
  //   },
  //   {
  //     order_id: 6,
  //     date_of_order: '01/26/2022',
  //     customer_id: '99',
  //     supplier_id: '4',
  //     order_Details: 'nothing',
  //   },
  //   {
  //     order_id: 7,
  //     date_of_order: '01/26/2022',
  //     customer_id: '1',
  //     supplier_id: '5',
  //     order_Details: 'nothing',
  //   },
  // ];

  displayedColumns = [
    'order_id',
    'date_of_order',
    'order_type',
    'order_Details',
  ];
}
