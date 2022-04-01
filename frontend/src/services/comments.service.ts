import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Constants as constants } from 'src/constants/constants';
import { GlobalFunctions as globals } from 'src/constants/globals';
import { PostId } from 'src/models/data.model';

@Injectable({
  providedIn: 'root'
})
export class CommentsService {

  constructor(public http: HttpClient) { }

  getComments(filter: PostId) {
    const params = globals.generateParams(filter);
    const url = `${constants.baseURL}${constants.comments}?${params}`
    return this.http.get(url);
  }
}
