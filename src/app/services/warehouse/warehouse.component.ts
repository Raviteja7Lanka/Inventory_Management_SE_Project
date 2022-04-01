import { Component, OnInit } from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';

@Component({
  selector: 'app-warehouse',
  templateUrl: './warehouse.component.html',
  styleUrls: ['./warehouse.component.css']
})

export class WarehouseComponent implements OnInit {
  pageSizeValue = [3, 5, 7, 10];


  constructor(private route: ActivatedRoute,private router: Router,private http: HttpClient ) {

   }
  
  

  ngOnInit(): void {
   
  }
  
  

}
export interface ResponseBody {
  message: string;
}
