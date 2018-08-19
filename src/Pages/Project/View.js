import React, { Component } from 'react'
import { css, StyleSheet } from 'aphrodite/no-important'
import moment from 'moment'
import { withScriptjs, withGoogleMap, GoogleMap, Marker } from 'react-google-maps'
import getWeb3 from '../../util/getweb3'
import store from 'Store'
import config from 'Config'
import contract from 'truffle-contract'
import ICOContract from "../../../build/contracts/BREMICO.json"
import TokenContract from "../../../build/contracts/BREMToken.json"
import BREMContract from "../../../build/contracts/BREM.json"
import axios from 'axios'

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
<<<<<<< HEAD

=======
import getWeb3 from '../../util/getweb3'
import store from 'Store'
import config from 'Config'
import contract from 'truffle-contract'
import ICOContract from "../../../build/contracts/BREMICO.json"
import TokenContract from "../../../build/contracts/BREMToken.json"
import BREMContract from "../../../build/contracts/BREM.json"
import axios from 'axios';
>>>>>>> 6be69b7dc9a89c4957dc334f7a2b3fafaecb81c9

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
    await getWeb3();
    const { web3Instance, web3Account } = store
    const { currentProvider, utils, eth } = web3Instance

    const coinbase = await web3Instance.getCoinbase()

    const ico = contract(ICOContract)
    ico.setProvider(currentProvider)
    const icoInstance = await ico.at(this.state.projectId)

    const icoBalance = await eth.betBalance(this.state.projectId)
    const tokenAddress = await icoInstance.token()
    const wallet = await icoInstance.wallet()
    const rate = await icoInstance.rate()
    const weiRaised = await icoInstance.weiRaised()
    const cap = await icoInstance.cap()
    const docHash = await icoInstance.docHash()
    const closingTimeEpochs = await icoInstance.closingTime()
    const weiInvested = await icoInstance.getBalance(coinbase, {from: coinbase})

    const token = contract(TokenContract)
    token.setProvider(currentProvider)
    const tokenInstance = await token.at(tokenAddress)

    const name = await tokenInstance.name()
    const symbol = await tokenInstance.symbol()
    const tokenBalance = await tokenInstance.balanceOf(coinbase)

    const brem = contract(BREMContract)
    brem.setProvider(currentProvider)
    const bremInstance = await brem.deployed() 

    const devName = await bremInstance.login(wallet)

    this.setState({
      name: name,
      tokenAddress: tokenAddress,
      icoBalance: utils.fromWei(icoBalance, "ether"),
      devName: devName,
      wallet: wallet,
      rate: rate.toNumber(),
      weiRaised: utils.fromWei(weiRaised, "ether"),
      cap: utils.fromWei(cap, "ether"),
      docsURL: 'https://ipfs.infura.io/ipfs/' + docHash,
      closingTime: new Date(closingTimeEpochs.toNumber() * 1000).toString(),
      invested: utils.fromWei(weiInvested, "ether"),
      symbol: symbol,
      tokenBalance: utils.fromWei(tokenBalance, "ether")
    })

    // Read status
    const hasClosed = await icoInstance.hasClosed()
    const capReached = await icoInstance.capReached()
    const auditSelected = await icoInstance.auditSelected()
    const isRequested = await icoInstance.isRequested()
    const isWithdrawn = await icoInstance.isWithdrawn()
    const isOverdue = await icoInstance.isOverdue()

    this.setState({
      isCreated: !auditSelected,
      isOpened: auditSelected && !hasClosed,
      isSuccess: hasClosed && capReached && !isRequested && !isWithdrawn,
      isFailed: auditSelected && hasClosed && !capReached,
      isRequested: isRequested,
      isWithdrawn: isWithdrawn,
      isOverdue: isOverdue
    })

    // Read current auditors
    const auditorsAmount = icoInstance.auditorsAmount()
    const auditors = []
    for (let i = 0; i < auditorsAmount; i++) {
      const auditorAddress = await icoInstance.getAuditor(i)
      const auditorUsername = await bremInstance.login(auditorAddress)
      auditors.push({
        address: auditorAddress,
        username: auditorUsername
      })
    }
    this.setState({
      auditorsList: auditors
    })

    // Get user type
    const isSuperuser = await bremInstance.isSuperuser(coinbase)
    const isAuditor = await icoInstance.isAuditor(coinbase)
    const isDeveloper = coinbase === wallet

    this.setState({
      isSuperuser: isSuperuser,
      isAuditor: isAuditor,
      isDeveloper: isDeveloper
    })

    if (isRequested) {
      const request = await icoInstance.request()
      const requestedValue = utils.fromWei(request[0], "ether")
      this.setState({
        requestedValue: requestedValue
      })
    }

    if (this.state.isSuccess && isDeveloper) {
      const withdrawFeePercent = icoInstance.withdrawFeePercent()
      this.setState({
        currentICOFee: withdrawFeePercent.toNumber() // Показывать при запросе вывода застройщиком
      })
    }
  }

  depositETH = async e => {
    // Здесь нужна отдельная форма с расчетом количества токенов
    e.preventDefault()
    const depositValue = this.state.depositValue
    if (depositValue <= 0) {
      // Ошибка
      return
    }

    await getWeb3
    const { web3Instance, web3Account } = store
    const { currentProvider, utils, eth } = web3Instance

    const coinbase = await eth.getCoinbase()

    const ico = contract(ICOContract)
    ico.setProvider(currentProvider)
    const icoInstance = await ico.at(this.state.projectId)

    // Check for status
    const hasClosed = await icoInstance.hasClosed()
    const auditSelected = await icoInstance.auditSelected()
    if (!auditSelected || hasClosed) {
      // Ошибка, нельзя купить
      return
    }

    // Invest
    const depositInWei = utils.toWei(depositValue, "ether")
    try {
      const txRes = await icoInstance.transfer({from: coinbase, to: ico.address, value: depositInWei})
      const isSuccessTX = txRes.receipt.status === "0x1"
      const tx = txRes.tx;
      if (isSuccessTX) {
        // Все ок, можно посмотреть на etherscan https://rinkeby.etherscan.io/tx/ + tx
        // Предложить добавить адрус токена в metamask
      } else {
        // Ошибка, можно посмотреть на etherscan https://rinkeby.etherscan.io/tx/ + tx
      }
    } catch(err) {
      console.error(err)
    }
  }

  withDrawETH = async e => {
    // Кнопка доступна только если isFailed && инвестированый баланс больше 0
    e.preventDefault()

    await getWeb3
    const { web3Instance, web3Account } = store
    const { currentProvider, utils, eth } = web3Instance

    const coinbase = await eth.getCoinbase()

    const ico = contract(ICOContract)
    ico.setProvider(currentProvider)
    const icoInstance = await ico.at(this.state.projectId)

    // Check for status
    const hasClosed = await icoInstance.hasClosed()
    const capReached = await icoInstance.capReached()
    const weiInvestedBN = await icoInstance.getBalance({from: coinbase})
    const weiInvested = weiInvestedBN.toNumber()
    if (hasClosed && !capReached && weiInvested <= 0) {
      // Ошибка, нельзя купить
      return
    }

    // Refund Ether
    try {
      const txRes = await icoInstance.refund({from: coinbase})
      const isSuccessTX = txRes.receipt.status === "0x1"
      const tx = txRes.tx;
      if (isSuccessTX) {
        // Все ок, можно посмотреть на etherscan https://rinkeby.etherscan.io/tx/ + tx
        // Предложить добавить адрус токена в metamask
      } else {
        // Ошибка, можно посмотреть на etherscan https://rinkeby.etherscan.io/tx/ + tx
      }
    } catch(err) {
      console.error(err)
    }
  }

  auditorConfirmRequest = async e => {

  }

  devWithdrawETH = async e => {
        // Кнопка доступна только если isSuccess && !isWithdrawn && wallet == coinbase && !isRequested
        e.preventDefault()

        const requestEthValue = this.state.devRequestValue

        await getWeb3
        const { host } = config
        const { web3Instance, web3Account } = store
        const { currentProvider, utils, eth } = web3Instance

        const coinbase = await eth.getCoinbase()

        const ico = contract(ICOContract)
        ico.setProvider(currentProvider)
        const icoInstance = await ico.at(this.state.projectId)

        // Check for status
        const hasClosed = await icoInstance.hasClosed()
        const capReached = await icoInstance.capReached()
        const isRequested = await icoInstance.isRequested()
        const isWithdrawn = await icoInstance.isWithdrawn()
        const wallet = await icoInstance.wallet()
        let requestValue = utils.toWei(requestEthValue, "ether")
        const contractBalance = await eth.getBalance(ico.address)
        if (!hasClosed || !capReached || isRequested || isWithdrawn || contractBalance < requestValue || coinbase !== wallet) {
          // Ошибка, нельзя сделать запрос
          return
        }

        if (contractBalance - requestValue < 100) {
          requestValue = contractBalance; // При ввода значения для вывода сделать такую проверку
        }

        // Make request
        try {
          const txRes = await icoInstance.withdraw(requestValue, {from: coinbase})
          const isSuccessTX = txRes.receipt.status === "0x1"
          const tx = txRes.tx;
          if (isSuccessTX) {
            // Все ок, можно посмотреть на etherscan https://rinkeby.etherscan.io/tx/ + tx
            // Предложить добавить адрус токена в metamask
            axios.put(host)

          } else {
            // Ошибка, можно посмотреть на etherscan https://rinkeby.etherscan.io/tx/ + tx
          }
        } catch(err) {
          console.error(err)
        }
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
