import React, { Component } from 'react'
import { css, StyleSheet } from 'aphrodite/no-important'
import { Editor } from 'react-draft-wysiwyg'
import 'react-draft-wysiwyg/dist/react-draft-wysiwyg.css'
import draftToHtml from 'draftjs-to-html'
import { convertToRaw, convertFromRaw, EditorState } from 'draft-js'
import DatePicker from 'react-datepicker'
import 'react-datepicker/dist/react-datepicker.css';
import moment from 'moment'
import { withScriptjs, withGoogleMap, GoogleMap, Marker } from 'react-google-maps'

import getWebThree from './../../util/getweb3'
import store from 'Store'
import config from 'Config'
import contract from 'truffle-contract'
import axios from 'axios'
import BREMContract from "../../../build/contracts/BREM.json"
import ipfs from "../../Config/ipfs";


import Alert from '../../Components/Alert'

const style = StyleSheet.create({
  row: {
    display: 'flex',
    flexDirection: 'column',
    marginBottom: 20
  },
  input: {
    height: 35,
    borderRadius: 3,
    border: 'solid 1px lightgray',
    padding: '0 15px'
  },
  form: {
    maxWidth: 500,
    marginBottom: 75
  },
  label: {
    marginBottom: 4,
    fontSize: 14
  }
})

const Maps = withScriptjs(withGoogleMap((props) =>
  <GoogleMap
    onClick={props.clickHandler}
    defaultZoom={4}
    defaultCenter={{ lat: 50.632138, lng: 10.396202 }}
  >
    {props.marker && <Marker position={{ lat: props.marker.lat, lng: props.marker.lng }} />}
  </GoogleMap>
))

const FormRow = ({ setRef, label, required, placeholder}) => (
  <div className={css(style.row)}>
    <label className={css(style.label)} htmlFor={label}>{ label }{required ? <span style={{color: 'red'}}>*</span> : ''}</label>
    <input className={css(style.input)} required={required} ref={setRef} type="text" placeholder={placeholder} name={label}/>
  </div>
)

async function createNewBREMICO(
  name,
  symbol,
  rate,
  cap,
  closingTime,
  description,
  files,
  image,
  location,
  locAddress
) {
  console.log('createNewBREMICO');
  const { host } = config
  const { web3Instance, web3Account } = store
  const { currentProvider, utils, eth } = web3Instance
  const brem = contract(BREMContract);
  brem.setProvider(currentProvider)
  const coinbase = await eth.getCoinbase()
  store.update({
    web3Coinbase: coinbase
  })
  
  const bremInstance = await brem.deployed();

  const isSuperuser = await bremInstance.isSuperuser(coinbase);
  const isAuditor   = await bremInstance.isAuditor(coinbase);

  if(isSuperuser || isAuditor){
    alert('Not for superuser or auditor');
    return;
  }

  if(cap < 100){
    alert('Invalid cap');
    return;
  }
  if(name.length == 0){
    alert('Invalid name');
    return;
  }
  if(symbol.length == 0){
    alert('Invalid symbol');
    return;
  }

  const isValidName = await bremInstance.isValidName(name);
  if (!isValidName){
    alert('Name already exists');
    return;
  }

  const feePercent = await bremInstance.withdrawFeePercent();


  ipfs.files.add(files, (error, result) => {
    if (error) {
      console.log(error);
      return;
    }
    const docHash = result[result.length - 1].hash;


    bremInstance
      .createBREMICO(
        name,
        symbol,
        rate,
        utils.toWei(cap.toString(), "ether"),
        closingTime.toString(),
        docHash,
        { from: coinbase }
      )
      .then(async TXres => {
        try {
          const ico = {
            address: TXres.logs[0].args.icoAddress,
            developer: {
              address: coinbase
            },
            description: description,
            closing_time: closingTime, //format?
            fee_percent: feePercent,
            token_address: TXres.logs[0].args.tokenAddress,
            name: name,
            symbol: symbol,
            location: location,
            loc_address: locAddress
          };
          let res = await axios.post(host + "dev/ico", ico);  //add authorization
          console.log(res);

          let formData = new FormData();
          formData.append(
            "address",
            TXres.logs[0].args.icoAddress
          );
          formData.append("image", image);
          const config = {
            headers: {
              "content-type": "multipart/form-data"
            }
          };
          axios.post(                                         //add authorization
                  host + "dev/image",
                  formData,
                  config
                );

        } catch(err){
          console.log(err);
        }
            

        alert(
          "TX: " +
          TXres.tx +
          " ICO: " +
          TXres.logs[0].args.icoAddress +
          " Token: " +
          TXres.logs[0].args.tokenAddress
        );

      });
                               
    });
    
}

class Create extends Component {
  state = {
    selectedDate: moment().add(1, 'days'),
    shouldAlert: '',
    marker: '',
    editorState: EditorState.createEmpty()
  }
  setRef = name => elem => this[name] = elem

