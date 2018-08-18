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
          name: 'Login',
          component: LoginTab,
          default: true
        },
        {
          name: 'Register',
          component: SignupTab
        }
      ]
    })
  }

  render() {
    const { tabs, needToShowSignup } = this.state
    return (
      <div>
        <Tabs data={needToShowSignup ? tabs.map(tab => tab.name === 'Register' ? {...tab, default: true} : {...tab, default: false}) : tabs} />
      </div>
    )
  }
}

export default Login
