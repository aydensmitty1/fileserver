import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { AppComponent } from './app.component';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {MatButtonModule, MatIconModule} from '@angular/material';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatInputModule} from '@angular/material/input';
import {MatSlideToggleModule} from '@angular/material/slide-toggle';
import {MatTabsModule} from '@angular/material/tabs';
import { Page2Component } from './page2/page2.component';
import { Page3Component } from './page3/page3.component';
import { HomepageComponent } from './homepage/homepage.component';
import {MatAutocompleteModule} from '@angular/material/autocomplete';
import {MatSliderModule} from '@angular/material/slider';
import {MatBadgeModule} from '@angular/material/badge';
import { HttpClientModule } from '@angular/common/http';
import { FilekeeperComponent } from './filekeeper/filekeeper.component';
@NgModule({
  declarations: [
    AppComponent,
    Page2Component,
    Page3Component,
    HomepageComponent,
    FilekeeperComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,

    MatButtonModule,
    MatIconModule,
    MatToolbarModule,
    MatInputModule,
    MatSlideToggleModule,
    MatTabsModule,
    MatAutocompleteModule,
    MatSliderModule,
    MatBadgeModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
