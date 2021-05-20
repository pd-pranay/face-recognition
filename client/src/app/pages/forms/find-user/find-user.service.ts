import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpBackend } from '@angular/common/http';

import { environment } from '../../../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class FindUserService {

  private httpClient: HttpClient;

  constructor(handler: HttpBackend) {
    this.httpClient = new HttpClient(handler);
  }

  url = environment.hostUrl;

  sendMlCall(data) {
    return this.httpClient.post(environment.mlUrl + 'face_match', data);
  }

  fetchMatchUsers(ids) {
    return this.httpClient.get(this.url + 'users/ml/' + ids);
  }

}
