
![ITRDSN](https://user-images.githubusercontent.com/58651329/79336006-8e9fb380-7f55-11ea-84bf-553a86126906.png)
The **itrdsn** package is a small wrapper to get the **hard disk serial number**, currently supported **Windows** and **Linux** operating systems.

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
		itrlog.Fatalw("error getting block storage serial number: %v", err)
	}
	fmt.Println("hddSerialNo: ", hddSerialNo)
}
```

# License
Code is distributed under MIT license, feel free to use it in your proprietary projects as well.
