import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { AddorderComponent } from './addorder/addorder.component';
import { EditorderComponent } from './editorder/editorder.component';
import { DeleteorderComponent } from './deleteorder/deleteorder.component';

@Component({
  selector: 'app-orders',
  templateUrl: './orders.component.html',
  styleUrls: ['./orders.component.css'],
})
export class OrdersComponent implements OnInit {
  constructor(private dialog: MatDialog) {}

  ngOnInit(): void {}

  add_order() {
    
    this.dialog.open(AddorderComponent, {
      width: '40%',
    });
  }
  delete_order() {
    this.dialog.open(DeleteorderComponent, {
      width: '40%',
    });
  }
  edit_order() {
    this.dialog.open(EditorderComponent, {
      width: '40%',
    });
  }
}
