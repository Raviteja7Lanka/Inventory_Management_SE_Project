import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { ApiService } from '../../services/api.service';
import { DialogRole, MatDialogRef } from '@angular/material/dialog';

@Component({
  selector: 'app-add-supplier',
  templateUrl: './add-supplier.component.html',
  styleUrls: ['./add-supplier.component.css'],
})
export class AddSupplierComponent implements OnInit {
  constructor(
    private formbuilder: FormBuilder,
    private api: ApiService,
    private dialogRef: MatDialogRef<AddSupplierComponent> // private dialogRef: MatDialogRef,
  ) {}
  addSupplierForm!: FormGroup;

  ngOnInit(): void {
    this.addSupplierForm = this.formbuilder.group({
      firstName: ['', Validators.required],
      lastName: ['', Validators.required],
      address: ['', Validators.required],
      phoneNum: ['', Validators.required],
      emailId: ['', Validators.required],
    });
  }

  addSupplier() {
    if (this.addSupplierForm.valid) {
      console.log(this.addSupplierForm.value);
      this.api.postSupplier(this.addSupplierForm.value).subscribe({
        next: (res) => {
          alert('Supplier Added Successfully');
          this.addSupplierForm.reset();
          this.dialogRef.close();
        },
        error: () => {
          alert(
            'There was an error occured while adding the Supplier, Please try again !'
          );
          this.dialogRef.close();
        },
      });
    }
  }
}
