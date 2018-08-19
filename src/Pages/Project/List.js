import React, { Component } from 'react'
import { css, StyleSheet } from 'aphrodite/no-important'

import ProjectCard from '../../Components/ProjectCard'

const style = StyleSheet.create({
	list: {
		
	},
	more: {
		
	}
})

class List extends Component {
  

	render() {
	const { struct = {} } = this.props
    return (
      <div className={css(style.list)}>
		<h1>Все проекты</h1>

		{
			struct.data.map((item,num) => (<ProjectCard struct={item} key={'ProjectCard'+num} />))
		}
		
		{struct.showMore ? (<div>321</div>) : ''}

      </div>
    )
  }
}

export default List
