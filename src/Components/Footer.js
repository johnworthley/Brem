import React, { Component } from 'react'
import { css, StyleSheet } from 'aphrodite/no-important'
import config from 'Config'

const footerHeight = 170

const style = StyleSheet.create({
  main: {
    background: '#262b2e',
    color: '#fff',
    height: footerHeight,
  },
  holder: {
    width: '100%',
    height: footerHeight,
    maxWidth: config.pageWidth,
    margin: '0 auto',
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
  }
})



class Footer extends Component {
  state = {

  }

  render = () => {

    return (
      <footer className={css(style.main)}>
        <div className={css(style.holder)}>
          <div>
            Мы в социальных сетях
          </div>
          <div>
            "Brem" (c) 2000-2018
          </div>
        </div>
      </footer>
    )
  }
}
export default Footer
