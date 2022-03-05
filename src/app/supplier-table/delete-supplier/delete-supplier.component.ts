import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators, Form } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { ApiService } from '../../services/api.service';
@Component({
  selector: 'app-delete-supplier',
  templateUrl: './delete-supplier.component.html',
  styleUrls: ['./delete-supplier.component.css'],
})
export class DeleteSupplierComponent implements OnInit {
  constructor(
    private formbuilder: FormBuilder,
    private dialogref: MatDialogRef<DeleteSupplierComponent>,
    private api: ApiService
  ) {}

  ngOnInit(): void {}

  deleteSupplierForm!: FormGroup;

  delete_supplier() {}
}
