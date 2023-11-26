package domain

type SetupField struct {
	Identifier   string
	Name         string
	Type         string
	SelectValues []string
	Value        string
}

type SetupStep struct {
	StepIdentifier string
	VerifyDone     func() bool
	RequiredFields func() []SetupField
}
