package medicalData

import (
	"fmt"
)

type Workflow2 struct {
	next Workflow
}

func (w *Workflow2) execute(medicalData *MedicalData) {
	fmt.Println("Starting Workflow2")
	fmt.Println(medicalData.data)
}

func (w *Workflow2) setNext(next Workflow) {
	w.next = next
}
