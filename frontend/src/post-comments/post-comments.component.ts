import { Component, Input, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { CommentsService } from 'src/services/comments.service';

@Component({
  selector: 'app-post-comments',
  templateUrl: './post-comments.component.html',
  styleUrls: ['./post-comments.component.scss']
})
export class PostCommentsComponent implements OnInit {
  @Input() postId: any;
  comments$: Observable<any> | undefined;
  comments: any;
  constructor(private commentsService: CommentsService) { }

  ngOnInit(): void {
    this.comments$ = this.commentsService.getComments({postId: this.postId});
    this.comments$.subscribe((data: any) => {
      this.comments = data;
      console.log(this.comments);
    })
  }

}
