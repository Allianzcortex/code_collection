import styled from "styled-components";
import style from "../../assets/global-style"

export const Top=styled.div`
display: flex;
flex-direction: row;
justify-content: space-between;
background:${style["theme-color"]};
//border:.1px solid;
padding:.5px 10px;
// TODO &>span 是类似 SASS 的用法，自己需要看
&>span {
color:white;
line-height: 18px;
font-size:10px;
&.iconfont {
font-size: 10px;
}
}

`;

export const Tab=styled.div`
display: flex;
// TODO 这里 flex:1 的作用点是在于什么
//flex: 1;
flex-direction: row;
justify-content: space-around;
background:${style["theme-color"]};
font-size:8px;
//border:.5px solid;
//height:20px;
line-height: 10px;
padding-bottom: 4px;
//a:link{color:#fff}  未访问时的状态（鼠标点击前显示的状态）
//  2、a:hover{color:#fff}  鼠标悬停时的状态
//  3、a:visited{color:#fff}  已访问过的状态（鼠标点击后的状态）
//  4、a:active{color:#fff}  鼠标点击时的状态
//  5、a:focus{color:#fff}  点击后鼠标移开保持鼠标点击时的状态（只有在<a href="#"></a>时标签中有效）
// 所以要实现点击后一直是某种状态只能在 a 标签里选择
// TODO 还是需要比较一下和原版实现的不同
a:focus{
    font-weight: 700;
    border-bottom: 1px solid white;
}

// 我觉得用这种方式来加高亮下划线总不是很对
// &.selected {
//       span {
`;

export const TabItem = styled.div`
height:100%;
display: flex;
//border:.5px solid;
justify-content: center;
align-items: center;
//span {
//    &:focus {
//      font-weight: 700;
//    border-bottom: 1px solid white;
//}
}
`;
