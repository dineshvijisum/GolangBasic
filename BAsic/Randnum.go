package main

import "fmt"
import "time"
import "math/rand"
import "runtime"

func main() {
  numThreads := runtime.NumCPU()
  runtime.GOMAXPROCS(numThreads)
  var numIntsToGenerate = 100000000
  var numIntsPerThread = numIntsToGenerate / numThreads
  ch := make(chan int, numIntsToGenerate)
  singleThreadIntSlice := make([]int, numIntsToGenerate, numIntsToGenerate)
  multiThreadIntSlice := make([]int, numIntsToGenerate, numIntsToGenerate)
  fmt.Printf("Initiating single-threaded random number generation.\n")
  startSingleRun := time.Now()
  go makeRandomNumbers(numIntsToGenerate, ch)
  for i := 0; i < numIntsToGenerate; i++ {
    singleThreadIntSlice = append(singleThreadIntSlice,(<-ch))
  }
  elapsedSingleRun := time.Since(startSingleRun)
  fmt.Printf("Single-threaded run took %s\n", elapsedSingleRun)
  fmt.Printf("Initiating multi-threaded random number generation.\n")
  startMultiRun := time.Now()
  for i := 0; i < numThreads; i++ {
    go makeRandomNumbers(numIntsPerThread, ch)
  }
  for i := 0; i < numIntsToGenerate; i++ {
    multiThreadIntSlice = append(multiThreadIntSlice,(<-ch))
  }
  elapsedMultiRun := time.Since(startMultiRun)
  fmt.Printf("Multi-threaded run took %s\n", elapsedMultiRun)
}
func makeRandomNumbers(numInts int, ch chan int) {
  source := rand.NewSource(time.Now().UnixNano())
  generator := rand.New(source)
  for i := 0; i < numInts; i++ {
        ch <- generator.Intn(numInts*100)
    }
}