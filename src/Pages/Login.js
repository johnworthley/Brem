import React, { Component } from 'react'
import Tabs from '../Components/Tabs'

import LoginTab from './Tabs/LoginTab'
import SignupTab from './Tabs/SignupTab'

import store from 'Store'

class Login extends Component {
  state = {
    tabs: []
  }

  componentWillMount = async () => {
    store.watchState(this, ['needToShowSignup'])
    this.setState({
      tabs: [
        {
          name: 'Register',
          component: SignupTab,
          default: true
        }
      ]
    })
  }

  render() {
    const { tabs, needToShowSignup } = this.state
    return (
      <div>
        <Tabs data={tabs} />
      </div>
    )
  }
}

export default Login
