import { TestBed } from '@angular/core/testing';

import { CrudService } from './crud.service';
import {HttpClientTestingModule} from '@angular/common/http/testing';

describe('CrudService', () => {
  let service: CrudService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
    });
    service = TestBed.inject(CrudService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
