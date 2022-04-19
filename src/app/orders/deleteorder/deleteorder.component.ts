import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators, Form } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { ApiService } from 'src/app/services/api.service';

@Component({
  selector: 'app-deleteorder',
  templateUrl: './deleteorder.component.html',
  styleUrls: ['./deleteorder.component.css'],
})
export class DeleteorderComponent implements OnInit {
  constructor(
    private formbuilder: FormBuilder,
    private dialogRef: MatDialogRef<DeleteorderComponent>,
    private api: ApiService
  ) {}

  ngOnInit(): void {}

  deleteOrderForm!: FormGroup;

  orderDelete() {}
}
