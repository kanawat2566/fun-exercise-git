package main

import (
	"github.com/orgs/KKGo-Software-engineering/fun-ex-git/contact"
	"github.com/orgs/KKGo-Software-engineering/fun-ex-git/notification"
	"github.com/orgs/KKGo-Software-engineering/fun-ex-git/toggle"
)

func main() {
	isEmailEnabled := true
	toggle := toggle.New(isEmailEnabled)

	contact := contact.New("kanawat2566@gmail.com", "0812345678")
	notification := notification.New(toggle)
	notification.Send("Hello, Go!", contact)
}
