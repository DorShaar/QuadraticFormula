import {disassamble} from "./Disassembler.js";

const equationCoefficients = disassamble("-36y^2+385y+3=0", "y");
console.log("A: " + equationCoefficients.a);