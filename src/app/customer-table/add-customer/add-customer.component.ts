import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators, Form } from '@angular/forms';

@Component({
  selector: 'app-add-customer',
  templateUrl: './add-customer.component.html',
  styleUrls: ['./add-customer.component.css'],
})
export class AddCustomerComponent implements OnInit {
  constructor(private formbuilder: FormBuilder) {}
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
    console.log(this.addCustomerForm.value);
  }
}
