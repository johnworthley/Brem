import React, { Component } from 'react'
import { css, StyleSheet } from 'aphrodite/no-important'

const style = StyleSheet.create({
  table : {
    borderCollapse: '1px #ff0000 solid'
  },
  th : {
    border: '1px #ff0000 solid'
  },
  td : {
    border: '1px #999999 solid'
  }
})

export default class Table extends Component {
	render = () => {
		const { struct } = this.props
		const headers = []
		const rows = []
		let colNum = 0;
		
		struct.data.forEach(((column,key) => {
			headers.push(key)
			column.map((value,rowNum) => {
				rows[rowNum] ? false : rows[rowNum]=[]
				rows[rowNum][colNum] = value
			})
			colNum++
		}))

		return (
		<div>
			{
			struct && struct.data ? 

				<table className={css(style.th)}>
				<tbody>
					<tr>
					{
						headers.map((item,thNum) => (
							<th key={'TableHead'+thNum} className={css(style.th)}>
								{item}
							</th>
						))
					}
					</tr>
					{
						rows.map((item,rowNum) => (
							<tr key={ 'TableRow' + rowNum}>
								{
									item.map((item,cellNum) => (
										<td key={ 'TableCell' + cellNum + rowNum} className={css(style.td)}>
											{item}
										</td>
									))
				
								}
							</tr>
						))
					}
				</tbody>
				</table>				
			: 
				'No data found'
			}

		</div>
		)
	}
}