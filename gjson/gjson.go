package gjson

import "github.com/tidwall/gjson"

// GetInt returns an integer representation.
func GetInt(json, path string) int64 {
	ss := gjson.Get(json, path)
	return ss.Int()
}

// GetString returns a string representation of the value.
func GetString(json, path string) string {
	ss := gjson.Get(json, path)
	return ss.String()
}

// GetInterface returns one of these types:
//
//	bool, for JSON booleans
//	float64, for JSON numbers
//	Number, for JSON numbers
//	string, for JSON string literals
//	nil, for JSON null
//	map[string]interface{}, for JSON objects
//	[]interface{}, for JSON arrays
//
func GetInterface(json, path string) interface{} {
	ss := gjson.Get(json, path)
	return ss.Value()
}

// GetBool returns an boolean representation.
func GetBool(json, path string) bool {
	ss := gjson.Get(json, path)
	return ss.Bool()
}

// GetFloat searches json for the specified path.
// A path is in dot syntax, such as "name.last" or "age".
// When the value is found it's returned immediately.
//
// A path is a series of keys separated by a dot.
// A key may contain special wildcard characters '*' and '?'.
// To access an array value use the index as the key.
// To get the number of elements in an array or to access a child path, use
// the '#' character.
// The dot and wildcard character can be escaped with '\'.
//
//  {
//    "name": {"first": "Tom", "last": "Anderson"},
//    "age":37,
//    "children": ["Sara","Alex","Jack"],
//    "friends": [
//      {"first": "James", "last": "Murphy"},
//      {"first": "Roger", "last": "Craig"}
//    ]
//  }
//  "name.last"          >> "Anderson"
//  "age"                >> 37
//  "children"           >> ["Sara","Alex","Jack"]
//  "children.#"         >> 3
//  "children.1"         >> "Alex"
//  "child*.2"           >> "Jack"
//  "c?ildren.0"         >> "Sara"
//  "friends.#.first"    >> ["James","Roger"]
//
// This function expects that the json is well-formed, and does not validate.
// Invalid json will not panic, but it may return back unexpected results.
// If you are consuming JSON from an unpredictable source then you may want to
// use the Valid function first.
func GetFloat(json, path string) float64 {
	ss := gjson.Get(json, path)
	return ss.Float()
}

// GetMap returns back an map of values. The result should be a JSON array.
func GetMap(json, path string) map[string]gjson.Result {
	ss := gjson.Get(json, path)
	return ss.Map()
}
