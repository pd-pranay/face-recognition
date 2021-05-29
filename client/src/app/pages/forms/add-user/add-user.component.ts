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
  pic = false;
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
          this.pic = true;
          console.log(params['image_uid']);
          this.addUserService.getUserById(params['image_uid']).subscribe(
            (response: any) => {
              if (response.code == 200) {
                console.log('User Fetched');
                this.user.id = response.data['id'];
                this.user.image_uid = response.data['image_uid'];
                this.user.image_path = response.data['image_path'];
                console.log(this.user.image_path, "=====");
                this.user.name = response.data['name'].String;
                this.user.college_name = response.data['college_name'].String;
                this.user.address = response.data['address'].String;
                this.user.mobile_no = response.data['mobile_no'].Int32;
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
          this.router.navigateByUrl('/pages/forms/list-user');
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
    this.alerts = [];

    const data: any = new FormData();
    data.append('name', this.user.name);
    data.append('college_name', this.user.college_name);
    data.append('address', this.user.address);
    data.append('mobile_no', this.user.mobile_no);
    console.log('===> ', this.SelectedFile);
    if (this.SelectedFile == undefined) { // uid no image update
      data.append('is_change', "no");
      data.append('image_uid', this.user.image_uid);
      data.append('image_path', this.user.image_path);
      console.log('image not updated');
    } else {
      console.log('image updated');
      data.append('is_change', "yes");
      data.append('file', this.SelectedFile);

    }

    this.addUserService.putData(this.user.image_uid, data).subscribe(
      (response: any) => {
        if (response.code == 200) {
          console.log('User Updated');
          this.router.navigateByUrl('/pages/forms/list-user');
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

  submit() {
    console.log('id ===> ', this.user.id);
    if (this.user.id) {
      this.putData()
      console.log('update block');
    } else {
      console.log('post block');

      this.postData()
    }
  }

  SelectedFile: File;

  onFileChanged(event) {
    this.SelectedFile = event.target.files[0];
    console.log("FILE CHANGE");
    console.log(event);
    console.log(this.SelectedFile.name);
  }

}
