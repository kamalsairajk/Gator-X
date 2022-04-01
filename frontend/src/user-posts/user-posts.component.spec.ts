import { ComponentFixture, TestBed } from '@angular/core/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { UserPostsComponent } from './user-posts.component';
import { FilterPipe } from '../pipes/filter.pipe'

xdescribe('UserPostsComponent', () => {
  let component: UserPostsComponent;
  let fixture: ComponentFixture<UserPostsComponent>;
  let pipe: FilterPipe
  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ UserPostsComponent ],
      imports: [HttpClientTestingModule],
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(UserPostsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();

  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
