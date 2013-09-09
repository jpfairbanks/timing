/* timing.go: a struct for performance timing
Author: James Fairbanks
Date: 2012-09-02
Liscence: BSD
*/
package timing

import(
	"time"
	"fmt"
	"strings"
)

//Timing: This is a struct for capturing performance numbers.
//Use Tic and Toc to start and stop the timer.
type Timing struct {
    ts []time.Time
    te []time.Time
    Td []time.Duration
    length int
}

//New: create a new Timer
func New(n int) Timing {
    ts := make([]time.Time, n)
    te := make([]time.Time, n)
    Td := make([]time.Duration, n)
	tg := Timing{ts, te, Td, n}
	return tg
}

//Tic: Start the ith timer.
func (tg Timing) Tic(i int) {
    tg.ts[i] = time.Now()
}

//Toc: Stop the ith timer.
func (tg Timing) Toc(i int) {
    tg.te[i] = time.Now()
}

//Resolve: compute the deltas all at once.
func (tg Timing) Resolve() {
    for i:=0; i < tg.length; i++ {
		tg.Td[i] = tg.te[i].Sub(tg.ts[i])
    }
}

//String: Basic representation of the deltas.
func (tg Timing) String() string {
	arr := make([]string, len(tg.Td))
	for i,t := range tg.Td{
		arr[i] = fmt.Sprintf("%d", t)
	}
	str := strings.Join(arr, " ")
    return str
}

//KeyString: Json friendly representation of the deltas.
func (tg Timing) KeyString(key string) string {
    return fmt.Sprintf("%s:%d,", key, tg.Td)
}

//TupleString: Columnar representation of the deltas for printing to a file.
func (tg Timing) TupleString(sep string) string {
	var lines []string
	lines = make([]string, tg.length)
	for index, dura := range tg.Td{
		lines[index] = fmt.Sprintf("%d %v", index, dura.Nanoseconds())
	}
	bigstring := strings.Join(lines, sep)
    return bigstring
}
