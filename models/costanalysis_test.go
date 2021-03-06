package costanalysis

import (
	"encoding/json"
	"testing"
	"io/ioutil"
	"reflect"
	"fmt"
)

const (
	inputsPath = "test/sample_inputs.json"
	resultPath = "test/sample_results.json"
)

func TestPredict(t *testing.T) {
	queryJson, _ := ioutil.ReadFile(inputsPath)
	inputs := new(Inputs)
	json.Unmarshal(queryJson, inputs)

	resultsJson, _ := ioutil.ReadFile(resultPath)
	results := new(Results)
	json.Unmarshal(resultsJson, results)

	fmt.Printf("%+v\n", *results)
	fmt.Printf("%+v", *Predict(inputs))

	if reflect.DeepEqual(*results, *Predict(inputs)) == false {
		t.Errorf("Failed")
	}
}