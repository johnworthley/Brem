import React, { Component } from 'react'
import { css, StyleSheet } from 'aphrodite/no-important'
import config from 'Config'
import store from 'Store'

const style = StyleSheet.create({
  holder: {
    maxWidth: config.pageWidth,
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'space-between',
    width: '100%',
    margin: '0 auto'
  },
  main: {
    height: 70,
    display: 'flex',
    backgroundColor: '#fff',
    color: '#1e1e1e',
    marginBottom: 30,
    padding: '0 20px'
  },
  left: {
    marginLeft: 10,
    display: 'flex',
    alignItems: 'center',
  },
  right: {

  },
  button: {
    height: 50,
    borderRadius: '3px',
    backgroundColor: '#f0723b',
    padding: '0 40px',
    border: 'none',
    color: '#fff'
  }
})

const leftStyle = StyleSheet.create({
  part: {
    marginLeft: 30,
    textTransform: 'uppercase',
    ':first-child': {
      marginLeft: 0
    }
  },
  icon: {

  },
  orange: {
    color: '#f0723b'
  }
})

const LeftPart = ({ value, name, delim, additional, orange }) => (
  <div className={css(leftStyle.part)}>
    <div className={css(leftStyle.icon)} />
    <span>
      <span style={{display: 'flex'}}>
        <span className={orange === 'left' || orange === 'all' ? css(leftStyle.orange) : ''}>
          { name }{delim ? ':' : ''}
        </span>

        <span style={{marginLeft: 6, display: 'inline-block', maxWidth: 100, overflow: 'hidden', textOverflow: 'ellipsis'}} className={orange === 'right' || orange === 'all' ? css(leftStyle.orange) : ''}>
          { ' ' + value }
        </span>
        {
          additional ? (
            <span className={css(leftStyle.additional)}>
              {' '} / { additional }
            </span>
          ) : ''
        }
      </span>
    </span>
  </div>
)



class Header extends Component {
  state = {
    web3Instance: undefined,
    web3Account: ''
  }

  componentWillMount = () =>
    store.watchState(this, [
      'web3Account',
      'web3Instance',
      'web3EthAmount'
    ])

  render = () => {
    const { web3Instance: web3, web3Account, web3EthAmount } = this.state
    let amount = ''
    const account = web3Account || 'Loading'
    if(web3) {
      amount = web3EthAmount.substring(0, 4)
    }
    return (
      <header className={css(style.main)}>
        <div className={css(style.holder)}>
          <div className={css(style.left)}>
            <LeftPart orange="all" value="ETH" name={amount} />
            <LeftPart orange="right" value={account} name="WALLET" delim />
          </div>
          <div className={css(style.right)}>
            <button className={css(style.button)}>
              Marketplace
            </button>
          </div>
        </div>
      </header>
    )
  }
}
export default Header
