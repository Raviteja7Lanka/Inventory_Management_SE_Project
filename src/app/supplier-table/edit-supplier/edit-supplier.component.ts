import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { DialogRole, MatDialogRef } from '@angular/material/dialog';
import { ApiService } from '../../services/api.service';

@Component({
  selector: 'app-edit-supplier',
  templateUrl: './edit-supplier.component.html',
  styleUrls: ['./edit-supplier.component.css'],
})
export class EditSupplierComponent implements OnInit {
  constructor(
    private formbuilder: FormBuilder,
    private api: ApiService,
    private dialogref: MatDialogRef<EditSupplierComponent>
  ) {}

  editSupplierForm!: FormGroup;

  ngOnInit(): void {
    this.editSupplierForm = this.formbuilder.group({
      supplier_Id: ['', Validators.required],
      // name: ['', Validators.required],
      // address: ['', Validators.required],
      // phone: ['', Validators.required],
      // email: ['', Validators.required],
      // other_details: [''],
    });
  }

  editSupplier() {
    if (this.editSupplierForm.valid) {
      console.log(this.editSupplierForm.value);
      this.api.updateSupplier(this.editSupplierForm.value).subscribe({
        next: (res) => {
          alert('Supplier Updated Successfully');
          this.editSupplierForm.reset();
          this.dialogref.close();
        },
        error: () => {
          alert(
            'There was an error occurred while updating the supplier, Please Try Again'
          );
          this.dialogref.close();
        },
      });
    }
  }
}
