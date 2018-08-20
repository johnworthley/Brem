import React, { Component } from 'react'
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'
import { css, StyleSheet } from 'aphrodite/no-important'
import store from 'Store'
import config from 'Config'

import Login from '../Pages/Login'
import Cabinet from '../Pages/Cabinet'
import ProjectCreate from '../Pages/Project/Create'
import ProjectView from '../Pages/Project/View'
import ProjectList from '../Pages/Project/List'
import Header from '../Components/Header'
import Footer from '../Components/Footer'
import Alert from '../Components/Alert'
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
    const { logged, instance } = web3Status
    console.log('routing login:', logged, instance)
    return (
      <Router>
        <Route render={({ location }) => (
          <div>
            <Header />
            <div className={css(style.main)}>
              {
                logged === 'pending' ? (
                  <Page>
                    <h3>Loading.</h3>
                    <p>Check your MetaMask plugin</p>
                  </Page>

                ) : (
                  logged ? (
                    <Page>
                      <Switch location={location}>
                        <Route exact path="/" component={ProjectList} />
                        <Route exact path="/project/list" component={ProjectList} />
                        <Route exact path="/project/new" component={ProjectCreate} />
                        <Route exact path="/cabinet" component={Cabinet} />
                        <Route path="/project/:id" component={ProjectView} />
                      </Switch>
                    </Page>
                  ) : (
                    <div>
                      <Login />
                      {
                        instance ? '' : (
                          <Alert heading="Plugin error" text="No MetaMask session or plugin presented. Please use latest version of Chrome with MetaMask plugin installed with an active MetaMask session." />
                        )
                      }
                    </div>
                  )
                )
              }


            </div>
            <Footer />
          </div>
        )} />
      </Router>
    )
  }
}
export default Routing
