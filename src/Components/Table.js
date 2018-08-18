import React, { Component } from 'react'
import { css, StyleSheet } from 'aphrodite/no-important'

const style = StyleSheet.create({
  table : {
    borderCollapse: 'collapse',
	boxSizing: 'border-box',
	width: '100%',
	fontFamily: 'OpenSans',
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
	borderWidth: '1px 0',
	boxSizing: 'border-box',
	textAlign: 'left',
	height: 31,
	padding: '0 20px'
  },
  tdRight : {
	textAlign: 'right'
  },
  tdRed : {
	color: '#ff0000'
  },
  tdEven : {
	backgroundColor: '#f8f8f8'
  }
})

export default class Table extends Component {
	render = () => {
		const { struct = {} } = this.props
		const { options = {} } = struct
		const { columns = {} } = options 
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
							<tr key={'TableRow' + rowNum}>
								{
									item.map((item,cellNum) => (
										<td key={ 'TableCell' + cellNum + rowNum} className={css(
										style.td, 
										(typeof(item) === 'number') ? style.tdRight : '',
										(rowNum % 2 === 0) ? style.tdEven : '',
										(columns[headers[cellNum]] && columns[headers[cellNum]].red) ? style.tdRed : '',
										)}>
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