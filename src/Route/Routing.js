import React, { Component } from 'react'
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'
import { css, StyleSheet } from 'aphrodite/no-important'
import store from 'Store'

import Cabinet from '../Pages/Cabinet'
import Page from '../Components/Page'



const style = StyleSheet.create({
  main: {
    minHeight: 'calc(100vh - 290px)'
  }
})

class Routing extends Component {
  componentWillMount = () => store.watchState(this, ['loginStatus'])
  state = {

  }

  render = () => {

    return (
      <Router>
        <Route render={({ location }) => (
          <div className={css(style.main)}>
            <Page>
				<Route exact path="/cabinet" component={Cabinet} />
            </Page>
          </div>
        )} />
      </Router>
    )
  }
}
export default Routing
