import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpClientModule } from '@angular/common/http';
import { FormBuilder, FormGroup, NgForm } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'app-search',
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.css']
})
export class SearchComponent implements OnInit{

  form: FormGroup;

  constructor(private formBuilder: FormBuilder, private http: HttpClient, private router: Router) {
  }

  ngOnInit(): void {
    this.form = this.formBuilder.group({
    })
  };

  submit(): void {
    console.log(this.form.getRawValue());
    //this.http.post('http://localhost:8080/create', this.form.getRawValue()).subscribe(()=>this.router.navigate(['/login']));
    this.router.navigate(['/database']);

  }
}
