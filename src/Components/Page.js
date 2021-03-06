import React, { Component } from 'react'
import config from 'config'
import { css, StyleSheet } from 'aphrodite/no-important'

const style = StyleSheet.create({
  main: {
    width: '100%',
	  boxSizing: 'border-box'
  }
})

class Page extends Component {
  render = () => {

    return (
      <div className={css(style.main)}>
        {
          this.props.children
        }
      </div>
    )
  }
}
export default Page
