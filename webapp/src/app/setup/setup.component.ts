import {Component} from '@angular/core';
import {FormControl, Validators} from '@angular/forms';
import {ApiService} from '../api.service';

@Component({
  selector: 'app-setup',
  templateUrl: './setup.component.html',
  styleUrls: ['./setup.component.css']
})
export class SetupComponent {

  response: any;
  appIDControl = new FormControl('', [Validators.required]);
  pathControl = new FormControl('', [Validators.required]);

  constructor(
    private appService: ApiService,
  ) {
  }

  submit() {
    this.appService.setup({
      application_id: this.appIDControl.value,
      game_path: this.pathControl.value,
    }).subscribe(
      response => {
        this.response = response;
        if (this.response.ok) {
          window.location.reload();
        } else {
          if (!this.response.app_id_ok) {
            this.appIDControl.setErrors({});
          }
          if (!this.response.path_ok) {
            this.pathControl.setErrors({});
          }
        }
      }
    );
  }
}
