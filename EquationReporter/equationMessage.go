package equationreporter

// EquationMessage holds equation information.
type EquationMessage struct {
	correlationID     string
	originalEquation  string
	organizedEquation string
	a                 int
	b                 int
	c                 int
	root1             float32
	root2             float32
	tracedMethodsList []TracedMethod
}
