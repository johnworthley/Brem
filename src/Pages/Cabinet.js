import React, { Component } from 'react'
import { css, StyleSheet } from 'aphrodite/no-important'
import { Link } from 'react-router-dom'

import Tabs from '../Components/Tabs'
import Table from '../Components/Table'

import getWebThree from '../util/getweb3'
import store from 'Store'
import contract from 'truffle-contract'
import BREMContract from "../../build/contracts/BREM.json"
import ICOContract from '../../build/contracts/BREMICO.json'
import config from 'Config'
import axios from 'axios'

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
		width: 'calc(50% - 50px)',
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
		minHeight: 192,
		display: 'flex',
		justifyContent: 'space-between',
		flexDirection: 'column',
	},
	input: {
		height: 35,
    borderRadius: 3,
    border: 'solid 1px lightgray',
    padding: '0 15px',
		marginRight: 15
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
			data : new Map(
				// ['First row', ['Row 1 first value','Second one','And last']],
				// ['Second row', [10,33,30]],
				// ['Last row', [1101.786,44444.99,'550.2']]
			)
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
    const bremICO = contract(ICOContract)
    bremICO.setProvider(currentProvider)

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

      // Get all created ICOs
      try {
      const res = await axios.get(host + 'super/ico', authConfig)
      const icos = res.data
     
      const { struct } = {...this.state}
      struct.data = icos.reduce((map, ico) => {
        map.set('Name', [...(map.get('Name') || []), ico.name])
        map.set('Address', [...(map.get('Address') || []), <Link to={`/project/${ico.address}`}>{ico.address}</Link>])
        map.set('Developer', [...(map.get('Developer') || []), ico.developer.username])
        map.set('Deadline', [...(map.get('Deadline') || []), new Date(ico.closing_time).toString()])
        map.set('Status', [...(map.get('Status') || []), ico.status])
					return map
				}, new Map()
			)
			this.setState({
				struct
			})
      } catch(err) {
        console.error(err)
      }
    }

    const isAuditor = await bremInstance.isAuditor(coinbase)
    console.log(isAuditor)
    if (isAuditor) {
      // Get all to ico to confirm
      const res = await axios.get(host + 'audit/ico', authConfig)
      const icos = res.data
      const { struct } = {...this.state}
      const icosWithWeb3 = await Promise.all(icos.map(ico => new Promise(async resolve => {
        try {
          const icoInstance = await bremICO.at(ico.address)
				  const isConfirmed = await icoInstance.isConfirmed(coinbase)
				  resolve({
            isConfirmed,
            name: ico.name,
            symbol: ico.symbol,
            address: ico.address,
            closing_time: ico.closing_time,
            status: ico.status
				  })
          resolve(ico)
        } catch(err) {
          console.error(err)
          resolve({
					  isConfirmed: true
				  })
          resolve(ico)
        } 
			})))
      struct.data = icosWithWeb3.reduce((map, ico) => {
        if (!ico.isConfirmed) {
          map.set('Name', [...(map.get('Name') || []), ico.name])
          map.set('Address', [...(map.get('Address') || []), <Link to={`/project/${ico.address}`}>{ico.address}</Link>])
          map.set('Developer', [...(map.get('Developer') || []), ico.developer.username])
          map.set('Deadline', [...(map.get('Deadline') || []), new Date(ico.closing_time).toString()])
          map.set('Status', [...(map.get('Status') || []), ico.status])
          return map
        }
				}, new Map()
			)
			this.setState({
				struct
			})
    }

    if (!isAuditor && !isSuperuser) {
      // Get current dev ico's
      const res = await axios.get(host + 'dev/ico', authConfig)
      const icos = res.data //  Read cap and raised for all icos (example in ico form)
			// console.log(icos)
			const { struct } = {...this.state}
			const icosWithWeb3 = await Promise.all(icos.map(ico => new Promise(async resolve => {
        try {
				  const icoInstance = await bremICO.at(ico.address)
				  const weiRaised = await icoInstance.weiRaised()
				  const capWei = await icoInstance.cap()

				  // Convert values
				  const ethRaised = utils.fromWei(weiRaised.toFixed(), "ether")
				  const capEth = utils.fromWei(capWei.toFixed(), "ether")
				  resolve({
					  ethRaised,
            capEth,
            name: ico.name,
            symbol: ico.symbol,
            address: ico.address,
            closing_time: ico.closing_time,
            status: ico.status
				  })
          resolve(ico)
        } catch(err) {
          console.error(err)
          resolve({
					  ethRaised: 'NA',
            capEth: 'NA',
            name: ico.name,
            symbol: ico.symbol,
            address: ico.address,
            closing_time: ico.closing_time,
            status: ico.status
				  })
          resolve(ico)
        } 
			})))
			struct.data = icosWithWeb3.reduce((map, ico) => {
        map.set('Name', [...(map.get('Name') || []), ico.name])
        map.set('Address', [...(map.get('Address') || []), <Link to={`/project/${ico.address}`}>{ico.address}</Link>])
        map.set('Token', [...(map.get('Token') || []), ico.symbol])
        map.set('Cap', [...(map.get('Cap') || []), ico.capEth])
        map.set('Collected', [...(map.get('Collected') || []), ico.ethRaised])
        map.set('Deadline', [...(map.get('Deadline') || []), new Date(ico.closing_time).toString()])
        map.set('Status', [...(map.get('Status') || []), ico.status])
					return map
				}, new Map()
			)
			this.setState({
				struct
			})
		}
}

  withdrawFees = async e => {
    e.preventDefault()
    const withdrawValue = this.withdrawFees.value; // Значение в Ether

    await getWebThree()
    const { web3Instance } = store
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
        alert('Withdraw fees: done')
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
				alert('Change fee: done')
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

    const { web3Instance } = store
    const { currentProvider, utils, eth } = web3Instance

    if (!utils.isAddress(auditor)) {
      return alert("Wrong auditor name")
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
        alert('addAuditors: done')
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
		const { struct = false } = this.state
		const { accountType } = store
		const data = []
		if(accountType === 'developer') data.push({
			name: 'Your projects',
			default: true,
			component: () => (
				<div>
					<div className={css(style.projectsTop)}>
						<Link to="/project/new">
							<button className={css(style.button)} style={{marginRight: 25}} onClick={this.addNewProject}>
								Add new
							</button>
						</Link>
						{/* <button className={css(style.button)} onClick={this.loadMoreProjects}>
							Load more
						</button> */}
					</div>
					<Table struct={struct} />
				</div>
			)
		})
		if(accountType !== 'developer') data.push({
			name: 'Settings',
			default: true,
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
