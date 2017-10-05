
# Guides

## Instalations

```
go get -u github.com/u007/lib/...
```

## Dates

```
import (
	"github.com/u007/lib/tools"
	"fmt"
)

//parsing iso time and iso date to time.Time
thetime, err := tools.TimeFromISOString(val)
if err != nil {
	fmt.Errorf("Unable to parse time: %v %s", val, err.Error())
}


### Date only (will be deprecated when golang 2.0 is out)

```
import (
	"github.com/u007/lib/tools"
)

type MyStruct struct {
	Name   string     `db:"name" json:"name"`
	Created    tools.Date `db:"created" json:"created"`
}

```
