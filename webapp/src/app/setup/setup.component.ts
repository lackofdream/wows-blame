import { Component } from '@angular/core';
import { FormControl, Validators } from '@angular/forms';
import { WowsBlameSetupParam, WowsBlameSetupResponse } from '../models';
import { AppService } from '../app.service';

@Component({
  selector: 'app-setup',
  templateUrl: './setup.component.html',
  styleUrls: ['./setup.component.css']
})
export class SetupComponent {

  param: WowsBlameSetupParam = {
    application_id: '',
    game_path: '',
  };

  appIDControl = new FormControl('', [Validators.required]);
  pathControl = new FormControl('', [Validators.required]);

  response: WowsBlameSetupResponse;

  constructor(
    private appService: AppService,
  ) { }

  submit() {
    this.appService.setup(this.param).subscribe(
      response => {
        this.response = WowsBlameSetupResponse.createFrom(response.json());
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
