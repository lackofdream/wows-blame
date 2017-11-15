import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpModule } from '@angular/http';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';
import { MatCardModule, MatInputModule, MatButtonModule, NoConflictStyleCompatibilityMode } from '@angular/material';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { AppComponent } from './app.component';
import { AppService } from './app.service';
import { SetupComponent } from './setup/setup.component';
import { BlameComponent } from './blame/blame.component';

@NgModule({
  declarations: [
    AppComponent,
    SetupComponent,
    BlameComponent
  ],
  imports: [
    BrowserModule,
    HttpModule,
    BrowserAnimationsModule,
    FormsModule,
    ReactiveFormsModule,
    MatCardModule,
    MatInputModule,
    MatButtonModule,
    NoConflictStyleCompatibilityMode,
  ],
  providers: [AppService],
  bootstrap: [AppComponent]
})
export class AppModule { }
