package medicalData

type Workflow interface {
	execute(*MedicalData)
	setNext(Workflow)
}
