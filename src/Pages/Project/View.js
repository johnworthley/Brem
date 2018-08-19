import React, { Component } from 'react'
import { css, StyleSheet } from 'aphrodite/no-important'
import moment from 'moment'
import { withScriptjs, withGoogleMap, GoogleMap, Marker } from 'react-google-maps'

const Maps = withScriptjs(withGoogleMap((props) =>
  <GoogleMap
    onClick={props.clickHandler}
    defaultZoom={4}
    defaultCenter={{ lat: 50.632138, lng: 10.396202 }}
  >
    {props.marker && <Marker position={{ lat: props.marker.lat, lng: props.marker.lng }} />}
  </GoogleMap>
))

const style = StyleSheet.create({
  main: {

  },
  top: {
    display: 'flex',
    minHeight: 300,
    marginBottom: 75
  },
  topPart: {
    width: '50%',
    minHeight: 300
  },
  topRight: {
    display: 'flex',
    flexDirection: 'column',
    padding: '0 20px',
    justifyContent: 'space-between'
  },
  topRightInner: {
    display: 'flex',
    flexDirection: 'column',
    ':not(:active) > span': { // fun hack
      marginBottom: 10
    },
  },
  topLeft: {
    backgroundColor: 'lightgray'
  },
  progressBar: {
    backgroundColor: 'lightgray',
    width: '100%',
    height: 8,
    borderRadius: 4,
    overflow: 'hidden'
  },
  progress: {
    backgroundColor: 'red',
    height: '100%',
    width: '50%'
  },
  description: {
    marginBottom: 30,
    lineHeight: 1.4
  }
})

class View extends Component {
  state = {
    projectId: '',
    project: {
      name: 'Neva Towers',
      company: 'Company name',
      address: 'Ul pidora, 12c4',
      closingDate: moment(),
      myEmphasis: 50, // How much eth was given to project by user
      softCap: 1000,
      status: 'Opened',
      icoBalance: 400,
      withDrawRequest: 0,
      goal: 2000,
      auditorsList: [],
      description: '<div style="color: green;">HTML EEEE</div>',
      latLng: {
        lat: 2, lng: 50
      }
    }
  }

  componentDidMount = async () => {
    this.setState({
      projectId: this.props.match.params.id
    })


    // Get all data here and push it to state.project object
  }
  depositETH = async () => {

  }

  withDrawETH = async () => {

  }

  auditorConfirmRequest = async () => {

  }

  devWithdrawETH = async () => {

  }

  superAddAuditor = () => {
    const auditorKey = this.newAuditorName.value
    console.log(auditorKey)
  }

  render() {
    const {
      projectId,
      name,
      company,
      address,
      closingDate,
      myEmphasis,
      softCap,
      status,
      icoBalance,
      withDrawRequest,
      goal,
      auditorsList,
      description
    } = this.state.project
    const html = {__html: description}
    return (
      <div className={css(style.main)}>
        <section className={css(style.top)}>
          <div className={css(style.topLeft, style.topPart)}>

          </div>
          <div className={css(style.topRight, style.topPart)}>
            <div className={css(style.topRightInner)}>
              <span>{ company }</span>
              <span style={{textTransform: 'uppercase', fontSize: 18}}>{ name }</span>
              <span> { closingDate + '' } </span>
              <span>My funds: <span style={{color: 'red', }}>{ myEmphasis }</span><span style={{fontSize: 12, color: 'red', marginLeft: 3}}>ETH</span></span>
              <span>
                Soft-cap: { softCap } ETH
              </span>
              <div className={css(style.progressBar)}>
                <div className={css(style.progress)} />
              </div>
            </div>
            <div className={css(style.topButtons)}>
              <button style={{marginRight: 25}} onClick={this.depositETH}>
                 Deposit ETH
              </button>
              <button onClick={this.withdrawETH}>
                Withdraw ETH
              </button>
            </div>
          </div>
        </section>
        <section className={css(style.description)}>
          <div dangerouslySetInnerHTML={html} />
        </section>
        <section className={css(style.map)}>
          <Maps
            googleMapURL="https://maps.googleapis.com/maps/api/js?key=AIzaSyCuwdVfYrdvKY_TSJKHIq27jSaz-i3f55Y&v=3.exp&libraries=geometry,drawing,places"
            loadingElement={<div style={{ height: '400px', background: 'gray', display: 'flex' }}>
              <span style={{margin: 'auto', fontSize: '14px'}}>Maps are loading</span>
            </div>}
            containerElement={<div style={{ height: `400px` }} />}
            mapElement={<div style={{ height: `100%` }} />}
          />
        </section>
        Project ID: { projectId }

        <button onClick={this.auditorConfirmRequest}>
          Auditor Confirm Request
        </button>
        <button onClick={this.devWithdrawETH}>
          Dev withdraw eth
        </button>
        <input ref={elem => this.newAuditorName = elem} type="text" placeholder="Auditor name"/>
        <button onClick={this.superAddAuditor}>
          SuperUser Add auditor
        </button>
      </div>
    )
  }
}

export default View
