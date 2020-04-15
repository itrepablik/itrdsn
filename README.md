
The **kopy** package is the Go's common backup files operation, you can copy an entire directory or big files recursively with low CPU memory usage.  This **kopy** package fully implemented in [gokopy](https://github.com/itrepablik/gokopy) backup files software.

# Installation
```
go get -u github.com/itrepablik/itrdsn
```

# Usage
```
package main

import (
  "log"
  "fmt"
  "github.com/itrepablik/itrdsn"
  "github.com/itrepablik/itrlog"
 )
 
func main() {
	hddSerialNo, err := itrdsn.GetDiskSerialNo()
	if err != nil {
		itrlog.Fatalw("error getting block storage info: %v", err)
	}
	fmt.Println("hddSerialNo: ", hddSerialNo)
}
```

# License
Code is distributed under MIT license, feel free to use it in your proprietary projects as well.
