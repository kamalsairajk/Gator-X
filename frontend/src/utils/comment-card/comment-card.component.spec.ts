import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CommentCardComponent } from './comment-card.component';

describe('CommentCardComponent', () => {
  let component: CommentCardComponent;
  let fixture: ComponentFixture<CommentCardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CommentCardComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CommentCardComponent);
    component = fixture.componentInstance;
    component.comment = {email: 'abc', name: 'def', body: 'ghi'};
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
