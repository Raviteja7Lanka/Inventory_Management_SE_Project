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
      firstName: ['', Validators.required],
      lastName: ['', Validators.required],
      address: ['', Validators.required],
      phoneNum: ['', Validators.required],
      emailId: ['', Validators.required],
    });
  }
  customerAdd() {
    if (this.addCustomerForm.valid) {
      this.api.postSupplier(this.addCustomerForm.value).subscribe({
        next: (res) => {
          alert('Customer Added Successfully');
          this.addCustomerForm.reset();
          this.dialogRef.close();
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
