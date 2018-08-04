const BREM = artifacts.require("BREM");
const BRM = artifacts.require("BRMToken");
const BREMICO = artifacts.require("BREMICO");
const BREMToken = artifacts.require("BREMToken");

const { assertRevert } = require("./helpers/assertRevert");
const { latestTime } = require("./helpers/latestTime");
const {
  increaseTime,
  increaseTimeTo,
  duration
} = require("./helpers/increaseTime");
const { ethGetBalance } = require("./helpers/web3");

contract("BREMICO", function(accounts) {
  let brem;
  let brm;
  let bremToken;
  let bremIco;

  let superuser = accounts[0];
  let dev = accounts[1];
  let auditor1 = accounts[2];
  let auditor2 = accounts[3];
  let auditor3 = accounts[4];
  let anyone = accounts[5];
  let anyone2 = accounts[6];

  beforeEach(async function() {
    this.bname = "test";
    this.bsymbol = "tst";
    this.brate = 1;
    this.bcap = 10000;
    this.bopeningTime = await latestTime();
    this.bduringTime = duration.weeks(2);
    this.bclosingTime = this.bopeningTime + this.bduringTime;
    this.bdescription = "sometext";
    this.bdocHash = "somedocHash123456789123456789123";
    this.icoCreationPrice = 0;
    this.withdrawFeePercent = 5;

    this.brm = await BRM.new({ from: superuser });
    this.brem = await BREM.new(
      this.brm.address,
      this.icoCreationPrice,
      this.withdrawFeePercent,
      { from: superuser }
    );
    await this.brm.transferOwnership(this.brem.address, { from: superuser });

    await this.brem.addDeveloper({ from: dev });

    res = await this.brem.createBREMICO(
      this.bname,
      this.bsymbol,
      this.brate,
      this.bcap,
      this.bclosingTime,
      this.bdescription,
      this.bdocHash,
      { from: dev }
    );

    this.bremToken = await BREMToken.at(res.logs[0].args.tokenAddress);
    this.bremIco = await BREMICO.at(res.logs[0].args.icoAddress);
  });

  context("parameters were initialized correctly", () => {
    it("name", async function() {
      res = await this.bremIco.name();
      assert.equal(res, this.bname);
    });
    it("brem", async function() {
      res = await this.bremIco.brem();
      assert.equal(res, this.brem.address);
    });
    it("wallet", async function() {
      res = await this.bremIco.wallet();
      assert.equal(res, dev);
    });

    it("rate", async function() {
      res = await this.bremIco.rate();
      assert.equal(res.toNumber(), this.brate);
    });
    it("cap", async function() {
      res = await this.bremIco.cap();
      assert.equal(res.toNumber(), this.bcap);
    });
    it("closingTime", async function() {
      res = await this.bremIco.closingTime();
      assert.equal(res.toNumber(), this.bclosingTime);
    });
    it("description", async function() {
      res = await this.bremIco.description();
      assert.equal(res, this.bdescription);
    });
    it("docHash", async function() {
      res = await this.bremIco.docHash();
      assert.equal(res, this.bdocHash);
    });
    it("withdrawFeePercent", async function() {
      res = await this.bremIco.withdrawFeePercent();
      assert.equal(res.toNumber(), this.withdrawFeePercent);
    });

    it("name token", async function() {
      res = await this.bremToken.name();
      assert.equal(res, this.bname);
    });
    it("symbol token", async function() {
      res = await this.bremToken.symbol();
      assert.equal(res, this.bsymbol);
    });
  });

  context("addAuditor", () => {
    it("allows the superuser to add a local Auditor", async function() {
      await this.brem.addAuditor(auditor1, { from: superuser });
      await this.bremIco.addAuditor(auditor1, { from: superuser });

      res = await this.bremIco.auditorsAmount();
      assert.equal(res.toNumber(), 1);
    });
    it("does not allow the superuser to add a local Auditor, if the Auditor does not exist globally", async function() {
      await assertRevert(this.bremIco.addAuditor(auditor1));
    });
    it("does not allow anyone to add a local Auditor", async function() {
      await this.brem.addAuditor(auditor1, { from: superuser });

      await assertRevert(this.bremIco.addAuditor(auditor1, { from: dev }));

      await assertRevert(this.bremIco.addAuditor(auditor1, { from: anyone }));
    });
    it("does not allow the superuser to add a local Auditor second time", async function() {
      await this.brem.addAuditor(auditor1, { from: superuser });
      await this.bremIco.addAuditor(auditor1, { from: superuser });

      await assertRevert(
        this.bremIco.addAuditor(auditor1, { from: superuser })
      );
    });
    it("does not allow the superuser to add a local Auditor after finishAuditorSelection", async function() {
      await this.brem.addAuditor(auditor1, { from: superuser });
      await this.bremIco.addAuditor(auditor1, { from: superuser });
      await this.brem.addAuditor(auditor2, { from: superuser });
      await this.bremIco.finishAuditorSelection({ from: superuser });

      await assertRevert(
        this.bremIco.addAuditor(auditor2, { from: superuser })
      );
    });
  });

  context("finishAuditorSelection", () => {
    it("does not allow the superuser finishAuditorSelection when auditorsAmount == 0", async function() {
      await assertRevert(
        this.bremIco.finishAuditorSelection({ from: superuser })
      );
    });
    it("does not allow anyone to finishAuditorSelection", async function() {
      await this.brem.addAuditor(auditor1, { from: superuser });
      await this.bremIco.addAuditor(auditor1, { from: superuser });

      await assertRevert(this.bremIco.finishAuditorSelection({ from: dev }));
      await assertRevert(this.bremIco.finishAuditorSelection({ from: anyone }));
    });
    it("allow the superuser to finishAuditorSelection", async function() {
      await this.brem.addAuditor(auditor1, { from: superuser });
      await this.bremIco.addAuditor(auditor1, { from: superuser });
      await this.bremIco.finishAuditorSelection({ from: superuser });

      res = await this.bremIco.auditSelected();

      assert.equal(res, true);
    });
  });

  context("withdraw", () => {
    it("does not allow the superuser to withdraw", async function() {
      await this.brem.addAuditor(auditor1, { from: superuser });
      await this.bremIco.addAuditor(auditor1, { from: superuser });
      await this.bremIco.finishAuditorSelection({ from: superuser });

      await this.bremIco.buyTokens(anyone, {
        from: anyone,
        value: this.bcap * 2
      });

      await increaseTimeTo(this.bclosingTime + 3600);

      await assertRevert(this.bremIco.withdraw(this.bcap, { from: superuser }));
    });

    it("does not allow anyone to withdraw", async function() {
      await this.brem.addAuditor(auditor1, { from: superuser });
      await this.bremIco.addAuditor(auditor1, { from: superuser });
      await this.bremIco.finishAuditorSelection({ from: superuser });

      await this.bremIco.buyTokens(anyone, {
        from: anyone,
        value: this.bcap * 2
      });

      await increaseTimeTo(this.bclosingTime + 3600);

      await assertRevert(this.bremIco.withdraw(this.bcap, { from: anyone }));
    });
    it("allows the dev to withdraw", async function() {
      await this.brem.addAuditor(auditor1, { from: superuser });
      await this.bremIco.addAuditor(auditor1, { from: superuser });
      await this.bremIco.finishAuditorSelection({ from: superuser });

      await this.bremIco.buyTokens(anyone, {
        from: anyone,
        value: this.bcap * 2
      });

      await increaseTimeTo(this.bclosingTime + 3600);

      this.bremIco.withdraw(this.bcap, { from: dev });
    });

    it("does not allow creation of a new withdrawal, if the decision on the previous one was not made", async function() {
      await this.brem.addAuditor(auditor1, { from: superuser });
      await this.bremIco.addAuditor(auditor1, { from: superuser });
      await this.bremIco.finishAuditorSelection({ from: superuser });

      await this.bremIco.buyTokens(anyone, {
        from: anyone,
        value: this.bcap * 3
      });

      await increaseTimeTo(this.bclosingTime + 3600);

      await this.bremIco.withdraw(this.bcap, { from: dev });

      await assertRevert(this.bremIco.withdraw(this.bcap, { from: dev }));
    });

    it("allows creation of a new withdrawal, if the decision on the previous one was made", async function() {
      await this.brem.addAuditor(auditor1, { from: superuser });
      await this.bremIco.addAuditor(auditor1, { from: superuser });
      await this.bremIco.finishAuditorSelection({ from: superuser });

      await this.bremIco.buyTokens(anyone, {
        from: anyone,
        value: this.bcap * 3
      });

      await increaseTimeTo(this.bclosingTime + 3600);

      await this.bremIco.withdraw(this.bcap, { from: dev });

      await this.bremIco.confirmWithdraw({ from: auditor1 });

      await this.bremIco.withdraw(this.bcap, { from: dev });
    });

    it("does not allow creation of a new withdrawal during the ico", async function() {
      await this.brem.addAuditor(auditor1, { from: superuser });
      await this.bremIco.addAuditor(auditor1, { from: superuser });
      await this.bremIco.finishAuditorSelection({ from: superuser });

      await this.bremIco.buyTokens(anyone, {
        from: anyone,
        value: this.bcap * 2
      });

      await assertRevert(this.bremIco.withdraw(this.bcap, { from: dev }));
    });

    it("does not allow creation of a new withdrawal if the cap has not been reached", async function() {
      await this.brem.addAuditor(auditor1, { from: superuser });
      await this.bremIco.addAuditor(auditor1, { from: superuser });
      await this.bremIco.finishAuditorSelection({ from: superuser });

      await this.bremIco.buyTokens(anyone, {
        from: anyone,
        value: this.bcap / 2
      });

      await increaseTimeTo(this.bclosingTime + 3600);

      await assertRevert(this.bremIco.withdraw(this.bcap, { from: dev }));
    });
  });

  context("confirmWithdraw", () => {
    it("works with 1 auditor", async function() {
      await this.brem.addAuditor(auditor1, { from: superuser });
      await this.bremIco.addAuditor(auditor1, { from: superuser });
      await this.bremIco.finishAuditorSelection({ from: superuser });

      await this.bremIco.buyTokens(anyone, {
        from: anyone,
        value: this.bcap * 3
      });

      await increaseTimeTo(this.bclosingTime + 3600);

      await this.bremIco.withdraw(this.bcap, { from: dev });

      beforeBalance = await ethGetBalance(dev);

      await this.bremIco.confirmWithdraw({ from: auditor1 });

      res = await ethGetBalance(this.brem.address);
      assert.equal(res.toString(), (this.bcap / 100) * this.withdrawFeePercent);

      res = (await ethGetBalance(dev)).c[1] - beforeBalance.c[1];
      assert.equal(
        res,
        this.bcap - (this.bcap / 100) * this.withdrawFeePercent
      );
    });

    it("works with several auditors", async function() {
      await this.brem.addAuditor(auditor1, { from: superuser });
      await this.bremIco.addAuditor(auditor1, { from: superuser });

      await this.brem.addAuditor(auditor2, { from: superuser });
      await this.bremIco.addAuditor(auditor2, { from: superuser });

      await this.bremIco.finishAuditorSelection({ from: superuser });
      await this.bremIco.buyTokens(anyone, {
        from: anyone,
        value: this.bcap * 3
      });
      await increaseTimeTo(this.bclosingTime + 3600);

      await this.bremIco.withdraw(this.bcap, { from: dev });

      beforeBalance = await ethGetBalance(dev);

      await this.bremIco.confirmWithdraw({ from: auditor1 });

      res = await ethGetBalance(this.brem.address);
      assert.equal(res.toString(), 0);

      res = await ethGetBalance(dev);
      assert.equal(res.c[1], beforeBalance.c[1]);

      await this.bremIco.confirmWithdraw({ from: auditor2 });

      res = await ethGetBalance(this.brem.address);
      assert.equal(res.toString(), (this.bcap / 100) * this.withdrawFeePercent);

      res = (await ethGetBalance(dev)).c[1] - beforeBalance.c[1];
      assert.equal(
        res,
        this.bcap - (this.bcap / 100) * this.withdrawFeePercent
      );
    });

    it("does not allow 1 auditor confirmWithdraw second time", async function() {
      await this.brem.addAuditor(auditor1, { from: superuser });
      await this.bremIco.addAuditor(auditor1, { from: superuser });

      await this.brem.addAuditor(auditor2, { from: superuser });
      await this.bremIco.addAuditor(auditor2, { from: superuser });

      await this.bremIco.finishAuditorSelection({ from: superuser });
      await this.bremIco.buyTokens(anyone, {
        from: anyone,
        value: this.bcap * 3
      });
      await increaseTimeTo(this.bclosingTime + 3600);

      await this.bremIco.withdraw(this.bcap, { from: dev });

      beforeBalance = await ethGetBalance(dev);

      await this.bremIco.confirmWithdraw({ from: auditor1 });

      await assertRevert(this.bremIco.confirmWithdraw({ from: auditor1 }));
    });

    it("does not allow anyone to confirmWithdraw", async function() {
      await this.brem.addAuditor(auditor1, { from: superuser });
      await this.bremIco.addAuditor(auditor1, { from: superuser });

      await this.bremIco.finishAuditorSelection({ from: superuser });
      await this.bremIco.buyTokens(anyone, {
        from: anyone,
        value: this.bcap * 3
      });
      await increaseTimeTo(this.bclosingTime + 3600);

      await this.bremIco.withdraw(this.bcap, { from: dev });

      await assertRevert(this.bremIco.confirmWithdraw({ from: anyone }));

      await assertRevert(this.bremIco.confirmWithdraw({ from: superuser }));
    });
  });

  context("confirmWithdraw", () => {
    it("allows the user with non-zero balance to refund if the cap has not been reached", async function() {
      await this.brem.addAuditor(auditor1, { from: superuser });
      await this.bremIco.addAuditor(auditor1, { from: superuser });
      await this.bremIco.finishAuditorSelection({ from: superuser });

      beforeBalance = await ethGetBalance(anyone);

      await this.bremIco.buyTokens(anyone, {
        from: anyone,
        value: this.bcap / 4,
        gasPrice: 0
      });
      await this.bremIco.buyTokens(anyone2, {
        from: anyone2,
        value: this.bcap / 4
      });
      await increaseTimeTo(this.bclosingTime + 3600);

      this.bremIco.refund({ from: anyone, gasPrice: 0 });

      await increaseTimeTo(this.bclosingTime + 4000);

      res = await ethGetBalance(anyone);
      assert.equal(res.c[1], beforeBalance.c[1]);

      res = await this.bremIco.balances(anyone);
      assert.equal(res, 0);
    });

    it("does not allow confirmWithdraw if capReached", async function() {
      await this.brem.addAuditor(auditor1, { from: superuser });
      await this.bremIco.addAuditor(auditor1, { from: superuser });
      await this.bremIco.finishAuditorSelection({ from: superuser });

      beforeBalance = await ethGetBalance(anyone);

      await this.bremIco.buyTokens(anyone, {
        from: anyone,
        value: this.bcap * 2
      });
      await this.bremIco.buyTokens(anyone2, {
        from: anyone2,
        value: this.bcap / 4
      });
      await increaseTimeTo(this.bclosingTime + 3600);

      await assertRevert(this.bremIco.refund({ from: anyone }));
    });

    it("does not allow confirmWithdraw during the ico", async function() {
      await this.brem.addAuditor(auditor1, { from: superuser });
      await this.bremIco.addAuditor(auditor1, { from: superuser });
      await this.bremIco.finishAuditorSelection({ from: superuser });

      beforeBalance = await ethGetBalance(anyone);

      await this.bremIco.buyTokens(anyone, {
        from: anyone,
        value: this.bcap / 4
      });
      await this.bremIco.buyTokens(anyone2, {
        from: anyone2,
        value: this.bcap / 4
      });

      await assertRevert(this.bremIco.refund({ from: anyone }));
    });
  });

  context("buyTokens", () => {
    it("allow anyone buyTokens", async function() {
      await this.brem.addAuditor(auditor1, { from: superuser });
      await this.bremIco.addAuditor(auditor1, { from: superuser });
      await this.bremIco.finishAuditorSelection({ from: superuser });

      purchaseAmount = this.bcap / 4;
      await this.bremIco.buyTokens(anyone, {
        from: anyone,
        value: purchaseAmount
      });
      await this.bremIco.buyTokens(anyone2, {
        from: anyone2,
        value: this.bcap / 4
      });

      res = await this.bremToken.balanceOf(anyone);
      assert.equal(res, purchaseAmount * this.brate);

      res = await this.bremIco.balancesInToken(anyone);
      assert.equal(res, purchaseAmount * this.brate);

      res = await this.bremIco.balances(anyone);
      assert.equal(res, purchaseAmount);
    });
  });
});
