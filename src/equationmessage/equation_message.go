package equationmessage

import (
	"encoding/json"
	)

type EquationMessage struct {
	CorrelationId 		string
	OriginalEquation 	string
	ArrangedEquation 	string
	A					string
	B 					string
	C					string
	Root1             	float32
	Root2             	float32
	TracedMethodsList 	[]TracedMethod
}

func (equationMessage *EquationMessage) SerializeToJson() (string, error) {
	serializedMessage, err := json.Marshal(equationMessage)
	if err != nil {
		return "error", err
	}

	return string(serializedMessage), nil
}

func DeserializeFromJson(serializedMessage string) (*EquationMessage ,error) {
	equationMessage := EquationMessage{}
	err := json.Unmarshal([]byte(serializedMessage), &equationMessage)

	if err != nil {
		return nil, err
	}

	return &equationMessage, nil
}