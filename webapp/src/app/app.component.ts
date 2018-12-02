import {Component} from '@angular/core';
import {ApiService} from './api.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {

  public setup: boolean = null;

  constructor(private api: ApiService) {
    this.api.getToken().subscribe(
      (token) => this.setup = true,
      (_) => this.setup = false,
    );
  }

}
