import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-homepage',
  templateUrl: './homepage.component.html',
  styleUrls: ['./homepage.component.css']
})
export class HomepageComponent implements OnInit {

  constructor() { }
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
  ngOnInit() {
  }

}
