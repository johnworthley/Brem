import React, { Component } from "react";
import store from "../../../store";
import axios from "axios";
import { browserHistory } from "react-router";
import ICOContract from "../../../../build/contracts/BREMICO.json";
import mHost from "../../../config";

const contract = require("truffle-contract");

class BremDevForm extends Component {
  constructor(props) {
    super(props);

    this.state = {
      address: props.value.address,
      status: props.value.Status,
      withdrawValue: 0
    };

    const web3 = store.getState().web3.web3Instance;
    if (typeof web3 !== "undefined" && web3 !== null) {
      const ico = contract(ICOContract);
      ico.setProvider(web3.currentProvider);
      ico.at(this.state.address).then(icoInstance => {
        icoInstance.name().then(name => {
          this.setState({ name: name });
        });

        // Get ICO image
        axios
          .get("http://" + mHost + "/ico/image", {
            params: {
              address: this.state.address
            },
            responseType: "arraybuffer"
          })
          .then(res => {
            const base64 = Buffer.from(res.data, "binary").toString("base64");
            this.setState({ img: "data:image/jpeg;base64," + base64 });
          })
          .catch(err => console.error(err));

        icoInstance.description().then(description => {
          this.setState({ description: description });
        });

        icoInstance.withdrawFeePercent().then(fee => {
          this.setState({ fee: fee.toNumber() });
        });

        web3.eth.getBalance(this.state.address).then(balance => {
          this.setState({ balance: web3.utils.fromWei(balance, "ether") });
        });

        if (this.state.status === "requested") {
          icoInstance.request().then(request => {
            this.setState({
              requestedValue: web3.utils.fromWei(request[0], "ether")
            });
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
    const eth = e.target.value;
    if (this.state.balance - eth <= 0) {
      return alert("Error, withdraw value must be less than ICO balance");
    }
    const web3 = store.getState().web3.web3Instance;
    if (typeof web3 !== "undefined" && web3 !== null) {
      const wei = web3.utils.toWei(eth, "ether");
      if (wei < 100) {
        return alert("Error, withdraw value must by bigger or equal 100 Wei");
      }
      const icoBalance = web3.utils.toWei(this.state.balance, "ether");
      const delta = icoBalance - wei;
      if (delta > 0 && delta < 100) {
        this.setState({ withdrawValue: this.state.balance });
      } else {
        this.setState({ withdrawValue: eth });
      }
    } else {
      console.error("Web3 is not initialized.");
    }
  }

  handleSubmitWithrawRequest(e) {
    e.preventDefault();

    if (this.state.withdrawValue <= 0) {
      return alert("Withdraw value must be bigger than 0");
    }

    this.props.onSubmitWithdrawRequest(
      this.state.address,
      this.state.withdrawValue,
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
              {this.state &&
                this.state.img !== undefined && (
                  <p>
                    <img src={this.state.img} alt="Image loading error" />
                  </p>
                )}
              <p>Address: {this.state.address}</p>
              <p>{this.state.description}</p>
              <p>Status: {this.state.status}</p>
              <p>ICO balance: {this.state.balance} Eth</p>
              {this.state.status === "success" && (
                <form
                  className="pure-form pure-form-ctacked"
                  onSubmit={this.handleSubmitWithrawRequest.bind(this)}
                >
                  <fieldset>
                    <legend>Request withdraw</legend>
                    <input
                      type="number"
                      value={this.state.withdrawValue}
                      step="0.000000000000000001"
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
                  <p>Requested value: {this.state.requestedValue} Eth </p>
                )}

              <p>
                <button
                  className="pure-button pure-button-primary"
                  onClick={this.handleOpenICOPage.bind(this)}
                >
                  Open ICO Page
                </button>
                <span className="pure-form-message">
                  <br />Withdraw fee: {this.state.fee} %
                </span>
              </p>
            </fieldset>
          )}
      </div>
    );
  }
}

export default BremDevForm;
