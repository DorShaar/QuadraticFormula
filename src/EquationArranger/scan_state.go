package equationscanner

type ScanState int

const (
    DuringStart =      	iota
    DuringNumber
    DuringVariable 
    DuringCoefficientSign
    DuringExponentSign
    DuringExponentNumber
    DuringEqualSign
    DuringMultiplySign
)