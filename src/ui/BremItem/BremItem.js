import React, { Component } from "react";
import { browserHistory } from "react-router";
import store from "../../store";
import BREMContract from "../../../build/contracts/BREM.json";
import ICOContract from "../../../build/contracts/BREMICO.json";

const contract = require("truffle-contract");

class BremItem extends Component {
  constructor(props) {
    super(props);

    this.state = {
      index: props.value
    };

    const web3 = store.getState().web3.web3Instance;
    if (typeof web3 !== "undefined" && web3 !== null) {
      const brem = contract(BREMContract);
      brem.setProvider(web3.currentProvider);
      brem.deployed().then(bremInstance => {
        bremInstance.getProject(this.state.index).then(icoAddress => {
          this.setState({ address: icoAddress });

          const ico = contract(ICOContract);
          ico.setProvider(web3.currentProvider);
          ico.at(this.state.address).then(icoInstance => {
            icoInstance.auditSelected().then(res => {
              this.setState({ selected: res });

              if (res) {
                icoInstance.name().then(name => {
                  this.setState({ name: name });
                });

                icoInstance.description().then(descr => {
                  this.setState({ descr: descr });
                });

                icoInstance.wallet().then(developer => {
                  bremInstance.login({ from: developer }).then(username => {
                    this.setState({ user: username });
                  });
                });
              }
            });
          });
        });
      });
    } else {
      console.error("Web3 is not initialized.");
    }
  }

  handleOpenICOForm(e) {
    browserHistory.push({
      pathname: "/ico/" + this.state.address
    });
  }

  render() {
    return (
      <div>
        {this.state &&
          this.state.selected !== null &&
          this.state.selected &&
          this.state.name !== null &&
          this.state.descr !== null &&
          this.state.user !== null && (
            <fieldset>
              <legend>{this.state.name}</legend>
              <label> {this.state.descr}</label>
              <span className="pure-form-message">by {this.state.user} </span>
              <button
                type="button"
                className="pure-button pure-button-primary"
                onClick={this.handleOpenICOForm.bind(this)}
              >
                Open ICO Form
              </button>
            </fieldset>
          )}
      </div>
    );
  }
}

export default BremItem;
