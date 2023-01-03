package vidio

import (
	"bytes"
	"fmt"
)

func (BS *BinarySquare) Solver(n int) {
	for i := 0; i <= n; i++ {
		var b bytes.Buffer
		for j := 0; j <= n; j++ {

			if i%2 == 0 {
				if j%2 == 0 {
					b.WriteString("0")
				} else {
					b.WriteString("1")
				}
			} else {
				if j%2 == 0 {
					b.WriteString("1")
				} else {
					b.WriteString("0")
				}
			}
		}

		fmt.Println(b.String())
	}
}
