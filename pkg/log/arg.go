package log

// Arg represents a key-value pair.
type Arg struct {
	Key   string
	Value interface{}
}

// flatten converts a list of Arg into a flat list of interface{}.
func flatten(args []Arg) []interface{} {
	flattened := make([]interface{}, 0, 2*len(args))
	for _, arg := range args {
		flattened = append(flattened, arg.Key, arg.Value)
	}
	return flattened
}

// Any - handy shortcut function to build Arg.
// It can be useful when you want to compact somehow log call code.
// Example:
// log.Debug("example message string", Any("key", "value"))
// log.Debug("example message string", Arg{Key: "key", Value: "value"})
func Any(key string, any any) Arg {
	arg := Arg{
		Key:   key,
		Value: any,
	}

	return arg
}

// Err - is shortcut function which inits Arg with the predefined key: 'error' and value is error's string.
// Example:
// log.Error("example error", Err(io.EOF))
// log.Error("example error", Arg{Key: "error", Value: io.EOF})
func Err(err error) Arg {
	arg := Arg{
		Key:   "error",
		Value: err,
	}

	return arg
}
