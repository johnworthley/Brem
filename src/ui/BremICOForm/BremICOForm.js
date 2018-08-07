import React, { Component } from "react";
import store from "../../store";
import axios from "axios";
import getWeb3 from "../../util/web3/getWeb3";
import BREMContract from "../../../build/contracts/BREM.json";
import ICOContract from "../../../build/contracts/BREMICO.json";
import BREMTokenContract from "../../../build/contracts/BREMToken";

const contract = require("truffle-contract");

class BremICOForm extends Component {
  constructor(props) {
    super(props);

    getWeb3
      .then(results => {
        console.log("Web3 initialized!");
        this.setState({ loaded: true });
      })
      .catch(() => {
        console.log("Error in web3 initialization.");
      });

    this.state = {
      address: props.value,
      etherAmount: 0,
      tokensAmount: 0,
      userCurrentBalance: 0
    };

    this.icoAuditors.bind(this);

    const web3 = store.getState().web3.web3Instance;
    if (typeof web3 !== "undefined" && web3 !== null) {
      const ico = contract(ICOContract);
      ico.setProvider(web3.currentProvider);
      ico.at(this.state.address).then(icoInstance => {
        web3.eth.getCoinbase((err, coinbase) => {
          if (err) {
            console.error(err);
          }

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
                this.setState({ tokenDecimals: tokenDecimals.toNumber() });
              });

              tokenInstance.balanceOf(coinbase).then(userCurrentBalance => {
                this.setState({
                  userCurrentBalance: userCurrentBalance.toNumber()
                });
              });
            });
          });

          icoInstance.closingTime().then(closingTime => {
            this.setState({
              closingTime: new Date(closingTime.toNumber() * 1000).toString()
            });
          });

          icoInstance.rate().then(rate => {
            this.setState({ rate: rate.toNumber() });
          });

          web3.eth.getBalance(this.state.address).then(balance => {
            this.setState({ balanceInWei: balance });
            this.setState({ balance: web3.utils.fromWei(balance, "ether") });
          });

          icoInstance.balances(coinbase).then(userBalance => {
            this.setState({ userBalance: userBalance.toNumber() });
          });

