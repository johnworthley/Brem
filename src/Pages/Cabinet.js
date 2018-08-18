import React, { Component } from 'react'
import { css, StyleSheet } from 'aphrodite/no-important'

import Table from '../Components/Table'

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
				['Last row', [101,444,'550']]
			])
		}
	}
	render = () => {
		//const { struct } = this.props
		const { struct } = this.state
		return (
			<div>
				Cabinet
				
				<Table struct={struct} />
			</div>
		)
	}
}