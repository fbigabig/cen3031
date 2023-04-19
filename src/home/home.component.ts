import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { EventEmitter } from "@angular/core";
import { HomeServiceService } from '../home/home-service.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  static authEmitter = new EventEmitter<boolean>();
  message = 'You are not logged in';

  constructor(private http: HttpClient, private HomeServiceService: HomeServiceService) { }
  
  libary: any;

  ngOnInit(): void {

    this.HomeServiceService.getLibrary().subscribe(data => {
      this.libary = data;
    })

    this.http.get('http://localhost:8080/home', { withCredentials: true }).subscribe((res: any) => {
      this.message = 'Login successful';
      HomeComponent.authEmitter.emit(true);
    },
      err => {
        console.log(err);
        HomeComponent.authEmitter.emit(false);
      });
  }



}
