package registers

import (
	"fmt"
	"log"
)

func ParseLine(input string) (action *Instruction, condition *Condition) {
	action = new(Instruction)
	condition = new(Condition)

	_, err := fmt.Sscanf(input, "%s %s %d if %s %s %d", &action.Register, &action.Operator, &action.Value, &condition.Register, &condition.Comparator, &condition.Value)
	if err != nil {
		log.Fatal(err)
	}

	return action, condition
}
