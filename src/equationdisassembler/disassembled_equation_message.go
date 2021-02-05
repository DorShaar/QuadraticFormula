package equationdisassembler

type DisassembledEquationMessage struct {
	CorrelationId 		string
	Equation 			string
	A					string
	B 					string
	C					string
	IsDisassembleFailed bool
}