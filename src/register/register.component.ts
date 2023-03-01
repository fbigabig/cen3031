import { Component, OnInit } from '@angular/core';
import { Form, FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit{
  form: FormGroup;

  constructor(private formBuilder: FormBuilder, private http: HttpClient, private router: Router) {
  }

  ngOnInit(): void {
    this.form = this.formBuilder.group({
      username: '',
      password: ''
    })
  };

  submit(): void {

    let tempObj = this.form.getRawValue();

    let formattedInfo = JSON.stringify(tempObj);

    //console.log(this.form.getRawValue());
    this.http.post('http://localhost:8080/create', formattedInfo).subscribe(()=>this.router.navigate(['/login']));
    //this.router.navigate(['/login']);

  }
}
