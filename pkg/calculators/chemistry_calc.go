package calculators

type ChemicalEquation struct {
	Compounds         []string `yaml:"conpounds"`
	LeftCompundsCount uint16   `yaml:"left_compounds_count"`
}
