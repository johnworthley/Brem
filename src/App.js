import React, { Component } from 'react'
import Routing from './Route/Routing'
import Header from './Components/Header'
import Footer from './Components/Footer'
import store from 'Store'

import getweb3 from './util/getweb3'

import './App.css'

getweb3()

console.log(store)


class App extends Component {
  render = () => (
    <div>
      <Header />
      <Routing />
      <Footer />
    </div>
  )
}

export default App
