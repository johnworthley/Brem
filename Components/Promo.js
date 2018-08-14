import React, { Component } from 'react'
import { CSSTransitionGroup } from 'react-transition-group'

import store from '../lib/store'


import {Motion, spring} from 'react-motion'


import UserLogReg from './UserLogReg'

import Signin from './Signin'
import Signup from './Signup'

import events from 'dobroevents'


class Promo extends Component {
  constructor() {
    super()
    this.state = {
      promopic: '',
      height: 300,
    }
  }
  componentWillMount() {
    this.userstatusListener = store.watchState(this, ['promopic', 'promotype', 'promofullscreen'])
  }
  render() {
    const { promopic, promofullscreen, height, promotype } = this.state
    const backgroundImage = `url(${ promopic })`
    const backgroundPosition = 'center'
    const classn = this.props.blogPost === true ?
    'promo_inner promo_inner-active' : 'promo_inner'
    const heightParam = promotype ? {
      minHeight: promofullscreen ? '100vh' : height,
      maxHeight: 'auto'
    } : {
      maxHeight: height,
      minHeight: 'auto'
     }
    return (
      <div className="promo" style={heightParam}>
        <CSSTransitionGroup
          transitionName="PromoTransition"
          transitionEnterTimeout={300}
          transitionLeaveTimeout={300}>
          <div key={promopic + '' + promotype} className={classn}>
            <div className="promo_background" style={{ backgroundImage, backgroundPosition }}></div>
            {
              promotype ? <UserLogReg
                comp={promotype=== 'signin' ? Signin : Signup}
                type={promotype} /> : ''
            }


          </div>
        </CSSTransitionGroup>
      </div>
    )
  }
}

export default Promo
