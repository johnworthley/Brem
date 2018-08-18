import React, { Component } from 'react'

class SignupTab extends Component {
  state = {

  }

  render() {

    return (
      <div>
        <h3>Register</h3>
        <form onSubmit={this.handleSubmit}>
          <input type="text" placeholder="Username"/>
          <button>Register me!</button>
        </form>
      </div>
    )
  }
}

export default SignupTab
