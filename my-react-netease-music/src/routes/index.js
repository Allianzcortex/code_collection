import React from "react";
import {Redirect} from "react-router-dom"
import Home from "../application/Home"
import Recommend from '../application/Recommend';
import Singers from '../application/Singers';
import Rank from '../application/Rank';


export default [
    {
        // 这里的 component 在数组第一层，所以 Home 可以渲染
        // 但是其他的 component 无法渲染
        path:"/",
        component:Home,
        routes:[
            //TODO 下面这四个组件位置显然是平行的
            //我想知道为什么对 "/" 要用 Redirect 去渲染
            {
                path:"/recommend",
                component:Recommend
            },
            {
                path:"/singers",
                exact:true,
                component:Singers
            },
            {
                path:"/rank",
                component:Rank
            }
        ]
    }
]
