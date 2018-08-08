import React, { Component } from "react";
import { browserHistory } from "react-router";
import axios from "axios";
import store from "../../store";
import BREMContract from "../../../build/contracts/BREM.json";
import ICOContract from "../../../build/contracts/BREMICO.json";
import mHost from "../../../config";

const contract = require("truffle-contract");

class BremItem extends Component {
  constructor(props) {
    super(props);

    this.state = {
      address: props.value
    };

    const web3 = store.getState().web3.web3Instance;
    if (typeof web3 !== "undefined" && web3 !== null) {
      const ico = contract(ICOContract);
      ico.setProvider(web3.currentProvider);
      ico.at(this.state.address).then(icoInstance => {
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
            const base64 = Buffer.from(res.data, "binary").toString("base64");
            this.setState({ img: "data:image/jpeg;base64," + base64 });
          })
          .catch(err => console.error(err));

        icoInstance.description().then(descr => {
          this.setState({ descr: descr });
        });

        icoInstance.wallet().then(developer => {
          const brem = contract(BREMContract);
          brem.setProvider(web3.currentProvider);
          brem.deployed().then(bremInstance => {
            bremInstance.login({ from: developer }).then(username => {
              this.setState({ user: username });
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
          this.state.name !== null &&
          this.state.descr !== null &&
          this.state.user !== null && (
            <fieldset>
              <legend>{this.state.name}</legend>
              {this.state &&
                this.state.img !== undefined && (
                  <p>
                    <img src={this.state.img} alt="Image loading error" />
                  </p>
                )}
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
