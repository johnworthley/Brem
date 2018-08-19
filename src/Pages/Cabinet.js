import React, { Component } from 'react'
import { css, StyleSheet } from 'aphrodite/no-important'

import Table from '../Components/Table'

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
		}
	}

	changeFee = e => {
		e.preventDefault()
		const fee = this.fee.value
		console.log(fee)
	}
	addAuditors = e => {
		e.preventDefault()
		const auditors = this.auditors.value
		console.log(auditors)
	}

	render = () => {
		//const { struct } = this.props
		const { struct } = this.state
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
			</div>
		)
	}
}
