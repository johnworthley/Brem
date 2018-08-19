import React, { Component } from 'react'
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'
import { css, StyleSheet } from 'aphrodite/no-important'
import store from 'Store'
import config from 'Config'

import Login from '../Pages/Login'
import Cabinet from '../Pages/Cabinet'
import ProjectCreate from '../Pages/Project/Create'


import Page from '../Components/Page'



const style = StyleSheet.create({
  main: {
    minHeight: 'calc(100vh - 290px)',
    maxWidth: config.pageWidth,
    boxSizing: 'border-box',
    padding: '0 20px',
    margin: '0 auto'
  }
})

class Routing extends Component {
  state = {
    web3Status: {
      logged: false
    }
  }

  componentWillMount = () => store.watchState(this, ['loginStatus', 'web3Status'])

  render = () => {
    const { loginStatus, web3Status } = this.state
    console.log(loginStatus, web3Status)
    return (
      <Router>
        <Route render={({ location }) => (
          <div className={css(style.main)}>
            {/* <Login /> */}
            <Page>
              <Route exact path="/project/new" component={ProjectCreate} />
				      <Route exact path="/cabinet" component={Cabinet} />
            </Page>
          </div>
        )} />
      </Router>
    )
  }
}
export default Routing
