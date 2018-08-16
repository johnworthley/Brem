import React, { Component } from 'react'
import { css, StyleSheet } from 'aphrodite/no-important'

export default class Table extends Component {
	render = () => {
		//const { struct } = this.props
		const { struct } = this.props
		return (
			<div qwe="qe">
			{
				[10, 20, 30].map(item => (
					<div>
						{item}
					</div>
				))
			}
			</div>
		)
	}
}