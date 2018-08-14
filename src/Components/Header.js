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
    marginBottom: 30
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

const LeftPart = ({ value, name, delim, additional }) => (
  <div className={css(leftStyle.part)}>
    <div className={css(leftStyle.icon)} />
    <span>
      <span>
        <span className={css(leftStyle.orange)}>
          { name }{delim ? ':' : ''}
        </span>
        { ' ' + value }
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
    if(web3) {
      console.log(web3)
      const amount = web3EthAmount.substring(0, 4)
      console.log(amount)
    }
    return (
      <header className={css(style.main)}>
        <div className={css(style.holder)}>
          <div className={css(style.left)}>
            <LeftPart value="ETH" name="0.3" additional={`${'139.63'} USD`} />
            <LeftPart value="0.3" name="WALLET" delim />
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
