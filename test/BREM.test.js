const BREM = artifacts.require("BREM");
const BRM = artifacts.require("BRMToken");

const { assertRevert } = require("./helpers/assertRevert");
const { ethGetBalance } = require("./helpers/web3");
const {
  increaseTime,
  increaseTimeTo,
  duration
} = require("./helpers/increaseTime");

contract("BREM", function(accounts) {
  let brem;

  let superuser = accounts[0];
  let dev = accounts[1];
  let auditor1 = accounts[2];
  let auditor2 = accounts[3];
  let anyone = accounts[4];
  let anyone2 = accounts[5];

  let withdrawFeePercent;

  beforeEach(async function() {
    this.withdrawFeePercent = 15;
    this.brem = await BREM.new(
      this.withdrawFeePercent,
      { from: superuser }
    );
  });

  // context("parameters were initialized correctly", () => {
  //   it("brmAddress", async function() {
  //     res = await this.brem.BRM();
  //     assert.equal(res, this.brm.address);
  //   });
  //   it("icoCreationPrice", async function() {
  //     res = await this.brem.icoCreationPrice();
  //     assert.equal(res.toNumber(), this.icoCreationPrice);
  //   });
  //   it("withdrawFeePercent", async function() {
  //     res = await this.brem.withdrawFeePercent();
  //     assert.equal(res.toNumber(), this.withdrawFeePercent);
  //   });
  // });

  // context("addAuditor", () => {
  //   it("allows the superuser to add an Audition", async function() {
  //     await this.brem.addAuditor(auditor1, { from: superuser });

  //     res = await this.brem.isAuditor(auditor1);

  //     assert.equal(res, true);
  //   });
  //   it("does not allow anyone to add an Audition", async function() {
  //     await assertRevert(this.brem.addAuditor(auditor1, { from: anyone }));
  //   });
  // });

  // // context('removeAuditor', () => {
  // //     it('allows the superuser to remove the Auditor', async function () {
  // //         await this.brem.addAuditor(auditor1, {from: superuser});

  // //         await this.brem.removeAuditor(auditor1, {from: superuser});

  // //         res = await this.brem.isAuditor(auditor1);

  // //         assert.equal(res, false);
  // //     });

  // //     it('does not allow anyone to remove the Auditor', async function () {
  // //         await this.brem.addAuditor(auditor1, {from: superuser});
  // //         res = await this.brem.isAuditor(auditor1);
  // //         assert.equal(res, true);

  // //         await assertRevert(
  // //             this.brem.removeAuditor(auditor1, {from: anyone})
  // //         );
  // //     });
  // // });

  // context("addDeveloper", () => {
  //   it("allows anyone to become a developer", async function() {
  //     await this.brem.addDeveloper({ from: dev });

  //     res = await this.brem.isDeveloper(dev, { from: dev });

  //     assert.equal(res, true);
  //   });
  // });

  // context("setIcoCreationPrice", () => {
  //   it("allows the superuser to change the icoCreationPrice", async function() {
  //     this.icoCreationPrice++;
  //     await this.brem.setIcoCreationPrice(this.icoCreationPrice, {
  //       from: superuser
  //     });

  //     res = await this.brem.icoCreationPrice();

  //     assert.equal(res, this.icoCreationPrice);
  //   });
  //   it("does not allow the superuser to reset the same icoCreationPrice", async function() {
  //     res = await this.brem.icoCreationPrice();

  //     await assertRevert(
  //       this.brem.setIcoCreationPrice(res.toNumber(), { from: superuser })
  //     );
  //   });

  //   it("does not allow anyone to change the icoCreationPrice", async function() {
  //     this.icoCreationPrice++;
  //     await assertRevert(
  //       this.brem.setIcoCreationPrice(this.icoCreationPrice, { from: anyone })
  //     );
  //   });
  // });

  // context("setWithdrawFeePercent", () => {
  //   it("allows the superuser to change the withdrawFeePercent", async function() {
  //     this.withdrawFeePercent++;
  //     await this.brem.setWithdrawFeePercent(this.withdrawFeePercent, {
  //       from: superuser
  //     });

  //     res = await this.brem.withdrawFeePercent();

  //     assert.equal(res, this.withdrawFeePercent);
  //   });
  //   it("allows the superuser to set the correct withdrawFeePercent(0<=withdrawFeePercent<=100)", async function() {
  //     this.withdrawFeePercent = 0;
  //     await this.brem.setWithdrawFeePercent(this.withdrawFeePercent, {
  //       from: superuser
  //     });
  //     res = await this.brem.withdrawFeePercent();
  //     assert.equal(res, this.withdrawFeePercent);

  //     this.withdrawFeePercent = 1;
  //     await this.brem.setWithdrawFeePercent(this.withdrawFeePercent, {
  //       from: superuser
  //     });
  //     res = await this.brem.withdrawFeePercent();
  //     assert.equal(res, this.withdrawFeePercent);

  //     this.withdrawFeePercent = 99;
  //     await this.brem.setWithdrawFeePercent(this.withdrawFeePercent, {
  //       from: superuser
  //     });
  //     res = await this.brem.withdrawFeePercent();
  //     assert.equal(res, this.withdrawFeePercent);

  //     this.withdrawFeePercent = 100;
  //     await this.brem.setWithdrawFeePercent(this.withdrawFeePercent, {
  //       from: superuser
  //     });
  //     res = await this.brem.withdrawFeePercent();
  //     assert.equal(res, this.withdrawFeePercent);
  //   });
  //   it("does not allow the superuser to set the incorrect withdrawFeePercent(withdrawFeePercent>100)", async function() {
  //     await assertRevert(
  //       this.brem.setWithdrawFeePercent(101, { from: superuser })
  //     );
  //   });
  //   it("does not allow the superuser to reset the same withdrawFeePercent", async function() {
  //     res = await this.brem.withdrawFeePercent();

  //     await assertRevert(
  //       this.brem.setWithdrawFeePercent(res.toString(), { from: superuser })
  //     );
  //   });

  //   it("does not allow anyone to change the withdrawFeePercent", async function() {
  //     this.withdrawFeePercent++;
  //     await assertRevert(
  //       this.brem.setWithdrawFeePercent(this.withdrawFeePercent, {
  //         from: anyone
  //       })
  //     );
  //   });
  // });

  // context("signUp", () => {
  //   it("allows anyone to add their own name", async function() {
  //     this.userName = "qwerty";
  //     await this.brem.signUp(this.userName, { from: anyone });

  //     res = await this.brem.login({ from: anyone });

  //     assert.equal(res.toString(), this.userName);
  //   });
  //   it("signUp adds a developer role", async function() {
  //     this.userName = "qwerty";
  //     await this.brem.signUp(this.userName, { from: anyone });

  //     res = await this.brem.isDeveloper(anyone);

  //     assert.equal(res, true);
  //   });
  //   it("does not allow to change the name", async function() {
  //     this.userName = "qwerty";
  //     await this.brem.signUp(this.userName, { from: anyone });

  //     this.newUserName = this.userName + "2";
  //     await this.brem.signUp(this.newUserName, { from: anyone });

  //     res = await this.brem.login({ from: anyone });

  //     assert.notEqual(res.toString(), this.newUserName);
  //   });
  // });

  // context("withdrawFees", () => {
  //   it("allows the superuser withdrawFees", async function() {
  //     val = 500;
  //     await this.brem.send(val, { from: anyone2 });

  //     beforeBalance = await ethGetBalance(superuser);
  //     await this.brem.withdrawFees(val, { from: superuser, gasPrice: 0 });
  //     await increaseTime(3600);

  //     res = await ethGetBalance(superuser);
  //     assert.equal(res.c[1] - beforeBalance.c[1], val);
  //   });

  //   it("does not allow anyone withdrawFees", async function() {
  //     val = 500;
  //     await this.brem.send(val, { from: anyone2 });

  //     await assertRevert(this.brem.withdrawFees(val, { from: anyone }));
  //   });
  // });
});
