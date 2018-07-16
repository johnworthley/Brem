const BREM = artifacts.require('BREM');
const BRM = artifacts.require('BRMToken');

contract('BREM', function(accounts) {

    let brem;
    let brm;

    before(async function() {
        this.brm = await BRM.new();
        this.brem = await BREM.new(this.brm.address);
        await this.brm.transferOwnership(this.brem.address);
    });

    it('Check superuser address', async function() {
        console.log(accounts);
    });
});