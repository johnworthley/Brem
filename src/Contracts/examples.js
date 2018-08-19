import BREMContract from "../../build/contracts/BREM.json";
import web3Provider from "../util/getweb3";
import axios from "axios";

const contract = require("truffle-contract");

export function signUpUser(name) {
  const web3 = web3Provider.getWeb3();
  // Using truffle-contract we create the authentication object.
  const brem = contract(BREMContract);
  brem.setProvider(web3.currentProvider);

  // Get current ethereum wallet.
  web3.eth.getCoinbase((error, coinbase) => {
    // Log errors, if any.
    if (error) {
      console.error(error);
      return
    }

    web3.currentProvider.sendAsync(
      {
        method: "personal_sign",
        params: [web3.utils.utf8ToHex(coinbase), coinbase],
        from: coinbase
      },
      (err, res) => {
        if (err) {
          console.error(err);
          return
        }
        const sign = res.result;
        
        brem.deployed().then(bremIntance => {
          // Checking fro superuser address
          bremIntance.isSuperuser(coinbase).then(isSuperuser => {
            if (isSuperuser) {
              // Try to login
              bremIntance.login({from: coinbase}).then(name => {
                // Name to login
                return
              })
              .catch(err => {
                // Sign up
                bremIntance.signUp(name, {from: coinbase}).then(name => {
                  console.log(name)
                  return
                })
                .catch(err => {
                  console.error(err)
                })
              })
            } else {
              // Checking for auditor address
              bremIntance.isAuditor(coinbase).then(isAuditor => {
                if (isAuditor) {
                  bremIntance.login({from: coinbase}.then(name => {
                    console.log(name)
                    return
                  }))
                  .catch(err => {
                    bremIntance.signUp(name, {from: coinbase}).then(name => {
                      console.log(name)
                      return
                    })
                    .catch(err => {
                      console.error(err)
                    })
                  })
                } else {
                 // Register developer
                 const host = "Здесь должна быть переменная с используменым сервером"
                 const developer = {
                   address: coinbase,
                   username: name,
                   sign: sign
                 }
                 axios.post(host + "/signup", developer)
                  .then(res => {
                    console.log(res);
                  })
                  .catch(err => {
                    console.error(err);
                  });
                }
              })
            }
          })
        })
      })
    })
  }

  // Login
  export function loginUser() {
    const web3 = web3Provider.getWeb3();
    // Using truffle-contract we create the authentication object.
    const brem = contract(BREMContract);
    brem.setProvider(web3.currentProvider);
  
    // Get current ethereum wallet.
    web3.eth.getCoinbase((error, coinbase) => {
      if (error) {
        console.error(error);
        return
      }

      web3.currentProvider.sendAsync(
        {
          method: "personal_sign",
          params: [web3.utils.utf8ToHex(coinbase), coinbase],
          from: coinbase
        },
        (err, res) => {
          if (err) {
            console.error(err);
            return
          }
          const sign = res.result;

          brem.deployed().then(bremInstance => {
            // Attempt to login user.
            bremInstance.login({ from: coinbase })
            .then(function(userName) {
            // If no error, login user. (Superuser of auditor)
            // Write to cokkie address and sign
          })
          .catch(res => {
            const host = "Здесь должна быть переменная с используменым сервером"
            const developer = {
              address: coinbase,
              sign: sign
            }
            axios.post(host + "/signup", developer)
             .then(res => {
               console.log(res); // В респонсе будет структура застройщика
             })
             .catch(err => {
               console.error(err);
             });
          })
        })
      });
    });
  };

  


