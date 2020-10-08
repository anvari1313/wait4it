package model

type Checker interface {
	BuildContext(cx CheckContext)
	Validate() (bool, error)
	Check() (bool, bool, error)
}
