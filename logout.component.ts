import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import {FormBuilder, FormGroup} from '@angular/forms';
import { HttpClient, HttpClientModule } from '@angular/common/http';
@Component({
  selector: 'app-logout',
  templateUrl: './logout.component.html',
  styleUrls: ['./logout.component.css']
})
export class LogoutComponent implements OnInit{

  constructor(private formBuilder: FormBuilder, 
    private http: HttpClient, private router: Router){
    
  }
  ngOnInit(): void {
    throw new Error('Method not implemented.');
  }

 logout() : void{
    this.http.post('http://localhost:8080/api/logout', {}, {withCredentials: true}).subscribe(() => this.authenticated = false);
    this.router.navigate(['/']);
  }

  nologout():void{
    this.router.navigate(['/']);
  }
}