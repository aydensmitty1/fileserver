import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  sliderOn = false;

  slider() {
    if (this.sliderOn === true) {
      this.sliderOn = false;
    } else {
      this.sliderOn = true;
    }
  }

  getColor() {
    return this.sliderOn === true ? 'black' : 'white';
  }

  getTextColor() {
    return this.sliderOn === true ? 'white' : 'black';
  }
}


