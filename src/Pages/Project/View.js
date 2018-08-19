import React, { Component } from 'react'
import getWeb3 from '../../util/getweb3'
import store from 'Store'
import config from 'Config'
import contract from 'truffle-contract'
import ICOContract from "../../../build/contracts/BREMICO.json"
import TokenContract from "../../../build/contracts/BREMToken.json"
import BREMContract from "../../../build/contracts/BREM.json"
import axios from 'axios';

class View extends Component {
  state = {
    projectId: '',
    project: {}
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

    this.setState({
      name: name,
      tokenAddress: tokenAddress,
      icoBalance: utils.fromWei(icoBalance, "ether"),
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

    const brem = contract(BREMContract)
    brem.setProvider(currentProvider)
    const bremInstance = await brem.deployed()

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
    const hasClosed = icoInstance.hasClosed()
    const auditSelected = icoInstance.auditSelected()
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
    const hasClosed = icoInstance.hasClosed()
    const capReached = icoInstance.capReached()
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
        const hasClosed = icoInstance.hasClosed()
        const capReached = icoInstance.capReached()
        const isRequested = icoInstance.isRequested()
        const isWithdrawn = icoInstance.isWithdrawn()
        let requestValue = utils.toWei(requestEthValue, "ether")
        const contractBalance = eth.getBalance(ico.address)
        if (!hasClosed || !capReached || isRequested || isWithdrawn || contractBalance < requestValue) {
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
    const { projectId } = this.state
    return (
      <div>
        Project ID: { projectId }
        <button onClick={this.depositETH}>
           Deposit ETH
        </button>
        <button onClick={this.withdrawETH}>
          Withdraw ETH
        </button>
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
