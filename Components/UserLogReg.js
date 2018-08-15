import React, { Component } from 'react'
import Transition from 'react-motion-ui-pack'
import { spring } from 'react-motion'
import { CSSTransitionGroup } from 'react-transition-group'
import csscuts from 'dobrocss'



export default class UserLogReg extends Component  {
  render() {
    const { type, comp: Comp } = this.props
    return (
      <Transition
        component="ul"
        enter={{
          opacity: spring(1, {stiffness: 40, damping: 10}),
        }}
        leave={{
          opacity: spring(0, {stiffness: 40, damping: 10}),
        }}
      >
      <div key={type} style={{...csscuts.fill, backgroundColor: 'rgba(30,30,30,.6)'}}>
        <Comp />
      </div>
      </Transition>

    )
  }
}
