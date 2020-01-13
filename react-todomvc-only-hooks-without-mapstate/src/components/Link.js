import React from 'react'
import PropTypes from 'prop-types'
import classnames from 'classnames'
import { setVisibilityFilter } from '../actions'
import { useSelector,useDispatch} from 'react-redux'

const Link = ({ filter,children }) =>{

  const active = useSelector(state=>state.visibilityFilter===filter)
  const dispatch=useDispatch();
  // const mapDispatchToProps = (dispatch, ownProps) => ({
  //   setFilter: () => {
  //     dispatch(setVisibilityFilter(ownProps.filter))
  //   }
  // })

  return (
    // eslint-disable-next-line jsx-a11y/anchor-is-valid
    <a
      className={classnames({ selected: active })}
      style={{ cursor: 'pointer' }}
      // onClick={() => {setFilter()}}
      onClick={() =>{dispatch(setVisibilityFilter(filter))}}
    >
      {children}
    </a>
  )}


Link.propTypes = {
  active: PropTypes.bool,
  children: PropTypes.node.isRequired,
  setFilter: PropTypes.func
}

export default Link
