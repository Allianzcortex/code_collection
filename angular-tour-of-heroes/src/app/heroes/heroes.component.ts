import { Component, OnInit } from '@angular/core';
// 刚刚新建了一个 Hero 类，现在要导入
import { Hero } from '../hero';
// 导入所创建的 fake 记录
import {HEROES} from '../mock-heroes';

/**
 * @Component 是一个装饰器，会为 export 出的 class 指定所需的元数据
 * selector 是组件的选择器
 * templateUrl 模板文件的位置
 * styleUrls 组件 CSS 样式的位置
 * ngOnInit 放置生命周期
 * 始终要 export 这个 class ，从而方便在其他地方导入
 */
@Component({
  selector: 'app-heroes',
  templateUrl: './heroes.component.html',
  styleUrls: ['./heroes.component.css']
})
export class HeroesComponent implements OnInit {

  heroes = HEROES;

  // 接下来因为当用户选择一个 hero 时

  hero: Hero = {
    id: 1,
    name: 'windstorm'
  };

  constructor() { }

  ngOnInit() {
  }

}
