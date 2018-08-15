import { Component, ViewChild, Input, Directive } from '@angular/core';
import {FormControl} from '@angular/forms';
import { HttpClient, } from '@angular/common/http';

@Component({
  selector: 'app-page2',
  templateUrl: './page2.component.html',
  styleUrls: ['./page2.component.css'],
})

@Directive({selector: '[App-Filer]' })
export class Filers {
  @Input() file: File;
}

export class Page2Component {
  fileId: File;
  @ViewChild(Filers)
  set filer(v: Filers) {
   v.file = this.fileId ;
  }
  myControl = new FormControl();
  options: string[] = ['/Public/', '/Intake/', '/Public/Lol/rip/', '/Public/Lol'];
  sliderOn = false;
  file: string;
  upload() {
    console.log(this.fileId);
    return this.http.post('10.12.181.59:9993/upload/', this.fileId);
  }
  public constructor(private http: HttpClient) {
  }
}



