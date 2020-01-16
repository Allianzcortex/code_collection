import React from 'react';
import {IconStyle} from "./assets/iconfont/iconfont";
import {GlobalStyle} from "./style";
// 下面是路由配置
import {renderRoutes} from "react-router-config";
// TODO 看下面引用方式的命名问题,default 与 const
import routes from "./routes/index"
import {BrowserRouter} from "react-router-dom";

function App() {
  return(
      <BrowserRouter>

        <GlobalStyle></GlobalStyle>
        <IconStyle></IconStyle>
          {/*<i className="iconfont">&#xe62b;</i>*/}
          {renderRoutes(routes)}
      </BrowserRouter>
  )
}

export default App;
