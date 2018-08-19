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
    const tabsToRender = tabs.map(tab => {
      if(needToShowSignup && tab.name === 'Register') return {...tab, default: true}
      else if(needToShowSignup) return {...tab, default: false}
      else if(!needToShowSignup && tab.name === 'Login') return {...tab, default: true}
      else return {...tab, default: false}
    })
    return (
      <div>
        <Tabs data={tabsToRender} />
      </div>
    )
  }
}

export default Login
