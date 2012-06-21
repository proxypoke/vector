package vector

import (
	"strconv"
)

type IndexError int

func (e IndexError) Error() string {
	return "Index out of range: " + strconv.Itoa(int(e))
}
