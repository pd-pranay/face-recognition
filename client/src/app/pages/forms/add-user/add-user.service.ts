import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpBackend } from '@angular/common/http';

import { environment } from '../../../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class AddUserService {

  private httpClient: HttpClient;

  constructor(handler: HttpBackend) {
    this.httpClient = new HttpClient(handler);
  }

  url = environment.hostUrl;

  postData(data) {
    return this.httpClient.post(this.url + 'users', data);
  }

  putData(id, data) {
    return this.httpClient.put(this.url + 'users/' + id, data);
  }
  getUserById(id) {
    return this.httpClient.get(this.url + 'users/read/' + id);
  }
}
