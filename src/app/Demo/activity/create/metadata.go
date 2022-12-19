package create

// Input corresponds to activity.json inputs
type Input struct {
	name   string `md:"name"`
	salary string `md:"salary"`
	age    string `md:"age"`
}

// Output corresponds to activity.json outputs
type Output struct {
	result map[string]interface{} `md:"result"`
}

// ToMap converts Input struct to map
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name":   i.name,
		"salary": i.salary,
		"age":    i.age,
	}
}

// FromMap converts a map to Input struct
// func (i *Input) FromMap(values map[string]interface{}) error {
// 	var err error
// 	i.name, err = coerce.ToString(values["name"])
// 	if err != nil {
// 		return err
// 	}

// 	i.salary, err = coerce.ToMap(values["salary"])
// 	if err != nil {
// 		return err
// 	}

// 	i.age, err = coerce.ToMap(values["age"])
// 	if err != nil {
// 		return err
// 	}

// 	// i.Input, err = coerce.ToObject(values["input"])
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	return nil
// }

// ToMap converts Output struct to map
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"result": o.result,
	}
}

// // FromMap converts a map to Output struct
// func (o *Output) FromMap(values map[string]interface{}) error {
// 	var err error
// 	o.result, err = coerce.ToObject(values["result"])
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
