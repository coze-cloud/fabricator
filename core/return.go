package core

type Return struct {
	Values []string
}

func (r Return) String() string {
	return "return " + r.listValues()
}


func (r Return) listValues() string {
	valueList := r.Values[0]
	for _, value := range r.Values[1:] {
		valueList = valueList + ", " + value
	}
	return valueList
}

