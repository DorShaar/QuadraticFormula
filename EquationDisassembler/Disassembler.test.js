import {disassamble} from "./Disassembler.js";

test("Disassamble valid equations as expected", () => {

    const equationCoefficients = disassamble("-36y^2+385y+3=0", "y");
    expect(equationCoefficients.a).toBe("-36");
    expect(equationCoefficients.b).toBe("385");
    expect(equationCoefficients.c).toBe("3");
    expect(equationCoefficients.equation).toBe("-36y^2+385y+3=0")
});

 // tODO  throws
 // disassamble("-36y^2+385y+3=0", "y");
// disassamble("-36y^2 +385y+3=0", "y");
// disassamble("-36y^2+385y+3= 0", "y");