import { Component, OnInit } from '@angular/core';
import { Router, NavigationExtras } from "@angular/router";

import { ListUserService } from './list-user.service';

@Component({
  selector: 'ngx-list-user',
  templateUrl: './list-user.component.html',
  styleUrls: ['./list-user.component.scss']
})
export class ListUserComponent implements OnInit {

  alerts: any = [];
  users: any = [];
  constructor(
    private listUsersService: ListUserService,
    private router: Router,
  ) { }

  ngOnInit(): void {
    this.getUsers();
  }

  getUsers() {
    this.alerts = [];
    this.listUsersService.getUsers().subscribe(
      (response: any) => {
        if (response.code == 200) {
          console.log(response.data);
          this.users = response.data;
        } else {
          console.log("Error: ", response.error);
          this.alerts.push(response.error)
        }
      },
      (err) => {
        console.log(err);
        this.alerts.push(err);
      }
    );
  }
  editUser(id) {
    // this.listUsersService.getUserById(id).subscribe(
    //   (response: any) => {
    //     if (response.code == 200) {
    //       console.log(response.data);
    let navigationExtras: NavigationExtras = {
      queryParams: { image_uid: id },
    };
    this.router.navigate(["/pages/forms", "add-user"], navigationExtras);
    //     } else {
    //       console.log("Error: ", response.error);
    //       this.alerts.push(response.error);
    //     }
    //   },
    //   err => {
    //     console.error(err);
    //     this.alerts.push(err);
    //   }
    // );
  }

  deleteUser(id) {
    if (!confirm("Are you sure ?")) {
      return;
    }
    this.listUsersService.delete(id).subscribe(
      (response: any) => {
        if (response.code == 200) {
          location.reload();
        } else {
          console.log("Error: ", response.error);
          this.alerts.push(response.error)
        }
      },
      (err) => {
        console.log(err);
        this.alerts.push(err);
      }
    )
  }
}
