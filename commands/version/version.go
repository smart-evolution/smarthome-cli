package version

import (
	"fmt"
	"github.com/smart-evolution/smarthome-cli/utils"
)

func Handler() {
	fmt.Println("smarthome-cli version is " + utils.VERSION)
}
