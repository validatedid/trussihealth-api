package importData

import "github.com/validatedid/trussihealth-api/src/packages/dataTransformer"

type ImportData struct{}

func (id ImportData) Execute(inJsonData string) {

	healthData := dataTransformer.DataTransformer{}.Extract(inJsonData)

}
