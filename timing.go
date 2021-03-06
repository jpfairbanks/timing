/* timing.go: a struct for performance timing
Author: James Fairbanks
Date: 2012-09-02
Liscence: BSD
*/
package timing

import (
	"fmt"
	"strings"
	"time"
)

type tcollection map[string]time.Time
type dcollection map[string]time.Duration

//Timing: This is a struct for capturing performance numbers.
//Use Tic and Toc to start and stop the timer.
type MapTiming struct {
	ts     tcollection
	te     tcollection
	Td     dcollection
	length int
}

//New: create a new Timer
func NewMapTiming(n int) MapTiming {
	ts := make(tcollection, n)
	te := make(tcollection, n)
	Td := make(dcollection, n)
	tg := MapTiming{ts, te, Td, n}
	return tg
}

//Tic: Start the ith timer.
func (tg MapTiming) Tic(i string) {
	tg.ts[i] = time.Now()
}

//Toc: Stop the ith timer.
func (tg MapTiming) Toc(i string) {
	tg.te[i] = time.Now()
}

//Resolve: compute the deltas all at once.
func (tg MapTiming) Resolve() {
	for key, _ := range tg.Td {
		tg.Td[key] = tg.te[key].Sub(tg.ts[key])
	}
}

//String: Basic representation of the deltas.
func (tg MapTiming) String() string {
	arr := make([]string, len(tg.Td))
	i := 0
	for _, value := range tg.Td {
		arr[i] = fmt.Sprintf("%d", value)
		i++
	}
	str := strings.Join(arr, " ")
	return str
}

//KeyString: Json friendly representation of the deltas.
func (tg MapTiming) KeyString(key string) string {
	return fmt.Sprintf("%s:%d,", key, tg.Td)
}

//TupleString: Columnar representation of the deltas for printing to a file.
func (tg MapTiming) TupleString(sep string) string {
	var lines []string
	lines = make([]string, tg.length)
	i := 0
	for key, dura := range tg.Td {
		lines[i] = fmt.Sprintf("%d %v", key, dura.Nanoseconds())
		i++
	}
	bigstring := strings.Join(lines, sep)
	return bigstring
}
