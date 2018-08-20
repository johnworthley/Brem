import React, { Component } from 'react'
import { css, StyleSheet } from 'aphrodite/no-important'

import store from 'Store'
import login from '../../util/login'

const style = StyleSheet.create({
  form: {
    display: 'flex',

    width: 220,
    flexDirection: 'column'
  },
  input: {
    borderRadius: 3,
    border: 'solid 1px lightgray',
    padding: '0 15px',
    marginBottom: 7,
    boxSizing: 'border-box',
    height: 40
  },
  button: {
    height: 40,
    boxSizing: 'border-box'
  }
})


class SignupTab extends Component {
  state = {

  }
  handleSubmit = async e => {
    e.preventDefault()
    const name = this.name.value
    store.update({ regName: name })
    await login()
    console.log('logged in')
  }
  render() {

    return (
      <div>
        <h3>Register</h3>
        <form className={css(style.form)} onSubmit={this.handleSubmit}>
          <input required ref={elem => this.name = elem} className={css(style.input)} type="text" placeholder="Username"/>
          <button className={css(style.button)}>Register me!</button>
        </form>
      </div>
    )
  }
}

export default SignupTab
