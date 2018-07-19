var BREM = artifacts.require("./BREM.sol");
var BRM = artifacts.require("./BRMToken.sol");

module.exports = function(deployer) {
    deployer.deploy(BRM).then(function() {
        deployer.link(BRM, BREM);
    }).then(function() {
        deployer.deploy(BREM, BRM.address);
    });
};