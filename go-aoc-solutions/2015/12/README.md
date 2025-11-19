# 2015 Day 12

Delved into some json marshalling and unmarshalling, it's actually pretty
slick in GO.  You can use a defined type, but you can just specify
`interface{}` as the type which will unmarshal to anything.  In that case,
JS objects are marshalled to `map[string]interface{}` and arrays are
marshalled to `[]interface{}`.

Reading [the docs](https://pkg.go.dev/encoding/json#Unmarshal) it
uses these types:

* bool, for JSON booleans
* float64, for JSON numbers
* string, for JSON strings
* []interface{}, for JSON arrays
* map[string]interface{}, for JSON objects
* nil for JSON null

I was able to unmarshal to a `data` variable of type `interface{}`:

```go
var data interface{}
contentBytes := []byte(contents)
_ := json.Unmarshal(contentBytes, &data)
```

And created a function to remove any objects with a property that has a value
of "red":

```go
func removeRedObjects(obj interface{}) interface{} {
	switch value := obj.(type) {
	case []interface{}:
		for i := range value {
			value[i] = removeRedObjects(value[i])
		}
	case map[string]interface{}:
		// it's an object, it it has a property with the VALUE 'red', return nil
		for _, v := range value {
			switch s := v.(type) {
			case string:
				if s == "red" {
					return nil
				}
			}
		}
		// ok, it doesn't have a 'red' property itself, but check each property
		// for them
		for k, v := range value {
			value[k] = removeRedObjects(v)
		}
	}
	return obj
}
```

And then I just marshalled it back to JSON and used my code for part 1 to get
the numbers and total them:

```go
removed := removeRedObjects(data)
newJsonBytes, _ := json.Marshal(removed)
newJsonString := string(newJsonBytes)
```
