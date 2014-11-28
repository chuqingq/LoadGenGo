package loadgen

func NewVarRandom(varconfig *LsVarConfig) error {
	varconfig.Var = &VarRandom{}
	return nil
}

// VarRandom

type VarRandom struct {
	// LsVar
}

func (v *VarRandom) NewSessionVar() (LsSessionVar, error) {
	return &VarRandomSession{}, nil
}

func (v *VarRandom) Terminate() {

}

// VarRandomSession

type VarRandomSession struct {
	LsSessionVar
}

func (v *VarRandomSession) SessionIterate() error {
	return nil
}

func (v *VarRandomSession) SessionTerminate() {
}

func (v *VarRandomSession) GetNextValue() (string, error) {
	return "1", nil
}
