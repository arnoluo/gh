package t

// go helper types
type (
	// int
	I = int

	// large int
	LI = int64

	// tiny int
	TI = int8

	// small int
	SI = int16

	// string
	S = string

	// bool
	B = bool

	// float64
	F = float64

	// complex64
	C = complex64

	// uintptr
	P = uintptr

	// interface{}
	Itf = interface{}

	// array of interface{}
	AItf = []interface{}

	// array of int
	AI = []int

	// array of string
	AS = []string

	// array of float
	AF = []float64

	// map of string -> interface{}
	MSItf = map[string]interface{}

	// map of int -> interface{}
	MIItf = map[int]interface{}

	// map of interface{} -> interface{}
	MItf = map[interface{}]interface{}
)

type (
	StrType   struct{}
	IntType   struct{}
	BoolType  struct{}
	FloatType struct{}
	ItfType   struct{}
)
