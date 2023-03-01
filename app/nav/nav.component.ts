import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { HomeComponent } from '../home/home.component';
import { EventEmitter } from "@angular/core";

@Component({
  selector: 'app-nav',
  templateUrl: './nav.component.html',
  styleUrls: ['./nav.component.css']
})
export class NavComponent implements OnInit {
  static authEmitter = new EventEmitter<boolean>();
  authenticated = false;

  constructor(private http: HttpClient) {

  }

  ngOnInit() : void{
    NavComponent.authEmitter.subscribe(
      (auth: boolean) => {
        this.authenticated = auth;
      }
    );
  }

  logout() : void{
    this.http.post('http://localhost:8080/api/logout', {}, {withCredentials: true}).subscribe(() => this.authenticated = false);
  }
}