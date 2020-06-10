import { TestBed } from '@angular/core/testing';

import { UserService } from './user.service';
import {HttpClientTestingModule} from '@angular/common/http/testing';
import {UserRegister} from '../models/user-register';
import {User} from '../models/user';

describe('UserService', () => {
  let service: UserService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
    });
    service = TestBed.inject(UserService);
    spyOn(service, 'getUser');
    spyOn(service, 'register');
    spyOn(service, 'login');
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should call getUser()', () => {
    service.getUser();
    expect(service.getUser).toHaveBeenCalled();
  });

  it('should call register()', () => {
    service.register(new UserRegister('test', 'a@b.c'));
    expect(service.register).toHaveBeenCalledWith(new UserRegister('test', 'a@b.c'));
  });

  it('should call login()', () => {
    service.login(new User('test', 'asdf'));
    expect(service.login).toHaveBeenCalledWith(new User('test', 'asdf'));
  });

});
