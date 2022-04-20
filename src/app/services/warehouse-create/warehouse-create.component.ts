import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'app-warehouse-create',
  templateUrl: './warehouse-create.component.html',
  styleUrls: ['./warehouse-create.component.css']
})
export class WarehouseCreateComponent implements OnInit {
  
  public  addWarehouseForm !: FormGroup;
  constructor(private formBuilder:FormBuilder,private http:HttpClient, private router:Router) { }

  ngOnInit(): void {

    this.addWarehouseForm = this.formBuilder.group(
      {
        name:[''],
        location: [''],
        capacity: [''],
        description: [''],
        warehouse_id: parseInt((Math.random()*10000).toString()).toString()
      }
    )
  }

  addWarehouse()
  {
   
    this.addWarehouseForm.value.capacity= parseInt(this.addWarehouseForm.value.capacity)
    this.http.post<any>("http://localhost:8085/warehouse/add",this.addWarehouseForm.value)
    .subscribe(res=>{
      alert("Added New Warehouse");
      this.addWarehouseForm.reset();
      this .router.navigate(['warehouse']);
    },err=>{
      alert("something went wrong")
    }
    )
  }

}
