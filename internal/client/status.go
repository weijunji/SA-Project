package client

var _locked bool = true

func LockTheDoor() {
	_locked = true
}

func UnlockTheDoor() {
	_locked = false
}
