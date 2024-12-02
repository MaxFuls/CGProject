package parsing

import "errors"

func GetMolarMass(name string) (float64, error) {
	chem_table := map[string]float64{
		"H":  1.00797,
		"He": 4.0026,
		"Li": 6.939,
		"Be": 9.0122,
		"B":  10.811,
		"C":  12.01115,
		"N":  14.0067,
		"O":  15.9994,
		"F":  18.9984,
		"Ne": 20.179,
		"Na": 22.9898,
		"Mg": 24.305,
		"Al": 26.9815,
		"Si": 28.086,
		"P":  30.9738,
		"S":  32.064,
		"Cl": 35.453,
		"Ar": 39.948,
		"K":  39.102,
		"Ca": 40.08,
		"Sc": 44.956,
		"Ti": 47.90,
		"V":  50.942,
		"Cr": 51.996,
		"Mn": 54.9380,
		"Fe": 55.847,
		"Co": 58.9330,
		"Ni": 58.71,
		"Cu": 63.546,
		"Zn": 65.37,
		"Ga": 69.72,
		"ge": 72.59,
		"As": 74.9216,
		"Se": 78.96,
		"Br": 79.904,
		"Kr": 83.80,
	}
	mass, is_exist := chem_table[name]
	if is_exist {
		return mass, nil
	} else {
		return 0.0, errors.New("incorrect input")
	}
}
