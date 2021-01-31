package equationscanner

type ScanState int

const (
    DuringStart ScanState =      	iota
    DuringNumber
    DuringVariable 
    DuringCoefficientSign
    DuringExponentSign
    DuringExponentNumber
    DuringEqualSign
    DuringMultiplySign
)