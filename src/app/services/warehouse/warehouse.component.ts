import { Component, OnInit } from '@angular/core';
import {ActivatedRoute, NavigationExtras, Router} from '@angular/router';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import { AnyCatcher } from 'rxjs/internal/AnyCatcher';

@Component({
  selector: 'app-warehouse',
  templateUrl: './warehouse.component.html',
  styleUrls: ['./warehouse.component.css']
})

export class WarehouseComponent implements OnInit {
  pageSizeValue = [3, 5, 7, 10];
  warehouses=null;
  constructor(private route: ActivatedRoute,private router: Router,private http: HttpClient ) {

   }
  
  

  ngOnInit(): void {
    this.http.get<any>(`http://localhost:8085/warehouse/all`).subscribe(res => {
    // console.log(res);
    this.warehouses=res;
  });
      
  }
  
  viewCategories(warehouse:any)
  {
    // console.log(warehouse);
    let navigationExtras: NavigationExtras = {
      state: {
        warehouse: warehouse
      }
    };
    // this.router.navigate(['/user/raja'], navigationExtras);
    this.router.navigate(["/warehouse-categories"], navigationExtras);
  }
  

}
export interface ResponseBody {
  message: string;
}
