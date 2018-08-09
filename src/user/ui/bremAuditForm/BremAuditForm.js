import React, { Component } from "react";
import store from "../../../store";
import axios from "axios";
import { browserHistory } from "react-router";
import BREMContract from "../../../../build/contracts/BREM.json";
import ICOContract from "../../../../build/contracts/BREMICO.json";
import mHost from "../../../config";

const contract = require("truffle-contract");

class BremAuditForm extends Component {
  constructor(props) {
    super(props);

    this.state = {
      address: props.value.address,
      status: props.value.status,
      visible: true
    };
    const web3 = store.getState().web3.web3Instance;
    if (typeof web3 !== "undefined" && web3 !== null) {
      const ico = contract(ICOContract);
      ico.setProvider(web3.currentProvider);
      web3.eth.getCoinbase((error, coinbase) => {
        if (error) {
          console.error(error);
        }

        ico.at(this.state.address).then(icoInstance => {
          icoInstance.request().then(request => {
            this.setState({
              requestedValue: web3.utils.fromWei(request[0], "ether")
            });
          });

          icoInstance.isConfirmed(coinbase).then(isConfirmed => {
            this.setState({ visible: !isConfirmed });
            this.setState({ isConfirmed: isConfirmed });
            if (!isConfirmed) {
              icoInstance.name().then(name => {
                this.setState({ name: name });
              });

              axios
                .get("http://" + mHost + "/ico/image", {
                  params: {
                    address: this.state.address
                  },
                  responseType: "arraybuffer"
                })
                .then(res => {
                  const base64 = Buffer.from(res.data, "binary").toString(
                    "base64"
                  );
                  this.setState({ img: "data:image/jpeg;base64," + base64 });
                })
                .catch(err => console.error(err));

              icoInstance.description().then(description => {
                this.setState({ description: description });
              });

              icoInstance.wallet().then(wallet => {
                this.setState({ wallet: wallet });

                const brem = contract(BREMContract);
                brem.setProvider(web3.currentProvider);

                brem.deployed().then(bremInstance => {
                  bremInstance.login({ from: wallet }).then(username => {
                    this.setState({ username: username });
                  });
                });
              });
            }
          });
        });
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

  handleConfirm(e) {
    e.preventDefault();

    this.props.onConfirmSubmit(this.state.address, this);
  }

  render() {
    return (
      <div>
        {this.state &&
          this.state.visible === true &&
          this.state.isConfirmed === false &&
          this.state.name &&
          this.state.description &&
          this.state.requestedValue &&
          this.state.username && (
            <fieldset>
              <legend>{this.state.name}</legend>
              {this.state &&
                this.state.img !== undefined && (
                  <p>
                    <img src={this.state.img} alt="Image loading error" />
                  </p>
                )}
              <span className="pure-form-message">
                by {this.state.username}
              </span>
              <p>Address: {this.state.address}</p>
              <p>{this.state.description}</p>
              <p>Wallet: {this.state.wallet}</p>
              <p>Requested value: {this.state.requestedValue} Wei</p>

              <button
                className="pure-button pure-button-primary"
                onClick={this.handleConfirm.bind(this)}
              >
                Confirm withdraw
              </button>

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

export default BremAuditForm;
