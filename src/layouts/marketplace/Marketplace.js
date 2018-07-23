import React, { Component } from "react";
import MarketplaceForm from "../../ui/MarketplaceForm/MarketplaceForm";

class Marketplace extends Component {
  render() {
    return (
      <main className="container">
        <div className="pure-g">
          <div className="pure-u-1-1">
            <h1>Marketplace</h1>
            <MarketplaceForm />
          </div>
        </div>
      </main>
    );
  }
}

export default Marketplace;