          icoInstance.balancesInToken(coinbase).then(userTotalTokens => {
            this.setState({ userTotalTokens: userTotalTokens.toNumber() });
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

          icoInstance.hasClosed().then(hasClosed => {
            this.setState({ hasClosed: hasClosed });
          });

          icoInstance.capReached().then(capReached => {
            this.setState({ capReached: capReached });
          });

          icoInstance.auditSelected().then(auditSelected => {
            this.setState({ auditSelected: auditSelected });
          });

          icoInstance.isRequested().then(isRequested => {
            this.setState({ isRequested: isRequested });
          });

          icoInstance.isWithdrawn().then(isWithdrawn => {
            this.setState({ isWithdrawn: isWithdrawn });
          });

          const brem = contract(BREMContract);
          brem.setProvider(web3.currentProvider);
          brem.deployed().then(bremInstance => {
            bremInstance.isSuperuser(coinbase).then(isSuperuser => {
              if (isSuperuser) {
                this.setState({ role: "superuser" });
                this.setState({ newAuditorAddress: "" });
                // Get current ico auditors
                axios
                  .get("http://127.0.0.1:8080/ico/audit", {
                    params: {
                      address: this.state.address
                    }
                  })
                  .then(res => {
                    this.setState({ auditors: res.data });
                  })
                  .catch(err => console.log(err));
              }
            });

            bremInstance.isDeveloper(coinbase).then(isDeveloper => {
              if (isDeveloper) {
                this.setState({ role: "developer" });
                this.setState({ withdrawValueInWei: 100 });
              }
            });
          });
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

    this.props.onBuyTokenSubmit(
      this.state.address,
      this.state.tokenAddress,
      etherAmount,
      this
    );
  }

  handleRefund(e) {
    e.preventDefault();

    this.props.onRefundSubmit(
      this.state.address,
      this.state.tokenAddress,
      this
    );
  }

  icoAuditors() {
    const amount = this.state.auditors.length;
    if (amount === 0) {
      return (
        <div>
          <p>Auditors didn't add</p>
        </div>
      );
    }
    return this.state.auditors.map(auditor => (
      <p key={auditor.address}>{auditor.address}</p>
    ));
  }

  handleAddNewAuditorChange(e) {
    this.setState({ newAuditorAddress: e.target.value });
  }

  handleAddNewAuditor(e) {
    e.preventDefault();

    const web3 = store.getState().web3.web3Instance;
    if (typeof web3 !== "undefined" && web3 !== null) {
      const auditorAddress = this.state.newAuditorAddress;
      if (!web3.utils.isAddress(auditorAddress)) {
        return alert(auditorAddress + " is not Ethereum address");
      }

      this.props.onAddNewAuditorSubmit(
        this.state.address,
        auditorAddress,
        this
      );
    } else {
      console.error("Web3 is not initialized.");
    }
  }

  handlePublishProject(e) {
    e.preventDefault();
    const web3 = store.getState().web3.web3Instance;
    if (typeof web3 !== "undefined" && web3 !== null) {
      this.props.onPublishProjectSubmit(this.state.address, this);
    } else {
      console.error("Web3 is not initialized.");
    }
  }

  // Developer form
  handleChangeWithdrawValue(e) {
    const wei = e.target.value;
    if (this.state.balanceInWei - wei < 0) {
      return alert("Error, withdraw value must be less than ICO balance");
    }
    if (wei < 100) {
      return alert("Error, withdraw value must by bigger or equal 100 Wei");
    }

    const delta = this.state.balanceInWei - wei;
    if (delta > 0 && delta < 100) {
      alert(
        "WARNING, if you will accept tx, you loss " + delta + " in contract"
      );
    }
    this.setState({ withdrawValueInWei: e.target.value });
  }

  handleSubmitWithrawRequest(e) {
    e.preventDefault();
    if (this.state.withdrawValueInWei < 100) {
      return alert("Error, set up withdraw value more or equal 100 Wei");
    }

    this.props.onSubmitWithdrawRequest(
      this.state.address,
      this.state.withdrawValueInWei,
      this
    );
  }

  render() {
    const SuperuserForm = (
      <div>
        <h2>Current ICO Auditors</h2>
        {this.state && this.state.auditors !== undefined && this.icoAuditors()}
        {this.state &&
          this.state.auditSelected === false &&
          this.state.newAuditorAddress !== undefined && (
            <div>
              <form
                className="pure-form pure-form-ctacked"
                onSubmit={this.handleAddNewAuditor.bind(this)}
              >
                <fieldset>
                  <legend>Add new auditor to project</legend>
                  <input
                    type="text"
                    value={this.state.newAuditorAddress}
                    onChange={this.handleAddNewAuditorChange.bind(this)}
                    placeholder="New Auditor address"
                  />

                  <button
                    type="submit"
                    className="pure-button pure-button-primary"
                  >
                    Add auditor to project
                  </button>
                </fieldset>
              </form>

              <button
                className="pure-button pure-button-primary"
                onClick={this.handlePublishProject.bind(this)}
              >
                Publish project
              </button>
            </div>
          )}
      </div>
    );

    const DeveloperForm = (
      <div>
        {this.state &&
          this.state.withdrawValueInWei !== undefined &&
          this.state.balanceInWei !== undefined &&
          this.state.hasClosed === true &&
          this.state.capReached === true &&
          this.state.isWithdrawn === false &&
          this.state.isRequested === false && (
            <form
              className="pure-form pure-form-ctacked"
              onSubmit={this.handleSubmitWithrawRequest.bind(this)}
            >
              <fieldset>
                <legend>Request withdraw</legend>
                <input
                  type="number"
                  value={this.state.withdrawValueInWei}
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
      </div>
    );

    return (
      <div className="pure-u-1-1">
        {this.state && this.state.icoName && <p>{this.state.icoName}</p>}
        <p>BREM ICO Address: {this.state.address}</p>
        <p>Wallet: {this.state.wallet}</p>
        <p>Token address: {this.state.tokenAddress}</p>
        <p>Closing time: {this.state.closingTime}</p>
        <p>Rate: {this.state.rate}</p>
        <p>Contract balance: {this.state.balance} Ether</p>
        <p>
          Docs:{" "}
          <a href={`https://ipfs.infura.io/ipfs/${this.state.docHash}`}>
            ipfs.infura.io/ipfs/{this.state.docHash}
          </a>
        </p>
        <p>You spend: {this.state.userBalance} Wei </p>
        <p>
          You bought: {this.state.userTotalTokens} {this.state.tokenSymbol}
        </p>
        <p>
          Current Balance: {this.state.userCurrentBalance}{" "}
          {this.state.tokenSymbol}
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

        {this.state &&
          this.state.auditSelected === true &&
          this.state.etherAmount !== undefined &&
          this.state.tokensAmount !== undefined &&
          this.state.hasClosed === false && (
            <form
              className="pure-form pure-form-ctacked"
              onSubmit={this.handleBuyTokens.bind(this)}
            >
              <fieldset>
                <legend>Buy Tokens</legend>
                <input
                  type="number"
                  min="0"
                  value={this.state.etherAmount}
                  step="0.000000000000000001"
                  onChange={this.handleEtherValueCahnge.bind(this)}
                  placeholder="Amount in Ether"
                />
                <p>
                  You will buy: {this.state.tokensAmount}
                  {this.state.tokenSymbol}
                </p>
                <button
                  type="submit"
                  className="pure-button pure-button-primary"
                >
                  Buy
                </button>
              </fieldset>
            </form>
          )}

        {this.state &&
          this.state.hasClosed === true &&
          this.state.capReached === false &&
          this.state.userBalance !== undefined &&
          this.state.tokenSymbol !== undefined && (
            <form
              className="pure-form pure-form-ctacked"
              onSubmit={this.handleRefund.bind(this)}
            >
              <fieldset>
                <legend>Refund</legend>
                {this.state.userBalance === 0 && (
                  <p>You haven't Ether to refund</p>
                )}
                {this.state.userBalance > 0 && (
                  <div>
                    <p>You can refund {this.state.userBalance} Wei</p>
                    <span className="pure-form-message">
                      Important: you {this.state.tokenSymbol} tokens balance
                      must be equal to total bought tokens amount
                    </span>
                    <button
                      type="submit"
                      className="pure-button pure-button-primary"
                    >
                      Refund
                    </button>
                  </div>
                )}
              </fieldset>
            </form>
          )}

        {this.state &&
          this.state.role &&
          this.state.role === "superuser" &&
          SuperuserForm}

        {this.state &&
          this.state.role &&
          this.state.role === "developer" &&
          DeveloperForm}
      </div>
    );
  }
}

export default BremICOForm;
