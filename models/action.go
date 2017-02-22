package models

type Action struct {
	Name        string
	Publish     bool
	Annotations Annotations
	version     string
	Namespace   string
}

type Annotation struct {
	Key   string
	Value interface{}
}

type Annotations []Annotation

type Actions []Action

type Parameters map[string]string

type ActionInvocation struct {
	Name     string
	Params   Parameters
	Blocking bool
	Result   bool
	Timeout  int
}

func NewActionInvocation(actionName string, options ...func(*ActionInvocation)) *ActionInvocation {

	i := &ActionInvocation{Name: actionName, Params: Parameters{}}

	for _, option := range options {
		option(i)
	}

	return i
}
func (i *ActionInvocation) AddParameter(key string, value string) {
	i.Params[key] = value
}

func Blocking(blocking bool) func(*ActionInvocation) {
	return func(i *ActionInvocation) {
		i.Blocking = blocking
	}
}

func ResultOnly(result bool) func(*ActionInvocation) {
	return func(i *ActionInvocation) {
		i.Result = result
	}
}

func Timeout(ms int) func(*ActionInvocation) {
	return func(i *ActionInvocation) {
		i.Timeout = ms
	}
}
