import React, { Component } from 'react'
import Routing from './Route/Routing'
import Header from './Components/Header'
import Footer from './Components/Footer'

import store from 'Store'


import login from './util/login'

import './App.css'



class App extends Component {
  render = () => (
    <div>
      <Header />
      <Routing />
      <Footer />
    </div>
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
