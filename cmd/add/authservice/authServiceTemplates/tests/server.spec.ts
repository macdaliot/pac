import 'mocha';
    import * as chai from 'chai';
    const should = chai.should();
    const expect = chai.expect;
    const assert = chai.assert;
    
    
    describe('Array', function () {
        describe('#indexOf()', function () {
            it('should return -1 when the value is not present', function () {
                assert.equal([1, 2, 3].indexOf(4), -1);
            });
        });
    });
    
