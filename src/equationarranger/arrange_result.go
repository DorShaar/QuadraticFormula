package equationarranger

type ArrangeResult interface {
	OriginalEquation() 		string
	ArrangedEquation()		string
	IsArrangeSucceeded() 	bool
}

type SuccessArrangeResult struct {
	originalEquation 	string
	arrangedEquation	string
}

func (successArrangeResult *SuccessArrangeResult) OriginalEquation() string {
	return successArrangeResult.originalEquation
}

func (successArrangeResult *SuccessArrangeResult) ArrangedEquation() string {
	return successArrangeResult.arrangedEquation
}

func (successArrangeResult *SuccessArrangeResult) IsArrangeSucceeded() bool {
	return true
}

type FailedArrangeResult struct {
	originalEquation 	string
	isArrangeSucceeded	bool
}

func (failedArrangeResult *FailedArrangeResult) OriginalEquation() string {
	return failedArrangeResult.originalEquation
}

func (failedArrangeResult *FailedArrangeResult) ArrangedEquation() string {
	return ""
}

func (failedArrangeResult *FailedArrangeResult) IsArrangeSucceeded() bool {
	return false
}