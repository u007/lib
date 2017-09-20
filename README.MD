
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
```
