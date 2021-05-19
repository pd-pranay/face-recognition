import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpBackend } from '@angular/common/http';

import { environment } from '../../../../environments/environment';

@Injectable({
  providedIn: 'root'
})

export class ListUserService {

  private httpClient: HttpClient;

  constructor(handler: HttpBackend) {
    this.httpClient = new HttpClient(handler);
  }

  url = environment.hostUrl;

  getUsers() {
    return this.httpClient.get(this.url + 'users')
  }

  getUserById(id) {
    return this.httpClient.get(this.url + 'users/read/' + id);
  }

  delete(id) {
    return this.httpClient.delete(this.url + 'users/' + id);
  }

}
