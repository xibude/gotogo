package logic

import (
	"log"
	"os"
	"runtime"
	"testing"
)

var lg *log.Logger = log.New(os.Stdout, "-->", log.Lshortfile|log.LstdFlags)

func TestNodeSlice(t *testing.T) {
	hostname, _ := os.Hostname()
	lg.Printf("OStype = %s, GOver=%s, HostName=%v",
		runtime.GOOS, runtime.Version(), hostname)

	var personSlice PersonSlice
	var a Person = 13
	var b Person = 21
	personSlice.Add(&a)

	personSlice.Add(&b)
	lg.Printf("nSlice=%d", personSlice)

	c := personSlice.Find(&a)
	lg.Printf("pOrg=%v, pResult=%v, Result=%v", &personSlice[0], c, *c)

	return
}
