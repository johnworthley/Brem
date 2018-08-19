import React, { Component } from 'react'

class View extends Component {
  state = {
    projectId: '',
    project: {}
  }

  componentDidMount = async () => {
    this.setState({
      projectId: this.props.match.params.id
    })


    // Get all data here and push it to state.project object
  }
  depositETH = async () => {

  }

  withDrawETH = async () => {

  }

  auditorConfirmRequest = async () => {

  }

  devWithdrawETH = async () => {

  }

  superAddAuditor = () => {
    const auditorKey = this.newAuditorName.value
    console.log(auditorKey)
  }

  render() {
    const { projectId } = this.state
    return (
      <div>
        Project ID: { projectId }
        <button onClick={this.depositETH}>
           Deposit ETH
        </button>
        <button onClick={this.withdrawETH}>
          Withdraw ETH
        </button>
        <button onClick={this.auditorConfirmRequest}>
          Auditor Confirm Request
        </button>
        <button onClick={this.devWithdrawETH}>
          Dev withdraw eth
        </button>
        <input ref={elem => this.newAuditorName = elem} type="text" placeholder="Auditor name"/>
        <button onClick={this.superAddAuditor}>
          SuperUser Add auditor
        </button>
      </div>
    )
  }
}

export default View
