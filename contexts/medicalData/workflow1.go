package medicalData

import (
	"fmt"
)

type Workflow1 struct {
	next Workflow
}

func (w *Workflow1) execute(medicalData *MedicalData) {
	fmt.Println("Starting Workflow1")
	fmt.Println(medicalData.data)
	w.next.execute(medicalData)
}

func (w *Workflow1) setNext(next Workflow) {
	w.next = next
}
