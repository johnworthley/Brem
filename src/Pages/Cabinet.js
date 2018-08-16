import React, { Component } from 'react'
import { css, StyleSheet } from 'aphrodite/no-important'

import Table from '../Components/Table'

export default class Cabinet extends Component {
	state = {
		struct: {
			data : new Map([
				['a', [1,2,3]],
				['b', [10,'33',30]]
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