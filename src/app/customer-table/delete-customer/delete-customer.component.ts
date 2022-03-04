import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators, Form } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { ApiService } from 'src/app/services/api.service';

@Component({
  selector: 'app-delete-customer',
  templateUrl: './delete-customer.component.html',
  styleUrls: ['./delete-customer.component.css'],
})
export class DeleteCustomerComponent implements OnInit {
  constructor(
    private formbuilder: FormBuilder,
    private dialogref: MatDialogRef<DeleteCustomerComponent>,
    private api: ApiService
  ) {}
  deleteCustomerForm!: FormGroup;
  ngOnInit(): void {
    this.deleteCustomerForm = this.formbuilder.group({
      firstName: ['', Validators.required],
      lastName: ['', Validators.required],
      emailId: ['', Validators.required],
    });
  }
  delete_customer() {
    if (this.deleteCustomerForm.valid) {
      this.dialogref.close();
    }
  }
}
