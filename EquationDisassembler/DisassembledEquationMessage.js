export class DisassembledEquationMessage {
  constructor(equation, a, b, c, startTime, endTime) {
    this.equation = equation;
    this.a = a;
    this.b = b;
    this.c = c;
    this.startTime = startTime;
    this.endTime = endTime;
  }
}