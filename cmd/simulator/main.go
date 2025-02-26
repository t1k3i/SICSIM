package main

import (
	"github.com/t1k3i/SICSIM/cmd/simulator/gui"
	"github.com/t1k3i/SICSIM/pkg/machine"
)

func main() {
	machine := machine.NewMachine()

	gui.StartGUI(machine)
}