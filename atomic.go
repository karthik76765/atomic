package main
import (
	"fmt"
	"sync"
	"sync/atomic"
)
func main() {
	var ops uint64
	var i := 0; c < 1000; c++ {
		atomic.addUint64(&ops, 1)

	}
	wg.done()
}()
}
wg.wait()
fmt.Printlv("ops:", ops)
}