package services

import (
	"ChemistryPR/internal/models"
	"fmt"
)

// MolarMassService is a service that provides functionalities
// related to calculating the molar mass of chemical compounds.
// It embeds the ChemicalService to utilize its methods and
// properties.
type MolarMassService struct {
	ChemicalService
}

// MolarMassElementInfo holds information about an element in
// a compound, including its name, weight in the compound,
// the count of atoms of that element, and the weight percentage
// of the element in the compound.
type MolarMassElementInfo struct {
	Symbol           string
	Name             string // Name of the element
	WeightInCompound string // Weight of the element in the compound
	AtomsCount       string // Number of atoms of the element in the compound
	WeightPercent    string // Weight percentage of the element in the compound
}

// MolarMassResponse encapsulates the result of a molar mass
// calculation, including the total general weight and a list
// of element information.
type MolarMassResponse struct {
	Formula  string
	Total    float64                // Total weight of the compound
	Elements []MolarMassElementInfo // Slice of element information
}

// GetResponse processes the provided requestedData string to
// calculate the molar mass response. It attempts to parse
// a chemical compound from the string, retrieve its elements,
// and compute the molar mass details. It returns a MolarMassResponse
// containing the calculated data and an error if anything goes
// wrong during the process.
//
// Parameters:
//   - requestedData: A string representing the chemical compound
//     for which the molar mass is to be calculated.
//
// Returns:
//   - MolarMassResponse: The response containing general weight
//     and detailed information about the elements in the compound.
//   - error: An error indicator, nil if no errors occurred.
func (service MolarMassService) GetResponse(requestedData string) (MolarMassResponse, error) {
	response := MolarMassResponse{}
	response.Elements = nil
	compound, err := service.ParseCompound(requestedData)
	if err != nil {
		return response, err
	}
	elements, err := service.Store.GetElements(compound)
	if err != nil {
		return response, err
	}

	response = service.ComputeData(compound, elements)
	response.Formula = requestedData

	return response, nil
}

// computeData processes the given compound and elements to
// compute the molar mass data. It calculates the general
// weight of the compound and provides detailed information
// about each element in relation to the compound.
//
// Parameters:
//   - compound: The chemical compound whose molar mass needs
//     to be calculated.
//   - elements: A slice of elements involved in the compound.
//
// Returns:
//   - MolarMassResponse: A response containing the general weight
//     and detailed information about the elements in the compound.
func (service MolarMassService) ComputeData(compound models.Compound, elements []models.Element) MolarMassResponse {
	var (
		generalWeight float64                = 0.0
		elementsInfo  []MolarMassElementInfo = make([]MolarMassElementInfo, len(elements))
	)

	sumWeight := 0.0
	for _, element := range elements {
		elementWeight := element.AtomicWeight * float64(compound.Data[element.Symbol])
		sumWeight += elementWeight
	}
	generalWeight = sumWeight

	for i, element := range elements {
		elementsInfo[i].Name = element.Name
		elementsInfo[i].Symbol = element.Symbol
		elementsInfo[i].AtomsCount = fmt.Sprint(compound.Data[element.Symbol])
		weigth := element.AtomicWeight * float64(compound.Data[element.Symbol])
		elementsInfo[i].WeightInCompound = fmt.Sprintf("%.3f", weigth)
		elementsInfo[i].WeightPercent = fmt.Sprintf("%.3f", weigth/generalWeight*100)
	}

	return MolarMassResponse{
		Total:    generalWeight,
		Elements: elementsInfo,
	}
}
