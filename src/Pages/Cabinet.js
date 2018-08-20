import React, { Component } from 'react'
import { css, StyleSheet } from 'aphrodite/no-important'

import Tabs from '../Components/Tabs'
import Table from '../Components/Table'

import getWebThree from '../util/getweb3'
import store from 'Store'
import contract from 'truffle-contract'
import BREMContract from "../../build/contracts/BREM.json"
import config from 'Config'
import axios from 'axios';

const style = StyleSheet.create({
	main: {
		marginBottom: 75
	},
	settingsTop: {
		display: 'flex',
		justifyContent: 'space-between',
		width: '100%'
	},
	settingsTopPart: {
		width: 'calc(50% - 15px)',
		minHeight: 150,
		display: 'flex',
		justifyContent: 'space-between',
		flexDirection: 'column',
	},
	settingsTopLeft: {
		border: 'solid 1px lightgray',
		borderRadius: 4,
		background: '#fff',
		padding: 20
	},
	settingsTopForm: {
		minHeight: 150,
		display: 'flex',
		justifyContent: 'space-between',
		flexDirection: 'column',
	},
	input: {
		height: 35,
    borderRadius: 3,
    border: 'solid 1px lightgray',
    padding: '0 15px'
	},
	button: {
		width: 150,
		height: 35
	},
	projectsTop: {
		display: 'flex',
		marginBottom: 40
	}
})

export default class Cabinet extends Component {
	state = {
		struct: {
			options: {
				columns : {
					'Second row' : {
						red: true
					}
				}
			},
			data : new Map([
				['First row', ['Row 1 first value','Second one','And last']],
				['Second row', [10,33,30]],
				['Last row', [1101.786,44444.99,'550.2']]
			])
		}
  }

	componentDidMount = async () => {
    store.update({
		  currentLocation: 'cabinet'
    })

    const { host, authConfig } = config
    const { web3Instance, web3Account } = store
    const { currentProvider, utils, eth } = web3Instance
    const brem = contract(BREMContract)
    brem.setProvider(currentProvider)

    const coinbase = await eth.getCoinbase()

    const bremInstance = await brem.deployed()

    const isSuperuser = await bremInstance.isSuperuser(coinbase)
    console.log(isSuperuser)
    if (isSuperuser) {
      const factoryAddress = brem.address
      console.log(factoryAddress)

      const collectedWei = await eth.getBalance(brem.address)
      console.log(collectedWei)
      const collectedEth = utils.fromWei(collectedWei, "ether")
      console.log(collectedEth)

      const feePercentBN = await bremInstance.withdrawFeePercent()
      const feePercent = await feePercentBN.toNumber()
      console.log(feePercent)

      // Get all auditors
      const res = await axios.get(host + 'super/audit', authConfig)
      const auditors = res.data 
    }

    const isAuditor = await bremInstance.isAuditor(coinbase)
    console.log(isAuditor)
    if (isAuditor) {
      // Get all to ico to confirm
      const res = await axios.get(host + 'audit/ico')
      // Check each for isConfirmed(coinbase) on ico instance
    }

    if (!isAuditor && !isSuperuser) {
      // Get current dev ico's
      const res = await axios.get(host + '/dev/ico')
      const icos = res.data //  Read cap and raised for all icos (example in ico form)
    }
}

  withdrawFees = async e => {
    e.preventDefault()
    const withdrawValue = this.withdrawFees.value; // Значение в Ether

    await getWebThree()
    const { web3Instance, web3Account } = store
    const { currentProvider, utils, eth } = web3Instance
    const brem = contract(BREMContract)
    brem.setProvider(currentProvider)

    const coinbase = await eth.getCoinbase()
    // TODO: check coinbase with state

    const bremInstance = await brem.deployed()
    const isSuperuser = await bremInstance.isSuperuser(coinbase)
    if (!isSuperuser) {
      // Ошибка, пользователь не админ
      return
    }

    const withdrawValueInWei = utils.toWei(withdrawValue, "ether")

    // Check for contract's balance
    const currentBalanceBN = await eth.getBalance(brem.address)
    const currentBalance = currentBalanceBN.toNumber()
    if (withdrawValueInWei > currentBalance) {
      // Нет такой суммы на контракте
      return
    }
    // Withdraw collected fee
    try {
      const txRes = await bremInstance.withdrawFees(withdrawValueInWei, {from: coinbase})
      const tx = txRes.tx
      const status = txRes.receipt.status
      if (status === "0x1") {
        alert(tx)
        // Можно обновить текущее значение баланса контракта из web3d
      } else{
        alert("Error tx")
      }
    } catch(err) {
      console.error(err)
    }
  }

