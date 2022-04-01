import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { UsersService } from '../services/users.service';
import { MatDialog } from '@angular/material/dialog';


@Component({
  selector: 'app-users-list',
  templateUrl: './users-list.component.html',
  styleUrls: ['./users-list.component.scss'],
  encapsulation: ViewEncapsulation.None,
})
export class UsersListComponent implements OnInit {
  users: any = null
  constructor(
    public router: Router,
    private route: ActivatedRoute,
    private userService: UsersService,
    public dialog: MatDialog,
    ) {}

  ngOnInit(): void {
    if(!this.userService.users) {
      this.userService.getUsers().subscribe((data: any) => {
        this.users = data;
        this.userService.users = this.users;
      })
    } else this.users = this.userService.users;
  }

  navigate(user: any) {
    this.userService.user = user;
    this.router.navigate([user.id, 'details'], {relativeTo: this.route});
  }

  getToolTipInfo(user: any) {
    const tooltipInfo = `Username:${user.username}  Email:${user.email}`;
    return tooltipInfo;
  }
}
