import React, { Component } from 'react'

class PromoPic extends Component {
  componentWillMount() {
  }
  render() {
    const backgroundImage = `url(${this.props.url})`
    return (
      <div key="backgroundImage" className="promo_background" style={{backgroundImage}}></div>
    )
  }
}

export default PromoPic
