import { Component, OnInit, Input } from '@angular/core';
import { Article } from './article.model';

@Component({
  selector: 'app-article',
  templateUrl: './article.component.html',
  styleUrls: ['./article.component.css'],
  // 之后我再用 hostbinding 来取代好不好...
  host: {
    class: 'row'
  }
})
export class ArticleComponent implements OnInit {
  // 已经可以用一个 Article Model 来表示下面三个 field
  // votes: number;
  // title: string;
  // link: string;
  // 接下来加入 @Input()，不必一定要从构造函数中建立
  @Input() article: Article;


  constructor() {
    // this.article = new Article(
    //   'Angular 2',
    //   'http://angular.io',
    //   10
    // );

  }

  // 这里用返回 boolean false 而不是 void() 函数
  // 如果是空的话点击 upvote 或者 downvote 会导向空白页
  voteUp(): boolean {
    this.article.voteUp();
    return false;
  }

  voteDown(): boolean {
    this.article.voteDown();
    return false;
  }

  ngOnInit() {
  }

}
