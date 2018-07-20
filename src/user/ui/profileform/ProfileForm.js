import React, { Component } from 'react';
import store from '../../../store';
import BREMContract from '../../../../build/contracts/BREM.json';

const contract = require('truffle-contract');

class ProfileForm extends Component {

    constructor(props) {
        super(props)

        this.state = {
            BREMInstance: null,
            name: this.props.name,
            address: null,
            isDeveloper: false,
            isAuditor: false,
            isSuperuser: false
        }

        this.init();
    }

    init() {
        const web3 = store.getState().web3.web3Instance;
        if (typeof web3 !== 'undefined') {
            const brem = contract(BREMContract);
            brem.setProvider(web3.currentProvider);
            brem.deployed().then((instance) => {
                this.setState({BREMInstance: instance});
                // Get coinbase
                web3.eth.getCoinbase((err, coinbase) => {
                    if (err) {
                        console.error(err);
                    }
                    this.setState({address: coinbase});
                    instance.isDeveloper(coinbase).then((res) => {
                        if (res) {
                            this.setState({
                                isDeveloper: true,
                                isAuditor: false,
                                isSuperuser: false
                            })
                        }
                    })
                    instance.isAuditor(coinbase).then((res) => {
                        if (res) {
                            this.setState({
                                isAuditor: true,
                                isDeveloper: false,
                                isSuperuser: false
                            })
                        }
                    })
                    instance.isSuperuser(coinbase).then((res) => {
                        this.setState({
                            isSuperuser: true,
                            isDeveloper: false,
                            isAuditor: false
                        })
                    })
                })
            })
        } else {
            console.error('Web3 is not initialized.');
        }
    }

//   handleSubmit(event) {
//     event.preventDefault()

//     if (this.state.name.length < 2)
//     {
//       return alert('Please fill in your name.')
//     }

//     this.props.onProfileFormSubmit(this.state.name)
//   }

    render() {

        

    return(
      <form className="pure-form pure-form-stacked" onSubmit={this.handleSubmit.bind(this)}>
        <fieldset>
          <label htmlFor="name">Username: {this.state.name}</label>
          <label htmlFor="address">Address: {this.state.address}</label>

          <button type="submit" className="pure-button pure-button-primary">Update</button>
        </fieldset>
      </form>
    )
  }
}

export default ProfileForm
