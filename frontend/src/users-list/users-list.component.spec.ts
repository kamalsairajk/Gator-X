import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { UsersListComponent } from './users-list.component';
import { MatDialog } from '@angular/material/dialog';

describe('UsersListComponent', () => {
  let component: UsersListComponent;
  let fixture: ComponentFixture<UsersListComponent>;
  let matDialogService: jasmine.SpyObj<MatDialog>;
    matDialogService = jasmine.createSpyObj<MatDialog>('MatDialog', ['open']);
  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ UsersListComponent ],
      imports: [
        RouterTestingModule,
        HttpClientTestingModule,
      ],
      providers: [
        {
          provide: MatDialog,
          useValue: matDialogService,
        },
      ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(UsersListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  describe('navigate', ()=> {
    it('should call navigate method on router when called',()=>{
      const spy = spyOn(component.router, 'navigate');
      component.navigate({id: 1});
      expect(spy).toHaveBeenCalled();
    })
  })

  describe('getToolTipInfo',() =>{
    it('should return tooltip format string when called',() => {
      expect(component.getToolTipInfo({username: 'abc', email: 'def'})).toEqual('Username:abc  Email:def');
    })
  })
});
