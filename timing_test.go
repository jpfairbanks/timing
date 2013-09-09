/* timing_test.go: test a struct for performance timing
Author: James Fairbanks
Date: 2012-09-02
Liscence: BSD
*/

package timing

import (
	"testing"
	"time"
	"fmt"
	"errors"
	"strings"
	"math"
)

//validate: Make sure that the timer is valid
//not necessarily correct. this is useful for tests.
func (tg Timing) validate() error {
	var err error
	if tg.length != len(tg.ts){
		err = errors.New("ts is wrong length")
	}
	if tg.length != len(tg.te){
		err = errors.New("te is wrong length")
	}
	if tg.length != len(tg.Td){
		err = errors.New("Td is wrong length")
	}
	for i, _ := range tg.te{
		if tg.te[i].Sub(tg.ts[i]) < 0{
			err = errors.New("time travel detected")
		}
	}
	return err
}


func TestTime(t * testing.T) {
	n:=5
	indexer := make([]int, n)
	var err error
	for i:=0; i<n; i++{
		indexer[i] = i
	}
	tg := New(n)
	for i,_ := range indexer{
		tg.Tic(i)
	}
	for i,_ := range indexer{
		tg.Toc(i)
	}
	tg.Resolve()
	err = tg.validate()
	if err != nil {
		t.Error(err)
	}
}

func TestValidate(t * testing.T) {
	n:=5
	indexer := make([]int, n)
	var err error
	for i:=0; i<n; i++{
		indexer[i] = i
	}
	tg := New(n)
	for i,_ := range indexer{
		tg.Tic(i)
	}
	for i,_ := range indexer{
		tg.Toc(i)
	}
	var epoch time.Time
	tg.te[0] = epoch
	tg.Resolve()
	err = tg.validate()
	if err == nil {
		t.Log("false negative")
		t.Error(err)
	}
	tg.Toc(0)
	err = tg.validate()
	if err != nil{
		t.Log("false positive")
		t.Error(err)
	}
	tg.Tic(n-1)
	err = tg.validate()
	if err == nil{
		t.Log("reticking")
		t.Error(err)
	}
}

//dummyTiming: Make a timer that has the known answers
//It should not valid. It is for testing the string representations.
func dummyTiming() Timing {
	n:=5
	tg := New(n)
	for i:=0; i<n; i++{
		tg.Td[i] = time.Duration(i)*time.Second
	}
	return tg
}

func TestString(t *testing.T){
	tg := dummyTiming()
	out := tg.String()
	answer := "0 1000000000 2000000000 3000000000 4000000000"
	if !strings.EqualFold(out, answer) {
		t.Error(answer)
		t.Error(out)
	}
}

func TestKeyString(t *testing.T){
	tg := dummyTiming()
	out := tg.KeyString("testkey")
	answer := "testkey:[0 1000000000 2000000000 3000000000 4000000000],"
	if !strings.EqualFold(out, answer) {
		t.Error(answer)
		t.Error(out)
	}
}

func TestTupleString(t *testing.T){
	tg := dummyTiming()
	out := tg.TupleString("\n")
	answer := "0 0\n1 1000000000\n2 2000000000\n3 3000000000\n4 4000000000"
	if !strings.EqualFold(out, answer) {
		t.Error(answer)
		t.Error(out)
	}
}

func TestResolution(t *testing.T){
	var count int64
	var i int64
	count = 100000
	var tstart time.Time
	var tend   time.Time
	var td     int64
	var tsum   int64
	var sumsqr int64
	tdarr := make([]int64, count)
	for i = 0; i < count; i++{
		tstart = time.Now()
		tend = time.Now()
		td = tend.Sub(tstart).Nanoseconds()
		//fmt.Printf("%d ", td)
		tdarr[i] = td
		tsum += td
		sumsqr += td * td
	}
	mean := tsum/count
	variance := (sumsqr/(count) - mean*mean)
	vari := float64(0)
	for i=0; i < count; i++{
		x := float64(tdarr[i]) - float64(mean)
		vari += x*x
	}
	fmt.Printf("mean: %v\n", mean)
	fmt.Printf("variance: %v\n", variance)
	fmt.Printf("stddev: %v\n", math.Sqrt(float64(variance)))
	//fmt.Println(tdarr)
}