  handleDateChange = date => {
    const min = moment().add(1, 'days')
    const diffTime = date.diff(min, 'days')
    if(diffTime < 0) return this.setState({
      shouldAlert: {
        text: `Please select date no less than 1 day more to the current date!`,
        heading: 'Wrong date!'
      }
    })
    this.setState({
      selectedDate: date
    })
  }

  handleProjectCreate = e => {
    e.preventDefault()
    const { state, name, thumbnail, files, time, street, token, cap, ratio } = this
    const { marker = {}, editorState } = state
    const { lat, lng } = marker
    const data = {
      name: name.value,
      thumbnail: thumbnail.files[0],
      desc: draftToHtml(convertToRaw(editorState.getCurrentContent())),
      // time: time.value,
      street: street.value,
      token: token.value,
      cap: cap.value,
      ratio: ratio.value,
      files: files.files,
      lat,
      lng
    }
    console.log(data.desc.length)
    if(data.desc.length < 45 && data.desc.includes('An actual description')) return this.setState({
      shouldAlert: {
        heading: 'Wow!',
        text: 'Uhh that was smart. But it won\'t trick me anyways ;)'
      }
    })
    if(data.desc.length < 200) return this.setState({
      shouldAlert: {
        heading: 'Description is too short!',
        text: 'Please write an actual description to your project!'
      }
    })
    if(!data.lat || !data.lng) return this.setState({
      shouldAlert: {
        heading: 'No address!',
        text: 'Please select valid address on the map!'
      }
    })
    if(data.token.length !== 3) return this.setState({
      shouldAlert: {
        heading: 'Wrong token name!',
        text: 'Token name have to contain three letters like this: "TKN"'
      }
    })
    /*
      SERVER AND CONTRACTS STUFF
    */
  }

  onEditorStateChange = editorState => this.setState({
    editorState
  })

  handleErrorResolved = () => this.setState({
    shouldAlert: ''
  })

  handleMarkerChange = e => {
    const { latLng } = e
    const { lat, lng } = latLng
    const marker = {
      lat: lat(),
      lng: lng()
    }
    this.setState({ marker })
  }

  render() {
    const { setRef, state } = this
    const { editorState, shouldAlert, selectedDate, marker } = state
    return (
      <div>
        {
          shouldAlert ? <Alert {...shouldAlert} callback={this.handleErrorResolved} /> : ''
        }
        <h3 className="page-heading">
          New Project
        </h3>
        <form className={css(style.form)} onSubmit={this.handleProjectCreate}>
          <FormRow setRef={setRef('name')} label="Name" placeholder="Object name" required />
          <div className={css(style.row)}>
            <label className={css(style.label)}>Description<span style={{color: 'red'}}>*</span></label>
              <Editor
                wrapperClassName="projecteditor-wrapper"
                editorClassName="projecteditor"
                editorState={editorState || null}
                onEditorStateChange={this.onEditorStateChange.bind(this)}
                toolbar={{ image: { uploadCallback: this.uploadImageEditor }}}
              />
          </div>
          <div className={css(style.row)}>
            <label className={css(style.label)}>
              Thumbnail<span style={{color: 'red'}}>*</span>
            </label>
            <input required style={{color: 'gray'}} ref={this.setRef('thumbnail')} type="file"/>
          </div>
          <div className={css(style.row)}>
            <label className={css(style.label)}>
              Ending date<span style={{color: 'red'}}>*</span>
            </label>
            <DatePicker
              className={css(style.input)}
              selected={selectedDate}
              onChange={this.handleDateChange}
            />
          </div>
          <div className={css(style.row)}>
            <label className={css(style.label)}>
              Map position<span style={{color: 'red'}}>*</span>
            </label>
            <Maps
              googleMapURL="https://maps.googleapis.com/maps/api/js?key=AIzaSyCuwdVfYrdvKY_TSJKHIq27jSaz-i3f55Y&v=3.exp&libraries=geometry,drawing,places"
              loadingElement={<div style={{ height: '400px', background: 'lightgray', display: 'flex' }}>
                <span style={{margin: 'auto', fontSize: '14px'}}>Maps are loading</span>
              </div>}
              containerElement={<div style={{ height: `400px` }} />}
              mapElement={<div style={{ height: `100%` }} />}
              marker={marker}
              clickHandler={this.handleMarkerChange}
            />
          </div>
          <FormRow setRef={setRef('street')} label="Address" placeholder="Street" required />
          <FormRow setRef={setRef('token')} label="Token Name" placeholder="Enter token name(3 letters)" required />
          <FormRow setRef={setRef('cap')} label="Soft-Cap" placeholder="Soft cap for your project" required />
          <FormRow setRef={setRef('ratio')} label="Rate" placeholder="token to WEI ratio" required />
          <div className={css(style.row)}>
            <label className={css(style.label)}>
              Additional files<span style={{color: 'red'}}></span>
            </label>
            <input multiple style={{color: 'gray'}} ref={this.setRef('files')} type="file"/>
          </div>
          <button>Submit</button>
        </form>
      </div>
    )
  }
}

export default Create
