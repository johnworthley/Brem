import BREMContract from "../../build/contracts/BREM.json";
import web3Provider from "../util/getweb3";

const contract = require("truffle-contract");

// Получение данных с контрактов для макета 1.3.1 (Настройки суперюзера)
export default async () => {
  const ab = async => {
    const web3 = web3Provider.getWeb3()

    const brem = contract(BREMContract)
    brem.setProvider(web3.currentProvider)

    // Адресс фактори контракта
    const bremAddress = brem.address
    // Баланс контракта
    let bremBalanceInEther
    web3.eth.getBalance(bremAddress).then(balanceInWei => {
      bremBalanceInEther = web3.utils.fromWei(balanceInWei.toNumber(), "ether")
    })
    
    brem.deployed().then(bremInstance => {
      // Текущий процент комисии
      bremInstance.withdrawFeePercent().then(fee => {
        const currentFee = fee.toNumber()
      })
    })
  }
}

// Изменение комиссии
// newFeePercent - новый процент (от 0 до 100)
export function changeFee(newFeePercent) {
  const web3 = web3Provider.getWeb3()

  const brem = contract(BREMContract)
  brem.setProvider(web3.currentProvider)

  web3.eth.getCoinbase((error, coinbase) => {
    if (error) return console.log(error)

    brem.deployed().then(bremInstance => {
      // Проверить на адрес чтобы лишний раз не дергать метамаск
      bremInstance.isSuperuser(coinbase).then(isSuperuser => {
        if (!isSuperuser) return

        // Проверить на старое значение, чтобы лишний раз не дерагать метамаск
        bremInstance.withdrawFeePercent().then(oldFee => {
          const oldFeePercent = oldFee.toNumber()
          if (oldFeePercent === newFeePercent) return;

          // Изменить значение
          bremInstance.setWithdrawFeePercent(newFeePercent, {from: coinbase}).then(res => {
            // Показываем хэш транзакции
            console.log(res.tx)
          })
          .catch(err => console.error(err))
        })
      })
    })
  })
}

// Добавление нового аудитора
// auditorAddress - адрес нового аудитора
// проверка ввода адреса: web3.web3.utils.isAddress(address)
export function addNewAuditor(auditorAddress) {
  const web3 = web3Provider.getWeb3()

  const brem = contract(BREMContract)
  brem.setProvider(web3.currentProvider)

  web3.eth.getCoinbase((error, coinbase) => {
    if (error) return console.log(error)

    brem.deployed().then(bremInstance => {
      // Проверить на адрес чтобы лишний раз не дергать метамаск
      bremInstance.isSuperuser(coinbase).then(isSuperuser => {
        if (!isSuperuser) return

        // Проверить на уже добавленного аудитора, чтобы лишний раз не дерагать метамаск
        bremInstance.isAuditor(auditorAddress).then(isAuditor => {
          if (isAuditor) return

          // Добавить аудитора
          bremInstance.addAuditor(auditorAddress, {from: coinbase}).then(res => {
            // Показываем хэш транзакции
            console.log(res.tx)
          })
          .catch(err => console.error(err))
        })
      })
    })
  })
}