async function createNewBREMICO(
  name,
  symbol,
  rate,
  cap,
  closingTime,
  description,
  files,
  image,
  location,
  locAddress
) {
  console.log('createNewBREMICO');
  const { host } = config
  const { web3Instance, web3Account } = store
  const { currentProvider, utils, eth } = web3Instance
  const brem = contract(BREMContract);
  brem.setProvider(currentProvider)
  const coinbase = await eth.getCoinbase()
  store.update({
    web3Coinbase: coinbase
  })
  
  const bremInstance = await brem.deployed();

  const isSuperuser = await bremInstance.isSuperuser(coinbase);
  const isAuditor   = await bremInstance.isAuditor(coinbase);

  if(isSuperuser || isAuditor){
    alert('Not for superuser or auditor');
    return;
  }

  if(cap < 100){
    alert('Invalid cap');
    return;
  }
  if(name.length == 0){
    alert('Invalid name');
    return;
  }
  if(symbol.length == 0){
    alert('Invalid symbol');
    return;
  }

  const isValidName = await bremInstance.isValidName(name);
  if (!isValidName){
    alert('Name already exists');
    return;
  }

  const feePercent = await bremInstance.withdrawFeePercent();


  ipfs.files.add(files, (error, result) => {
    if (error) {
      console.log(error);
      return;
    }
    const docHash = result[result.length - 1].hash;


    bremInstance
      .createBREMICO(
        name,
        symbol,
        rate,
        utils.toWei(cap.toString(), "ether"),
        closingTime.toString(),
        docHash,
        { from: coinbase }
      )
      .then(async TXres => {
        try {
          const ico = {
            address: TXres.logs[0].args.icoAddress,
            developer: {
              address: coinbase
            },
            description: description,
            closing_time: closingTime, //format?
            fee_percent: feePercent,
            token_address: TXres.logs[0].args.tokenAddress,
            name: name,
            symbol: symbol,
            location: location,
            loc_address: locAddress
          };
          let res = await axios.post(host + "dev/ico", ico);  
          console.log(res);

          let formData = new FormData();
          formData.append(
            "address",
            TXres.logs[0].args.icoAddress
          );
          formData.append("image", image);
          const config = {
            headers: {
              "content-type": "multipart/form-data"
            }
          };
          axios.post(
                  host + "dev/image",
                  formData,
                  config
                );

        } catch(err){
          console.log(err);
        }
            

        alert(
          "TX: " +
          TXres.tx +
          " ICO: " +
          TXres.logs[0].args.icoAddress +
          " Token: " +
          TXres.logs[0].args.tokenAddress
        );

      });
                               
    });
    
}


async function addNewAuditor(address) {
  console.log('addNewAuditor');
  const { host } = config
  const { web3Instance, web3Account } = store
  const { currentProvider, utils, eth } = web3Instance
  const brem = contract(BREMContract);
  brem.setProvider(currentProvider)
  const coinbase = await eth.getCoinbase()
  store.update({
    web3Coinbase: coinbase
  })
  
  const bremInstance = await brem.deployed();

  if (address.length == 0){
    alert('Invalid address');
    return;
  }

  const isSuperuser = await bremInstance.isSuperuser(coinbase);
  if(!isSuperuser){
    alert('Only for superuser');
    return;
  }

  const isAuditor = await bremInstance.isAuditor(address);
  if (isAuditor) {
    alert(address + " is already auditor");
    return;
  }

  bremInstance.addAuditor(address, { from: coinbase }).then(async txRes => {
    try {
      const auditor = {
        address: address
      };
    
      let res = await axios.post(host + "audit", auditor) //Нету post/audit
      console.log(res);
      
      
      alert("Success. TX: " + txRes.tx);

    } catch(err){
      console.log(err);
    }
  });
  
}

async function withdrawFees(withdrawAmount) { //not tested
  console.log('withdrawFees');
  const { host } = config
  const { web3Instance, web3Account } = store
  const { currentProvider, utils, eth } = web3Instance
  const brem = contract(BREMContract);
  brem.setProvider(currentProvider)
  const coinbase = await eth.getCoinbase()
  store.update({
    web3Coinbase: coinbase
  })

  const bremInstance = await brem.deployed();
  
  const isSuperuser = await bremInstance.isSuperuser(coinbase);
  if(!isSuperuser){
    alert('Only for superuser');
    return;
  }

  const balance = await eth.getBalance(bremInstance.address);
  if(balance < utils.toWei(withdrawAmount, "ether")){
    alert('Insufficient funds on the account');
    return;    
  }
  
          
  let res = await bremInstance.withdrawFees(
              utils.toWei(withdrawAmount, "ether"),
              { from: coinbase }
            );

  alert(res.tx);

}

async function changeWithdrawFee(feePercent) {
  console.log('withdrawFees');
  const { host } = config
  const { web3Instance, web3Account } = store
  const { currentProvider, utils, eth } = web3Instance
  const brem = contract(BREMContract);
  brem.setProvider(currentProvider)
  const coinbase = await eth.getCoinbase()
  store.update({
    web3Coinbase: coinbase
  })

  const bremInstance = await brem.deployed();
  
  if (feePercent < 0 || feePercent > 100) {
    alert("WithdrawFeePercent must be between 0 and 100");
    return;
  }

  const isSuperuser = await bremInstance.isSuperuser(coinbase);
  if(!isSuperuser){
    alert('Only for superuser');
    return;
  }
  

  const currentWithdrawFeePercent = await bremInstance.withdrawFeePercent();
  if (currentWithdrawFeePercent == feePercent) {
    alert("WithdrawFeePercent is already " + feePercent);
    return;
  }

  let res = await bremInstance.setWithdrawFeePercent(feePercent, { from: coinbase });
  return alert(res.tx);
        
}
