package loadgen

func NewVarDate(varconfig *LsVarConfig) error {
	varconfig.Var = &VarDate{}
	return nil
}

// VarDate

type VarDate struct {
	// LsVar
}

func (v *VarDate) NewSessionVar() (LsSessionVar, error) {
	return &VarDateSession{}, nil
}

func (v *VarDate) Terminate() {

}

// VarDateSession

type VarDateSession struct {
	LsSessionVar
}

func (v *VarDateSession) SessionIterate() error {
	return nil
}

func (v *VarDateSession) SessionTerminate() {
}

func (v *VarDateSession) GetNextValue() (string, error) {
	return "1", nil
}
