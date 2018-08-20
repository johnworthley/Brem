import React, { Component } from 'react'
import Routing from './Route/Routing'

import { BrowserRouter as Router } from 'react-router-dom'

import store from 'Store'


import login from './util/login'

import './App.css'



class App extends Component {
  render = () => (
    <Routing />
  )
}

export default App;


//Testing no-internet fumction
(() => {
  store.update({
    web3Status: {
      logged: true
    }
  })
})

login()
