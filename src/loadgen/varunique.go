package loadgen

func NewVarUnique(varconfig *LsVarConfig) error {
	varconfig.Var = &VarUnique{name: "123"}
	return nil
}

// VarUnique

type VarUnique struct {
	name string
	// LsVar
}

func (v *VarUnique) NewSessionVar() (LsSessionVar, error) {
	return &VarUniqueSession{}, nil
}

func (v *VarUnique) Terminate() {
}

// VarUniqueSession

type VarUniqueSession struct {
	LsSessionVar
}

func (v *VarUniqueSession) SessionIterate() error {
	return nil
}

func (v *VarUniqueSession) SessionTerminate() {
}

func (v *VarUniqueSession) GetNextValue() (string, error) {
	return "1", nil
}
