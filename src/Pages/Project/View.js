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

console.log(store)

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
    ':active > span': { // fun hack
      marginBottom: 15
    },
    ':not(:active) > span': { // fun hack
      marginBottom: 15
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
    backgroundColor: 'rgb(240, 114, 59)',
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
      // name: 'Neva Towers',
      // company: 'Company name',
      // address: 'Ul pidora, 12c4',
      // closingDate: moment(),
      // myEmphasis: 50, // How much eth was given to project by user
      // softCap: 1000,
      // status: 'Opened',
      // icoBalance: 400,
      // withDrawRequest: 0,
      // goal: 2000,
      // auditorsList: [],
      // description: '<div style="color: green;">HTML EEEE</div>',
      // latLng: {
      //   lat: 2, lng: 50
      // }
    }
  }

  componentDidMount = async () => {
    const icoAddress = this.props.match.params.id;

    this.setState({
      projectId: icoAddress
    })
    let serverData

    // Get description, image, location, address and dev name
    const { host } = config
    console.log(host)
    axios.get(host + 'ico/', {
      params: {
        address: icoAddress
      }
    })
    .then(res => {
      serverData = res.data
      console.log(serverData)
      // get image
      axios.get(host + 'ico/image', {
        params:{
          address: icoAddress,
        },
        responseType: "arraybuffer"
      })
      .then(res => {
          const base64 = Buffer.from(res.data, "binary").toString("base64");
          this.setState({ img: "data:image/jpeg;base64," + base64 });
        })
        .catch(err => console.error(err))
    })
    .catch(err => console.error(err))

    // Get all data here and push it to state.project object
    const { web3Instance, web3Account } = store
    const { currentProvider, utils, eth } = web3Instance

    const coinbase = await eth.getCoinbase()

    const ico = contract(ICOContract)
    ico.setProvider(currentProvider)
    const icoInstance = await ico.at(this.state.projectId)

    const icoBalance = await eth.getBalance(this.state.projectId)
    const tokenAddress = await icoInstance.token()
    const wallet = await icoInstance.wallet()
    const rate = await icoInstance.rate()
    const weiRaised = await icoInstance.weiRaised()
    const capWei = await icoInstance.cap()
    const docHash = await icoInstance.docHash()
    const closingTimeEpochs = await icoInstance.closingTime()
    const weiInvested = await icoInstance.getBalance(coinbase, {from: coinbase})

    const token = contract(TokenContract)
    token.setProvider(currentProvider)
    const tokenInstance = await token.at(tokenAddress)

    const name = await tokenInstance.name()
    const symbol = await tokenInstance.symbol()
    const tokenBalanceRaw = await tokenInstance.balanceOf(coinbase)

    const brem = contract(BREMContract)
    brem.setProvider(currentProvider)
    const bremInstance = await brem.deployed()


    // Convert values
    const ethRaised = utils.fromWei(weiRaised.toFixed(), "ether")
    const capEth = utils.fromWei(capWei.toFixed(), "ether")
    const ethInvested = utils.fromWei(weiInvested.toFixed(), "ether")
    const tokenBalance = utils.fromWei(tokenBalanceRaw.toFixed(), "ether")
    // projectId: '',
    // project: {
    //   name: 'Neva Towers',
    //   company: 'Company name',
    //   address: 'Ul pidora, 12c4',
    //   closingDate: moment(),
    //   myEmphasis: 50, // How much eth was given to project by user
    //   softCap: 1000,
    //   status: 'Opened',
    //   icoBalance: 400,
    //   withDrawRequest: 0,
    //   goal: 2000,
    //   auditorsList: [],
    //   description: '<div style="color: green;">HTML EEEE</div>',
    //   latLng: {
    //     lat: 2, lng: 50
    //   }
    // }
    console.log('SERVER:',serverData)
    console.log('INSTANCE:', icoInstance)
    const { fee_percent: feePercent, location, loc_address, developer, description } = serverData
    const project = {
      feePercent,
      symbol,
      description,
      name,
      company: developer.username,
      softCap: capEth,
      icoBalance,
      status,
      closingTime: moment(new Date(closingTimeEpochs.toNumber() * 1000)),
      docsUrl: 'https://ipfs.infura.io/ipfs/' + docHash,
      myEmphasis: ethInvested,
      raised: ethRaised,
      latLng: JSON.parse(location || '{}'),
      loc_address,
      auditorsList: []
    }
    console.log('VIEW:', this)

    this.setState({
      project,
      /*

      name: name,
      tokenAddress: tokenAddress,
      icoBalance: utils.fromWei(icoBalance, "ether"),
      wallet: wallet,
      rate: rate.toNumber(),
      ethRaised: ethRaised,
      capEth: capEth,
      docsURL: 'https://ipfs.infura.io/ipfs/' + docHash,
      closingTime: new Date(closingTimeEpochs.toNumber() * 1000).toString(),
      ethInvested: ethInvested,
      symbol: symbol,
      tokenBalance: tokenBalance

      */
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


    // TODO: Асинхронных цикл
    // Read current auditors
    const auditorsAmount = await icoInstance.auditorsAmount()
    const amount = auditorsAmount.toNumber()
    const auditors = [amount]
    const auditorsList = await auditors.map((val, i) => {
      icoInstance.getAuditor(i).then(auditorAddress => {
        return auditorAddress
        // Uncomment with new contracts
        // const username = await bremInstance.getUsername(auditorAddress)
      })
      .catch(err => console.error(err))
    })
    console.log(auditorsList)
    this.setState({
      auditorsList: auditorsList
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
      const requestedValue = utils.fromWei(request[0].toFixed(), "ether")
      this.setState({
        requestedValue: requestedValue
      })
    }

    if (this.state.isSuccess && isDeveloper) {
      const withdrawFeePercent = await icoInstance.withdrawFeePercent()
      this.setState({
        currentICOFee: withdrawFeePercent.toNumber() // Показывать при запросе вывода застройщиком
      })
    }

    if (isAuditor) {
      const isConfirmed = await icoInstance.isConfirmed(coinbase)
      this.setState({
        isConfirmed: isConfirmed
      })
    }
    console.log(store)
  }

  depositETH = async e => {
    // Здесь нужна отдельная форма с расчетом количества токенов
    e.preventDefault()

    // const depositValue = this.state.depositValue
    const depositValue = prompt('Please enter amount of wei you want to deposit')
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
    if (hasClosed) {
      if (hasClosed) {
        // #alertaction [ invest ]
        alert('Error, ico has closed')
        return
      }
    }
    const auditSelected = await icoInstance.auditSelected()
    if (!auditSelected) {
      // #alertaction [ invest ]
      alert('Error, ico not opened')
      return
    }

    // Invest
    const depositInWei = utils.toWei(depositValue, "ether")
    try {
      const txRes = await icoInstance.sendTransaction({from: coinbase, value: depositInWei})
      const isSuccessTX = txRes.receipt.status === "0x1"
      const tx = txRes.tx;
      alert(tx)
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

  refundETH = async e => {
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
    if (!hasClosed || capReached || weiInvested <= 0) {
      alert('You dont have value to refund')
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
    // Кнопка доступна только если isReqeusted && !isConfirmed
    e.preventDefault()

    await getWeb3
    const { web3Instance, web3Account } = store
    const { currentProvider, utils, eth } = web3Instance

    const coinbase = await eth.getCoinbase()

    const ico = contract(ICOContract)
    ico.setProvider(currentProvider)
    const icoInstance = await ico.at(this.state.projectId)

    // Check for auditor
    const isAuditor = await icoInstance.isAuditor(coinbase)
    if (!isAuditor) {
      // Ошибка
       return
    }

    // Check for status
    const isRequested = await icoInstance.isRequested()
    const isConfirmed = await icoInstance.isConfirmed(coinbase)
    if (!isRequested || isConfirmed) {
      // Ошибка, нельзя купить
      return
    }

    // Confirm
    try {
      const txRes = await icoInstance.confirmWithdraw({from: coinbase})
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

  devWithdrawETH = async e => {
        // Кнопка доступна только если isSuccess && !isWithdrawn && wallet == coinbase && !isRequested
        e.preventDefault()

        const withdrawValue = prompt('Please enter amount of wei you want to deposit')
      
        const { host, authConfig } = config
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
        let requestValue = utils.toWei(withdrawValue, "ether")
        const contractBalance = await eth.getBalance(icoInstance.address)
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
            axios.put(host + '/ico/request', {
              address: icoInstance.address
            }, authConfig)

          } else {
            // Ошибка, можно посмотреть на etherscan https://rinkeby.etherscan.io/tx/ + tx
          }
        } catch(err) {
          console.error(err)
        }
  }

  superAddAuditor = async () => {
    const auditorKey = this.newAuditorAddress.value

    await getWeb3
    const { host, authConfig } = config
    const { web3Instance, web3Account } = store
    const { currentProvider, utils, eth } = web3Instance

    const coinbase = await eth.getCoinbase()

    if (!utils.isAddress(auditorKey)) {
      alert(auditorKey + 'is not address')
      return
    }

    const brem = contract(BREMContract);
    brem.setProvider(currentProvider)
    const bremInstance = await brem.deployed()

    // Check for superuse
    const isSuperuser = await bremInstance.isSuperuser(coinbase)
    if (!isSuperuser) {
      alert('Is not superuser')
      return
    }

    // Check for global auditor
    const isAuditor = await bremInstance.isAuditor(auditorKey)
    if (!isAuditor) {
      alert(auditorKey + ' is not auditor')
      return
    }

    const ico = contract(ICOContract)
    ico.setProvider(currentProvider)
    const icoInstance = await ico.at(this.state.projectId)

    // Check for status
    const hasClosed = await icoInstance.hasClosed()
    const isCurrentAuditor = await icoInstance.isAuditor(auditorKey)
    if (hasClosed || isCurrentAuditor) {
      console.log('Error')
      return
    }

    // Add auditor to current ico
    try {
      const txRes = await icoInstance.addAuditor(auditorKey, {from: coinbase})
      const isSuccessTX = txRes.receipt.status === "0x1"
      const tx = txRes.tx;
      if (isSuccessTX) {
        // Все ок, можно посмотреть на etherscan https://rinkeby.etherscan.io/tx/ + tx
        // Предложить добавить адрус токена в metamask
        const res = await axios.post(host + 'super/ico/audit', {
          ico: {
            address: icoInstance.address
          },
          auditor: {
            address: auditorKey
          }
        }, authConfig)
        console.log(res)
      } else {
        // Ошибка, можно посмотреть на etherscan https://rinkeby.etherscan.io/tx/ + tx
      }
    } catch(err) {
      console.error(err)
    }
  }

  superOpenICO = async e => {
    e.preventDefault()

    await getWeb3
    const { host, authConfig } = config
    const { web3Instance } = store
    const { currentProvider, eth } = web3Instance

    const coinbase = await eth.getCoinbase()

    const brem = contract(BREMContract);
    brem.setProvider(currentProvider)
    const bremInstance = await brem.deployed()

    // Check for superuser
    const isSuperuser = await bremInstance.isSuperuser(coinbase)
    if (!isSuperuser) {
      // #alertaction [ open ICO ]
      alert('Error, you are not superuser')
      return
    }

    const ico = contract(ICOContract)
    ico.setProvider(currentProvider)
    const icoInstance = await ico.at(this.state.projectId)

    // Check for status
    const hasClosed = await icoInstance.hasClosed()
    if (hasClosed) {
      // #alertaction [ open ICO]
      alert('Error, ico has closed')
      return
    }
    const auditSelected = await icoInstance.hasClosed()
    if (auditSelected) {
      // #alertaction [ openICO ]
      alert('You already opened ico')
      return
    }
    const auditorsAmount = await icoInstance.auditorsAmount()
    if (auditorsAmount.toNumber() === 0) {
      // #alertaction [ Open ICO ]
      alert('Not enaught auditors')
      return
    }

    // Open ICO
    try {
      const txRes = await icoInstance.finishAuditorSelection({from: coinbase})
      const isSuccessTX = txRes.receipt.status === "0x1"
      const tx = txRes.tx;
      if (isSuccessTX) {
        // Все ок, можно посмотреть на etherscan https://rinkeby.etherscan.io/tx/ + tx
        // Предложить добавить адрус токена в metamask
        const res = await axios.put(host + 'super/ico/open', {
          address: icoInstance.address
        }, authConfig)
        console.log(res)
      } else {
        // Ошибка, можно посмотреть на etherscan https://rinkeby.etherscan.io/tx/ + tx
      }
    } catch(err) {
      console.error(err)
    }
  }

  render() {
    const {
      isAuditor,
      isCreated,
      isDeveloper,
      isFailed,
      isOpened,
      isOverdue,
      isRequested,
      isSuccess,
      isSuperuser,
      isWithdrawn,
      project
    } = this.state

    const {
      projectId,
      name = 'Loading...',
      company = '',
      address = '',
      closingDate = '',
      myEmphasis = 0,
      softCap = 0,
      status = '',
      icoBalance = 0,
      withDrawRequest = false,
      goal = 0,
      auditorsList = [],
      description = ''
    } = project



    console.log(this.state)

    const html = {__html: description}
    return (
      <div className={css(style.main)}>
        <section className={css(style.top)}>
          <div className={css(style.topLeft, style.topPart)}>
            <img src={this.state.img} />
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
              {isOpened && ( 
                <button style={{marginRight: 25}} onClick={this.depositETH}>
                 Deposit ETH
                </button>
              )}
              {isFailed && (
                <button onClick={this.refundETH}>
                  Refund ETH
                </button>
              )}
            </div>
          </div>
        </section>
        <section>
          <p>Project ID: { projectId }</p>
          {
              isSuperuser && isCreated && (
                <div>
                  <div>
                    <input ref={elem => this.newAuditorAddress = elem} type="text" placeholder="Auditor address"/>
                    <button onClick={this.superAddAuditor}>
                      Add auditor
                    </button>
                  </div>
                  <div>
                  <button onClick={this.superOpenICO}>
                    Open ICO
                  </button>
                </div>
               </div>
              )
          }
          {
            isAuditor && isRequested && this.state.isConfirmed === false && (
              <div>
                <button onClick={this.auditorConfirmRequest}>
                  Confirm Request
                </button>
              </div>
            )
          }
          {
            isDeveloper && isSuccess && (
              <button onClick={this.devWithdrawETH}>
                Withdraw eth
              </button>
            )
          }
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





      </div>
    )
  }
}

export default View
