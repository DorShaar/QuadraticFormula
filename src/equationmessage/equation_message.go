package equationmessage

type DisassembledEquationMessage struct {
	CorrelationId 		string
	OrigianlEquation 	string
	ArrangedEquation 	string
	A					string
	B 					string
	C					string
	Root1             	float32
	Root2             	float32
	TracedMethodsList 	[]TracedMethod
}