package steps

import (
	"fmt"

	"github.com/Samasource/jen/src/internal/evaluation"
	"github.com/Samasource/jen/src/internal/exec"
	logging "github.com/Samasource/jen/src/internal/logging"
)

// If represents a conditional step that executes its child executable only if
// a given condition evaluates to true
type If struct {
	Condition string
	Then      exec.Executables
}

func (i If) String() string {
	return "do"
}

// Execute executes a child action only when a given condition evaluates to true
func (i If) Execute(context exec.Context) error {
	result, err := evaluation.EvalBoolExpression(context.(evaluation.Context), i.Condition)
	if err != nil {
		return fmt.Errorf("evaluating if conditional: %w", err)
	}
	if !result {
		logging.Log("Skipping sub-steps because condition %q evaluates to false", i.Condition)
		return nil
	}
	logging.Log("Executing sub-steps because condition %q evaluates to true", i.Condition)
	return i.Then.Execute(context)
}
