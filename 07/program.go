package day07

import "fmt"

type Program struct {
	Name          string
	SelfWeight    int
	Children      []string
	Parent        *Program
	ChildPrograms ProgramSlice
}

type ProgramSlice []*Program
type WeightsMap map[int]ProgramSlice

type WeightError struct {
	Program    *Program
	Difference int
}

// AddChild sets the references to the parent and child program on both nodes
func (program *Program) AddChild(child *Program) {
	program.ChildPrograms = append(program.ChildPrograms, child)

	child.Parent = program
}

func (program *Program) Weight() (weight int, err error) {
	weight = program.SelfWeight

	if len(program.ChildPrograms) == 0 {
		return weight, nil
	}

	weights := make(WeightsMap)

	for _, child := range program.ChildPrograms {
		var weight2 int
		weight2, err = child.Weight()
		if err != nil {
			return 0, err
		}

		weights[weight2] = append(weights[weight2], child)
		weight = weight + weight2
	}

	var imbalance *Program
	var normal, outlier int
	for w, programs := range weights {
		if len(programs) == 1 {
			outlier = w
			imbalance = programs[0]
		} else if len(programs) != len(program.ChildPrograms) {
			normal = w
		}
	}

	if imbalance != nil {
		err = WeightError{Program:imbalance, Difference: outlier - normal}
	}

	return weight, err
}

func (err WeightError) Error() string {
	return fmt.Sprintf("weight anomaly detected, correct weight %d", err.CorrectWeight())
}

func (err WeightError) CorrectWeight() int {
	return err.Program.SelfWeight - err.Difference
}
