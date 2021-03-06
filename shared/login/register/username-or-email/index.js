// @flow
import React, {Component} from 'react'
import {connect} from 'react-redux'
import Render from './index.render'
import {cancelLogin} from '../../../actions/login'
import type {Props} from './index.render'

class UsernameOrEmail extends Component<void, Props, void> {
  render () {
    return <Render {...this.props} />
  }
}

export default connect(
  state => ({}),
  dispatch => ({
    onBack: () => dispatch(cancelLogin()),
    onSubmit: usernameOrEmail => console.log('TODO submit username: ', usernameOrEmail) // dispatch(submitUsername(UsernameOrEmail))
  })
)(UsernameOrEmail)
