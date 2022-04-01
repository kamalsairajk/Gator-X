import { HttpClientTestingModule } from '@angular/common/http/testing';
import { TestBed } from '@angular/core/testing';

import { UsersService } from './users.service';

describe('UsersService', () => {
  let service: UsersService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
    });
    service = TestBed.inject(UsersService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  describe('getUsers', () => {
    it('should call get on http when called',() => {
      const spy = spyOn(service.http, 'get');
      service.getUsers();
      expect(spy).toHaveBeenCalled();
    })
  })

  describe('getUserDetails', () => {
    it('should call get on http when called',() => {
      const spy = spyOn(service.http, 'get');
      service.getUserDetails({id: 1});
      expect(spy).toHaveBeenCalled();
    })
  })
});
