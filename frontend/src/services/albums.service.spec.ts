import { HttpClientTestingModule } from '@angular/common/http/testing';
import { TestBed } from '@angular/core/testing';

import { AlbumsService } from './albums.service';

describe('AlbumsService', () => {
  let service: AlbumsService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
    });
    service = TestBed.inject(AlbumsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  describe('getAlbumIds', () => {
    it('should call get on http when called',() => {
      const spy = spyOn(service.http, 'get');
      service.getAlbumIds({userId: 1});
      expect(spy).toHaveBeenCalled();
    })
  })

  describe('getPhotos', () => {
    it('should call get on http when called',() => {
      const spy = spyOn(service.http, 'get');
      service.getPhotos({albumId: 1});
      expect(spy).toHaveBeenCalled();
    })
  })
});
