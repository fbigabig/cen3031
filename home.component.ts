import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { EventEmitter } from "@angular/core";

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  static authEmitter = new EventEmitter<boolean>();
  message = 'You are not logged in';

  constructor(private http: HttpClient) {

  }

  ngOnInit(): void {
    this.http.get('http://localhost:8000/', { withCredentials: true }).subscribe((res: any) => {
      this.message = 'Welcome back';
      HomeComponent.authEmitter.emit(true);
    },
      err => {
        console.log(err);
        HomeComponent.authEmitter.emit(false);
      });
  }

}