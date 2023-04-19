import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class HomeServiceService {

  constructor(private httpClient: HttpClient) { }

  getLibrary(){

   return this.httpClient.get('http://localhost:8080/home');
    // needs to be url where value is given

  }

}
