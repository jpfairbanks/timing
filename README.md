timing
======

A go package for timing the execution of code at microscopic level.

Install
=======

You can install this package by importing this in your own go code.

  import "github.com/jpfairbanks/timing"

Then call go test in the repository to run the test cases.

Using
=====

A Timing is a struct that keeps timing records for each iteration of a loop.
You can allocate one with timing.New(length) where the length is the number
of iterations in the loop. Then Tic(i) will start the timer and Toc(i) will 
stop the timer. Then you can use Resolve() in order to find the deltas.
After resolving the timer can be Printed with fmt.Sprintf(tg)
