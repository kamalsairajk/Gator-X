import { Component, Input, OnInit } from '@angular/core';
import { UserData } from 'src/models/data.model';
import { UsersService } from 'src/services/users.service';

@Component({
  selector: 'app-profile-card',
  templateUrl: './profile-card.component.html',
  styleUrls: ['./profile-card.component.scss']
})
export class ProfileCardComponent implements OnInit {
  @Input() user: any;
  userData: UserData = {
    name: '',
    username: '',
    email: '',
    phone: '',
    id: undefined,
    website: '',
    address: '',
    company: ''
  };
  link = '';
  userDetails: any;
  constructor() {}

  ngOnInit(): void {
     this.userData.name = this.user.name;
     this.userData.username = this.user.username;
     this.userData.email = this.user.email;
     this.userData.phone = this.user.phone;
     this.userData.website = this.user.website;
     this.link = `https://www.${this.user.website}/`;
     this.userData.id = this.user.id;
     this.userData.address = `${this.user.address.suite}, ${this.user.address.street}, ${this.user.address.city}, ${this.user.address.zipcode}`;
     this.userData.company = this.user.company.name;
     this.userDetails = Object.entries(this.userData);
  }

}
