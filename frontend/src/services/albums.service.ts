import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { AlbumId, UserId } from 'src/models/data.model';
import { Constants as constants } from '../constants/constants';
import { GlobalFunctions as globals } from '../constants/globals';

@Injectable({
  providedIn: 'root'
})
export class AlbumsService {

  constructor(public http: HttpClient) { }

  getAlbumIds(filter: UserId) {
    const params = globals.generateParams(filter);
    const url = `${constants.baseURL}${constants.albums}?${params}`
    return this.http.get(url);
  }

  getPhotos(filter: AlbumId) {
    const params = globals.generateParams(filter);
    const url = `${constants.baseURL}${constants.photos}?${params}`
    return this.http.get(url);
  }
}
