import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Constants as constants } from '../constants/constants'
import { GlobalFunctions as globals } from '../constants/globals';
import { Id } from 'src/models/data.model';

@Injectable({
  providedIn: 'root'
})
export class UsersService {
  users: any;
  user: any;

  constructor(public http: HttpClient) { }

  getUsers() {
    const url = `${constants.baseURL}${constants.users}`;
    return this.http.get(url);
  }

  getUserDetails(filter: Id) {
    const params = globals.generateParams(filter);
    const url = `${constants.baseURL}${constants.users}?${params}`;
    return this.http.get(url)
  }
}
