import { Component, Input, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { NewPost } from 'src/models/data.model';
import { PostsService } from 'src/services/posts.service';

@Component({
  selector: 'app-user-posts',
  templateUrl: './user-posts.component.html',
  styleUrls: ['./user-posts.component.scss']
})
export class UserPostsComponent implements OnInit {

  @Input() userId: any;
  posts$: Observable<any> | undefined
  posts: any;
  filteredPosts: any;
  search: string = '';
  title: string = '';
  body: string = '';
  constructor(private postsService: PostsService) { }

  ngOnInit(): void {
    this.posts$ = this.postsService.getPosts({userId: this.userId});
    this.posts$.subscribe((data: any) => {
      this.posts = data;
      this.posts = this.posts.map((post: any) => {
        post.opened = false;
        post.edit = false;
        return post;
      });
    })
  }

  postComment(title: string, body: string) {
    const post: NewPost = {
      title: title,
      body: body,
      userId: this.userId,
    };
    this.postsService.addPost(post).subscribe((data: any) =>{
      if(data && data.id) this.ngOnInit();
    });
  }

  updatePost(post: any, flag: boolean, event: any) {
    if(!flag) {
      post.edit = true;
    }
    else {
      event.stopPropagation();
      post.edit = false;
      const body = {
        id: post.id,
        title: post.title,
        body: post.body,
        userId: post.userId,
      }
      this.postsService.updatePost(body, this.userId).subscribe((data: any)=> {
        if(data && data.id) this.ngOnInit();
      })
    }
  }

  deletePost(id: any, event: any) {
    event.stopPropagation();
    this.postsService.deletePost(id).subscribe((data: any) => {})
  }

}
