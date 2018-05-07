package decoprator

import (
	"testing"
	"fmt"
)

func TestCompany(t *testing.T){
	baseCom := BaseCompany{}
	baseCom.Show()
	fmt.Println("")

	sonCom := SonCompany{baseCom}
	sonCom.Show()
	fmt.Println("")

	parentCom := ParentCompany{sonCom}
	parentCom.Show()
	fmt.Println("")
}
