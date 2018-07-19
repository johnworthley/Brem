var HDWalletProvider = require("truffle-hdwallet-provider");
// Store MNEMONIC and Infura API KEY here is not good idea
const MNEMONIC = ""; // Fill metamask secret key
const INFURA_API_KEY = ""; // Fill INFURA API KEY

module.exports = {
    // See <http://truffleframework.com/docs/advanced/configuration>
    // to customize your Truffle configuration!
    networks: {
      development: {
        host: "127.0.0.1",
        port: 8545,
        network_id: "*",
      },
      ropsten: {
        provider: () => new HDWalletProvider(MNEMONIC, "https://ropsten.infura.io/" + INFURA_API_KEY),
        network_id: 3,
        gas: 4612388,
        gasPrice: 1000000000
      },
      kovan: {
        provider: () => new HDWalletProvider(MNEMONIC, "https://kovan.infura.io/" + INFURA_API_KEY),
        network_id: 42,
        gas: 4612388,
        gasPrice: 1000000000
      },
      rinkeby: {
        provider: () => new HDWalletProvider(MNEMONIC, "https://rinkeby.infura.io/" + INFURA_API_KEY),
        network_id: 4,
      },
      // main ethereum network(mainnet)
      main: {
        provider: () => new HDWalletProvider(MNEMONIC, "https://mainnet.infura.io/" + INFURA_API_KEY),
        network_id: 1,
        gas: 4612388,
        gasPrice: 1000000000
      }
    },
  };