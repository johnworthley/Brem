import React, { Component } from 'react'
import { css, StyleSheet } from 'aphrodite/no-important'

import Table from '../Components/Table'
import List from './Project/List'

import getWebThree from '../util/getweb3'
import store from 'Store'
import contract from 'truffle-contract'
import BREMContract from "../../build/contracts/BREM.json"

const style = StyleSheet.create({
	main: {
		marginBottom: 75
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
		},
		cards : {
			showMore: true,
			data: [{
				name: 'Каждой сове по дуплу!',
				company: 'ООО Ромашка',
				adress: 'Москва, 1-й Красногвардейкий проезд, участок 17-18',
				payment: 50000,
				progress: 50,
				limit: 500000,
				imageUrl: 'https://avatars.mds.yandex.net/get-pdb/199965/9baf77ea-a933-45ac-9cbd-f820d8e77a96/s1200'
			},{
				name: 'abc',
				company: 'Daisy Ltd',
				adress: '10 Downing Street Westminster London',
				payment: 100,
				progress: 25,
				limit: 1000,
				imageUrl: 'https://avatars.mds.yandex.net/get-pdb/939428/27374061-ba26-4c33-8b64-7bb569349414/s1200'
			}]
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
		const { struct, cards } = this.state
		return (
			<div className={css(style.main)}>
				<h3>Cabinet</h3>
				<div style={{marginBottom: 100}}>
					<form onSubmit={this.changeFee} style={{marginBottom: 10}}>
						<input required ref={elem => this.fee = elem} type="text" placeholder="Fee" />
						<button>Change fee</button>
					</form>
					<form onSubmit={this.addAuditors}>
						<input required ref={elem => this.auditors = elem} type="text" placeholder="Auditors" />
						<button>Add auditors</button>
					</form>
				</div>

				<Table struct={struct} />
				
				<List struct={cards} />
			</div>
		)
	}
}
