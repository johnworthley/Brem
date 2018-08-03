import React, { Component } from "react";
import { Link } from "react-router";
import getWeb3 from "../../util/web3/getWeb3";
import MarketplaceForm from "../../ui/MarketplaceForm/MarketplaceForm";

class Marketplace extends Component {
  constructor(props) {
    super(props);

    this.state = {
      status:
        this.props.location.query.status === undefined
          ? "opened"
          : this.props.location.query.status,
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
                to="/marketplace?status=withdrawn"
                className="pure-menu-heading pure-menu-link"
              >
                Withdrawn
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
              this.state.loaded === true &&
              this.state.status && (
                <div>
                  <h1>Marketplace</h1>
                  <MarketplaceForm
                    key={this.state.status}
                    value={this.state.status}
                  />
                </div>
              )}
          </div>
        </div>
      </main>
    );
  }
}

export default Marketplace;
