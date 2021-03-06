import React, { Component } from "react";
import store from "../../store";
import axios from "axios";
import getWeb3 from "../../util/web3/getWeb3";
import BREMContract from "../../../build/contracts/BREM.json";
import ICOContract from "../../../build/contracts/BREMICO.json";
import BREMTokenContract from "../../../build/contracts/BREMToken";
import mHost from "../../config";

const contract = require("truffle-contract");

class BremICOForm extends Component {
  constructor(props) {
    super(props);

    this.state = {
      address: props.value,
      etherAmount: 0,
      tokensAmount: 0,
      userCurrentBalance: 0
    };

    this.icoAuditors.bind(this);

    getWeb3
      .then(results => {
        console.log("Web3 initialized!");
        this.setState({ loaded: true });

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
                  closingTime: new Date(
                    closingTime.toNumber() * 1000
                  ).toString()
                });
              });

              icoInstance.rate().then(rate => {
                this.setState({ rate: rate.toNumber() });
              });

              web3.eth.getBalance(this.state.address).then(balance => {
                this.setState({
                  balance: web3.utils.fromWei(balance, "ether")
                });
              });

              icoInstance.balances(coinbase).then(userBalance => {
                this.setState({
                  userBalance: web3.utils.fromWei(userBalance, "ether")
                });
              });

              icoInstance.balancesInToken(coinbase).then(userTotalTokens => {
                this.setState({
                  userTotalTokens: web3.utils.fromWei(userTotalTokens, "ether")
                });
              });

              icoInstance.weiRaised().then(weiRaised => {
                this.setState({
                  raised: web3.utils.fromWei(weiRaised, "ether")
                });
              });

              icoInstance.cap().then(cap => {
                this.setState({ cap: web3.utils.fromWei(cap, "ether") });
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

              icoInstance.request().then(request => {
                this.setState({
                  requestedValue: web3.utils.fromWei(request[0], "ether")
                });
              });

              icoInstance.isRequested().then(isRequested => {
                this.setState({ isRequested: isRequested });
              });

              icoInstance.isWithdrawn().then(isWithdrawn => {
                this.setState({ isWithdrawn: isWithdrawn });
              });

              // Get current ico auditors
              axios
                .get("http://" + mHost + "/ico/audit", {
                  params: {
                    address: this.state.address
                  }
                })
                .then(res => {
                  this.setState({ auditors: res.data });
                })
                .catch(err => console.log(err));

              const brem = contract(BREMContract);
              brem.setProvider(web3.currentProvider);
              brem.deployed().then(bremInstance => {
                bremInstance.isSuperuser(coinbase).then(isSuperuser => {
                  if (isSuperuser) {
                    this.setState({ role: "superuser" });
                    this.setState({ newAuditorAddress: "" });
                    // Get current ico auditors
                    axios
                      .get("http://" + mHost + "/ico/audit", {
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
                  if (isDeveloper && coinbase === this.state.wallet) {
                    this.setState({ role: "developer" });
                    this.setState({ withdrawValue: 0 });

                    bremInstance.withdrawFeePercent().then(fee => {
                      this.setState({ fee: fee.toNumber() });
                    });
                  }
                });

                bremInstance.isAuditor(coinbase).then(isAuditor => {
                  if (isAuditor) {
                    this.setState({ role: "auditor" });
                    icoInstance.isConfirmed(coinbase).then(isConfirmed => {
                      this.setState({ isConfirmed: isConfirmed });
                    });
                  }
                });
              });
            });
          });
        } else {
          console.error("Web3 is not initialized.");
        }
      })
      .catch(() => {
        console.log("Error in web3 initialization.");
      });
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
    const eth = e.target.value;
    if (this.state.balance - eth < 0) {
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

  // Auditor form
  handleConfirm(e) {
    e.preventDefault();

    this.props.onConfirmSubmit(this.state.address, this);
  }

  render() {
    const SuperuserForm = (
      <div>
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
          this.state.withdrawValue !== undefined &&
          this.state.balance !== undefined &&
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
                <br />
                <span className="pure-form-message">
                  <br />Withdraw fee: {this.state.fee} %
                </span>
              </fieldset>
            </form>
          )}
      </div>
    );

    const AuditorForm = (
      <div>
        {this.state &&
          this.state.isRequested === true &&
          this.state.isConfirmed === false && (
            <form
              className="pure-form pure-form-ctacked"
              onSubmit={this.handleConfirm.bind(this)}
            >
              <fieldset>
                <legend>Auditor form</legend>
                <button
                  type="submit"
                  className="pure-button pure-button-primary"
                >
                  Submit withdraw
                </button>
              </fieldset>
            </form>
          )}
      </div>
    );

    return (
      <main className="container">
        <div className="pure-u-1-1">
          {this.state &&
            this.state.icoName !== undefined && <h1>{this.state.icoName}</h1>}
          {this.state &&
            this.state.img !== undefined && (
              <p>
                <img src={this.state.img} alt="Image loading error" />
              </p>
            )}
          <p>BREM ICO Address: {this.state.address}</p>
          <p>Wallet: {this.state.wallet}</p>
          <p>Token address: {this.state.tokenAddress}</p>
          <p>Closing time: {this.state.closingTime}</p>
          {this.state &&
            this.state.rate !== undefined &&
            this.state.tokenSymbol !== undefined && (
              <p>
                Rate: {this.state.rate}
                <span> {this.state.tokenSymbol} per Wei</span>
              </p>
            )}
          {this.state &&
            this.state.isRequested === true && (
              <p>Developer's Request: {this.state.requestedValue} Eth </p>
            )}
          <p>Contract balance: {this.state.balance} Ether</p>
          <p>
            Docs:{" "}
            <a href={`https://ipfs.infura.io/ipfs/${this.state.docHash}`}>
              ipfs.infura.io/ipfs/{this.state.docHash}
            </a>
          </p>
          <p>You spend: {this.state.userBalance} Eth </p>
          <p>
            You bought: {this.state.userTotalTokens} {this.state.tokenSymbol}
          </p>
          <p>
            Current Balance: {this.state.userCurrentBalance.toLocaleString()}{" "}
            {this.state.tokenSymbol}
          </p>
          {this.state &&
            this.state.weiRaised !== null &&
            this.state.cap !== null && (
              <fieldset>
                <legend>Progress</legend>
                <p>
                  <progress value={this.state.raised} max={this.state.cap} />
                </p>
                <p>
                  Wei raised: {this.state.raised} Eth / {this.state.cap} Eth
                </p>
              </fieldset>
            )}
          <h3>Current ICO Auditors</h3>
          {this.state &&
            this.state.auditors !== undefined &&
            this.icoAuditors()}

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
                  <span> Eth </span>
                  <p>
                    You will buy: {this.state.tokensAmount}{" "}
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
            this.state.tokenSymbol !== undefined &&
            this.state.userBalance > 0 && (
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

          {this.state &&
            this.state.role &&
            this.state.role === "auditor" &&
            AuditorForm}
        </div>
      </main>
    );
  }
}

export default BremICOForm;
