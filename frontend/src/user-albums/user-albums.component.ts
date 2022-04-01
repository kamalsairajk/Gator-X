import { Component, Input, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { AlbumsService } from 'src/services/albums.service';

@Component({
  selector: 'app-user-albums',
  templateUrl: './user-albums.component.html',
  styleUrls: ['./user-albums.component.scss']
})
export class UserAlbumsComponent implements OnInit {
  @Input() userId: any;
  albums: any;
  albums$: Observable<any> | undefined;
  albumIds: any;
  constructor(private albumsService: AlbumsService) { }

  ngOnInit(): void {
    this.albums$ = this.albumsService.getAlbumIds({userId: this.userId});
    this.albums$.subscribe(data =>{
      this.albums = data;
      this.albumIds = this.albums.map((album:any) => album.id);
    });
  }
}
