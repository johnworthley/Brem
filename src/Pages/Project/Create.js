import React, { Component } from 'react'
import { css, StyleSheet } from 'aphrodite/no-important'
import { Editor } from 'react-draft-wysiwyg'
import 'react-draft-wysiwyg/dist/react-draft-wysiwyg.css'
import draftToHtml from 'draftjs-to-html'
import { convertToRaw, EditorState } from 'draft-js'
// import DatePicker from 'react-datepicker'
// import 'react-datepicker/dist/react-datepicker.css'
import moment from 'moment'
import { withScriptjs, withGoogleMap, GoogleMap, Marker } from 'react-google-maps'

import store from 'Store'
import config from 'Config'
import contract from 'truffle-contract'
import axios from 'axios'
import BREMContract from "../../../build/contracts/BREM.json"
import ipfs from "../../Config/ipfs"


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

async function createNewBREMICO({
  name,
  symbol,
  rate,
  cap,
  closingTime,
  description,
  files,
  thumbnail: image,
  location,
  locAddress
}) {
  console.log('createNewBREMICO');
  const { host, authConfig } = config
  const { web3Instance } = store
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

  if(name.length === 0){
    alert('Invalid name');
    return;
  }

  if(symbol.length !== 3){
    alert('Invalid symbol');
    return;
  }

  const isValidName = await bremInstance.isValidName(name);
  if (!isValidName){
    alert('Name already exists');
    return;
  }

  const feePercentBN = await bremInstance.withdrawFeePercent();
  const feePercent = feePercentBN.toNumber()
  const closingTimeWeb3 = (closingTime.getTime() / 1000).toString();
  const closingTimeISO = closingTime.toISOString()
  const capWei = utils.toWei(cap, "ether")

  console.log("Waiting for IPFS")

  ipfs.files.add(files, async (error, result) => {
    let docHash = "123"
    if (error) {
      console.log(error);
      // return;
    } else {
      docHash = result[result.length - 1].hash
    }

    const TXres = await bremInstance.createBREMICO(name, symbol, rate, capWei, closingTimeWeb3, docHash, { from: coinbase })
        const tx = TXres.tx;
        const status = TXres.receipt.status === "0x1"
        if (!status) {
          console.log("Error, status " + TXres.receipt.status)
          // Look in etherscan https://rinkeby.etherscan.io/tx/ + tx
          return
        }

        const icoAddress = TXres.logs[0].args.tokenAddress
        const tokenAddress = TXres.logs[0].args.tokenAddress
        // TODO: message
        console.log(icoAddress)
        console.log(tokenAddress)
        console.log(tx)

        const ico = {
          address: icoAddress,
          developer: {
            address: coinbase
          },
          description: description,
          closing_time: closingTimeISO,
          fee_percent: feePercent,
          token_address: tokenAddress,
          name: name,
          symbol: symbol,
          location: location,
          loc_address: locAddress
        };

        try {
          let res = await axios.post(host + "dev/ico", ico, authConfig)
          console.log(res)

          const formData = new FormData();
          formData.append(
            'address',
            icoAddress
          );
          formData.append("image", image);

          const reqConfig = {
            headers: {
              "content-type": "multipart/form-data"
            },
            withCredentials: true,
          };
          res = await axios.post(
                  host + "dev/image",
                  formData,
                  reqConfig
                );
          console.log(res)

        } catch(err){
          console.log(err);
        }
    });

}

class Create extends Component {
  state = {
    selectedDate: moment().add(2, 'days'),
    shouldAlert: '',
    marker: '',
    editorState: EditorState.createEmpty()
  }
  setRef = name => elem => this[name] = elem

  handleDateChange = date => {
    // const min = moment().add(2, 'days')
    // const diffTime = date.diff(min, 'days')
    // if(diffTime < 0) return this.setState({
    //   shouldAlert: {
    //     text: `Please select date no less than 1 day more to the current date!`,
    //     heading: 'Wrong date!'
    //   }
    // })
    this.setState({
      selectedDate: date
    })
  }

  handleProjectCreate = async e => {
    e.preventDefault()
    const { state, name, thumbnail, files, street, token, cap, ratio, selectedDate, } = this
    const { marker = {}, editorState } = state
    const { lat, lng } = marker
    const data = {
      name: name.value,
      thumbnail: thumbnail.files[0],
      description: draftToHtml(convertToRaw(editorState.getCurrentContent())),
      // time: time.value,
      locAddress: street.value,
      symbol: token.value,
      cap: cap.value,
      rate: ratio.value,
      files: files.files,
      closingTime: new Date(selectedDate.value),
      location: JSON.stringify({
        lat,
        lng
      }),
    }
    if(data.description.length < 45 && data.description.includes('An actual description')) return this.setState({
      shouldAlert: {
        heading: 'Wow!',
        text: 'Uhh that was smart. But it won\'t trick me anyways ;)'
      }
    })
    if(data.description.length < 200) return this.setState({
      shouldAlert: {
        heading: 'Description is too short!',
        text: 'Please write an actual description to your project!'
      }
    })
    if(!lat || !lng) return this.setState({
      shouldAlert: {
        heading: 'No address!',
        text: 'Please select valid address on the map!'
      }
    })
    if(data.symbol.length !== 3) return this.setState({
      shouldAlert: {
        heading: 'Wrong token name!',
        text: 'Token name have to contain three letters like this: "TKN"'
      }
    })
    data.files = await Promise.all([...data.files].map(file => new Promise(resolve => {
      const reader = new window.FileReader()
      reader.readAsArrayBuffer(file)
      reader.onloadend = () => {
        resolve(Buffer(reader.result))
      }
    })))
    data.image = await new Promise(resolve => {
      const thumbnail= data.thumbnail
      const reader = new window.FileReader()
      reader.readAsArrayBuffer(thumbnail)
      reader.onloadend = () => {
        resolve(thumbnail)
      }
    })

    createNewBREMICO(data)
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
            <input required style={{color: 'gray'}} ref={this.setRef('thumbnail')} type="file" accept="image/*"/>
          </div>
          <div className={css(style.row)}>
            <label className={css(style.label)}>
              Ending date<span style={{color: 'red'}}>*</span>
            </label>
            <input
              type="datetime-local"
              ref={elem => this.selectedDate = elem}
              className={css(style.input)}
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
