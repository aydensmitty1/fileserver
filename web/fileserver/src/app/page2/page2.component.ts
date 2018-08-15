import {Component, ViewChild} from '@angular/core';
import {FormControl} from '@angular/forms';
import {HttpClient, HttpHeaders} from '@angular/common/http';

@Component({
  selector: 'app-page2',
  templateUrl: './page2.component.html',
  styleUrls: ['./page2.component.css']
})

// Directive({selector: '[filer]' })
// export class Filers {
 // @Input() file: File;
// }

export class Page2Component {

  @ViewChild('datfile')
  filers: File;
  // @ViewChild(Filers)
  // set filer(v: Filers) {
  //   v.file = this.filers;
  // }
  myControl = new FormControl();
  options: string[] = ['/Public/', '/Intake/', '/Public/Lol/rip/', '/Public/Lol'];
  sliderOn = false;
  file: string;
  fileToUpload: File = null;

  public constructor(private http: HttpClient) {
  }

  upload() {
    console.log(this.fileToUpload);
    const formdata = new FormData();
    formdata.append('uploadfile', this.fileToUpload);

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type':  'multipart/form-data'
      })
    };
    this.http.post('/api/upload/', formdata)
      .subscribe(res => {
        console.log(res);
      }, err => {
        console.log(err);
      });
  }

  handleFileInput(files: FileList) {
    this.fileToUpload = files[0];
  }
}



