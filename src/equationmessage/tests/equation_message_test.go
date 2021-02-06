package equationmessagetests

import (
	"equationmessage"
	"testing"
)

func TestSerializeToJson_SerializedAsExpected(t *testing.T) {
	tracedMethods := make([]equationmessage.TracedMethod, 0)
	tracedMethods = append(tracedMethods, equationmessage.TracedMethod { MethodName: "calculate equation", Duration:  5004 })
	tracedMethods = append(tracedMethods, equationmessage.TracedMethod { MethodName: "arrange equation", Duration:  3604 })

	equationMessage := equationmessage.EquationMessage {
		CorrelationId: "1234-5678-9abc-cdef-ghij",
		OriginalEquation:"x^2 + 3x + 5 = -3x",
		ArrangedEquation: "x^2+6x+5=0",
		A: "1",
		B: "6",
		C: "5",
		Root1: 3.5,
		Root2: -4.2,
		TracedMethodsList: tracedMethods,
	}

	expectedSerializedMessage := "{\"CorrelationId\":\"1234-5678-9abc-cdef-ghij\",\"OriginalEquation\":\"x^2 + 3x + 5 = -3x\",\"ArrangedEquation\":\"x^2+6x+5=0\",\"A\":\"1\",\"B\":\"6\",\"C\":\"5\",\"Root1\":3.5,\"Root2\":-4.2,\"TracedMethodsList\":[{\"MethodName\":\"calculate equation\",\"Duration\":5004},{\"MethodName\":\"arrange equation\",\"Duration\":3604}]}"

	serializedMessage, err := equationMessage.SerializeToJson()

	if err != nil {
		t.FailNow()
	}

	if expectedSerializedMessage != serializedMessage {
		t.FailNow()
	}
}

func TestDeserializeFromJson_DeserializedAsExpected(t *testing.T) {
	tracedMethods := make([]equationmessage.TracedMethod, 0)
	tracedMethods = append(tracedMethods, equationmessage.TracedMethod { MethodName: "calculate equation", Duration:  5004 })
	tracedMethods = append(tracedMethods, equationmessage.TracedMethod { MethodName: "arrange equation", Duration:  3604 })

	expectedEquationMessage := equationmessage.EquationMessage {
		CorrelationId: "1234-5678-9abc-cdef-ghij",
		OriginalEquation:"x^2 + 3x + 5 = -3x",
		ArrangedEquation: "x^2+6x+5=0",
		A: "1",
		B: "6",
		C: "5",
		Root1: 3.5,
		Root2: -4.2,
		TracedMethodsList: tracedMethods,
	}

	serializedMessage := "{\"CorrelationId\":\"1234-5678-9abc-cdef-ghij\",\"OriginalEquation\":\"x^2 + 3x + 5 = -3x\",\"ArrangedEquation\":\"x^2+6x+5=0\",\"A\":\"1\",\"B\":\"6\",\"C\":\"5\",\"Root1\":3.5,\"Root2\":-4.2,\"TracedMethodsList\":[{\"MethodName\":\"calculate equation\",\"Duration\":5004},{\"MethodName\":\"arrange equation\",\"Duration\":3604}]}"

	deserializedEquationMessage, err := equationmessage.DeserializeFromJson(serializedMessage)

	if err != nil {
		t.FailNow()
	}

	if expectedEquationMessage.CorrelationId != deserializedEquationMessage.CorrelationId {
		t.FailNow()
	}

	if expectedEquationMessage.OriginalEquation != deserializedEquationMessage.OriginalEquation {
		t.FailNow()
	}

	if expectedEquationMessage.ArrangedEquation != deserializedEquationMessage.ArrangedEquation {
		t.FailNow()
	}

	if expectedEquationMessage.A != deserializedEquationMessage.A {
		t.FailNow()
	}

	if expectedEquationMessage.B != deserializedEquationMessage.B {
		t.FailNow()
	}

	if expectedEquationMessage.C != deserializedEquationMessage.C {
		t.FailNow()
	}

	if expectedEquationMessage.Root1 != deserializedEquationMessage.Root1 {
		t.FailNow()
	}

	if expectedEquationMessage.Root2 != deserializedEquationMessage.Root2 {
		t.FailNow()
	}

	if expectedEquationMessage.TracedMethodsList[0].MethodName != deserializedEquationMessage.TracedMethodsList[0].MethodName {
		t.FailNow()
	}

	if expectedEquationMessage.TracedMethodsList[0].Duration != deserializedEquationMessage.TracedMethodsList[0].Duration {
		t.FailNow()
	}

	if expectedEquationMessage.TracedMethodsList[1].MethodName != deserializedEquationMessage.TracedMethodsList[1].MethodName {
		t.FailNow()
	}

	if expectedEquationMessage.TracedMethodsList[1].Duration != deserializedEquationMessage.TracedMethodsList[1].Duration {
		t.FailNow()
	}
}