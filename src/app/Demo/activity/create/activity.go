package create

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/project-flogo/core/activity"
)

// MyActivity is a stub for your Activity implementation
var activityMd = activity.ToMetadata(&Input{}, &Output{})

type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	// do eval
	input := &Input{}
	output := &Output{}

	//Get Input Object
	err = context.GetInputObject(input)

	name := input.name
	salary := input.salary
	age := input.age

	values := map[string]string{"name": name, "salary": salary, "age": age}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("https://dummy.restapiexample.com/api/v1/create", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	// var res map[string]interface{}

	// json.NewDecoder(resp.Body).Decode(&res)

	// fmt.Println(res["json"])

	json.NewDecoder(resp.Body).Decode(&output)

	fmt.Println(output)

	return true, nil
}
