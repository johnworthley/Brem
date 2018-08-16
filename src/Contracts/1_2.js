import axios from "axios";
import BREMICOContract from "../../build/contracts/BREMICO.json"
import web3Provider from "../util/getweb3";

const server = "http://46.101.99.139:8080"
const contract = require("truffle-contract");

// Получение данных для макета 1.2 (Маркетплейс)
export default async () => {
  const web3 = web3Provider.getWeb3()

  // Получение страницы c ico (роуты для фильтрации в свагере)
  const page = 0;
  axios.get(server + "/ico/all", {
    params: {
      page: page
    }
  }).then(icoListResponse => {
    const icoList = icoListResponse.data
    // Получили список ico. Если он null, то такой страницы пока нет
    // Что нужно вытащить из респонса посмотри в свагере
    if (icoList) {
      for (let i = 0; i < icoList.length; i++) {
        const icoAddress = icoList[i].address
        // Получить изображение для ICO
        axios.get(server + "/ico/image", {
          params: {
            address: icoAddress
          },
          responseType: "arraybuffer"
        })
        .then(res => {
          const base64 = Buffer.from(res.data, "binary").toString(
            "base64"
          );
          // Само изображение this.setState({ img: "data:image/jpeg;base64," + base64 });
        })
        // Вытаскивание данных с web3
        const ico = contract(BREMICOContract)
        ico.setProvider(web3.currentProvider)

        ico.deployed().then(icoInstance => {
          // Cap
          icoInstance.cap().then(cap => {
            const capEth = web3.utils.fromWei(cap, "ether")
          })
          // Сколько собрано на данный момент
          icoInstance.weiRaised(weiRaised => {
            const raised = web3.utils.fromWei(weiRaised, "ether")
          })
          // Для внесенных личных средств скоро добавим функцкии на контрактах
        })
      }
    }
  })
}