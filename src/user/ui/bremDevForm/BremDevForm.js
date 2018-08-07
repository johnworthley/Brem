import React, { Component } from "react";
import store from "../../../store";
import { browserHistory } from "react-router";
import ICOContract from "../../../../build/contracts/BREMICO.json";

const contract = require("truffle-contract");

class BremDevForm extends Component {
  constructor(props) {
    super(props);

    this.state = {
      address: props.value.address,
      status: props.value.Status,
      weiValue: 100
    };

    const web3 = store.getState().web3.web3Instance;
    if (typeof web3 !== "undefined" && web3 !== null) {
      const ico = contract(ICOContract);
      ico.setProvider(web3.currentProvider);
      ico.at(this.state.address).then(icoInstance => {
        icoInstance.name().then(name => {
          this.setState({ name: name });
        });

        icoInstance.description().then(description => {
          this.setState({ description: description });
        });

        web3.eth.getBalance(this.state.address).then(balance => {
          this.setState({ balance: balance });
        });

        if (this.state.status === "requested") {
          icoInstance.request().then(request => {
            this.setState({ requestedValue: request[0].toNumber() });
          });
        }
      });
    } else {
      console.error("Web3 is not initialized.");
    }
  }

  handleOpenICOPage() {
    browserHistory.push({
      pathname: "/ico/" + this.state.address
    });
  }

  handleChangeWithdrawValue(e) {
    const wei = e.target.value;
    if (this.state.balance - wei < 0) {
      return alert("Error, withdraw value must be less than ICO balance");
    }
    if (wei < 100) {
      return alert("Error, withdraw value must by bigger or equal 100 Wei");
    }
    const delta = this.state.balance - wei;
    if (delta > 0 && delta < 100) {
      alert(
        "WARNING, if you will accept tx, you loss " + delta + " in contract"
      );
    }
    this.setState({ weiValue: e.target.value });
  }

  handleSubmitWithrawRequest(e) {
    e.preventDefault();
    if (this.state.weiValue < 100) {
      return alert("Error, set up withdraw value more or equal 100 Wei");
    }

    this.props.onSubmitWithdrawRequest(
      this.state.address,
      this.state.weiValue,
      this
    );
  }

  render() {
    return (
      <div>
        {this.state &&
          this.state.name &&
          this.state.description && (
            <fieldset>
              <legend>{this.state.name}</legend>
              <p>Address: {this.state.address}</p>
              <p>{this.state.description}</p>
              <p>Status: {this.state.status}</p>
              <p>ICO balance: {this.state.balance} Wei</p>
              {this.state.status === "success" && (
                <form
                  className="pure-form pure-form-ctacked"
                  onSubmit={this.handleSubmitWithrawRequest.bind(this)}
                >
                  <fieldset>
                    <legend>Request withdraw</legend>
                    <input
                      type="number"
                      value={this.state.weiValue}
                      step="1"
                      onChange={this.handleChangeWithdrawValue.bind(this)}
                      placeholder="Request value in Wei"
                    />
                    <button
                      type="submit"
                      className="pure-button pure-button-primary"
                    >
                      Request withdraw
                    </button>
                  </fieldset>
                </form>
              )}

              {this.state.status === "requested" &&
                this.state.requestedValue !== undefined && (
                  <p>Requested value: {this.state.requestedValue} Wei </p>
                )}

              <p>
                <button
                  className="pure-button pure-button-primary"
                  onClick={this.handleOpenICOPage.bind(this)}
                >
                  Open ICO Page
                </button>
              </p>
            </fieldset>
          )}
      </div>
    );
  }
}

export default BremDevForm;
