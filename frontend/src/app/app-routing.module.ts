import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { UserDetailsComponent } from 'src/user-details/user-details.component';
import { UsersListComponent } from 'src/users-list/users-list.component';

const routes: Routes = [
  { path:'users/:id/details', component: UserDetailsComponent },
  { path: 'users', component: UsersListComponent },
  { path:'', redirectTo: '/users', pathMatch: 'full' },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
