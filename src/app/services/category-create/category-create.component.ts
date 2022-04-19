import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'app-category-create',
  templateUrl: './category-create.component.html',
  styleUrls: ['./category-create.component.css']
})
export class CategoryCreateComponent implements OnInit {

  public  addCategoryForm !: FormGroup;
  warehouses=null;
  constructor(private formBuilder:FormBuilder,private http:HttpClient, private router:Router) { }
  
  ngOnInit(): void {
    this.http.get<any>(`http://localhost:8085/warehouse/all`).subscribe(res => {
      this.warehouses=res;
      // console.log(res)
  });   

    this.addCategoryForm = this.formBuilder.group(
      {
        name:[''],
        description: [''],
        warehouse_id: [''],
        category_id: parseInt((Math.random()*10000).toString()).toString()
      }
    )
  }

  addCategory()
  {
    // console.log(this.addCategoryForm.value)
    this.http.post<any>("http://localhost:8085/category/add",this.addCategoryForm.value)
    .subscribe(res=>{
      alert("Added New Category");
      this.addCategoryForm.reset();
      this .router.navigate(['warehouse']);
    },err=>{
      alert("something went wrong")
    });
  }
}
