package services

import (
	"bytes"
	"os/exec"
	"strings"
)

// BalanceService handles chemical balancing operations.  It relies on a ChemicalService for underlying chemical information.
type BalanceService struct {
	ChemicalService
}

// BalanceCompoundInfo represents information about a compound involved in a balanced reaction.
type BalanceCompoundInfo struct {
	Formula    string // Chemical formula (e.g., "H2O")
	Name       string // Name of the compound (e.g., "Water")
	Appearance string // Physical appearance (e.g., "Clear liquid")
}

// BalanceResponse represents the response from a chemical balancing request.
type BalanceResponse struct {
	Reaction string
	Result   string                // Balanced equation
	Reagents []BalanceCompoundInfo // Array of reagent compound information.
	Products []BalanceCompoundInfo // Array of product compound information.
}

// fillCompoundInfo retrieves compound information from the data store and converts it to a slice of BalanceCompoundInfo structs.
//
// Args:
//
//	formulas: A slice of chemical formulas (e.g., ["H2O", "C6H12O6"]).
//
// Returns:
//
//	[]BalanceCompoundInfo: A slice of BalanceCompoundInfo structs containing the retrieved compound information.
//	Returns an empty slice if no compounds are found or an error occurs.
//	error: An error object if there's an issue during retrieval from the data store.
func (service BalanceService) fillCompoundInfo(formulas []string) ([]BalanceCompoundInfo, error) {
	compounds, err := service.Store.GetCompounds(formulas)
	if err != nil {
		return nil, err
	}

	compoundsInfo := make([]BalanceCompoundInfo, len(compounds))
	for i, compound := range compounds {
		newCompoundInfo := BalanceCompoundInfo{
			Formula:    compound.Formula,
			Name:       compound.Name,
			Appearance: compound.Appearance,
		}
		compoundsInfo[i] = newCompoundInfo
	}

	return compoundsInfo, nil
}

// GetResponse balances a chemical equation represented as a string and returns the balanced equation.
//
// It processes the input `requestedData`, which is expected to be in the form of
// "reagents => products", where both reagents and products are separated by a "+" sign.
// The function solves the system of linear equations formed by the stoichiometry of the chemical compounds
// involved and returns the balanced equation as a string in the form of "reactantSide => productSide",
// along with additional information about reagents and products.
//
// Arguments:
//
//	requestedData (string): The chemical equation to be balanced, formatted as "reagents => products".
//
// Returns:
//
//	BalanceResponse: A struct containing the balanced equation in `Result`, and additional compound information
//	in `Reagents` and `Products`.
//	error: An error, if any, that occurred during the balancing process.
//
// Example:
//
//	requestedData := "H2 + O2 => H2O"
//	response, err := service.GetResponse(requestedData)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(response.Result)
func (service BalanceService) GetResponse(requestedData string) (BalanceResponse, error) {
	var response BalanceResponse
	copy := requestedData
	cmd := exec.Command("python3", "balance.py", copy)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	_ = cmd.Run()

	response.Result = stdout.String()

	sides := strings.Split(requestedData, "=")
	reagents := strings.Split(strings.TrimSpace(sides[0]), "+")
	products := strings.Split(strings.TrimSpace(sides[1]), "+")

	response.Reagents, _ = service.fillCompoundInfo(reagents)
	response.Products, _ = service.fillCompoundInfo(products)
	response.Reaction = requestedData

	return response, nil
}
