package main

import "github.com/01-edu/z01"

const (
	CLOSE = iota
	OPEN
)

type Door struct {
	state int
}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...")
	ptrDoor.state = CLOSE
}

func OpenDoor(ptrdoor *Door) {
	PrintStr("Door Opening...")
	ptrdoor.state = OPEN
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door opened ?")
	return ptrDoor.state == OPEN
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	return ptrDoor.state == CLOSE
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func main() {
	door := &Door{}
	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}
