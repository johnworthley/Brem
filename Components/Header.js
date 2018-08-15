import React, { Component } from 'react'
import { CSSTransitionGroup } from 'react-transition-group'
import { Link } from 'react-router-dom'
import '../css/Header.css'

import store from '../lib/store'
import csscuts from 'dobrocss'
import events from 'dobroevents'

const topStyles = {
  display: 'flex',
  justifyContent: 'flex-end'
}

const styleGreeting = {
  ...csscuts.rel,
  fontWeight: 700,
  fontSize: 20,
  color: '#fff',
  display: 'flex'
}

const styleGreetingBg = {
  ...csscuts.abs,
  width: '150%',
  height: '150%',
  backgroundImage: 'linear-gradient(35deg, #5A224A, #1A345F)',
  zIndex: -1,
  ...csscuts.centerAbs,
  transform: 'translate(-50%, -50%) skewX(-35deg)',
  borderRadius: '2px 5px 2px 5px'
}


class Header extends Component {
  constructor(props) {
    super(props)
    this.state = {
      navItems: [],
      promotype: '',
      username: ''
    }
  }
  componentWillMount() {
    store.watchState(this, ['username', 'currentPosition', 'promotype'])
    //const watchCounter = store.watchState(this, ['qwe'])

    //setTimeout(() => watchCounter.detach(), 5000)


    this.setState({
      navItems: [
        {
          route: '/',
          name: 'Главная',
          routeName: 'home'
        },
        {
          route: '/blog',
          name: 'Блог',
          routeName: 'blog'
        },
        {
          route: '/reviews',
          name: 'Обзоры',
          routeName: 'reviews'
        },
        {
          route: '/jam',
          name: 'Джем',
          routeName: 'jam'
        },
      ]
    })

  }

  handleExit() {
    document.cookie = 'hash' + '=; expires=Thu, 01 Jan 1970 00:00:01 GMT;'
    document.cookie = 'userid' + '=; expires=Thu, 01 Jan 1970 00:00:01 GMT;'
    store.update({username: undefined})
  }

  handleLinkClick(e) {
    const { promoTransitionIsGoing } = store
    if(promoTransitionIsGoing) return e.preventDefault()
    store.update({promoTransitionIsGoing: true})
    setTimeout(() => store.update({promoTransitionIsGoing: false}), 300)
  }

  render() {
    const { navItems, promotype, username } = this.state
    const navItemsToRender = navItems && navItems.map ?
      this.state.navItems.map(item => (
        <Link onClick={this.handleLinkClick.bind(this)} key={item.route} className={`page_header_link ${this.state.currentPosition === item.routeName ? 'page_header_link-active' : ''}`} to={item.route}>{item.name}</Link>
      )) : ''
    return (
      <header style={{marginTop: promotype ? -205 : 0}} className="page_header">
        <div className="page_header_inner container">
          <div style={{...topStyles, height: 23}} className="container-text">
            <CSSTransitionGroup
              transitionName="PromoTransition"
              transitionEnterTimeout={300}
              transitionLeaveTimeout={300}>
              {
                username ? (
                  <div key="1" style={styleGreeting}>

                    <Link to={`/users/${username}`}>
                      <span style={{position: 'relative', marginLeft: 20, color: '#fff'}}>
                        <div style={styleGreetingBg} />
                        <div onClick={this.handleExit.bind(this)} className="page_header_exit" style={{color: '#000', fontSize: 13, lineHeight: '24px'}}>
                          Выйти
                        </div>
                        {username}
                      </span>
                    </Link>
                  </div>

                ) : (
                  <div key="2">
                    <Link to="/signup">
                      <span className="page_header_reglink" style={{marginRight: 20}}>
                        Регистрация
                      </span>
                    </Link>

                    <Link to="/signin">
                      <span className="page_header_reglink">
                        Вход
                      </span>
                    </Link>
                  </div>
                )
              }
            </CSSTransitionGroup>

          </div>
          <nav className="page_header_nav container-text">
            <Link className="page_header_user" to="/">
              <img
                style={{height: 100, width: 100}}
                src="https://pp.userapi.com/c631316/v631316029/243bc/JOvdlOw_Rts.jpg"
                alt="Логотип"/>
            </Link>
            {navItemsToRender}
          </nav>
        </div>
      </header>
    )
  }
}

export default Header
