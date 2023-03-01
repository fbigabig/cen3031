import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormGroup} from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit{
  form: FormGroup;

  constructor(private formBuilder: FormBuilder, 
    private http: HttpClient, private router: Router){
    
  }

  ngOnInit(): void{
    this.form = this.formBuilder.group({
      username: '',
      password: ''
    });
  }

  submit(): void{

    let tempObj = this.form.getRawValue();

    let formattedInfo = JSON.stringify(tempObj);
    //console.log(this.form.getRawValue());
    this.http.post('http://localhost:8080/login', formattedInfo, {withCredentials: true}).subscribe(()=> this.router.navigate(['/']));
    //this.router.navigate(['/']);
  }

}
