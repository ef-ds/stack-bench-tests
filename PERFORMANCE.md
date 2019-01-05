# Performance

Below compares the stack [benchmark tests](BENCHMARK_TESTS.md) results with the other tested stacks.

## Running the Tests
In the "testdata" directory, we have included the result of local test runs for all stacks. Below uses this run to compare the stacks, but it's possible and we highly encourage you to run the tests yourself to help validate the results.

To run the tests locally, clone the stack repo, cd to the stack main directory and run below command.

```
go test -benchmem -timeout 60m -bench=. -run=^$
```

This command will run all tests for all stacks locally once. This should be good enouh to give you a sense of the stacks performance, but to do a proper comparison, elimating test variations, we recommend you to run the tests as detailed [here](BENCHMARK_TESTS.md) by running the tests with multiple counts, splitting the files with [test-splitter](https://github.com/ef-ds/tools/tree/master/testsplitter) and using the [benchstat](https://godoc.org/golang.org/x/perf/cmd/benchstat) tool to aggregate the results.


## Bottom Line
As a general purpose LIFO stack, stack is the data structure that displays the most balanced performance, performing either very competitively or besting all other stacks in all the different test scenarios.


## Results
Given the enormous amount of test data, it can be difficult and time consuming to find out the net impact of all the tests,
so we generally spend most of the time on the results of the, arguably, most important test: the Microservice test.

Below results is for stack [v1.0.0](https://github.com/ef-ds/stack/blob/master/CHANGELOG.md).


### Microservice Test Results
stack vs [list](https://github.com/golang/go/tree/master/src/container/list) - [microservice tests](benchmark-microservice_test.go)
```
benchstat ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkMicroserviceStackv1.0.0.txt ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkMicroserviceList.txt
name        old time/op    new time/op    delta
/0-4          34.1ns ± 2%    39.7ns ± 4%   +16.66%  (p=0.000 n=10+9)
/1-4           344ns ± 1%     498ns ± 2%   +44.51%  (p=0.000 n=10+9)
/10-4         2.39µs ± 1%    4.98µs ± 6%  +108.09%  (p=0.000 n=10+10)
/100-4        22.2µs ± 0%    47.5µs ± 3%  +113.51%  (p=0.000 n=9+9)
/1000-4        212µs ± 1%     479µs ±10%  +126.56%  (p=0.000 n=10+10)
/10000-4      2.16ms ± 1%    5.18ms ± 9%  +139.92%  (p=0.000 n=10+10)
/100000-4     23.7ms ± 2%    82.7ms ±16%  +248.97%  (p=0.000 n=10+9)
/1000000-4     243ms ± 1%     897ms ± 8%  +269.36%  (p=0.000 n=9+9)

name        old alloc/op   new alloc/op   delta
/0-4           32.0B ± 0%     48.0B ± 0%   +50.00%  (p=0.000 n=10+10)
/1-4            256B ± 0%      496B ± 0%   +93.75%  (p=0.000 n=10+10)
/10-4         1.52kB ± 0%    4.53kB ± 0%  +197.89%  (p=0.000 n=10+10)
/100-4        16.7kB ± 0%    44.8kB ± 0%  +168.23%  (p=0.000 n=10+10)
/1000-4        134kB ± 0%     448kB ± 0%  +234.60%  (p=0.000 n=10+10)
/10000-4      1.29MB ± 0%    4.48MB ± 0%  +247.35%  (p=0.000 n=10+9)
/100000-4     12.8MB ± 0%    44.8MB ± 0%  +249.57%  (p=0.000 n=9+9)
/1000000-4     128MB ± 0%     448MB ± 0%  +249.84%  (p=0.000 n=8+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            10.0 ± 0%      15.0 ± 0%   +50.00%  (p=0.000 n=10+10)
/10-4           74.0 ± 0%     141.0 ± 0%   +90.54%  (p=0.000 n=10+10)
/100-4           706 ± 0%      1401 ± 0%   +98.44%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%    14.00k ± 0%   +99.81%  (p=0.000 n=10+10)
/10000-4       70.0k ± 0%    140.0k ± 0%   +99.93%  (p=0.000 n=10+10)
/100000-4       700k ± 0%     1400k ± 0%   +99.94%  (p=0.000 n=10+10)
/1000000-4     7.00M ± 0%    14.00M ± 0%   +99.94%  (p=0.000 n=10+10)
```

stack vs [CustomSliceStack](testdata.go) - [microservice tests](benchmark-microservice_test.go)
```
benchstat ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkMicroserviceStackv1.0.0.txt ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkMicroserviceSlice.txt
name        old time/op    new time/op    delta
/0-4          34.1ns ± 2%    39.3ns ± 2%  +15.45%  (p=0.000 n=10+10)
/1-4           344ns ± 1%     342ns ± 3%     ~     (p=0.222 n=10+10)
/10-4         2.39µs ± 1%    2.16µs ± 3%   -9.86%  (p=0.000 n=10+9)
/100-4        22.2µs ± 0%    18.4µs ± 0%  -17.44%  (p=0.000 n=9+9)
/1000-4        212µs ± 1%     175µs ± 0%  -17.18%  (p=0.000 n=10+10)
/10000-4      2.16ms ± 1%    1.87ms ± 1%  -13.26%  (p=0.000 n=10+10)
/100000-4     23.7ms ± 2%    26.1ms ± 2%  +10.07%  (p=0.000 n=10+9)
/1000000-4     243ms ± 1%     298ms ±14%  +22.86%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/0-4           32.0B ± 0%     32.0B ± 0%     ~     (all equal)
/1-4            256B ± 0%      200B ± 0%  -21.88%  (p=0.000 n=10+10)
/10-4         1.52kB ± 0%    1.40kB ± 0%   -7.89%  (p=0.000 n=10+10)
/100-4        16.7kB ± 0%    13.3kB ± 0%  -20.62%  (p=0.000 n=10+10)
/1000-4        134kB ± 0%     128kB ± 0%   -4.10%  (p=0.000 n=10+10)
/10000-4      1.29MB ± 0%    1.51MB ± 0%  +16.79%  (p=0.000 n=10+10)
/100000-4     12.8MB ± 0%    15.9MB ± 0%  +23.71%  (p=0.000 n=9+10)
/1000000-4     128MB ± 0%     157MB ± 0%  +22.75%  (p=0.000 n=8+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/1-4            10.0 ± 0%      11.0 ± 0%  +10.00%  (p=0.000 n=10+10)
/10-4           74.0 ± 0%      76.0 ± 0%   +2.70%  (p=0.000 n=10+10)
/100-4           706 ± 0%       709 ± 0%   +0.42%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.01k ± 0%   +0.07%  (p=0.000 n=10+10)
/10000-4       70.0k ± 0%     70.0k ± 0%   -0.01%  (p=0.000 n=10+10)
/100000-4       700k ± 0%      700k ± 0%   -0.02%  (p=0.000 n=10+10)
/1000000-4     7.00M ± 0%     7.00M ± 0%   -0.03%  (p=0.000 n=10+10)
```

stack vs [cookiejar](https://github.com/karalabe/cookiejar/blob/master/collections/stack/stack.go) - [microservice tests](benchmark-microservice_test.go)
```
benchstat ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkMicroserviceStackv1.0.0.txt ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkMicroserviceCookiejar.txt
name        old time/op    new time/op    delta
/0-4          34.1ns ± 2%  9907.4ns ± 3%   +28996.62%  (p=0.000 n=10+10)
/1-4           344ns ± 1%   10040ns ± 0%    +2815.96%  (p=0.000 n=10+9)
/10-4         2.39µs ± 1%   12.40µs ± 6%     +417.95%  (p=0.000 n=10+10)
/100-4        22.2µs ± 0%    29.7µs ± 1%      +33.52%  (p=0.000 n=9+8)
/1000-4        212µs ± 1%     203µs ± 1%       -3.93%  (p=0.000 n=10+9)
/10000-4      2.16ms ± 1%    1.98ms ± 1%       -8.31%  (p=0.000 n=10+10)
/100000-4     23.7ms ± 2%    23.1ms ± 2%       -2.75%  (p=0.000 n=10+9)
/1000000-4     243ms ± 1%     242ms ± 4%         ~     (p=0.447 n=9+10)

name        old alloc/op   new alloc/op   delta
/0-4           32.0B ± 0%  65648.0B ± 0%  +205050.00%  (p=0.000 n=10+10)
/1-4            256B ± 0%    65760B ± 0%   +25587.50%  (p=0.000 n=10+10)
/10-4         1.52kB ± 0%   66.77kB ± 0%    +4292.63%  (p=0.000 n=10+10)
/100-4        16.7kB ± 0%    76.8kB ± 0%     +359.62%  (p=0.000 n=10+10)
/1000-4        134kB ± 0%     178kB ± 0%      +32.67%  (p=0.000 n=10+10)
/10000-4      1.29MB ± 0%    1.32MB ± 0%       +2.10%  (p=0.000 n=10+10)
/100000-4     12.8MB ± 0%    12.8MB ± 0%       +0.19%  (p=0.000 n=9+10)
/1000000-4     128MB ± 0%     128MB ± 0%       +0.01%  (p=0.000 n=8+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      3.00 ± 0%     +200.00%  (p=0.000 n=10+10)
/1-4            10.0 ± 0%      10.0 ± 0%         ~     (all equal)
/10-4           74.0 ± 0%      73.0 ± 0%       -1.35%  (p=0.000 n=10+10)
/100-4           706 ± 0%       703 ± 0%       -0.42%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.00k ± 0%       -0.06%  (p=0.000 n=10+10)
/10000-4       70.0k ± 0%     70.0k ± 0%       -0.03%  (p=0.000 n=10+10)
/100000-4       700k ± 0%      700k ± 0%       -0.02%  (p=0.000 n=10+10)
/1000000-4     7.00M ± 0%     7.00M ± 0%       -0.02%  (p=0.000 n=10+10)
```

stack vs [deque](https://github.com/ef-ds/deque) - [microservice tests](benchmark-microservice_test.go)
```
benchstat ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkMicroserviceStackv1.0.0.txt ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkMicroserviceDeque.txt
name        old time/op    new time/op    delta
/0-4          34.1ns ± 2%    38.1ns ± 0%   +12.02%  (p=0.000 n=10+9)
/1-4           344ns ± 1%     357ns ± 1%    +3.72%  (p=0.000 n=10+9)
/10-4         2.39µs ± 1%    2.51µs ± 0%    +4.74%  (p=0.000 n=10+10)
/100-4        22.2µs ± 0%    23.4µs ± 0%    +5.23%  (p=0.000 n=9+9)
/1000-4        212µs ± 1%     222µs ± 0%    +4.87%  (p=0.000 n=10+10)
/10000-4      2.16ms ± 1%    2.28ms ± 1%    +5.48%  (p=0.000 n=10+10)
/100000-4     23.7ms ± 2%    24.9ms ± 1%    +5.01%  (p=0.000 n=10+10)
/1000000-4     243ms ± 1%     262ms ± 1%    +7.75%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/0-4           32.0B ± 0%     64.0B ± 0%  +100.00%  (p=0.000 n=10+10)
/1-4            256B ± 0%      288B ± 0%   +12.50%  (p=0.000 n=10+10)
/10-4         1.52kB ± 0%    1.55kB ± 0%    +2.11%  (p=0.000 n=10+10)
/100-4        16.7kB ± 0%    16.8kB ± 0%    +0.48%  (p=0.000 n=10+10)
/1000-4        134kB ± 0%     130kB ± 0%    -2.89%  (p=0.000 n=10+10)
/10000-4      1.29MB ± 0%    1.42MB ± 0%   +10.40%  (p=0.000 n=10+10)
/100000-4     12.8MB ± 0%    14.4MB ± 0%   +12.53%  (p=0.000 n=9+10)
/1000000-4     128MB ± 0%     144MB ± 0%   +12.73%  (p=0.000 n=8+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            10.0 ± 0%      10.0 ± 0%      ~     (all equal)
/10-4           74.0 ± 0%      74.0 ± 0%      ~     (all equal)
/100-4           706 ± 0%       707 ± 0%    +0.14%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.01k ± 0%    +0.09%  (p=0.000 n=10+10)
/10000-4       70.0k ± 0%     70.2k ± 0%    +0.18%  (p=0.000 n=10+10)
/100000-4       700k ± 0%      702k ± 0%    +0.19%  (p=0.000 n=10+10)
/1000000-4     7.00M ± 0%     7.02M ± 0%    +0.20%  (p=0.000 n=10+10)
```

### Other Test Results
#### stack vs [list](https://github.com/golang/go/tree/master/src/container/list)
stack vs list - [fill tests](benchmark-fill_test.go)
```
benchstat ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkFillStackv1.0.0.txt ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkFillList.txt
name        old time/op    new time/op    delta
/0-4          33.5ns ± 1%    41.7ns ±10%   +24.31%  (p=0.000 n=9+10)
/1-4           172ns ± 2%     114ns ± 4%   -33.39%  (p=0.000 n=10+9)
/10-4          574ns ± 3%     747ns ± 4%   +30.13%  (p=0.000 n=10+9)
/100-4        4.71µs ± 0%    7.00µs ± 6%   +48.82%  (p=0.000 n=8+10)
/1000-4       36.3µs ± 2%    71.9µs ± 7%   +97.97%  (p=0.000 n=10+10)
/10000-4       352µs ± 2%     726µs ± 6%  +106.41%  (p=0.000 n=9+10)
/100000-4     3.84ms ± 1%   20.65ms ± 8%  +438.35%  (p=0.000 n=9+10)
/1000000-4    43.4ms ± 3%   158.5ms ± 8%  +265.31%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/0-4           32.0B ± 0%     48.0B ± 0%   +50.00%  (p=0.000 n=10+10)
/1-4            160B ± 0%      112B ± 0%   -30.00%  (p=0.000 n=10+10)
/10-4           560B ± 0%      688B ± 0%   +22.86%  (p=0.000 n=10+10)
/100-4        7.12kB ± 0%    6.45kB ± 0%    -9.44%  (p=0.000 n=10+10)
/1000-4       37.9kB ± 0%    64.0kB ± 0%   +68.97%  (p=0.000 n=10+10)
/10000-4       330kB ± 0%     640kB ± 0%   +94.08%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    6.40MB ± 0%   +99.02%  (p=0.000 n=10+10)
/1000000-4    32.1MB ± 0%    64.0MB ± 0%   +99.63%  (p=0.000 n=10+9)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            4.00 ± 0%      3.00 ± 0%   -25.00%  (p=0.000 n=10+10)
/10-4           14.0 ± 0%      21.0 ± 0%   +50.00%  (p=0.000 n=10+10)
/100-4           106 ± 0%       201 ± 0%   +89.62%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     2.00k ± 0%   +98.71%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     20.0k ± 0%   +99.51%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      200k ± 0%   +99.60%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     2.00M ± 0%   +99.61%  (p=0.000 n=10+10)
```

stack vs list - [refill tests](benchmark-refill_test.go)
```
benchstat ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkRefillStackv1.0.0.txt ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkRefillList.txt
name       old time/op    new time/op    delta
/1-4         3.50µs ± 1%    7.09µs ± 0%  +102.64%  (p=0.000 n=10+8)
/10-4        33.6µs ± 1%    70.8µs ± 5%  +110.44%  (p=0.000 n=9+10)
/100-4        324µs ± 1%     707µs ±11%  +118.08%  (p=0.000 n=10+10)
/1000-4      3.22ms ± 1%    7.21ms ± 5%  +123.92%  (p=0.000 n=10+10)
/10000-4     32.6ms ± 2%    72.5ms ± 5%  +122.66%  (p=0.000 n=10+10)
/100000-4     378ms ± 2%    2059ms ±14%  +445.19%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    6.40MB ± 0%  +299.99%  (p=0.000 n=10+10)
/10000-4     16.0MB ± 0%    64.0MB ± 0%  +299.92%  (p=0.000 n=9+9)
/100000-4     161MB ± 0%     640MB ± 0%  +298.66%  (p=0.000 n=10+9)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     2.00M ± 0%  +100.00%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     20.0M ± 0%  +100.00%  (p=0.000 n=10+9)
```

stack vs list - [refill full tests](benchmark-refill-full_test.go)
```
benchstat ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkRefillFullStackv1.0.0.txt ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkRefillFullList.txt
name       old time/op    new time/op    delta
/1-4         3.38µs ± 2%    8.19µs ± 2%  +142.13%  (p=0.000 n=10+10)
/10-4        34.4µs ± 3%    86.4µs ± 6%  +150.77%  (p=0.000 n=10+10)
/100-4        336µs ± 2%     876µs ± 6%  +160.29%  (p=0.000 n=10+10)
/1000-4      3.33ms ± 3%    9.16ms ±15%  +174.74%  (p=0.000 n=10+9)
/10000-4     32.6ms ± 2%    91.3ms ± 8%  +180.33%  (p=0.000 n=10+10)
/100000-4     359ms ± 2%    1729ms ± 6%  +381.87%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=10+10)
/10000-4     16.0MB ± 0%    64.0MB ± 0%  +300.00%  (p=0.000 n=10+10)
/100000-4     160MB ± 0%     640MB ± 0%  +300.00%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     2.00M ± 0%  +100.00%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     20.0M ± 0%  +100.00%  (p=0.000 n=10+10)
```

stack vs list - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkSlowIncreaseStackv1.0.0.txt ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkSlowIncreaseList.txt
name        old time/op    new time/op    delta
/1-4           198ns ± 2%     176ns ± 1%   -11.02%  (p=0.000 n=10+9)
/10-4          867ns ± 1%    1427ns ± 3%   +64.54%  (p=0.000 n=10+10)
/100-4        7.48µs ± 0%   13.44µs ± 0%   +79.56%  (p=0.000 n=9+10)
/1000-4       64.9µs ± 1%   135.8µs ± 1%  +109.42%  (p=0.000 n=10+10)
/10000-4       636µs ± 0%    1414µs ± 0%  +122.29%  (p=0.000 n=10+9)
/100000-4     7.46ms ± 1%   26.66ms ± 7%  +257.47%  (p=0.000 n=9+10)
/1000000-4    78.6ms ± 2%   338.6ms ±10%  +330.53%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            176B ± 0%      176B ± 0%      ~     (all equal)
/10-4           720B ± 0%     1328B ± 0%   +84.44%  (p=0.000 n=10+10)
/100-4        8.72kB ± 0%   12.85kB ± 0%   +47.34%  (p=0.000 n=10+10)
/1000-4       53.9kB ± 0%   128.0kB ± 0%  +137.55%  (p=0.000 n=10+10)
/10000-4       490kB ± 0%    1280kB ± 0%  +161.35%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%   12.80MB ± 0%  +165.79%  (p=0.000 n=10+10)
/1000000-4    48.1MB ± 0%   128.0MB ± 0%  +166.34%  (p=0.000 n=10+8)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      5.00 ± 0%      ~     (all equal)
/10-4           24.0 ± 0%      41.0 ± 0%   +70.83%  (p=0.000 n=10+10)
/100-4           206 ± 0%       401 ± 0%   +94.66%  (p=0.000 n=10+10)
/1000-4        2.01k ± 0%     4.00k ± 0%   +99.35%  (p=0.000 n=10+10)
/10000-4       20.0k ± 0%     40.0k ± 0%   +99.76%  (p=0.000 n=10+10)
/100000-4       200k ± 0%      400k ± 0%   +99.80%  (p=0.000 n=10+10)
/1000000-4     2.00M ± 0%     4.00M ± 0%   +99.80%  (p=0.000 n=10+10)
```

stack vs list - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkSlowDecreaseStackv1.0.0.txt ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkSlowDecreaseList.txt
name        old time/op    new time/op    delta
/1-4          34.8ns ± 1%    69.4ns ± 2%   +99.75%  (p=0.000 n=10+8)
/10-4          349ns ± 1%     689ns ± 0%   +97.32%  (p=0.000 n=10+9)
/100-4        3.41µs ± 1%    6.94µs ± 6%  +103.69%  (p=0.000 n=9+9)
/1000-4       33.8µs ± 1%    69.2µs ± 4%  +104.98%  (p=0.000 n=9+9)
/10000-4       337µs ± 1%     688µs ± 3%  +104.23%  (p=0.000 n=8+10)
/100000-4     3.37ms ± 1%    6.84ms ± 1%  +102.87%  (p=0.000 n=10+9)
/1000000-4    33.6ms ± 0%    68.5ms ± 1%  +103.68%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     64.0B ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      640B ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=10+8)
/1000000-4    16.0MB ± 0%    64.0MB ± 0%  +300.00%  (p=0.000 n=8+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      2.00 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4           10.0 ± 0%      20.0 ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4           100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     2.00M ± 0%  +100.00%  (p=0.000 n=10+10)
```

stack vs list - [stable tests](benchmark-stable_test.go)
```
benchstat ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkStableStackv1.0.0.txt ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkStableList.txt
name        old time/op    new time/op    delta
/1-4          31.1ns ± 1%    79.9ns ± 2%  +157.20%  (p=0.000 n=10+10)
/10-4          312ns ± 1%     822ns ± 9%  +162.94%  (p=0.000 n=10+10)
/100-4        3.07µs ± 2%    8.01µs ± 2%  +161.25%  (p=0.000 n=10+9)
/1000-4       30.4µs ± 1%    79.9µs ± 2%  +162.95%  (p=0.000 n=10+9)
/10000-4       305µs ± 1%     799µs ± 1%  +162.13%  (p=0.000 n=9+9)
/100000-4     3.08ms ± 5%    7.98ms ± 1%  +159.25%  (p=0.000 n=9+10)
/1000000-4    31.6ms ± 2%    79.6ms ± 1%  +151.97%  (p=0.000 n=10+8)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     64.0B ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      640B ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000000-4    16.0MB ± 0%    64.0MB ± 0%  +300.00%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      2.00 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4           10.0 ± 0%      20.0 ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4           100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     2.00M ± 0%  +100.00%  (p=0.000 n=10+10)
```

#### stack vs [CustomSliceStack](testdata.go)
stack vs CustomSliceStack - [fill tests](benchmark-fill_test.go)
```
benchstat ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkFillStackv1.0.0.txt ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkFillSlice.txt
name        old time/op    new time/op    delta
/0-4          33.5ns ± 1%    39.0ns ± 1%   +16.49%  (p=0.000 n=9+10)
/1-4           172ns ± 2%      93ns ± 1%   -45.99%  (p=0.000 n=10+10)
/10-4          574ns ± 3%     570ns ± 1%      ~     (p=0.830 n=10+9)
/100-4        4.71µs ± 0%    3.58µs ± 1%   -23.99%  (p=0.000 n=8+10)
/1000-4       36.3µs ± 2%    30.0µs ± 1%   -17.42%  (p=0.000 n=10+9)
/10000-4       352µs ± 2%     368µs ± 1%    +4.48%  (p=0.000 n=9+10)
/100000-4     3.84ms ± 1%    8.11ms ± 2%  +111.49%  (p=0.000 n=9+9)
/1000000-4    43.4ms ± 3%   105.8ms ±10%  +143.74%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           32.0B ± 0%     32.0B ± 0%      ~     (all equal)
/1-4            160B ± 0%       56B ± 0%   -65.00%  (p=0.000 n=10+10)
/10-4           560B ± 0%      440B ± 0%   -21.43%  (p=0.000 n=10+10)
/100-4        7.12kB ± 0%    3.67kB ± 0%   -48.43%  (p=0.000 n=10+10)
/1000-4       37.9kB ± 0%    32.4kB ± 0%   -14.50%  (p=0.000 n=10+10)
/10000-4       330kB ± 0%     546kB ± 0%   +65.66%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    6.25MB ± 0%   +94.49%  (p=0.000 n=10+10)
/1000000-4    32.1MB ± 0%    61.2MB ± 0%   +90.86%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            4.00 ± 0%      3.00 ± 0%   -25.00%  (p=0.000 n=10+10)
/10-4           14.0 ± 0%      16.0 ± 0%   +14.29%  (p=0.000 n=10+10)
/100-4           106 ± 0%       109 ± 0%    +2.83%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     1.01k ± 0%    +0.50%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     10.0k ± 0%    -0.04%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      100k ± 0%    -0.17%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     1.00M ± 0%    -0.19%  (p=0.000 n=10+10)
```

stack vs CustomSliceStack - [refill tests](benchmark-refill_test.go)
```
benchstat ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkRefillStackv1.0.0.txt ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkRefillSlice.txt
name       old time/op    new time/op    delta
/1-4         3.50µs ± 1%    2.85µs ± 0%  -18.47%  (p=0.000 n=10+9)
/10-4        33.6µs ± 1%    27.9µs ± 1%  -17.01%  (p=0.000 n=9+9)
/100-4        324µs ± 1%     282µs ± 9%  -12.83%  (p=0.000 n=10+10)
/1000-4      3.22ms ± 1%    2.73ms ± 7%  -15.02%  (p=0.000 n=10+10)
/10000-4     32.6ms ± 2%    32.5ms ±12%     ~     (p=0.720 n=10+9)
/100000-4     378ms ± 2%     388ms ± 4%   +2.80%  (p=0.002 n=10+8)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%   -0.00%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    1.60MB ± 0%   -0.00%  (p=0.000 n=10+9)
/10000-4     16.0MB ± 0%    16.0MB ± 0%   +0.03%  (p=0.000 n=9+9)
/100000-4     161MB ± 0%     162MB ± 0%   +0.63%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.00M ± 0%     1.00M ± 0%     ~     (all equal)
/100000-4     10.0M ± 0%     10.0M ± 0%   -0.00%  (p=0.000 n=10+10)
```

stack vs CustomSliceStack - [refill full tests](benchmark-refill-full_test.go)
```
benchstat ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkRefillFullStackv1.0.0.txt ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkRefillFullSlice.txt
name       old time/op    new time/op    delta
/1-4         3.38µs ± 2%    3.18µs ± 7%  -6.06%  (p=0.000 n=10+10)
/10-4        34.4µs ± 3%    32.3µs ± 6%  -6.24%  (p=0.000 n=10+10)
/100-4        336µs ± 2%     309µs ± 4%  -8.13%  (p=0.000 n=10+10)
/1000-4      3.33ms ± 3%    3.11ms ± 6%  -6.85%  (p=0.000 n=10+10)
/10000-4     32.6ms ± 2%    31.1ms ± 6%  -4.56%  (p=0.009 n=10+10)
/100000-4     359ms ± 2%     390ms ± 9%  +8.66%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%    ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%    ~     (all equal)
/10000-4     16.0MB ± 0%    16.0MB ± 0%    ~     (all equal)
/100000-4     160MB ± 0%     160MB ± 0%    ~     (all equal)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%    ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%    ~     (all equal)
/10000-4      1.00M ± 0%     1.00M ± 0%    ~     (all equal)
/100000-4     10.0M ± 0%     10.0M ± 0%    ~     (all equal)
```

stack vs CustomSliceStack - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkSlowIncreaseStackv1.0.0.txt ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkSlowIncreaseSlice.txt
name        old time/op    new time/op    delta
/1-4           198ns ± 2%     153ns ± 1%  -22.76%  (p=0.000 n=10+9)
/10-4          867ns ± 1%     863ns ± 2%     ~     (p=0.591 n=10+10)
/100-4        7.48µs ± 0%    6.09µs ± 2%  -18.57%  (p=0.000 n=9+10)
/1000-4       64.9µs ± 1%    55.7µs ± 3%  -14.13%  (p=0.000 n=10+10)
/10000-4       636µs ± 0%     614µs ± 4%   -3.49%  (p=0.003 n=10+9)
/100000-4     7.46ms ± 1%   11.10ms ± 2%  +48.80%  (p=0.000 n=9+8)
/1000000-4    78.6ms ± 2%   140.3ms ± 9%  +78.43%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            176B ± 0%       88B ± 0%  -50.00%  (p=0.000 n=10+10)
/10-4           720B ± 0%      600B ± 0%  -16.67%  (p=0.000 n=10+10)
/100-4        8.72kB ± 0%    5.27kB ± 0%  -39.54%  (p=0.000 n=10+10)
/1000-4       53.9kB ± 0%    48.4kB ± 0%  -10.20%  (p=0.000 n=10+10)
/10000-4       490kB ± 0%     706kB ± 0%  +44.21%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    7.85MB ± 0%  +63.10%  (p=0.000 n=10+10)
/1000000-4    48.1MB ± 0%    77.2MB ± 0%  +60.61%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      5.00 ± 0%     ~     (all equal)
/10-4           24.0 ± 0%      26.0 ± 0%   +8.33%  (p=0.000 n=10+10)
/100-4           206 ± 0%       209 ± 0%   +1.46%  (p=0.000 n=10+10)
/1000-4        2.01k ± 0%     2.01k ± 0%   +0.25%  (p=0.000 n=10+10)
/10000-4       20.0k ± 0%     20.0k ± 0%   -0.02%  (p=0.000 n=10+10)
/100000-4       200k ± 0%      200k ± 0%   -0.08%  (p=0.000 n=10+10)
/1000000-4     2.00M ± 0%     2.00M ± 0%   -0.10%  (p=0.000 n=10+10)
```

stack vs CustomSliceStack - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkSlowDecreaseStackv1.0.0.txt ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkSlowDecreaseSlice.txt
name        old time/op    new time/op    delta
/1-4          34.8ns ± 1%    32.4ns ± 3%  -6.85%  (p=0.000 n=10+10)
/10-4          349ns ± 1%     323ns ± 1%  -7.60%  (p=0.000 n=10+8)
/100-4        3.41µs ± 1%    3.18µs ± 4%  -6.64%  (p=0.000 n=9+10)
/1000-4       33.8µs ± 1%    31.3µs ± 4%  -7.39%  (p=0.000 n=9+10)
/10000-4       337µs ± 1%     314µs ± 4%  -6.62%  (p=0.000 n=8+10)
/100000-4     3.37ms ± 1%    3.13ms ± 3%  -7.27%  (p=0.000 n=10+10)
/1000000-4    33.6ms ± 0%    31.6ms ± 2%  -5.90%  (p=0.000 n=9+9)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%    ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%    ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%    ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%    ~     (all equal)
/100-4           100 ± 0%       100 ± 0%    ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%    ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%    ~     (all equal)
```

stack vs CustomSliceStack - [stable tests](benchmark-stable_test.go)
```
benchstat ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkStableStackv1.0.0.txt ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkStableSlice.txt
name        old time/op    new time/op    delta
/1-4          31.1ns ± 1%    30.8ns ± 6%    ~     (p=0.224 n=10+10)
/10-4          312ns ± 1%     308ns ± 7%  -1.40%  (p=0.037 n=10+9)
/100-4        3.07µs ± 2%    3.00µs ± 3%  -2.00%  (p=0.014 n=10+10)
/1000-4       30.4µs ± 1%    29.9µs ± 5%    ~     (p=0.063 n=10+10)
/10000-4       305µs ± 1%     300µs ± 4%    ~     (p=0.268 n=9+10)
/100000-4     3.08ms ± 5%    2.97ms ± 5%  -3.61%  (p=0.002 n=9+9)
/1000000-4    31.6ms ± 2%    30.0ms ± 4%  -4.99%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%    ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%    ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%    ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%    ~     (all equal)
/100-4           100 ± 0%       100 ± 0%    ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%    ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%    ~     (all equal)
```

#### stack vs [cookiejar](https://github.com/karalabe/cookiejar/blob/v2/collections/stack/stack.go)
stack vs cookiejar - [fill tests](benchmark-fill_test.go)
```
benchstat ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkFillStackv1.0.0.txt ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkFillCookiejar.txt
name        old time/op    new time/op     delta
/0-4          33.5ns ± 1%  10035.1ns ± 7%   +29835.70%  (p=0.000 n=9+9)
/1-4           172ns ± 2%    10756ns ±16%    +6160.65%  (p=0.000 n=10+10)
/10-4          574ns ± 3%    10329ns ± 4%    +1698.58%  (p=0.000 n=10+9)
/100-4        4.71µs ± 0%    12.82µs ± 3%     +172.45%  (p=0.000 n=8+9)
/1000-4       36.3µs ± 2%     39.3µs ± 2%       +8.19%  (p=0.000 n=10+9)
/10000-4       352µs ± 2%      324µs ± 1%       -8.02%  (p=0.000 n=9+9)
/100000-4     3.84ms ± 1%     3.61ms ± 5%       -5.84%  (p=0.000 n=9+10)
/1000000-4    43.4ms ± 3%     43.6ms ± 3%         ~     (p=0.481 n=10+10)

name        old alloc/op   new alloc/op    delta
/0-4           32.0B ± 0%   65648.0B ± 0%  +205050.00%  (p=0.000 n=10+10)
/1-4            160B ± 0%     65664B ± 0%   +40940.00%  (p=0.000 n=10+10)
/10-4           560B ± 0%     65808B ± 0%   +11651.43%  (p=0.000 n=10+10)
/100-4        7.12kB ± 0%    67.25kB ± 0%     +844.49%  (p=0.000 n=10+10)
/1000-4       37.9kB ± 0%     81.6kB ± 0%     +115.41%  (p=0.000 n=10+10)
/10000-4       330kB ± 0%      357kB ± 0%       +8.21%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%     3.24MB ± 0%       +0.75%  (p=0.000 n=10+10)
/1000000-4    32.1MB ± 0%     32.1MB ± 0%       +0.03%  (p=0.000 n=10+9)

name        old allocs/op  new allocs/op   delta
/0-4            1.00 ± 0%       3.00 ± 0%     +200.00%  (p=0.000 n=10+10)
/1-4            4.00 ± 0%       4.00 ± 0%         ~     (all equal)
/10-4           14.0 ± 0%       13.0 ± 0%       -7.14%  (p=0.000 n=10+10)
/100-4           106 ± 0%        103 ± 0%       -2.83%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%      1.00k ± 0%       -0.40%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%      10.0k ± 0%       -0.18%  (p=0.000 n=10+10)
/100000-4       100k ± 0%       100k ± 0%       -0.17%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%      1.00M ± 0%       -0.17%  (p=0.000 n=10+10)
```

stack vs cookiejar - [refill tests](benchmark-refill_test.go)
```
benchstat ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkRefillStackv1.0.0.txt ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkRefillCookiejar.txt
name       old time/op    new time/op    delta
/1-4         3.50µs ± 1%    3.37µs ± 3%  -3.71%  (p=0.000 n=10+10)
/10-4        33.6µs ± 1%    31.4µs ± 2%  -6.64%  (p=0.000 n=9+10)
/100-4        324µs ± 1%     306µs ± 2%  -5.42%  (p=0.000 n=10+10)
/1000-4      3.22ms ± 1%    3.06ms ± 3%  -4.92%  (p=0.000 n=10+10)
/10000-4     32.6ms ± 2%    30.7ms ± 5%  -5.68%  (p=0.000 n=10+10)
/100000-4     378ms ± 2%     367ms ± 5%  -2.80%  (p=0.004 n=10+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%  +0.01%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     160kB ± 0%  +0.01%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    1.60MB ± 0%  +0.01%  (p=0.000 n=10+10)
/10000-4     16.0MB ± 0%    16.0MB ± 0%  +0.00%  (p=0.000 n=9+9)
/100000-4     161MB ± 0%     161MB ± 0%  +0.01%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%    ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%    ~     (all equal)
/10000-4      1.00M ± 0%     1.00M ± 0%    ~     (all equal)
/100000-4     10.0M ± 0%     10.0M ± 0%  -0.00%  (p=0.000 n=10+10)
```

stack vs cookiejar - [refill full tests](benchmark-refill-full_test.go)
```
benchstat ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkRefillFullStackv1.0.0.txt ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkRefillFullCookiejar.txt
name       old time/op    new time/op    delta
/1-4         3.38µs ± 2%    3.08µs ± 2%  -8.88%  (p=0.000 n=10+10)
/10-4        34.4µs ± 3%    31.5µs ± 4%  -8.51%  (p=0.000 n=10+10)
/100-4        336µs ± 2%     308µs ± 3%  -8.36%  (p=0.000 n=10+10)
/1000-4      3.33ms ± 3%    3.13ms ± 8%  -6.05%  (p=0.002 n=10+10)
/10000-4     32.6ms ± 2%    32.6ms ± 7%    ~     (p=1.000 n=10+10)
/100000-4     359ms ± 2%     349ms ± 3%  -2.66%  (p=0.000 n=10+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%    ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%    ~     (all equal)
/10000-4     16.0MB ± 0%    16.0MB ± 0%    ~     (all equal)
/100000-4     160MB ± 0%     160MB ± 0%    ~     (all equal)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%    ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%    ~     (all equal)
/10000-4      1.00M ± 0%     1.00M ± 0%    ~     (all equal)
/100000-4     10.0M ± 0%     10.0M ± 0%    ~     (all equal)
```

stack vs cookiejar - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkSlowIncreaseStackv1.0.0.txt ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkSlowIncreaseCookiejar.txt
name        old time/op    new time/op    delta
/1-4           198ns ± 2%    9932ns ± 0%   +4921.28%  (p=0.000 n=10+10)
/10-4          867ns ± 1%   10644ns ± 1%   +1127.67%  (p=0.000 n=10+9)
/100-4        7.48µs ± 0%   16.01µs ± 4%    +113.92%  (p=0.000 n=9+10)
/1000-4       64.9µs ± 1%    69.3µs ± 4%      +6.89%  (p=0.000 n=10+10)
/10000-4       636µs ± 0%     618µs ± 1%      -2.82%  (p=0.000 n=10+10)
/100000-4     7.46ms ± 1%    7.84ms ± 2%      +5.11%  (p=0.000 n=9+10)
/1000000-4    78.6ms ± 2%    85.8ms ± 3%      +9.07%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4            176B ± 0%    65680B ± 0%  +37218.18%  (p=0.000 n=10+10)
/10-4           720B ± 0%    65968B ± 0%   +9062.22%  (p=0.000 n=10+10)
/100-4        8.72kB ± 0%   68.85kB ± 0%    +689.54%  (p=0.000 n=10+10)
/1000-4       53.9kB ± 0%    97.6kB ± 0%     +81.15%  (p=0.000 n=10+10)
/10000-4       490kB ± 0%     517kB ± 0%      +5.53%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    4.84MB ± 0%      +0.50%  (p=0.000 n=10+10)
/1000000-4    48.1MB ± 0%    48.1MB ± 0%      +0.02%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      5.00 ± 0%        ~     (all equal)
/10-4           24.0 ± 0%      23.0 ± 0%      -4.17%  (p=0.000 n=10+10)
/100-4           206 ± 0%       203 ± 0%      -1.46%  (p=0.000 n=10+10)
/1000-4        2.01k ± 0%     2.00k ± 0%      -0.20%  (p=0.000 n=10+10)
/10000-4       20.0k ± 0%     20.0k ± 0%      -0.09%  (p=0.000 n=10+10)
/100000-4       200k ± 0%      200k ± 0%      -0.08%  (p=0.000 n=10+10)
/1000000-4     2.00M ± 0%     2.00M ± 0%      -0.09%  (p=0.000 n=10+10)
```

stack vs cookiejar - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkSlowDecreaseStackv1.0.0.txt ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkSlowDecreaseCookiejar.txt
name        old time/op    new time/op    delta
/1-4          34.8ns ± 1%    33.8ns ± 2%  -2.62%  (p=0.000 n=10+10)
/10-4          349ns ± 1%     349ns ± 4%    ~     (p=0.928 n=10+10)
/100-4        3.41µs ± 1%    3.36µs ± 7%    ~     (p=0.075 n=9+10)
/1000-4       33.8µs ± 1%    33.1µs ± 2%  -2.02%  (p=0.000 n=9+10)
/10000-4       337µs ± 1%     328µs ± 1%  -2.48%  (p=0.000 n=8+10)
/100000-4     3.37ms ± 1%    3.28ms ± 1%  -2.60%  (p=0.000 n=10+8)
/1000000-4    33.6ms ± 0%    32.8ms ± 1%  -2.41%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%    ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%    ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%    ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%    ~     (all equal)
/100-4           100 ± 0%       100 ± 0%    ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%    ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%    ~     (all equal)
```

stack vs cookiejar - [stable tests](benchmark-stable_test.go)
```
benchstat ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkStableStackv1.0.0.txt ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkStableCookiejar.txt
name        old time/op    new time/op    delta
/1-4          31.1ns ± 1%    29.9ns ± 3%  -3.67%  (p=0.000 n=10+9)
/10-4          312ns ± 1%     333ns ±11%  +6.43%  (p=0.024 n=10+10)
/100-4        3.07µs ± 2%    3.12µs ± 6%    ~     (p=0.160 n=10+10)
/1000-4       30.4µs ± 1%    31.5µs ± 8%    ~     (p=0.190 n=10+10)
/10000-4       305µs ± 1%     318µs ± 8%    ~     (p=0.156 n=9+10)
/100000-4     3.08ms ± 5%    3.08ms ± 7%    ~     (p=0.968 n=9+10)
/1000000-4    31.6ms ± 2%    30.1ms ± 8%  -4.76%  (p=0.003 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%    ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%    ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%    ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%    ~     (all equal)
/100-4           100 ± 0%       100 ± 0%    ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%    ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%    ~     (all equal)
```

#### stack vs [deque](https://github.com/ef-ds/deque)
stack vs Deque - [fill tests](benchmark-fill_test.go)
```
benchstat ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkFillStackv1.0.0.txt ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkFillDeque.txt
name        old time/op    new time/op    delta
/0-4          33.5ns ± 1%    37.4ns ± 0%   +11.67%  (p=0.000 n=9+9)
/1-4           172ns ± 2%     171ns ± 1%      ~     (p=0.574 n=10+9)
/10-4          574ns ± 3%     583ns ± 0%    +1.53%  (p=0.039 n=10+9)
/100-4        4.71µs ± 0%    4.73µs ± 0%    +0.61%  (p=0.000 n=8+8)
/1000-4       36.3µs ± 2%    36.6µs ± 0%    +0.72%  (p=0.017 n=10+9)
/10000-4       352µs ± 2%     363µs ± 0%    +3.14%  (p=0.000 n=9+9)
/100000-4     3.84ms ± 1%    3.85ms ± 2%      ~     (p=0.400 n=9+10)
/1000000-4    43.4ms ± 3%    44.0ms ± 3%    +1.44%  (p=0.043 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           32.0B ± 0%     64.0B ± 0%  +100.00%  (p=0.000 n=10+10)
/1-4            160B ± 0%      192B ± 0%   +20.00%  (p=0.000 n=10+10)
/10-4           560B ± 0%      592B ± 0%    +5.71%  (p=0.000 n=10+10)
/100-4        7.12kB ± 0%    7.20kB ± 0%    +1.12%  (p=0.000 n=10+10)
/1000-4       37.9kB ± 0%    34.0kB ± 0%   -10.22%  (p=0.000 n=10+10)
/10000-4       330kB ± 0%     323kB ± 0%    -2.04%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    3.22MB ± 0%    +0.19%  (p=0.000 n=10+10)
/1000000-4    32.1MB ± 0%    32.2MB ± 0%    +0.40%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            4.00 ± 0%      4.00 ± 0%      ~     (all equal)
/10-4           14.0 ± 0%      14.0 ± 0%      ~     (all equal)
/100-4           106 ± 0%       107 ± 0%    +0.94%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     1.01k ± 0%    +0.60%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     10.1k ± 0%    +0.58%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      101k ± 0%    +0.58%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     1.01M ± 0%    +0.58%  (p=0.000 n=10+10)
```

stack vs Deque - [refill tests](benchmark-refill_test.go)
```
benchstat ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkRefillStackv1.0.0.txt ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkRefillDeque.txt
name       old time/op    new time/op    delta
/1-4         3.50µs ± 1%    3.65µs ± 2%   +4.33%  (p=0.000 n=10+10)
/10-4        33.6µs ± 1%    35.9µs ± 1%   +6.66%  (p=0.000 n=9+10)
/100-4        324µs ± 1%     345µs ± 1%   +6.34%  (p=0.000 n=10+9)
/1000-4      3.22ms ± 1%    3.39ms ± 2%   +5.32%  (p=0.000 n=10+10)
/10000-4     32.6ms ± 2%    36.5ms ± 1%  +11.98%  (p=0.000 n=10+10)
/100000-4     378ms ± 2%     395ms ± 2%   +4.62%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%   -0.00%  (p=0.000 n=10+10)
/10000-4     16.0MB ± 0%    30.1MB ± 0%  +88.02%  (p=0.000 n=9+9)
/100000-4     161MB ± 0%     320MB ± 0%  +99.31%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.00M ± 0%     1.01M ± 0%   +0.68%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     10.1M ± 0%   +0.77%  (p=0.000 n=10+10)
```

stack vs Deque - [refill full tests](benchmark-refill-full_test.go)
```
benchstat ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkRefillFullStackv1.0.0.txt ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkRefillFullDeque.txt
name       old time/op    new time/op    delta
/1-4         3.38µs ± 2%    3.42µs ± 1%     ~     (p=0.085 n=10+10)
/10-4        34.4µs ± 3%    34.4µs ± 1%     ~     (p=0.842 n=10+9)
/100-4        336µs ± 2%     340µs ± 2%   +1.07%  (p=0.015 n=10+10)
/1000-4      3.33ms ± 3%    3.35ms ± 0%     ~     (p=0.529 n=10+10)
/10000-4     32.6ms ± 2%    37.2ms ± 0%  +14.29%  (p=0.000 n=10+10)
/100000-4     359ms ± 2%     393ms ± 3%   +9.51%  (p=0.000 n=10+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/10000-4     16.0MB ± 0%    30.1MB ± 0%  +88.06%  (p=0.000 n=10+10)
/100000-4     160MB ± 0%     320MB ± 0%  +99.97%  (p=0.000 n=10+9)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.00M ± 0%     1.01M ± 0%   +0.68%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     10.1M ± 0%   +0.77%  (p=0.000 n=10+10)
```

stack vs Deque - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkSlowIncreaseStackv1.0.0.txt ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkSlowIncreaseDeque.txt
name        old time/op    new time/op    delta
/1-4           198ns ± 2%     205ns ± 1%   +3.79%  (p=0.000 n=10+10)
/10-4          867ns ± 1%     915ns ± 0%   +5.52%  (p=0.000 n=10+9)
/100-4        7.48µs ± 0%    7.98µs ± 0%   +6.58%  (p=0.000 n=9+10)
/1000-4       64.9µs ± 1%    68.8µs ± 0%   +6.03%  (p=0.000 n=10+10)
/10000-4       636µs ± 0%     688µs ± 0%   +8.09%  (p=0.000 n=10+10)
/100000-4     7.46ms ± 1%    7.90ms ± 1%   +5.91%  (p=0.000 n=9+10)
/1000000-4    78.6ms ± 2%    83.2ms ± 1%   +5.75%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            176B ± 0%      208B ± 0%  +18.18%  (p=0.000 n=10+10)
/10-4           720B ± 0%      752B ± 0%   +4.44%  (p=0.000 n=10+10)
/100-4        8.72kB ± 0%    8.80kB ± 0%   +0.92%  (p=0.000 n=10+10)
/1000-4       53.9kB ± 0%    50.0kB ± 0%   -7.18%  (p=0.000 n=10+10)
/10000-4       490kB ± 0%     483kB ± 0%   -1.37%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    4.82MB ± 0%   +0.12%  (p=0.000 n=10+10)
/1000000-4    48.1MB ± 0%    48.2MB ± 0%   +0.28%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      5.00 ± 0%     ~     (all equal)
/10-4           24.0 ± 0%      24.0 ± 0%     ~     (all equal)
/100-4           206 ± 0%       207 ± 0%   +0.49%  (p=0.000 n=10+10)
/1000-4        2.01k ± 0%     2.01k ± 0%   +0.30%  (p=0.000 n=10+10)
/10000-4       20.0k ± 0%     20.1k ± 0%   +0.29%  (p=0.000 n=10+10)
/100000-4       200k ± 0%      201k ± 0%   +0.29%  (p=0.000 n=10+10)
/1000000-4     2.00M ± 0%     2.01M ± 0%   +0.29%  (p=0.000 n=10+10)
```

stack vs Deque - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkSlowDecreaseStackv1.0.0.txt ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkSlowDecreaseDeque.txt
name        old time/op    new time/op    delta
/1-4          34.8ns ± 1%    35.8ns ± 2%  +2.91%  (p=0.000 n=10+10)
/10-4          349ns ± 1%     359ns ± 1%  +2.69%  (p=0.000 n=10+10)
/100-4        3.41µs ± 1%    3.51µs ± 1%  +2.97%  (p=0.000 n=9+10)
/1000-4       33.8µs ± 1%    34.9µs ± 1%  +3.33%  (p=0.000 n=9+9)
/10000-4       337µs ± 1%     350µs ± 1%  +3.99%  (p=0.000 n=8+9)
/100000-4     3.37ms ± 1%    3.49ms ± 0%  +3.60%  (p=0.000 n=10+8)
/1000000-4    33.6ms ± 0%    34.8ms ± 2%  +3.64%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%    ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%    ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%    ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%    ~     (all equal)
/100-4           100 ± 0%       100 ± 0%    ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%    ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%    ~     (all equal)
```

stack vs Deque - [stable tests](benchmark-stable_test.go)
```
benchstat ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkStableStackv1.0.0.txt ./../../../ef-ds/stack-bench-tests/testdata/BenchmarkStableDeque.txt
name        old time/op    new time/op    delta
/1-4          31.1ns ± 1%    33.6ns ± 1%  +8.21%  (p=0.000 n=10+9)
/10-4          312ns ± 1%     342ns ± 2%  +9.54%  (p=0.000 n=10+10)
/100-4        3.07µs ± 2%    3.36µs ± 2%  +9.50%  (p=0.000 n=10+10)
/1000-4       30.4µs ± 1%    33.2µs ± 1%  +9.22%  (p=0.000 n=10+9)
/10000-4       305µs ± 1%     333µs ± 4%  +9.40%  (p=0.000 n=9+10)
/100000-4     3.08ms ± 5%    3.32ms ± 2%  +7.94%  (p=0.000 n=9+10)
/1000000-4    31.6ms ± 2%    33.3ms ± 2%  +5.32%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%    ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%    ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%    ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%    ~     (all equal)
/100-4           100 ± 0%       100 ± 0%    ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%    ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%    ~     (all equal)
```
