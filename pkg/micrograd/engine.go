package micrograd

type Value struct {
	Data float64
	Grad float64
	//Prev []*Value
}

func New(data float64) *Value {
	value := &Value{
		Data: data,
		Grad: 0.0
	}

	return value
}
