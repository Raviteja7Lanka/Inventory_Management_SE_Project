import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators, Form } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { ApiService } from '../../services/api.service';

@Component({
  selector: 'app-edit-customer',
  templateUrl: './edit-customer.component.html',
  styleUrls: ['./edit-customer.component.css'],
})
export class EditCustomerComponent implements OnInit {
  constructor(
    private dialogref: MatDialogRef<EditCustomerComponent>,
    private api: ApiService,
    private formbuilder: FormBuilder
  ) {}
  editCustomerForm!: FormGroup;

  ngOnInit(): void {}

  edit_customer() {}
}
