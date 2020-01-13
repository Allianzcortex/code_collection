import React from 'react'
import PropTypes from 'prop-types'
import TodoTextInput from './TodoTextInput'

import {addTodo} from '../actions'
import {useDispatch} from 'react-redux'

const Header = () => {
  const dispatch = useDispatch()
  return (
  <header className="header">
    <h1>todos</h1>
    <TodoTextInput
      newTodo
      onSave={(text) => {
        if (text.length !== 0) {
          dispatch(addTodo(text))
        }
      }}
      placeholder="What needs to be done?"
    />
  </header>
)}

Header.propTypes = {
  addTodo: PropTypes.func
}

export default Header
