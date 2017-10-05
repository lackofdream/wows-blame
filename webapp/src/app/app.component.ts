import { Component } from '@angular/core';
import { AppService } from './app.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  setup: boolean;
  constructor(
    private appService: AppService,
  ) {
    this.appService.isSetup().subscribe(result => this.setup = result);
  }


}
