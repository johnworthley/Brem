import React, { Component } from 'react'
import { css, StyleSheet } from 'aphrodite/no-important'

const style = StyleSheet.create({
  main: {
    position: 'fixed',
    left: 0,
    top: 0,
    right: 0,
    bottom: 0,
    display: 'flex',
    zIndex: 50
  },
  black: {
    position: 'fixed',
    left: 0,
    top: 0,
    right: 0,
    bottom: 0,
    zIndex: 1,
    background: 'rgba(0,0,0,.2)'
  },
  alert: {
    display: 'flex',
    flexDirection: 'column',
    background: '#fff',
    margin: 'auto',
    width: 500,
    zIndex: 2,
    boxSizing: 'border-box',
    padding: 20,
    border: 'solid 1px lightgray',
    borderRadius: 1,
    boxShadow: '5px 5px 10px 0 rgba(0,0,0,.3)'
  }
})

export default ({ buttonText = 'Got it!', heading = 'Heading', text = 'Text', callback, buttons = {} }) => (
  <div className={css(style.main)}>
    <div className={css(style.black)} />
    <div className={css(style.alert)}>
      <h3 style={{
        margin: 0,
        fontSize: '25px',
        color: '#050000'
      }}>
      { heading }
      </h3>
      <p style={{marginBottom: 15}}>
        { text }
      </p>
      {
        callback ? (
          <button style={{alignSelf: 'flex-end'}} onClick={callback}>{ buttonText }</button>
        ) : ''
      }
    </div>
  </div>
)
