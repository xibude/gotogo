# 终局视角
想象一下，现在是2041年，你已经花甲之年之年了。如果让你重新来过，你会怎么选择？


## interface

```
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	//NOTE: ignoring potential errors form input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

```

##