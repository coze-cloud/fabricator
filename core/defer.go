package core

type Defer struct {
	Call FuncCall
}

func (d Defer) String() string {
	return "defer " + d.Call.String()
}



