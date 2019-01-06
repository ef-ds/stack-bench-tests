# Benchmark Tests

## Tests
The benchmark tests are composed of all the tests implemented in the [benchmark package](https://github.com/ef-ds/benchmark).


## Tested Stacks

Besides stack, the tests also probe a few high quality open source stack implementations, alongside the standard list package as well as using simple slice as a stack.

- List based stack: uses the standard [list](https://github.com/golang/go/tree/master/src/container/list) package as a FIFO stack as well as a LIFO stack.
- [CustomSlice](testdata.go): uses a simple, dynamically growing slice as its underlying data structure.
- [cookiejar](https://github.com/karalabe/cookiejar/blob/master/collections/stack/stack.go): the stack implemented in this package uses a  uses a dynamically growing slice of blocks to store the elements.
- [deque](https://github.com/ef-ds/deque): deque stores the elements in a dynamic growing circular doubly linked list of arrays. The stack implemented on this package is actually based on this deque.

We're actively looking for other, high quality stacks to add to our tests. Due to the large volume of open source stacks available, it is not possible to add all of them to the tests. However, all the new tested ones we're adding to this [issue](https://github.com/ef-ds/stack/issues/1).


### Efficient Data Structures stack vs deque

Efficient Data Structures implements this stack package as well as the [deque](https://github.com/ef-ds/deque) package which can also be used as a LIFO stack.

The stack implementated in this stack package is a simplified version of this deque package. When it comes to using the packages as a LIFO stack, the main differences are:

1) Stack is a simpler version of deque that performs better than deque on most, if not all, LIFO stack tests
2) Differently from deque which attempts to reuse slices, stack doesn't implement such logic making it faster, especially for small data sets, but also causes more allocations for larger data sets (but with similar performance)


## Results

The raw results of a local run are stored under the [testdata](testdata) directory.

Refer [here](PERFORMANCE.md) for curated results.


## How To Run

From the package main directory, the tests can be run with below command.

```
go test -benchmem -timeout 60m -bench=. -run=^$
```

To run the test for a single stack, below command can be used.

```
go test -benchmem -timeout 60m -bench="STACK_NAME*" -run=^$
```

Replace the STACK_NAME with the desired stack such as "List", "Slice", "Cookiejar", "Deque", "Stack".


To run only a specific test suite, below command can be used.

```
go test -benchmem -timeout 60m -bench="TEST_SUITE_NAME*" -run=^$
```

Replace the TEST_SUITE_NAME with the desired test suite such as "Microservice", "Fill", "Refill", "RefillFull", "SlowIncrease", "SlowDecrease", "Stable".


## Test Variations

It is common to see significant variations in the test numbers with different test runs due to different reasons such as processes running in the hosting computer while the tests are running.

It is recommended to run the test multiple times and compare the aggregated results in order to help eliminate/smooth the test variations.

To run the tests multiple times, use the "go test" count parameter as below.

```
go test -benchmem -count 10 -timeout 600m -bench=. -run=^$
```

As the number of tests and now, test runs as well, is very large, it becomes very difficult to analyze and understand the results. In order to be able to analyze and compare the results between the different stacks, the [benchstat](https://godoc.org/golang.org/x/perf/cmd/benchstat) tool can be used to aggregate the test results. But as benchstat was designed to compare the same set of tests, it is necessary to first split all the different tests into separate test files renaming each
test with the same name, so benchstat will be able to match the different tests.

First step is to run the test and output the results in a file. Below command can be used to run all tests 10 times.

```
go test -benchmem -count 10 -timeout 600m -bench=. -run=^$ > testdata/results.txt
```

Next step is to split the "results.txt" file into separate test files. The [test-splitter](https://github.com/ef-ds/tools/tree/master/testsplitter) tool can be used for this purpose. To run the tool, clone the repo and run test-splitter from the "testsplitter" directory as follow.

```
go run *.go --file PATH_TO_RESULTS.TXT
```

Test-splitter should output one file per test name in the tests results file. The file names are named after each test name.

The last step is to run the [benchstat](https://godoc.org/golang.org/x/perf/cmd/benchstat) tool to aggregate and compare the results.

Below are the set of benchstat commands that can be used to compare deque against the other tested stacks.

Stack vs list
```
benchstat testdata/BenchmarkMicroserviceStackv1.0.0.txt testdata/BenchmarkMicroserviceList.txt
benchstat testdata/BenchmarkFillStackv1.0.0.txt testdata/BenchmarkFillList.txt
benchstat testdata/BenchmarkRefillStackv1.0.0.txt testdata/BenchmarkRefillList.txt
benchstat testdata/BenchmarkRefillFullStackv1.0.0.txt testdata/BenchmarkRefillFullList.txt
benchstat testdata/BenchmarkSlowIncreaseStackv1.0.0.txt testdata/BenchmarkSlowIncreaseList.txt
benchstat testdata/BenchmarkSlowDecreaseStackv1.0.0.txt testdata/BenchmarkSlowDecreaseList.txt
benchstat testdata/BenchmarkStableStackv1.0.0.txt testdata/BenchmarkStableList.txt
```

Stack vs slice
```
benchstat testdata/BenchmarkMicroserviceStackv1.0.0.txt testdata/BenchmarkMicroserviceSlice.txt
benchstat testdata/BenchmarkFillStackv1.0.0.txt testdata/BenchmarkFillSlice.txt
benchstat testdata/BenchmarkRefillStackv1.0.0.txt testdata/BenchmarkRefillSlice.txt
benchstat testdata/BenchmarkRefillFullStackvv1.0.0.txt testdata/BenchmarkRefillFullSlice.txt
benchstat testdata/BenchmarkSlowIncreaseStackvv1.0.0.txt testdata/BenchmarkSlowIncreaseSlice.txt
benchstat testdata/BenchmarkSlowDecreaseStackvv1.0.0.txt testdata/BenchmarkSlowDecreaseSlice.txt
benchstat testdata/BenchmarkStableStackvv1.0.0.txt testdata/BenchmarkStableSlice.txt
```

Stack vs cookiejar
```
benchstat testdata/BenchmarkMicroserviceStackv1.0.0.txt testdata/BenchmarkMicroserviceCookiejar.txt
benchstat testdata/BenchmarkFillStackv1.0.0.txt testdata/BenchmarkFillCookiejar.txt
benchstat testdata/BenchmarkRefillStackv1.0.0.txt testdata/BenchmarkRefillCookiejar.txt
benchstat testdata/BenchmarkRefillFullStackvv1.0.0.txt testdata/BenchmarkRefillFullCookiejar.txt
benchstat testdata/BenchmarkSlowIncreaseStackvv1.0.0.txt testdata/BenchmarkSlowIncreaseCookiejar.txt
benchstat testdata/BenchmarkSlowDecreaseStackvv1.0.0.txt testdata/BenchmarkSlowDecreaseCookiejar.txt
benchstat testdata/BenchmarkStableStackvv1.0.0.txt testdata/BenchmarkStableCookiejar.txt
```

Stack vs deque
```
benchstat testdata/BenchmarkMicroserviceStackv1.0.0.txt testdata/BenchmarkMicroserviceDeque.txt
benchstat testdata/BenchmarkFillStackv1.0.0.txt testdata/BenchmarkFillDeque.txt
benchstat testdata/BenchmarkRefillStackv1.0.0.txt testdata/BenchmarkRefillDeque.txt
benchstat testdata/BenchmarkRefillFullStackvv1.0.0.txt testdata/BenchmarkRefillFullDeque.txt
benchstat testdata/BenchmarkSlowIncreaseStackvv1.0.0.txt testdata/BenchmarkSlowIncreaseDeque.txt
benchstat testdata/BenchmarkSlowDecreaseStackvv1.0.0.txt testdata/BenchmarkSlowDecreaseDeque.txt
benchstat testdata/BenchmarkStableStackvv1.0.0.txt testdata/BenchmarkStableDeque.txt
```
