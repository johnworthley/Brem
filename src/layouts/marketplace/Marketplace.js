import React, { Component } from "react";
import { Link } from "react-router";
import getWeb3 from "../../util/web3/getWeb3";
import MarketplaceForm from "../../ui/MarketplaceForm/MarketplaceForm";

class Marketplace extends Component {
  constructor(props) {
    super(props);

    let status = this.props.location.query.status;
    if (status === undefined) {
      status = "opened";
    }
    this.state = {
      status: status,
      loaded: false
    };

    // Initialize web3 and set in Redux.
    getWeb3
      .then(results => {
        console.log("Web3 initialized!");
        this.setState({ loaded: true });
      })
      .catch(() => {
        console.log("Error in web3 initialization.");
      });
  }

  render() {
    return (
      <main className="container">
        <div className="pure-g">
          <div className="pure-u-1-1">
            <nav className="navbar pure-menu pure-menu-horizontal">
              <Link
                onClick={this.forceUpdate}
                to="/marketplace"
                className="pure-menu-heading pure-menu-link"
              >
                Open
              </Link>
              <Link
                onClick={this.forceUpdate}
                to="/marketplace?status=success"
                className="pure-menu-heading pure-menu-link"
              >
                Success
              </Link>
              <Link
                onClick={this.forceUpdate}
                to="/marketplace?status=withdrawed"
                className="pure-menu-heading pure-menu-link"
              >
                Withdrawed
              </Link>
              <Link
                onClick={this.forceUpdate}
                to="/marketplace?status=failed"
                className="pure-menu-heading pure-menu-link"
              >
                Failed
              </Link>
            </nav>
            <br />
            <br />
            {this.state &&
              this.state.loaded && (
                <div>
                  <h1>Marketplace</h1>
                  <MarketplaceForm value={this.state.status} />
                </div>
              )}
          </div>
        </div>
      </main>
    );
  }
}

export default Marketplace;
