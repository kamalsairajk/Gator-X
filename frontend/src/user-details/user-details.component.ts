import { Component, OnInit } from '@angular/core';
import { Location } from '@angular/common';
import { UsersService } from '../services/users.service';
import { ActivatedRoute } from '@angular/router';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-user-details',
  templateUrl: './user-details.component.html',
  styleUrls: ['./user-details.component.scss']
})
export class UserDetailsComponent implements OnInit {
  user$: Observable<any> | undefined;
  user: any;
  userId: any;
  snapshot: any = null;
  constructor(
    private userService: UsersService,
    private route: ActivatedRoute,
    ) {}

  ngOnInit(): void {
    this.snapshot = this.route.snapshot;
    this.userId = this.snapshot.params.id;
    this.user$ = this.userService.getUserDetails({id: this.userId});
    this.user$.subscribe(data => {
      this.user = data[0];
    })
  }
}
