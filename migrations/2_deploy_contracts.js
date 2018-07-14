var BREM = artifacts.require("./BREM.sol");
var BRM = artifacts.require("./BRMToken.sol");

module.exports = function(deployer) {
    deployer.deploy(BRM).then(function() {
        console.log(BRM.address);
        deployer.deploy(BREM, BRM.address);
    })
};