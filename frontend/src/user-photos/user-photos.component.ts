import { Component, Input, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { AlbumsService } from 'src/services/albums.service';

@Component({
  selector: 'app-user-photos',
  templateUrl: './user-photos.component.html',
  styleUrls: ['./user-photos.component.scss']
})
export class UserPhotosComponent implements OnInit {
  @Input() albumId: any;
  photos$: Observable<any> | undefined;
  photos: any;
  constructor(private albumService: AlbumsService) { }

  ngOnInit(): void {
    this.photos$ = this.albumService.getPhotos({albumId: this.albumId});
    this.photos$.subscribe((data: any) => {
      this.photos = data;
    });
  }
}
