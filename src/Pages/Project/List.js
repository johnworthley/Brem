import React, { Component } from 'react'
import { css, StyleSheet } from 'aphrodite/no-important'
import axios from 'axios'
import config from 'Config'
import contract from 'truffle-contract'
import ICOContract from "../../../build/contracts/BREMICO.json"
import TokenContract from "../../../build/contracts/BREMToken.json"
import BREMContract from "../../../build/contracts/BREM.json"


import ProjectCard from '../../Components/ProjectCard'

import store from 'Store'

const style = StyleSheet.create({
	more: {
		textAlign: 'center',
		marginTop: 50,
		marginBottom: 60,
		width: 944
	},
	moreButton: {
		width: 232,
		height: 56,
		background: 'transparent',
		color: '#1e1e1e',
		fontSize: 13,
		border: 'solid 2px #d5d5d6',
		borderRadius: 3,
		':hover': {
			borderColor: '#f0723b',
			background: 'transparent'
		}
	}
})

class List extends Component {
	state = {
		page: 0,
		showMore: true,
		cards: [],
		loading: true
		//CARDS
		/*
		{
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
		}
		*/
	}
	getPageData = async () => {
		const { host } = config
		const { page } = this.state
		const storedCards = localStorage.getItem('icolistcards')
		console.log(storedCards)
		if(storedCards) this.setState({
			cards: JSON.parse(storedCards),
			loading: false
		})
		const res = await axios.get(`${host}ico/all?page=${this.state.page}`)
		const { data } = res
		console.log(data, this)

		const cards = await Promise.all(data.map(card => new Promise(async resolve => {
		const { web3Instance } = store
	    const { currentProvider, utils, eth } = web3Instance
	    const coinbasePromise = eth.getCoinbase()
	    const ico = contract(ICOContract)
		ico.setProvider(currentProvider)

	    const icoInstancePromise = ico.at(card.address)
		const icoBalancePromise = eth.getBalance(card.address)
		const [icoInstance, icoBalance, coinbase] = await Promise.all([icoInstancePromise, icoBalancePromise, coinbasePromise])

	    const tokenAddressPromise = icoInstance.token()
	    // const wallet = await icoInstance.wallet()
	    const ratePromise = icoInstance.rate()
	    const weiRaisedPromise = icoInstance.weiRaised()
	    const capWeiPromise = icoInstance.cap()
	    const docHashPromise = icoInstance.docHash()
	    const closingTimeEpochsPromise = icoInstance.closingTime()
		const weiInvestedPromise = icoInstance.getBalance(coinbase, {from: coinbase})
		const [ tokenAddress, rate, weiRaised, capWei, docHash, closingTimeEpochs, weiInvested ] = await Promise.all([
			tokenAddressPromise,
			ratePromise,
			weiRaisedPromise,
			capWeiPromise,
			docHashPromise,
			closingTimeEpochsPromise,
			weiInvestedPromise
		])
		
	    const token = contract(TokenContract)
	    token.setProvider(currentProvider)
	    const tokenInstance = await token.at(tokenAddress)

	    const np = tokenInstance.name()
	    const sp = tokenInstance.symbol()
		const tbrp = tokenInstance.balanceOf(coinbase)
		
		const [ name, symbol, tokenBalanceRaw ] = await Promise.all([np, sp, tbrp])

	    const brem = contract(BREMContract)
	    brem.setProvider(currentProvider)
	    const bremInstance = await brem.deployed()


	    // Convert values
	    const ethRaised = utils.fromWei(weiRaised.toFixed(), "ether")
	    const capEth = utils.fromWei(capWei.toFixed(), "ether")
	    const ethInvested = utils.fromWei(weiInvested.toFixed(), "ether")
	    const tokenBalance = utils.fromWei(tokenBalanceRaw.toFixed(), "ether")
			resolve({
				icoBalance,
				tokenAddress,
				rate,
				capWei,
				closingTimeEpochs,
				weiInvested,
				name,
				symbol,
				ethRaised,
				capEth,
				ethInvested,
				tokenBalance,
				//docHash,
				address: card.address
			})
		})))
		this.setState({
			cards: [...this.state.cards, ...cards.map(item => {
				return {
					...item,
					payment: item.ethInvested,
					progress: (item.ethRaised / item.capEth) * 100,
					limit: item.capEth,
					imageUrl: `${host}ico/image?address=${item.address}`,
				}
			})],
			page: this.state.page + 1,
			showMore: cards.length,
			loading: false
		})
		localStorage.setItem('icolistcards', JSON.stringify(this.state.cards))
	}
	componentDidMount = async () => {
		// const { host } = config
		store.update({
			currentLocation: 'marketplace'
		})
		this.getPageData()
	}

	render() {
	const { cards = [], showMore = false, loading } = this.state
    return (
      <div>
		<h3>All projects</h3>

		{
			loading ? (
				<p>Loading...</p>
			) : cards.length ? (
				cards.map((item, num) => <ProjectCard struct={item} key={'ProjectCard' + num} />)
			) : 'No projects here. Create one to be first!'
		}

		{showMore ? (
			<div className={css(style.more)}>
				<button onClick={this.getPageData} className={css(style.moreButton)}>Load more</button>
			</div>
		) : ''}

      </div>
    )
  }
}

export default List
