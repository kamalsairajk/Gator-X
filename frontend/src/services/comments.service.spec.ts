import { HttpClientTestingModule } from '@angular/common/http/testing';
import { TestBed } from '@angular/core/testing';

import { CommentsService } from './comments.service';

describe('CommentsService', () => {
  let service: CommentsService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
    });
    service = TestBed.inject(CommentsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  describe('getComments', () => {
    it('should call get on http when called',() => {
      const spy = spyOn(service.http, 'get');
      service.getComments({postId: 1});
      expect(spy).toHaveBeenCalled();
    })
  })
});
