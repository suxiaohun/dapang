package lesson

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"strings"
)

func GetUUID() {
	uid1 := uuid.NewV4().String()
	uid2 := strings.Replace(uid1, "-", "", -1)

	fmt.Printf("satori uuid v4: %s\n", uid1)
	fmt.Printf("satori uuid v4 format: %s\n", uid2)
}
