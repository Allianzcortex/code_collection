import React from 'react'
import PropTypes from 'prop-types'
import TodoItem from './TodoItem'
import { useSelector,useDispatch } from 'react-redux'
import { SHOW_ALL, SHOW_COMPLETED, SHOW_ACTIVE } from '../constants/TodoFilters'
import {completeTodo,deleteTodo,editTodo} from '../actions'



const TodoList = ({}) => {
  const visibilityFilter = useSelector(state => state.visibilityFilter)
  const todos = useSelector(state => state.todos)
  // 这里关键是用 switch 来赋值：
  // https://www.codekitchen.ca/assigning-a-variable-using-a-switch-statement-in-javascript-2/
  let filteredTodos=(()=>{
  switch (visibilityFilter) {
    case SHOW_ALL:
      return todos
    case SHOW_COMPLETED:
        return todos.filter(t => t.completed)
    case SHOW_ACTIVE:
      return todos.filter(t => !t.completed)
  }})()

  console.log(`filter todo is ${filteredTodos}`)

  const dispatch=useDispatch();

  const actions={
    completeTodo:(id)=>dispatch(completeTodo(id)),
    deleteTodo:(id)=>dispatch(deleteTodo(id)),
    editTodo:(id,text)=>dispatch(editTodo(id,text))
  }


  return (

    <ul className="todo-list">
      {filteredTodos.map(todo =>
        <TodoItem key={todo.id} todo={todo} {...actions} />
      )}
    </ul>
  )
}

TodoList.propTypes = {
  filteredTodos: PropTypes.arrayOf(PropTypes.shape({
    id: PropTypes.number.isRequired,
    completed: PropTypes.bool.isRequired,
    text: PropTypes.string.isRequired
  }).isRequired),
  actions: PropTypes.object
}

export default TodoList
