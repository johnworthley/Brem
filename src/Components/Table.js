import React, { Component } from 'react'
import { css, StyleSheet } from 'aphrodite/no-important'

const style = StyleSheet.create({
  table : {
    borderCollapse: 'collapse',
	  boxSizing: 'border-box',
	  width: '100%',
	  fontSize: 14
  },
  th : {
  	height: 51,
  	boxSizing: 'border-box',
  	backgroundColor: '#f0723b',
  	color: '#ffffff',
  	textAlign: 'left',
  	padding: '0 20px',
  	fontWeight: 'normal'
  },
  td : {
    borderColor: '#e5e5e5',
  	borderStyle: 'solid',
  	borderWidth: '0 0 1px 0',
  	boxSizing: 'border-box',
  	textAlign: 'left',
  	height: 31,
  	padding: '0 20px',
  	':first-child': {
  		borderLeftWidth: 1
  	},
  	':last-child': {
  		borderRightWidth: 1
  	}
  },
  tdRight : {
  	textAlign: 'right'
  },
  tdRed : {
  	color: '#ff0000'
  },
  tr: {
	  background: '#ffffff',
	  ':hover': {
	    background: 'rgba(255, 128, 59, 0.33)'
    }
  },
  trEven : {
    backgroundColor: '#f8f8f8'
  },

})

export default class Table extends Component {
	render = () => {
		const { struct = {} } = this.props
		const { options = {} } = struct
		const { columns = {} } = options
		const headers = []
		const rows = []
		let colNum = 0

		struct.data.forEach(((column,key) => {
			headers.push(key)
			column.forEach((value,rowNum) => {
				rows[rowNum] = rows[rowNum] || []
				rows[rowNum][colNum] = value
			})
			colNum++
		}))

		return (
		<div>
			{
			struct && struct.data ?

				<table className={css(style.table)}>
				<tbody>
					<tr>
					{
						headers.map((columnName,thNum) => (
							<th key={'TableHead'+thNum} className={css(style.th)}>
								{columnName}
							</th>
						))
					}
					</tr>
					{
						rows.map((item,rowNum) => (
							<tr key={'TableRow' + rowNum} className={css(
							style.tr,
							(rowNum % 2 === 0) ? style.trEven : ''
							)}>
								{
									item.map((item,cellNum) => (
										<td key={ 'TableCell' + cellNum + rowNum} className={css(
										style.td,
										(typeof(item) === 'number') ? style.tdRight : '',
										(columns[headers[cellNum]] && columns[headers[cellNum]].red) ? style.tdRed : '',
										)}>
											{(typeof(item) === 'number') ? item.toLocaleString() : item}
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
