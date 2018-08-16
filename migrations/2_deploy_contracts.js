var BREM = artifacts.require("./BREM.sol");
var BRM = artifacts.require("./BRMToken.sol");

module.exports = function(deployer) {
  const fee = 10
  deployer.deploy(BREM, fee);
};