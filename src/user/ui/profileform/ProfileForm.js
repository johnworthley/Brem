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
                        this.setState({ role: 'developer' })
                    }
                })
                instance.isAuditor(coinbase).then((res) => {
                    if (res) {
                        this.setState({ role: 'auditor' })
                    }
                })
                instance.isSuperuser(coinbase).then((res) => {
                    if (res) {
                        this.setState({ role: 'superuser'})
                    }
                })
            })
        })
    } else {
        console.error('Web3 is not initialized.');
    }
}

handleMintAddressChange(event) {
    this.setState({ mintAddress: event.target.value })
}

handleMintAmountChange(event) {
    this.setState({ mintAmount: event.target.value })
}

handleMint(event) {
    event.preventDefault()

    const web3 = store.getState().web3.web3Instance;
    if (typeof web3 !== 'undefined') {
        const recieverAddress = this.state.mintAddress;
        if(!web3.utils.isAddress(recieverAddress)) {
            return alert(recieverAddress + " is not Ethereum address");
        }

        const mintAmount = this.state.mintAmount;
        if (mintAmount <= 0) {
            return alert("Mint amount must be bigger than 0");
        }

        this.props.onMintFormSubmit(recieverAddress, mintAmount)
    }
}

render() {

    const SuperuserForm = <div>
            <h3>Superuser</h3>

            <h4>BRM Token</h4>
            <form className="pure-form pure-form-ctacked" onSubmit={this.handleMint.bind(this)}>
                <fieldset>
                    <legend>Mint BRM tokens</legend>
                    <input id="mintAddress" type="text" onChange={this.handleMintAddressChange.bind(this)} placeholder="Address"/>
                    <input id="mintValue" type="number" onChange={this.handleMintAmountChange.bind(this)} placeholder="Amount"/>

                    <button type="submit" className="pure-button pure-button-primary">Mint</button>
                    <span className="pure-form-message">To simplify the MVP, only manual BRM mint is available</span>
                </fieldset>
            </form>
        </div>

    const AuditorForm = <h1>Auditor</h1>

    const DeveloperForm = <h1>Developer</h1>

    return(
        <div>
        <form className="pure-form pure-form-stacked">
            <fieldset>
                <label htmlFor="name">Username: {this.state.name}</label>
                <label htmlFor="address">Address: {this.state.address}</label>
            </fieldset>
        </form>
        
        { this.state && this.state.role && this.state.role === 'superuser' &&
            SuperuserForm
        }

        { this.state && this.state.role && this.state.role === 'developer' &&
            DeveloperForm
        }

        { this.state && this.state.role && this.state.role === 'auditor' &&
            AuditorForm
        }

        </div>
)
}
}

export default ProfileForm