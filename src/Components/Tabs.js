import React, { Component } from 'react'
import { css, StyleSheet } from 'aphrodite/no-important'

const style = StyleSheet.create({
  main: {
    display: 'flex',
    flexDirection: 'column',
    margin: '50px auto 0 auto'
  },
  tabLabels: {
    display: 'inline-flex',
    position: 'relative',
    marginBottom: 70,
    width: 'auto',
    '::after': {
      content: '""',
      width: '100%',
      height: 1,
      bottom: -30,
      backgroundColor: '#d5d5d6',
      position: 'absolute'
    }
  },
  tabLabel: {
    marginLeft: 20,
    cursor: 'pointer',
    color: '#f0723b',
    textTransform: 'uppercase',
    position: 'relative',
    ':first-child': {
      marginLeft: 0
    },
    '::after': {
      content: '""',
      left: 0,
      zIndex: 2,
      width: '100%',
      height: 3,
      bottom: -30,
      backgroundColor: '#f0723b',
      position: 'absolute',
      transformOrigin: 'bottom',
      transform: 'scaleY(0)',
      transition: 'transform .2s ease'
    }
  },
  activeTabLabel: {
    '::after': {
      transform: 'scale(1)'
    }
  }
})

class Tabs extends Component {
  state = {
    tabs: []
  }
  componentWillMount = () => {
    const { data = [], maxWidth = 500 } = this.props
    const TabComponent = data.find(tab => tab.default).component ||
      data[0] ? data[0].component : <div>DEV: Your tab list is empty!</div>
    this.setState({
      data: data.map(
        tab => tab.default ?
          {...tab, isActive: true} :
          {...tab, isActive: false}
      ), maxWidth, TabComponent
    })
  }

  changeTab = name => this.setState({
    TabComponent: this.state.data.find(tab => tab.name === name).component,
    data: this.state.data.map(
      tab => tab.name === name ?
        {...tab, isActive: true} :
        {...tab, isActive: false}
    )
  })

  render() {
    const { data, TabComponent, maxWidth } = this.state

    return (
      <div className={css(style.main)} style={{maxWidth}}>
        <div>
          <div className={css(style.tabLabels)}>
            {
              data.map(tab => {
                const { name, isActive } = tab
                return (
                  <div
                    key={name}
                    className={
                      css(style.tabLabel, isActive ? style.activeTabLabel : '')
                    }
                    onClick={() => this.changeTab(name)}
                  >
                    { name }
                  </div>
                )
              })
            }
          </div>
        </div>
        <div>
          <TabComponent />
        </div>
      </div>
    )
  }
}

export default Tabs
