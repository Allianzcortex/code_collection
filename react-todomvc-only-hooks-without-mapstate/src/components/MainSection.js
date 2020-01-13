import React from 'react'
import PropTypes from 'prop-types'
import Footer from './Footer'
import TodoList from '../components/TodoList'
import {useSelector,useDispatch} from 'react-redux'
import { CLEAR_COMPLETED, COMPLETE_ALL_TODOS } from '../constants/ActionTypes'

const MainSection = ({actions }) =>{
  const todosCount=useSelector(state=>state.todos.length);

  const completedCount=useSelector(state=>state.todos.reduce((count, todo) =>
  todo.completed ? count + 1 : count,
  0)
)

const dispatch = useDispatch();
// 还是直接 dispatch 好，不要先定义再 dispatch...
  // const clearCompleted=dispatch({type:'CLEAR_COMPLETED'})
  // const completeAllTodos=dispatch({type:COMPLETE_ALL_TODOS})
  


  return (
    <section className="main">
      {
        !!todosCount && 
        <span>
          <input
            className="toggle-all"
            type="checkbox"
            checked={completedCount === todosCount}
            readOnly
          />
          {/* <label onClick={actions.completeAllTodos}/> */}
          <label onClick={()=>{dispatch({type:COMPLETE_ALL_TODOS})}}></label>
          
        </span>
      }
      <TodoList />
      {
        !!todosCount &&
        <Footer
          completedCount={completedCount}
          activeCount={todosCount - completedCount}
          onClearCompleted={()=>{dispatch({type:CLEAR_COMPLETED})}}
        />
      }
    </section>
  )
    }

MainSection.propTypes = {
  todosCount: PropTypes.number,
  completedCount: PropTypes.number,
  actions: PropTypes.object
}

export default MainSection;