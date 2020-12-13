import {disassamble, createMessage} from "./Disassembler.js";

// A temp js script to run manual tests on Disassembler.

var x = disassamble("3x^2+x+3=0", "x");

createMessage(x);