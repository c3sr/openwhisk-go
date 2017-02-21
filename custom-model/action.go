package model

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
