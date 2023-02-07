import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'CEN3032';
}

var message = 'Hello World!';

function log(message)
{
  console.log(message);
}

log(message);



