import React, { Component } from 'react'
import { css, StyleSheet } from 'aphrodite/no-important'

const style = StyleSheet.create({
	card: {
		height: 558,
		width: 472,
		display: 'inline-block',
		marginRight: 2,
		marginBottom: 10,
		boxSizing: 'border-box'
	},
	cardImage: {
		height: 240,
		width: 472,
		boxSizing: 'border-box',
		background: '#a8a8a8 center',
		backgroundSize: 'cover'
	},
	cardData: {
		padding: 20,
		height: 318,
		width: 472,
		boxSizing: 'border-box',
		background: '#ffffff'
	},
	progressBar: {
		height: 8,
		width: 424,
		borderRadius: 4,
		background: '#e8e5e5',
		margin: "20px 0"
	},
	cardButton: {
		minWidth: 121
	},
	company: {
		fontSize: 14
	},
	name: {
		fontSize: 20,
		padding: "5px 0 20px 0"
	},
	adress: {
		fontSize: 14
	},
	deadline: {
		fontSize: 14
	},
	deadline: {
		fontSize: 14
	},
	payment: {
		fontSize: 14
	},
	paymentSum: {
		fontSize: 24,
		color: '#ff0000'
	},
	progress: {
		fontSize: 14
	},
	progressNum: {
		fontSize: 24
	}

})

export default ({ struct }) => {
	const countStyle = StyleSheet.create({
		cardImage: {
			backgroundImage: 'url("'+struct.imageUrl+'")'
		},
		progressLevel: {
			background: "#e8e5e5 linear-gradient(to left, transparent "+(100-struct.progress)+"%, #f0723b "+struct.progress+"%)"
		}
	});

	return <div className={css(style.card)}>
		<div className={css(style.cardImage, countStyle.cardImage)}></div>
		<div className={css(style.cardData)}>
			<div className={css(style.company)}>{struct.company}</div>
			<div className={css(style.name)}>{struct.name}</div>
			<div className={css(style.adress)}>{struct.adress}</div>
			<div className={css(style.deadline)}>{struct.deadline}</div>
			<div className={css(style.payment)}>
				My impact &mdash; <span className={css(style.paymentSum)}>{struct.payment.toLocaleString()} ETH</span>
			</div>
			<div className={css(style.progress)}>
				Collected <span className={css(style.progressNum)}>{struct.progress}%</span> из {struct.limit.toLocaleString()} ETH
			</div>
			<div className={css(style.progressBar, countStyle.progressLevel)}></div>

			<button className={css(style.cardButton)}>Open</button>
		</div>
	</div>
}
