package engine

type Value struct {
	Data float64
	Grad float64
	Prev map[*Value]int
}

func New(data float64) *Value {
	value := &Value{
		Data: data,
		Grad: 0.0,
		Prev: nil,
		//_Backward: func(value *Value) {}
	}

	return value
}

func (v1 *Value) Add(v2 *Value) *Value {
	added := New(v1.Data + v2.Data)
	added.Prev = make(map[*Value]int)
	added.Prev[v1] = 1
	added.Prev[v2] = 2
	return added
}
