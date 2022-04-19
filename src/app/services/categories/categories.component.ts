import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router,Navigation,NavigationExtras } from '@angular/router';
import {HttpClient} from '@angular/common/http';
@Component({
  selector: 'app-categories',
  templateUrl: './categories.component.html',
  styleUrls: ['./categories.component.css']
})
export class CategoriesComponent implements OnInit {
  warehouse=null;
  categories=null;
  constructor(private router: Router,private http: HttpClient ) { 
    let p=this.router.getCurrentNavigation()
    if (p!=null && p.extras.state) {
      this.warehouse = p.extras.state['warehouse'];
    }
  }
  
  ngOnInit(): void {

    console.log("From Categories");
    console.log(this.warehouse)
    if(this.warehouse!= null)
    {
      let warehouse_id= this.warehouse['warehouse_id']==null?'0': this.warehouse['warehouse_id'];
      console.log(warehouse_id)
      this.http.get<any>(`http://localhost:8085/warehouse-categories/${warehouse_id}`).subscribe(res => {
      console.log(res);
      this.categories=res;
    });
    }
  }
  viewProducts(category:any)
  {
    
  }

}
