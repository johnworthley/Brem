import React, { Component } from 'react'
import { css, StyleSheet } from 'aphrodite/no-important'

import ProjectCard from '../../Components/ProjectCard'

const style = StyleSheet.create({
	more: {
		textAlign: 'center',
		marginTop: 50,
		marginBottom: 60
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
		struct : {
			showMore: true,
			cards: [{
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
			}]
		}
	}

	render() {
	const { struct = {} } = this.state
    return (
      <div>
		<h3>Все проекты</h3>

		{
			struct.cards.map((item,num) => (<ProjectCard struct={item} key={'ProjectCard'+num} />))
		}
		
		{struct.showMore ? (
			<div className={css(style.more)}>
				<button className={css(style.moreButton)}>СМОТРЕТЬ ВСЕ</button>
			</div>
		) : ''}

      </div>
    )
  }
}

export default List
