import { Component, OnInit } from '@angular/core';
import { FindUserService } from './find-user.service';
@Component({
  selector: 'ngx-find-user',
  templateUrl: './find-user.component.html',
  styleUrls: ['./find-user.component.scss']
})
export class FindUserComponent implements OnInit {
  user: any;
  data: any;
  alerts: any = []
  constructor(
    private service: FindUserService,
  ) { }

  ngOnInit(): void {
  }

  postData() {
    const data: any = new FormData();
    data.append('file1', this.SelectedFile);

    this.service.sendMlCall(data).subscribe(
      (response: any) => {
        if (response && response.match.length > 0) {
          console.log("ml", response);

          this.service.fetchMatchUsers(response.match.join()).subscribe(
            (res: any) => {
              if (res.code == 200) {
                console.log('success', res);
              } else {
                this.alerts.push(res.error);
              }
            }, (err) => {
              console.log("golang err", err);
            }
          );

        } else {
          this.alerts.push('no match');
        }
      },
      (err: any) => {
        this.alerts.push(err);
        console.log('Error ', err);
      }
    )
  }

  SelectedFile: File;

  onFileChanged(event) {
    this.SelectedFile = event.target.files[0];
    console.log("FILE CHANGE");
    console.log(event);
    console.log(this.SelectedFile.name);
  }

}