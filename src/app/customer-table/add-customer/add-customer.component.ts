import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators, Form } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { ApiService } from 'src/app/services/api.service';

@Component({
  selector: 'app-add-customer',
  templateUrl: './add-customer.component.html',
  styleUrls: ['./add-customer.component.css'],
})
export class AddCustomerComponent implements OnInit {
  constructor(
    private formbuilder: FormBuilder,
    private dialogRef: MatDialogRef<AddCustomerComponent>,
    private api: ApiService
  ) {}
  addCustomerForm!: FormGroup;
  ngOnInit(): void {
    this.addCustomerForm = this.formbuilder.group({
      customer_id: parseInt((Math.random()*1000).toString()).toString(),
      first_name: ['', Validators.required],
      last_name: ['', Validators.required],
      address: ['', Validators.required],
      phone: ['', Validators.required],
      email: ['', Validators.required],
    });
  }
  customerAdd() {
    if (this.addCustomerForm.valid) {
      this.api.postCustomer(this.addCustomerForm.value).subscribe({
        next: (res) => {
          alert('Customer Added Successfully');
          this.addCustomerForm.reset();
          this.dialogRef.close();
          location.reload()
        },
        error: () => {
          alert(
            'There was an error occured while adding the Customer, Please try again !'
          );
          this.dialogRef.close();
        },
      });
    }
  }
}
