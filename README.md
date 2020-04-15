![Kopy](https://user-images.githubusercontent.com/58651329/75626348-e4b4d380-5c01-11ea-83bf-19e721932f81.png)
The **kopy** package is the Go's common backup files operation, you can copy an entire directory or big files recursively with low CPU memory usage.  This **kopy** package fully implemented in [gokopy](https://github.com/itrepablik/gokopy) backup files software.

# Installation
```
go get -u github.com/itrepablik/itrdsn
```

# Usage
To copy the **entire directory or a folder** as an example usage using the **kopy.CopyDir()** method. This will not compress the directory or a folder.
```
package main

import (
	"log"
  "fmt"
  "github.com/itrepablik/itrdsn"
 )
 
func main() {
	hddSerialNo, err := itrdsn.GetDiskSerialNo()
	if err != nil {
		log.Fatal("error getting block storage info: %v", err)
	}
	fmt.Println("hddSerialNo: ", hddSerialNo)
}
```

# License
Code is distributed under MIT license, feel free to use it in your proprietary projects as well.