	changeFee = async e => {
		e.preventDefault()
    const fee = this.fee.value
    if (fee < 0 || fee > 100) {
      return alert('Error')
    }

    await getWebThree()
    const { web3Instance, web3Account } = store
    const { currentProvider, utils, eth } = web3Instance
    const brem = contract(BREMContract)
    brem.setProvider(currentProvider)

    const coinbase = await eth.getCoinbase()
    // TODO: check coinbase with state

    const bremInstance = await brem.deployed()
    const isSuperuser = await bremInstance.isSuperuser(coinbase)
    if (!isSuperuser) {
      // Ошибка, пользователь не админ
      return
    }
    // Check for contract's current fee value
    const currentFeeBN = await bremInstance.withdrawFeePercent()
    const currentFee = currentFeeBN.toNumber()
    if (currentFee === fee) {
        // Значение уже установлено
        return
    }
    // Change fee
    try {
      const txRes = await bremInstance.setWithdrawFeePercent(fee, {from: coinbase})
      const tx = txRes.tx
      const status = txRes.receipt.status
      if (status === "0x1") {
        alert(tx)
        // Можно обновить текущее значение из web3
      } else{
        alert("Error tx")
      }
    } catch(err) {
      console.error(err)
    }
	}

	addNewProject = async () => {

	}

	loadMoreProjects = async () => {

	}

	addAuditors = async e => {
		e.preventDefault()
    const auditor = this.auditors.value

    await getWebThree()
    const { web3Instance, web3Account } = store
    const { currentProvider, utils, eth } = web3Instance

    if (!utils.isAddress(auditor)) {
      return alert("Error")
    }

    const brem = contract(BREMContract)
    brem.setProvider(currentProvider)

    const coinbase = await eth.getCoinbase()
    // TODO: check coinbase with state

    const bremInstance = await brem.deployed()
    const isSuperuser = await bremInstance.isSuperuser(coinbase)
    if (!isSuperuser) {
      // Ошибка, пользователь не админ
      return
    }

    // Check for auditor
    const isAuditor = bremInstance.isAuditor(auditor)
    if (isAuditor) {
      // Уже аудитор
      return
    }
    // Add audtor
    try {
      const txRes = await bremInstance.addAuditor(auditor, {from: coinbase})
      const tx = txRes.tx
      const status = txRes.receipt.status
      if (status === "0x1") {
        alert(tx)
        // Сообщение, что аудитор появиться в списке после регистрации
      } else{
        alert("Error tx") // Можно добавить ссылку на etherscan
      }
    } catch(err) {
      console.error(err)
    }
	}

	render = () => {
		//const { struct } = this.props
		const { struct } = this.state
		const { accountType } = store
		const data = [
			{
				name: 'Your projects',
				default: true,
				component: () => (
					<div>
						<div className={css(style.projectsTop)}>
							<button className={css(style.button)} style={{marginRight: 25}} onClick={this.addNewProject}>
								Add new
							</button>
							<button className={css(style.button)} onClick={this.loadMoreProjects}>
								Load more
							</button>
						</div>
						<Table struct={struct} />
					</div>
				)
			}
		]
		if(accountType !== 'developer') data.push({
			name: 'Settings',
			component: () => (
				<div>
					<div className={css(style.settingsTop)} style={{marginBottom: 100}}>
						<div className={css(style.settingsTopPart, style.settingsTopLeft)}>
							<div>
								<div>
									Factory address
								</div>
								<div>
									qwkeqlwkqlkwmelqwe
								</div>
							</div>
							<div>
								Ether collected 1000 ETH
							</div>
						</div>
						<div className={css(style.settingsTopPart)}>
							<div className={css(style.settingsTopForm)}>
								<form onSubmit={this.changeFee} style={{marginBottom: 10}}>
									<input className={css(style.input)} required ref={elem => this.fee = elem} type="text" placeholder="Fee" />
									<button className={css(style.button)}>Change fee</button>
								</form>
								<form onSubmit={this.addAuditors}>
									<input className={css(style.input)} required ref={elem => this.auditors = elem} type="text" placeholder="Auditors" />
									<button className={css(style.button)}>Add auditors</button>
								</form>
							</div>
						</div>
					</div>

					<Table struct={struct} />
				</div>
			)
		})
		return (
			<div className={css(style.main)}>
				<h3>Cabinet</h3>
				<Tabs maxWidth="100%" data={data} />
			</div>
		)
	}
}
