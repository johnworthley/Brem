import React, { Component } from 'react'
import { Link, Redirect } from 'react-router-dom'

import Loader from './Loader'

import store from '../lib/store'

import csscuts from 'dobrocss'

const mainHolderStyle = {
  width: '100%',
  maxWidth: 500,
  ...csscuts.centerAbs,
  ...csscuts.rel,
  background: 'rgba(255,255,255,.7)',
  mixBlendMode: 'burn',
  zIndex: 20,
  borderRadius: 4,
  overflow: 'hidden',
  boxSizing: 'border-box',
  padding: '30px 20px 50px 20px'
}

const mainDecorHolderStyle = {
  ...csscuts.fill,
  zIndex: -1
}

const stripDecor = {
  width: 240,
  height: '120%',
  background: 'red',
  transform: 'skew(-25deg) translateX(-83%)',
  ...csscuts.abs,
}

const stripStd = {
  width: 260,
  backgroundImage: 'linear-gradient(to bottom, #731B41, #1A345F)'
}

const stripGrr = {
  bottom: 0,
  right: '-50%',
  backgroundImage: 'linear-gradient(to bottom, green, pink)'

}

const stripCenter = {
  bottom: '-20%',
  right: '-50%',
  width: '80%',
  backgroundImage: 'linear-gradient(to bottom, #383838, #282828)'
}


const mainFormStyle = {
  maxWidth: '70%',
  margin: '0 auto',
  textAlign: 'center',
  display: 'flex',
  flexFlow: 'column nowrap',
  color: '#fff'
}

const formItemInput = {
  margin: '5px auto 30px auto',
  maxWidth: 124,
  height: 27,
  padding: '0 30px'
}

const formItem = {
  textTransform: 'uppercase',
  fontFamily: 'Ubuntu, sans-serif'
}

const formButton = {
  background: 'rgba(255,255,255,.5)',
  backgroundImage: 'linear-gradient(-35deg, #731B41, #731B41)',
  border: 'none',
  height: 33,
  color: '#fff',
  cursor: 'pointer',
  boxShadow: '0 0 30px 0 rgba(0,0,0,.2)',
  margin: '0 auto',
  padding: '0 30px'
}


const bottomStyle = {
  ...csscuts.abs,
  color: '#fff',
  bottom: 18,
}

const regStyle = {
  left: 18,
}

const backStyle = {
  right: 18,
}

class Signup extends Component {
  constructor() {
    super()
    this.state = {
      error: null,
      successfulSignup: false,
      needToRedirect: false,
      loading: false
    }
  }
  preventWithError(error) {
    this.setState({
      error: true,
      errordetails: error
    })
    clearTimeout(this.errorTimeout)
    this.errorTimeout = setTimeout(() => this.setState({error: null}), 3000)
  }
  handleRegister(e) {
    e.preventDefault()
    const { loginRef, passwordRef, repeatPasswordRef } = this.refs
    const login = loginRef.value
    const password = passwordRef.value
    const repeatPassword = repeatPasswordRef.value
    if(!login) return this.preventWithError('Не заполнено поле логин')
    else if(login.length < 3) return this.preventWithError('Логин должен состоять хотя бы из трёх символов.')
    else if(!password) return this.preventWithError('Не заполнено поле пароль')
    else if(!repeatPassword) return this.preventWithError('Вы не повторили пароль')
    else if(password !== repeatPassword) return this.preventWithError('Пароли не совпадают')
    this.setState({
      loading: true
    })
    fetch('https://guitclub.ru/api/signup', {
      method: 'POST',
      credentials: 'same-origin',
      body: JSON.stringify({
        login,
        password,
        repeatPassword
      })
    })
    .then(data => data.json())
    .then(data => {
      this.setState({
        loading: false
      })
      if(data.error) return this.preventWithError(data.errorname)
      this.setState({successfulSignup: true})
      setTimeout(() => {
        store.disableFullscreen()

        this.setState({needToRedirect: true})
    }, 1500)
    })
    .catch(e => {
      this.setState({
        loading: false
      })
      alert('Произошла страшная ошибка. Попробуйте перезагрузить страницу.')
      console.log(e)
    })
  }
  render() {
    const { error, errordetails, successfulSignup, loading, needToRedirect } = this.state
    return (
      <div style={mainHolderStyle}>
        {
          needToRedirect ? (
            <Redirect to="/signin" />
          ) : ''
        }
        <div style={mainDecorHolderStyle}>
          <div style={{...stripDecor, ...stripStd}} />
          <div style={{...stripDecor, ...stripGrr}} />
          <div style={{...stripDecor, ...stripCenter}} />
        </div>
        <form style={mainFormStyle} action="">

            {
              loading || successfulSignup  ? (
                <div style={{...csscuts.fill, backgroundColor: '#fff', display: 'flex'}}>
                {
                  loading ? (
                    <div style={{margin: 'auto', width: 60, height: 60}}>
                      <Loader />
                    </div>
                  ) : (
                    <div style={{margin: 'auto', color: '#282828', fontSize: 20}}>
                      Вы успешно зарегистрировались!
                    </div>
                  )
                }
              </div>
              ) : ''
            }
          <h2>Регистрация</h2>
          <div style={{...formItem, marginTop: 0}}>Ваш логин</div>
          <input ref="loginRef" style={{...formItemInput}} type="text"/>
          <div style={{...formItem}}>Ваш пароль</div>
          <input ref="passwordRef" style={{...formItemInput}} type="text"/>
          <div style={{...formItem}}>Повторите пароль</div>
          <input ref="repeatPasswordRef" style={{...formItemInput}} type="text"/>
          <button onClick={this.handleRegister.bind(this)} style={{...formButton}}>Зарегистрироваться</button>
          <span ref="errorRef" style={{display: 'inline', marginTop: 9, backgroundColor: '#fff', color: '#c10f0f', transition: 'opacity .3s ease', opacity: error ? 1 : 0}}>
            {errordetails || 'Что-то пошло не так'}
          </span>
        </form>

        <Link to="/signin">
          <div style={{...bottomStyle, ...regStyle}}>
            Уже зарегистрированы?
          </div>
        </Link>
        <Link to="/" onClick={store.disableFullscreen}>
          <div style={{...bottomStyle, ...backStyle}}>
            Вернуться
          </div>
        </Link>
      </div>
    )
  }
}

export default Signup
