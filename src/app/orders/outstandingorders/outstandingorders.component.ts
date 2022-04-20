import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-outstandingorders',
  templateUrl: './outstandingorders.component.html',
  styleUrls: ['./outstandingorders.component.css'],
})
export class OutstandingordersComponent implements OnInit {
  constructor() {}

  ngOnInit(): void {}
  datasource: any = [
    {
      order_id: 1,
      date_of_order: '01/26/2022',
      customer_id: '1',
      supplier_id: '5',
      order_Details: 'nothing',
    },
    {
      order_id: 2,
      date_of_order: '01/26/2022',
      customer_id: '76',
      supplier_id: '2',
      order_Details: 'nothing',
    },
    {
      order_id: 3,
      date_of_order: '01/26/2022',
      customer_id: '46',
      supplier_id: '82',
      order_Details: 'nothing',
    },
    {
      order_id: 4,
      date_of_order: '01/26/2022',
      customer_id: '100',
      supplier_id: '33',
      order_Details: 'nothing',
    },
    {
      order_id: 5,
      date_of_order: '01/26/2022',
      customer_id: '22',
      supplier_id: '46',
      order_Details: 'nothing',
    },
    {
      order_id: 6,
      date_of_order: '01/26/2022',
      customer_id: '99',
      supplier_id: '4',
      order_Details: 'nothing',
    },
    {
      order_id: 7,
      date_of_order: '01/26/2022',
      customer_id: '1',
      supplier_id: '5',
      order_Details: 'nothing',
    },
  ];
  displayedColumns = [
    'order_id',
    'date_of_order',
    'customer_id',
    'supplier_id',
    'order_Details',
  ];
}