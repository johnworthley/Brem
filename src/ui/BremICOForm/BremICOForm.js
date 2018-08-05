import React, { Component } from "react";
import store from "../../store";
import getWeb3 from "../../util/web3/getWeb3";
import ICOContract from "../../../build/contracts/BREMICO.json";
import BREMTokenContract from "../../../build/contracts/BREMToken";

const contract = require("truffle-contract");

class BremICOForm extends Component {
  constructor(props) {
    super(props);
    this.state = {
      address: props.value,
      etherAmount: 0,
      tokensAmount: 0
    };

    const web3 = store.getState().web3.web3Instance;
    if (typeof web3 === "undefined" || web3 === null) {
      getWeb3
        .then(results => {
          console.log("Web3 initialized!");
          this.setState({ loaded: true });
        })
        .catch(() => {
          console.log("Error in web3 initialization.");
        });
    }
    if (typeof web3 !== "undefined" && web3 !== null) {
      const ico = contract(ICOContract);
      ico.setProvider(web3.currentProvider);
      ico.at(this.state.address).then(icoInstance => {
        icoInstance.name().then(icoName => {
          this.setState({ icoName: icoName });
        });

        icoInstance.wallet().then(wallet => {
          this.setState({ wallet: wallet });
        });

        icoInstance.token().then(tokenAddress => {
          this.setState({ tokenAddress: tokenAddress });

          const bremToken = contract(BREMTokenContract);
          bremToken.setProvider(web3.currentProvider);
          bremToken.at(tokenAddress).then(tokenInstance => {
            tokenInstance.symbol().then(tokenSymbol => {
              this.setState({ tokenSymbol: tokenSymbol });
            });

            tokenInstance.decimals().then(tokenDecimals => {
              this.setState({ tokenDecimals: tokenDecimals });
            });
          });
        });

        icoInstance.closingTime().then(closingTime => {
          this.setState({
            closingTime: new Date(closingTime.toNumber()).toString()
          });
        });

        icoInstance.rate().then(rate => {
          this.setState({ rate: rate.toNumber() });
        });

        icoInstance.weiRaised().then(weiRaised => {
          this.setState({ weiRaised: weiRaised.toNumber() });
        });

        icoInstance.cap().then(cap => {
          this.setState({ cap: cap.toNumber() });
        });

        icoInstance.docHash().then(docHash => {
          this.setState({ docHash: docHash });
        });
      });
    } else {
      console.error("Web3 is not initialized.");
    }
  }

  handleEtherValueCahnge(e) {
    this.setState({ etherAmount: e.target.value });

    const web3 = store.getState().web3.web3Instance;
    if (typeof web3 !== "undefined" && web3 !== null) {
      let weiAmount = web3.utils.toWei(e.target.value, "ether");
      let tokensAmount =
        (weiAmount * this.state.rate) / Math.pow(10, this.state.tokenDecimals);
      this.setState({ tokensAmount: tokensAmount });
    }
  }

  handleBuyTokens(e) {
    e.preventDefault();

    const etherAmount = this.state.etherAmount;
    if (etherAmount === null || etherAmount <= 0) {
      return alert("Error, amount in wei must be bigger than 0");
    }

    this.props.onBuyTokenSubmit(this.state.address, etherAmount);
  }

  render() {
    return (
      <div className="pure-u-1-1">
        {this.state && this.state.icoName && <p>{this.state.icoName}</p>}
        <p>BREM ICO Address: {this.state.address}</p>
        <p>Wallet: {this.state.wallet}</p>
        <p>Token address: {this.state.tokenAddress}</p>
        <p>Closing time: {this.state.closingTime}</p>
        <p>Rate: {this.state.rate}</p>
        <p>
          Docs:{" "}
          <a href={`https://ipfs.infura.io/ipfs/${this.state.docHash}`}>
            ipfs.infura.io/ipfs/{this.state.docHash}
          </a>
        </p>
        {this.state &&
          this.state.weiRaised !== null &&
          this.state.cap !== null && (
            <fieldset>
              <legend>Progress</legend>
              <p>
                <progress value={this.state.weiRaised} max={this.state.cap} />
              </p>
              <p>
                Wei raised: {this.state.weiRaised} / {this.state.cap}
              </p>
            </fieldset>
          )}

        <form
          className="pure-form pure-form-ctacked"
          onSubmit={this.handleBuyTokens.bind(this)}
        >
          <fieldset>
            <legend>Buy Tokens</legend>
            <input
              type="number"
              min="0"
              step="0.000000000000000001"
              onChange={this.handleEtherValueCahnge.bind(this)}
              placeholder="Amount in Ether"
            />
            <p>
              You will buy: {this.state.tokensAmount} {this.state.tokenSymbol}
            </p>
            <button type="submit" className="pure-button pure-button-primary">
              Buy
            </button>
          </fieldset>
        </form>
      </div>
    );
  }
}

export default BremICOForm;
