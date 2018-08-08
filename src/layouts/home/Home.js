import React, { Component } from "react";
import Iframe from "react-iframe";

class Home extends Component {
  render() {
    return (
      <main className="container">
        <div className="pure-g">
          <div className="pure-u-1-1">
            <Iframe
              url="https://bremtoken.io"
              // position="absolute"
              // width="100%"
              // id="myId"
              // className="myClassname"
              // height="100%"
              styles={{ left: "0px" }}
              allowFullScreen
            />
          </div>
        </div>
      </main>
    );
  }
}

export default Home;
