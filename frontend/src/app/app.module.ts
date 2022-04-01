import { CUSTOM_ELEMENTS_SCHEMA, NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { UserDetailsComponent } from '../user-details/user-details.component';
import { UsersListComponent } from '../users-list/users-list.component';
import { HttpClientModule } from '@angular/common/http';
import { MaterialExampleModule } from 'src/material.module';
import { ProfileCardComponent } from 'src/utils/profile-card/profile-card.component';
import { UserAlbumsComponent } from 'src/user-albums/user-albums.component';
import { UserPostsComponent } from 'src/user-posts/user-posts.component';
import { UserPhotosComponent } from 'src/user-photos/user-photos.component';
import { FormsModule } from '@angular/forms';
import { FilterPipe } from 'src/pipes/filter.pipe';
import { PostCommentsComponent } from 'src/post-comments/post-comments.component';
import { CommentCardComponent } from 'src/utils/comment-card/comment-card.component';

@NgModule({
  declarations: [
    AppComponent,
    UserDetailsComponent,
    UsersListComponent,
    ProfileCardComponent,
    UserPostsComponent,
    UserAlbumsComponent,
    UserPhotosComponent,
    FilterPipe,
    PostCommentsComponent,
    CommentCardComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    HttpClientModule,
    MaterialExampleModule,
    FormsModule,
  ],
  schemas: [CUSTOM_ELEMENTS_SCHEMA],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
