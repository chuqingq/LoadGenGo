package loadgen

func NewVarFile(varconfig *LsVarConfig) error {
	varconfig.Var = &VarFile{}
	return nil
}

// VarFile

type VarFile struct {
	name string
	// LsVar
}

func (v *VarFile) NewSessionVar() (LsSessionVar, error) {
	return &VarFileSession{}, nil
}

func (v *VarFile) Terminate() {

}

// VarFileSession

type VarFileSession struct {
	LsSessionVar
}

func (v *VarFileSession) SessionIterate() error {
	return nil
}

func (v *VarFileSession) SessionTerminate() {
}

func (v *VarFileSession) GetNextValue() (string, error) {
	return "1", nil
}
