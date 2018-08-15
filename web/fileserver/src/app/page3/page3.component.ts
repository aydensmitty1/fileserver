import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-page3',
  templateUrl: './page3.component.html',
  styleUrls: ['./page3.component.css']
})
export class Page3Component implements OnInit {

  constructor() { }
  sliderOn = false;
  getColor() {
    return this.sliderOn === true ? 'black' : 'white';
  }
  getTextColor() {
    return this.sliderOn === true ? 'white' : 'black';
  }
  ngOnInit() {
  }

}
