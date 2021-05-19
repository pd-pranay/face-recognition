import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';

import { AddUserService } from './add-user.service';
@Component({
  selector: 'ngx-add-user',
  templateUrl: './add-user.component.html',
  styleUrls: ['./add-user.component.scss']
})
export class AddUserComponent implements OnInit {

  user: any = {};
  alerts: any = [];

  constructor(
    private router: Router,
    private activatedRoute: ActivatedRoute,
    private addUserService: AddUserService,
  ) {

  }

  ngOnInit(): void {
    this.getSingleData();
  }

  getSingleData() {

    this.activatedRoute.queryParams.subscribe(
      (params: any) => {
        if (params["image_uid"] != undefined) {
          console.log(params);
          this.user.image_uid = params['image_uid'];
          this.user.name = params['name'].String;
          this.user.college_name = params['college_name'].String;
          this.user.address = params['address'].String;
          this.user.mobile_no = params['mobile_no'].String;
        }
      });

  }

  postData() {
    this.alerts = [];

    const data: any = new FormData();
    data.append('name', this.user.name);
    data.append('college_name', this.user.college_name);
    data.append('address', this.user.address);
    data.append('mobile_no', this.user.mobile_no);
    data.append('file', this.SelectedFile);

    // for (var v of data) {
    //   console.log(v);
    // }
    this.addUserService.postData(data).subscribe(
      (response: any) => {
        if (response.code == 200) {
          console.log('User Added');
        } else {
          this.alerts.push(response.error);
        }
      },
      (err) => {
        this.alerts.push(err);
        console.log('Error ', err);
      }
    )

  }

  putData() {

  }

  SelectedFile: File;

  onFileChanged(event) {
    this.SelectedFile = event.target.files[0];
    console.log("FILE CHANGE");
    console.log(event);
    console.log(this.SelectedFile.name);
  }

}
