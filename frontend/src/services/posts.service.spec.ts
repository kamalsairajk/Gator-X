import { HttpClientTestingModule } from '@angular/common/http/testing';
import { TestBed } from '@angular/core/testing';

import { PostsService } from './posts.service';

describe('PostsService', () => {
  let service: PostsService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
    });
    service = TestBed.inject(PostsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  describe('getPosts', () => {
    it('should call get on http when called',() => {
      const spy = spyOn(service.http, 'get');
      service.getPosts({userId: 1});
      expect(spy).toHaveBeenCalled();
    })
  })

  describe('updatePost', () => {
    it('should call get on http when called',() => {
      const spy = spyOn(service.http, 'put');
      service.updatePost({
        id:1,
        title: 'foo',
        body: 'bar',
        userId: 1,
      }, 1);
      expect(spy).toHaveBeenCalled();
    })
  })

  describe('addPost', () => {
    it('should call get on http when called',() => {
      const spy = spyOn(service.http, 'post');
      service.addPost({
          title: 'foo',
          body: 'bar',
          userId: 1,
        } );
      expect(spy).toHaveBeenCalled();
    })
  })

  describe('deletePost', () => {
    it('should call get on http when called',() => {
      const spy = spyOn(service.http, 'delete');
      service.deletePost(1);
      expect(spy).toHaveBeenCalled();
    })
  })
});
