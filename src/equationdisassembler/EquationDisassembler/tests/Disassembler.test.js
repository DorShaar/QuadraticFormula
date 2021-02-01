import {disassamble} from "../Disassembler.js";
import assert from 'assert';

describe("Disassambler tests", function() {
    it("Disassamble valid equations as expected", function() {
        let equationCoefficients = disassamble("-36y^2+385y+3=0", "y");
        assert.strictEqual("-36", equationCoefficients.a);
        assert.strictEqual("385", equationCoefficients.b);
        assert.strictEqual("3", equationCoefficients.c);
        assert.strictEqual("-36y^2+385y+3=0", equationCoefficients.equation);

        equationCoefficients = disassamble("3x^2+x+3=0", "x");
        assert.strictEqual("3", equationCoefficients.a);
        assert.strictEqual("1", equationCoefficients.b);
        assert.strictEqual("3", equationCoefficients.c);
        assert.strictEqual("3x^2+x+3=0", equationCoefficients.equation);

        equationCoefficients = disassamble("x^2+2x+3=0", "x");
        assert.strictEqual("1", equationCoefficients.a);
        assert.strictEqual("2", equationCoefficients.b);
        assert.strictEqual("3", equationCoefficients.c);
        assert.strictEqual("x^2+2x+3=0", equationCoefficients.equation);

        equationCoefficients = disassamble("-x^2+2x+3=0", "x");
        assert.strictEqual("-1", equationCoefficients.a);
        assert.strictEqual("2", equationCoefficients.b);
        assert.strictEqual("3", equationCoefficients.c);
        assert.strictEqual("-x^2+2x+3=0", equationCoefficients.equation);
    });

    it("Disassamble valid equations as expected with z variable", function() {
        let equationCoefficients = disassamble("-36z^2+385z+3=0", "z");
        assert.strictEqual("-36", equationCoefficients.a);
        assert.strictEqual("385", equationCoefficients.b);
        assert.strictEqual("3", equationCoefficients.c);
        assert.strictEqual("-36z^2+385z+3=0", equationCoefficients.equation);

        equationCoefficients = disassamble("36z^2-385z-30=0", "z");
        assert.strictEqual("36", equationCoefficients.a);
        assert.strictEqual("-385", equationCoefficients.b);
        assert.strictEqual("-30", equationCoefficients.c);
        assert.strictEqual("36z^2-385z-30=0", equationCoefficients.equation);
    });

    it("Disassamble invalid equations - throws an exception", function() {
        assert.throws(() => disassamble("-36x^2+385y+3=0", "x"), SyntaxError);
        assert.throws(() => disassamble("-36y^2+385y+3=0", "x"), SyntaxError);
        assert.throws(() => disassamble("-36y^2+385y+3", "y"), SyntaxError);
        assert.throws(() => disassamble("-36y^2 +385y+3=0", "y"), SyntaxError);
    });
});