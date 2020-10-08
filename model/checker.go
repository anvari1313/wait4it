package model

type Checker interface {
	BuildContext(cx CheckContext)
	Validate() error
	Check() (bool, bool, error)
}
