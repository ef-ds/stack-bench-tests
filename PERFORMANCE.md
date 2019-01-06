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

Below results is for stack [v1.0.1](https://github.com/ef-ds/stack/blob/master/CHANGELOG.md).




### Microservice Test Results
stack vs [list](https://github.com/golang/go/tree/master/src/container/list) - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceStackv1.0.1.txt testdata/BenchmarkMicroserviceList.txt
name        old time/op    new time/op    delta
/0-4          28.1ns ± 1%    39.7ns ± 4%   +41.51%  (p=0.000 n=10+9)
/1-4           322ns ± 0%     498ns ± 2%   +54.31%  (p=0.000 n=9+9)
/10-4         2.23µs ± 1%    4.98µs ± 6%  +123.15%  (p=0.000 n=10+10)
/100-4        20.6µs ± 0%    47.5µs ± 3%  +130.98%  (p=0.000 n=9+9)
/1000-4        203µs ± 0%     479µs ±10%  +136.67%  (p=0.000 n=8+10)
/10000-4      2.12ms ± 1%    5.18ms ± 9%  +144.49%  (p=0.000 n=10+10)
/100000-4     24.0ms ± 3%    82.7ms ±16%  +244.78%  (p=0.000 n=9+9)
/1000000-4     258ms ± 2%     897ms ± 8%  +248.01%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/0-4           16.0B ± 0%     48.0B ± 0%  +200.00%  (p=0.000 n=10+10)
/1-4            288B ± 0%      496B ± 0%   +72.22%  (p=0.000 n=10+10)
/10-4         1.55kB ± 0%    4.53kB ± 0%  +191.75%  (p=0.000 n=10+10)
/100-4        15.2kB ± 0%    44.8kB ± 0%  +194.74%  (p=0.000 n=10+10)
/1000-4        161kB ± 0%     448kB ± 0%  +177.95%  (p=0.000 n=10+10)
/10000-4      1.76MB ± 0%    4.48MB ± 0%  +154.36%  (p=0.000 n=10+9)
/100000-4     17.6MB ± 0%    44.8MB ± 0%  +154.10%  (p=0.000 n=10+9)
/1000000-4     176MB ± 0%     448MB ± 0%  +154.17%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            10.0 ± 0%      15.0 ± 0%   +50.00%  (p=0.000 n=10+10)
/10-4           74.0 ± 0%     141.0 ± 0%   +90.54%  (p=0.000 n=10+10)
/100-4           707 ± 0%      1401 ± 0%   +98.16%  (p=0.000 n=10+10)
/1000-4        7.02k ± 0%    14.00k ± 0%   +99.53%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%    140.0k ± 0%   +99.54%  (p=0.000 n=10+10)
/100000-4       702k ± 0%     1400k ± 0%   +99.55%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%    14.00M ± 0%   +99.55%  (p=0.000 n=10+10)
```

stack vs [CustomSliceStack](testdata_test.go) - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceStackv1.0.1.txt testdata/BenchmarkMicroserviceSlice.txt
name        old time/op    new time/op    delta
/0-4          28.1ns ± 1%    39.3ns ± 2%   +40.04%  (p=0.000 n=10+10)
/1-4           322ns ± 0%     342ns ± 3%    +6.03%  (p=0.000 n=9+10)
/10-4         2.23µs ± 1%    2.16µs ± 3%    -3.34%  (p=0.000 n=10+9)
/100-4        20.6µs ± 0%    18.4µs ± 0%   -10.68%  (p=0.000 n=9+9)
/1000-4        203µs ± 0%     175µs ± 0%   -13.49%  (p=0.000 n=8+10)
/10000-4      2.12ms ± 1%    1.87ms ± 1%   -11.61%  (p=0.000 n=10+10)
/100000-4     24.0ms ± 3%    26.1ms ± 2%    +8.75%  (p=0.000 n=9+9)
/1000000-4     258ms ± 2%     298ms ±14%   +15.76%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           16.0B ± 0%     32.0B ± 0%  +100.00%  (p=0.000 n=10+10)
/1-4            288B ± 0%      200B ± 0%   -30.56%  (p=0.000 n=10+10)
/10-4         1.55kB ± 0%    1.40kB ± 0%    -9.79%  (p=0.000 n=10+10)
/100-4        15.2kB ± 0%    13.3kB ± 0%   -12.78%  (p=0.000 n=10+10)
/1000-4        161kB ± 0%     128kB ± 0%   -20.34%  (p=0.000 n=10+10)
/10000-4      1.76MB ± 0%    1.51MB ± 0%   -14.48%  (p=0.000 n=10+10)
/100000-4     17.6MB ± 0%    15.9MB ± 0%   -10.08%  (p=0.000 n=10+10)
/1000000-4     176MB ± 0%     157MB ± 0%   -10.82%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            10.0 ± 0%      11.0 ± 0%   +10.00%  (p=0.000 n=10+10)
/10-4           74.0 ± 0%      76.0 ± 0%    +2.70%  (p=0.000 n=10+10)
/100-4           707 ± 0%       709 ± 0%    +0.28%  (p=0.000 n=10+10)
/1000-4        7.02k ± 0%     7.01k ± 0%    -0.07%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     70.0k ± 0%    -0.20%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      700k ± 0%    -0.22%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     7.00M ± 0%    -0.22%  (p=0.000 n=10+10)
```

stack vs [cookiejar](https://github.com/karalabe/cookiejar/blob/master/collections/deque/deque.go) - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceStackv1.0.1.txt testdata/BenchmarkMicroserviceCookiejar.txt
name        old time/op    new time/op    delta
/0-4          28.1ns ± 1%  9907.4ns ± 3%   +35195.33%  (p=0.000 n=10+10)
/1-4           322ns ± 0%   10040ns ± 0%    +3013.61%  (p=0.000 n=9+9)
/10-4         2.23µs ± 1%   12.40µs ± 6%     +455.44%  (p=0.000 n=10+10)
/100-4        20.6µs ± 0%    29.7µs ± 1%      +44.44%  (p=0.000 n=9+8)
/1000-4        203µs ± 0%     203µs ± 1%       +0.35%  (p=0.021 n=8+9)
/10000-4      2.12ms ± 1%    1.98ms ± 1%       -6.56%  (p=0.000 n=10+10)
/100000-4     24.0ms ± 3%    23.1ms ± 2%       -3.92%  (p=0.000 n=9+9)
/1000000-4     258ms ± 2%     242ms ± 4%       -6.06%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           16.0B ± 0%  65648.0B ± 0%  +410200.00%  (p=0.000 n=10+10)
/1-4            288B ± 0%    65760B ± 0%   +22733.33%  (p=0.000 n=10+10)
/10-4         1.55kB ± 0%   66.77kB ± 0%    +4202.06%  (p=0.000 n=10+10)
/100-4        15.2kB ± 0%    76.8kB ± 0%     +405.05%  (p=0.000 n=10+10)
/1000-4        161kB ± 0%     178kB ± 0%      +10.20%  (p=0.000 n=10+10)
/10000-4      1.76MB ± 0%    1.32MB ± 0%      -25.23%  (p=0.000 n=10+10)
/100000-4     17.6MB ± 0%    12.8MB ± 0%      -27.17%  (p=0.000 n=10+10)
/1000000-4     176MB ± 0%     128MB ± 0%      -27.34%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      3.00 ± 0%     +200.00%  (p=0.000 n=10+10)
/1-4            10.0 ± 0%      10.0 ± 0%         ~     (all equal)
/10-4           74.0 ± 0%      73.0 ± 0%       -1.35%  (p=0.000 n=10+10)
/100-4           707 ± 0%       703 ± 0%       -0.57%  (p=0.000 n=10+10)
/1000-4        7.02k ± 0%     7.00k ± 0%       -0.20%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     70.0k ± 0%       -0.22%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      700k ± 0%       -0.22%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     7.00M ± 0%       -0.22%  (p=0.000 n=10+10)
```

stack vs [deque](https://github.com/ef-ds/deque) - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceStackv1.0.1.txt testdata/BenchmarkMicroserviceDeque.txt
name        old time/op    new time/op    delta
/0-4          28.1ns ± 1%    38.1ns ± 0%   +35.89%  (p=0.000 n=10+9)
/1-4           322ns ± 0%     357ns ± 1%   +10.75%  (p=0.000 n=9+9)
/10-4         2.23µs ± 1%    2.51µs ± 0%   +12.32%  (p=0.000 n=10+10)
/100-4        20.6µs ± 0%    23.4µs ± 0%   +13.84%  (p=0.000 n=9+9)
/1000-4        203µs ± 0%     222µs ± 0%    +9.55%  (p=0.000 n=8+10)
/10000-4      2.12ms ± 1%    2.28ms ± 1%    +7.48%  (p=0.000 n=10+10)
/100000-4     24.0ms ± 3%    24.9ms ± 1%    +3.75%  (p=0.000 n=9+10)
/1000000-4     258ms ± 2%     262ms ± 1%    +1.52%  (p=0.009 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           16.0B ± 0%     64.0B ± 0%  +300.00%  (p=0.000 n=10+10)
/1-4            288B ± 0%      288B ± 0%      ~     (all equal)
/10-4         1.55kB ± 0%    1.55kB ± 0%      ~     (all equal)
/100-4        15.2kB ± 0%    16.8kB ± 0%   +10.41%  (p=0.000 n=10+10)
/1000-4        161kB ± 0%     130kB ± 0%   -19.33%  (p=0.000 n=10+10)
/10000-4      1.76MB ± 0%    1.42MB ± 0%   -19.15%  (p=0.000 n=10+10)
/100000-4     17.6MB ± 0%    14.4MB ± 0%   -18.20%  (p=0.000 n=10+10)
/1000000-4     176MB ± 0%     144MB ± 0%   -18.10%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            10.0 ± 0%      10.0 ± 0%      ~     (all equal)
/10-4           74.0 ± 0%      74.0 ± 0%      ~     (all equal)
/100-4           707 ± 0%       707 ± 0%      ~     (all equal)
/1000-4        7.02k ± 0%     7.01k ± 0%    -0.06%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     70.2k ± 0%    -0.01%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      702k ± 0%    -0.00%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     7.02M ± 0%    -0.00%  (p=0.000 n=10+10)
```

### Other Test Results
#### stack vs [list](https://github.com/golang/go/tree/master/src/container/list)
stack vs list - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillStackv1.0.1.txt testdata/BenchmarkFillList.txt
name        old time/op    new time/op    delta
/0-4          27.6ns ± 1%    41.7ns ±10%   +51.09%  (p=0.000 n=10+10)
/1-4           156ns ± 1%     114ns ± 4%   -26.54%  (p=0.000 n=10+9)
/10-4          539ns ± 0%     747ns ± 4%   +38.77%  (p=0.000 n=9+9)
/100-4        4.13µs ± 0%    7.00µs ± 6%   +69.72%  (p=0.000 n=8+10)
/1000-4       34.2µs ± 0%    71.9µs ± 7%  +110.10%  (p=0.000 n=9+10)
/10000-4       321µs ± 0%     726µs ± 6%  +126.35%  (p=0.000 n=9+10)
/100000-4     3.50ms ± 1%   20.65ms ± 8%  +490.71%  (p=0.000 n=10+10)
/1000000-4    40.6ms ± 3%   158.5ms ± 8%  +290.39%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/0-4           16.0B ± 0%     48.0B ± 0%  +200.00%  (p=0.000 n=10+10)
/1-4            192B ± 0%      112B ± 0%   -41.67%  (p=0.000 n=10+10)
/10-4           592B ± 0%      688B ± 0%   +16.22%  (p=0.000 n=10+10)
/100-4        5.62kB ± 0%    6.45kB ± 0%   +14.81%  (p=0.000 n=10+10)
/1000-4       40.5kB ± 0%    64.0kB ± 0%   +58.03%  (p=0.000 n=10+10)
/10000-4       333kB ± 0%     640kB ± 0%   +92.46%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    6.40MB ± 0%   +98.76%  (p=0.000 n=10+10)
/1000000-4    32.1MB ± 0%    64.0MB ± 0%   +99.52%  (p=0.000 n=10+9)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            4.00 ± 0%      3.00 ± 0%   -25.00%  (p=0.000 n=10+10)
/10-4           14.0 ± 0%      21.0 ± 0%   +50.00%  (p=0.000 n=10+10)
/100-4           107 ± 0%       201 ± 0%   +87.85%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     2.00k ± 0%   +97.92%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     20.0k ± 0%   +99.07%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      200k ± 0%   +99.21%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     2.00M ± 0%   +99.22%  (p=0.000 n=10+10)
```

stack vs list - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillStackv1.0.1.txt testdata/BenchmarkRefillList.txt
name       old time/op    new time/op    delta
/1-4         3.21µs ± 1%    7.09µs ± 0%  +120.80%  (p=0.000 n=10+8)
/10-4        30.8µs ± 2%    70.8µs ± 5%  +129.59%  (p=0.000 n=10+10)
/100-4        296µs ± 2%     707µs ±11%  +138.68%  (p=0.000 n=10+10)
/1000-4      3.09ms ± 1%    7.21ms ± 5%  +133.58%  (p=0.000 n=10+10)
/10000-4     33.0ms ± 0%    72.5ms ± 5%  +119.36%  (p=0.000 n=9+10)
/100000-4     350ms ± 1%    2059ms ±14%  +487.53%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4      2.42MB ± 0%    6.40MB ± 0%  +164.20%  (p=0.000 n=9+10)
/10000-4     31.6MB ± 0%    64.0MB ± 0%  +102.37%  (p=0.000 n=10+9)
/100000-4     320MB ± 0%     640MB ± 0%   +99.77%  (p=0.000 n=10+9)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      200k ± 0%   +99.60%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     2.00M ± 0%   +99.24%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     20.0M ± 0%   +99.22%  (p=0.000 n=10+9)
```

stack vs list - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullStackv1.0.1.txt testdata/BenchmarkRefillFullList.txt
name       old time/op    new time/op    delta
/1-4         3.06µs ± 2%    8.19µs ± 2%  +168.13%  (p=0.000 n=10+10)
/10-4        30.8µs ± 2%    86.4µs ± 6%  +180.25%  (p=0.000 n=10+10)
/100-4        298µs ± 1%     876µs ± 6%  +193.51%  (p=0.000 n=10+10)
/1000-4      3.34ms ± 1%    9.16ms ±15%  +174.55%  (p=0.000 n=10+9)
/10000-4     33.5ms ± 0%    91.3ms ± 8%  +172.42%  (p=0.000 n=9+10)
/100000-4     346ms ± 0%    1729ms ± 6%  +400.18%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4      3.24MB ± 0%    6.40MB ± 0%   +97.24%  (p=0.000 n=10+10)
/10000-4     32.4MB ± 0%    64.0MB ± 0%   +97.24%  (p=0.000 n=10+10)
/100000-4     320MB ± 0%     640MB ± 0%   +99.77%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      200k ± 0%   +99.20%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     2.00M ± 0%   +99.20%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     20.0M ± 0%   +99.22%  (p=0.000 n=10+10)
```

stack vs list - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseStackv1.0.1.txt testdata/BenchmarkSlowIncreaseList.txt
name        old time/op    new time/op    delta
/1-4           187ns ± 1%     176ns ± 1%    -5.93%  (p=0.000 n=10+9)
/10-4          824ns ± 1%    1427ns ± 3%   +73.15%  (p=0.000 n=9+10)
/100-4        6.92µs ± 0%   13.44µs ± 0%   +94.27%  (p=0.000 n=8+10)
/1000-4       63.2µs ± 1%   135.8µs ± 1%  +114.89%  (p=0.000 n=10+10)
/10000-4       633µs ± 1%    1414µs ± 0%  +123.36%  (p=0.000 n=9+9)
/100000-4     7.30ms ± 0%   26.66ms ± 7%  +265.34%  (p=0.000 n=8+10)
/1000000-4    83.1ms ± 2%   338.6ms ±10%  +307.64%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%      176B ± 0%   -15.38%  (p=0.000 n=10+10)
/10-4           752B ± 0%     1328B ± 0%   +76.60%  (p=0.000 n=10+10)
/100-4        7.22kB ± 0%   12.85kB ± 0%   +78.05%  (p=0.000 n=10+10)
/1000-4       64.8kB ± 0%   128.0kB ± 0%   +97.75%  (p=0.000 n=10+10)
/10000-4       649kB ± 0%    1280kB ± 0%   +97.29%  (p=0.000 n=10+10)
/100000-4     6.42MB ± 0%   12.80MB ± 0%   +99.26%  (p=0.000 n=10+10)
/1000000-4    64.1MB ± 0%   128.0MB ± 0%   +99.57%  (p=0.000 n=10+8)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      5.00 ± 0%      ~     (all equal)
/10-4           24.0 ± 0%      41.0 ± 0%   +70.83%  (p=0.000 n=10+10)
/100-4           207 ± 0%       401 ± 0%   +93.72%  (p=0.000 n=10+10)
/1000-4        2.01k ± 0%     4.00k ± 0%   +98.76%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     40.0k ± 0%   +99.16%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      400k ± 0%   +99.21%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     4.00M ± 0%   +99.22%  (p=0.000 n=10+10)
```

stack vs list - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseStackv1.0.1.txt testdata/BenchmarkSlowDecreaseList.txt
name        old time/op    new time/op    delta
/1-4          30.9ns ± 2%    69.4ns ± 2%  +124.85%  (p=0.000 n=10+8)
/10-4          314ns ± 1%     689ns ± 0%  +119.81%  (p=0.000 n=9+9)
/100-4        3.05µs ± 1%    6.94µs ± 6%  +127.29%  (p=0.000 n=10+9)
/1000-4       30.4µs ± 2%    69.2µs ± 4%  +127.78%  (p=0.000 n=10+9)
/10000-4       303µs ± 1%     688µs ± 3%  +127.11%  (p=0.000 n=9+10)
/100000-4     3.03ms ± 1%    6.84ms ± 1%  +125.40%  (p=0.000 n=9+9)
/1000000-4    30.4ms ± 1%    68.5ms ± 1%  +125.24%  (p=0.000 n=10+10)

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
benchstat testdata/BenchmarkStableStackv1.0.1.txt testdata/BenchmarkStableList.txt
name        old time/op    new time/op    delta
/1-4          28.9ns ± 2%    79.9ns ± 2%  +176.14%  (p=0.000 n=10+10)
/10-4          294ns ± 3%     822ns ± 9%  +179.49%  (p=0.000 n=9+10)
/100-4        2.88µs ± 2%    8.01µs ± 2%  +178.63%  (p=0.000 n=10+9)
/1000-4       28.5µs ± 3%    79.9µs ± 2%  +180.77%  (p=0.000 n=10+9)
/10000-4       285µs ± 1%     799µs ± 1%  +180.56%  (p=0.000 n=10+9)
/100000-4     2.85ms ± 2%    7.98ms ± 1%  +180.33%  (p=0.000 n=9+10)
/1000000-4    28.4ms ± 1%    79.6ms ± 1%  +179.80%  (p=0.000 n=9+8)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     64.0B ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      640B ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000000-4    16.0MB ± 0%    64.0MB ± 0%  +300.00%  (p=0.000 n=9+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      2.00 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4           10.0 ± 0%      20.0 ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4           100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     2.00M ± 0%  +100.00%  (p=0.000 n=10+10)
```

#### stack vs [CustomSliceStack](https://github.com/ef-ds/stack-bench-tests/blob/master/testdata.go)
stack vs CustomSliceStack - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillStackv1.0.1.txt testdata/BenchmarkFillSlice.txt
name        old time/op    new time/op    delta
/0-4          27.6ns ± 1%    39.0ns ± 1%   +41.59%  (p=0.000 n=10+10)
/1-4           156ns ± 1%      93ns ± 1%   -40.44%  (p=0.000 n=10+10)
/10-4          539ns ± 0%     570ns ± 1%    +5.78%  (p=0.000 n=9+9)
/100-4        4.13µs ± 0%    3.58µs ± 1%   -13.32%  (p=0.000 n=8+10)
/1000-4       34.2µs ± 0%    30.0µs ± 1%   -12.36%  (p=0.000 n=9+9)
/10000-4       321µs ± 0%     368µs ± 1%   +14.57%  (p=0.000 n=9+10)
/100000-4     3.50ms ± 1%    8.11ms ± 2%  +132.06%  (p=0.000 n=10+9)
/1000000-4    40.6ms ± 3%   105.8ms ±10%  +160.47%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           16.0B ± 0%     32.0B ± 0%  +100.00%  (p=0.000 n=10+10)
/1-4            192B ± 0%       56B ± 0%   -70.83%  (p=0.000 n=10+10)
/10-4           592B ± 0%      440B ± 0%   -25.68%  (p=0.000 n=10+10)
/100-4        5.62kB ± 0%    3.67kB ± 0%   -34.62%  (p=0.000 n=10+10)
/1000-4       40.5kB ± 0%    32.4kB ± 0%   -20.04%  (p=0.000 n=10+10)
/10000-4       333kB ± 0%     546kB ± 0%   +64.28%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    6.25MB ± 0%   +94.24%  (p=0.000 n=10+10)
/1000000-4    32.1MB ± 0%    61.2MB ± 0%   +90.75%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            4.00 ± 0%      3.00 ± 0%   -25.00%  (p=0.000 n=10+10)
/10-4           14.0 ± 0%      16.0 ± 0%   +14.29%  (p=0.000 n=10+10)
/100-4           107 ± 0%       109 ± 0%    +1.87%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     1.01k ± 0%    +0.10%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     10.0k ± 0%    -0.26%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      100k ± 0%    -0.37%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     1.00M ± 0%    -0.39%  (p=0.000 n=10+10)
```

stack vs CustomSliceStack - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillStackv1.0.1.txt testdata/BenchmarkRefillSlice.txt
name       old time/op    new time/op    delta
/1-4         3.21µs ± 1%    2.85µs ± 0%  -11.16%  (p=0.000 n=10+9)
/10-4        30.8µs ± 2%    27.9µs ± 1%   -9.46%  (p=0.000 n=10+9)
/100-4        296µs ± 2%     282µs ± 9%   -4.59%  (p=0.043 n=10+10)
/1000-4      3.09ms ± 1%    2.73ms ± 7%  -11.35%  (p=0.000 n=10+10)
/10000-4     33.0ms ± 0%    32.5ms ±12%     ~     (p=0.050 n=9+9)
/100000-4     350ms ± 1%     388ms ± 4%  +10.78%  (p=0.000 n=10+8)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      2.42MB ± 0%    1.60MB ± 0%  -33.95%  (p=0.000 n=9+9)
/10000-4     31.6MB ± 0%    16.0MB ± 0%  -49.38%  (p=0.000 n=10+9)
/100000-4     320MB ± 0%     162MB ± 0%  -49.57%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%   -0.20%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     1.00M ± 0%   -0.38%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     10.0M ± 0%   -0.39%  (p=0.000 n=10+10)
```

stack vs CustomSliceStack - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullStackv1.0.1.txt testdata/BenchmarkRefillFullSlice.txt
name       old time/op    new time/op    delta
/1-4         3.06µs ± 2%    3.18µs ± 7%   +4.03%  (p=0.002 n=10+10)
/10-4        30.8µs ± 2%    32.3µs ± 6%   +4.79%  (p=0.005 n=10+10)
/100-4        298µs ± 1%     309µs ± 4%   +3.59%  (p=0.000 n=10+10)
/1000-4      3.34ms ± 1%    3.11ms ± 6%   -6.92%  (p=0.000 n=10+10)
/10000-4     33.5ms ± 0%    31.1ms ± 6%   -7.25%  (p=0.000 n=9+10)
/100000-4     346ms ± 0%     390ms ± 9%  +12.79%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      3.24MB ± 0%    1.60MB ± 0%  -50.69%  (p=0.000 n=10+10)
/10000-4     32.4MB ± 0%    16.0MB ± 0%  -50.69%  (p=0.000 n=10+10)
/100000-4     320MB ± 0%     160MB ± 0%  -50.06%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%   -0.40%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     1.00M ± 0%   -0.40%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     10.0M ± 0%   -0.39%  (p=0.000 n=10+10)
```

stack vs CustomSliceStack - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseStackv1.0.1.txt testdata/BenchmarkSlowIncreaseSlice.txt
name        old time/op    new time/op    delta
/1-4           187ns ± 1%     153ns ± 1%  -18.34%  (p=0.000 n=10+9)
/10-4          824ns ± 1%     863ns ± 2%   +4.75%  (p=0.000 n=9+10)
/100-4        6.92µs ± 0%    6.09µs ± 2%  -11.90%  (p=0.000 n=8+10)
/1000-4       63.2µs ± 1%    55.7µs ± 3%  -11.89%  (p=0.000 n=10+10)
/10000-4       633µs ± 1%     614µs ± 4%   -3.03%  (p=0.004 n=9+9)
/100000-4     7.30ms ± 0%   11.10ms ± 2%  +52.07%  (p=0.000 n=8+8)
/1000000-4    83.1ms ± 2%   140.3ms ± 9%  +68.95%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%       88B ± 0%  -57.69%  (p=0.000 n=10+10)
/10-4           752B ± 0%      600B ± 0%  -20.21%  (p=0.000 n=10+10)
/100-4        7.22kB ± 0%    5.27kB ± 0%  -26.94%  (p=0.000 n=10+10)
/1000-4       64.8kB ± 0%    48.4kB ± 0%  -25.24%  (p=0.000 n=10+10)
/10000-4       649kB ± 0%     706kB ± 0%   +8.86%  (p=0.000 n=10+10)
/100000-4     6.42MB ± 0%    7.85MB ± 0%  +22.27%  (p=0.000 n=10+10)
/1000000-4    64.1MB ± 0%    77.2MB ± 0%  +20.34%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      5.00 ± 0%     ~     (all equal)
/10-4           24.0 ± 0%      26.0 ± 0%   +8.33%  (p=0.000 n=10+10)
/100-4           207 ± 0%       209 ± 0%   +0.97%  (p=0.000 n=10+10)
/1000-4        2.01k ± 0%     2.01k ± 0%   -0.05%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.0k ± 0%   -0.32%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      200k ± 0%   -0.38%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.00M ± 0%   -0.39%  (p=0.000 n=10+10)
```

stack vs CustomSliceStack - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseStackv1.0.1.txt testdata/BenchmarkSlowDecreaseSlice.txt
name        old time/op    new time/op    delta
/1-4          30.9ns ± 2%    32.4ns ± 3%  +4.86%  (p=0.000 n=10+10)
/10-4          314ns ± 1%     323ns ± 1%  +2.93%  (p=0.000 n=9+8)
/100-4        3.05µs ± 1%    3.18µs ± 4%  +4.18%  (p=0.000 n=10+10)
/1000-4       30.4µs ± 2%    31.3µs ± 4%  +2.91%  (p=0.002 n=10+10)
/10000-4       303µs ± 1%     314µs ± 4%  +3.84%  (p=0.002 n=9+10)
/100000-4     3.03ms ± 1%    3.13ms ± 3%  +3.03%  (p=0.000 n=9+10)
/1000000-4    30.4ms ± 1%    31.6ms ± 2%  +4.06%  (p=0.000 n=10+9)

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
benchstat testdata/BenchmarkStableStackv1.0.1.txt testdata/BenchmarkStableSlice.txt
name        old time/op    new time/op    delta
/1-4          28.9ns ± 2%    30.8ns ± 6%  +6.50%  (p=0.000 n=10+10)
/10-4          294ns ± 3%     308ns ± 7%  +4.80%  (p=0.005 n=9+9)
/100-4        2.88µs ± 2%    3.00µs ± 3%  +4.51%  (p=0.000 n=10+10)
/1000-4       28.5µs ± 3%    29.9µs ± 5%  +4.93%  (p=0.000 n=10+10)
/10000-4       285µs ± 1%     300µs ± 4%  +5.46%  (p=0.000 n=10+10)
/100000-4     2.85ms ± 2%    2.97ms ± 5%  +4.23%  (p=0.003 n=9+9)
/1000000-4    28.4ms ± 1%    30.0ms ± 4%  +5.51%  (p=0.000 n=9+10)

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
benchstat testdata/BenchmarkFillStackv1.0.1.txt testdata/BenchmarkFillCookiejar.txt
name        old time/op    new time/op     delta
/0-4          27.6ns ± 1%  10035.1ns ± 7%   +36285.46%  (p=0.000 n=10+9)
/1-4           156ns ± 1%    10756ns ±16%    +6803.59%  (p=0.000 n=10+10)
/10-4          539ns ± 0%    10329ns ± 4%    +1817.95%  (p=0.000 n=9+9)
/100-4        4.13µs ± 0%    12.82µs ± 3%     +210.71%  (p=0.000 n=8+9)
/1000-4       34.2µs ± 0%     39.3µs ± 2%      +14.82%  (p=0.000 n=9+9)
/10000-4       321µs ± 0%      324µs ± 1%       +0.87%  (p=0.004 n=9+9)
/100000-4     3.50ms ± 1%     3.61ms ± 5%       +3.32%  (p=0.000 n=10+10)
/1000000-4    40.6ms ± 3%     43.6ms ± 3%       +7.29%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op    delta
/0-4           16.0B ± 0%   65648.0B ± 0%  +410200.00%  (p=0.000 n=10+10)
/1-4            192B ± 0%     65664B ± 0%   +34100.00%  (p=0.000 n=10+10)
/10-4           592B ± 0%     65808B ± 0%   +11016.22%  (p=0.000 n=10+10)
/100-4        5.62kB ± 0%    67.25kB ± 0%    +1097.44%  (p=0.000 n=10+10)
/1000-4       40.5kB ± 0%     81.6kB ± 0%     +101.46%  (p=0.000 n=10+10)
/10000-4       333kB ± 0%      357kB ± 0%       +7.31%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%     3.24MB ± 0%       +0.62%  (p=0.000 n=10+10)
/1000000-4    32.1MB ± 0%     32.1MB ± 0%       -0.03%  (p=0.000 n=10+9)

name        old allocs/op  new allocs/op   delta
/0-4            1.00 ± 0%       3.00 ± 0%     +200.00%  (p=0.000 n=10+10)
/1-4            4.00 ± 0%       4.00 ± 0%         ~     (all equal)
/10-4           14.0 ± 0%       13.0 ± 0%       -7.14%  (p=0.000 n=10+10)
/100-4           107 ± 0%        103 ± 0%       -3.74%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%      1.00k ± 0%       -0.79%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%      10.0k ± 0%       -0.40%  (p=0.000 n=10+10)
/100000-4       100k ± 0%       100k ± 0%       -0.37%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%      1.00M ± 0%       -0.36%  (p=0.000 n=10+10)
```

stack vs cookiejar - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillStackv1.0.1.txt testdata/BenchmarkRefillCookiejar.txt
name       old time/op    new time/op    delta
/1-4         3.21µs ± 1%    3.37µs ± 3%   +4.92%  (p=0.000 n=10+10)
/10-4        30.8µs ± 2%    31.4µs ± 2%   +1.86%  (p=0.005 n=10+10)
/100-4        296µs ± 2%     306µs ± 2%   +3.51%  (p=0.000 n=10+10)
/1000-4      3.09ms ± 1%    3.06ms ± 3%     ~     (p=0.143 n=10+10)
/10000-4     33.0ms ± 0%    30.7ms ± 5%   -7.08%  (p=0.000 n=9+10)
/100000-4     350ms ± 1%     367ms ± 5%   +4.75%  (p=0.000 n=10+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%   +0.01%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     160kB ± 0%   +0.01%  (p=0.000 n=10+10)
/1000-4      2.42MB ± 0%    1.60MB ± 0%  -33.95%  (p=0.000 n=9+10)
/10000-4     31.6MB ± 0%    16.0MB ± 0%  -49.40%  (p=0.000 n=10+9)
/100000-4     320MB ± 0%     161MB ± 0%  -49.89%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%   -0.20%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     1.00M ± 0%   -0.38%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     10.0M ± 0%   -0.39%  (p=0.000 n=10+10)
```

stack vs cookiejar - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullStackv1.0.1.txt testdata/BenchmarkRefillFullCookiejar.txt
name       old time/op    new time/op    delta
/1-4         3.06µs ± 2%    3.08µs ± 2%   +0.91%  (p=0.045 n=10+10)
/10-4        30.8µs ± 2%    31.5µs ± 4%   +2.25%  (p=0.007 n=10+10)
/100-4        298µs ± 1%     308µs ± 3%   +3.34%  (p=0.000 n=10+10)
/1000-4      3.34ms ± 1%    3.13ms ± 8%   -6.12%  (p=0.002 n=10+10)
/10000-4     33.5ms ± 0%    32.6ms ± 7%     ~     (p=0.549 n=9+10)
/100000-4     346ms ± 0%     349ms ± 3%     ~     (p=0.050 n=9+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      3.24MB ± 0%    1.60MB ± 0%  -50.69%  (p=0.000 n=10+10)
/10000-4     32.4MB ± 0%    16.0MB ± 0%  -50.69%  (p=0.000 n=10+9)
/100000-4     320MB ± 0%     160MB ± 0%  -50.06%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%   -0.40%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     1.00M ± 0%   -0.40%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     10.0M ± 0%   -0.39%  (p=0.000 n=10+10)
```

stack vs cookiejar - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseStackv1.0.1.txt testdata/BenchmarkSlowIncreaseCookiejar.txt
name        old time/op    new time/op    delta
/1-4           187ns ± 1%    9932ns ± 0%   +5208.44%  (p=0.000 n=10+10)
/10-4          824ns ± 1%   10644ns ± 1%   +1191.91%  (p=0.000 n=9+9)
/100-4        6.92µs ± 0%   16.01µs ± 4%    +131.45%  (p=0.000 n=8+10)
/1000-4       63.2µs ± 1%    69.3µs ± 4%      +9.68%  (p=0.000 n=10+10)
/10000-4       633µs ± 1%     618µs ± 1%      -2.36%  (p=0.000 n=9+10)
/100000-4     7.30ms ± 0%    7.84ms ± 2%      +7.42%  (p=0.000 n=8+10)
/1000000-4    83.1ms ± 2%    85.8ms ± 3%      +3.27%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%    65680B ± 0%  +31476.92%  (p=0.000 n=10+10)
/10-4           752B ± 0%    65968B ± 0%   +8672.34%  (p=0.000 n=10+10)
/100-4        7.22kB ± 0%   68.85kB ± 0%    +854.10%  (p=0.000 n=10+10)
/1000-4       64.8kB ± 0%    97.6kB ± 0%     +50.80%  (p=0.000 n=10+10)
/10000-4       649kB ± 0%     517kB ± 0%     -20.34%  (p=0.000 n=10+10)
/100000-4     6.42MB ± 0%    4.84MB ± 0%     -24.65%  (p=0.000 n=10+10)
/1000000-4    64.1MB ± 0%    48.1MB ± 0%     -25.06%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      5.00 ± 0%        ~     (all equal)
/10-4           24.0 ± 0%      23.0 ± 0%      -4.17%  (p=0.000 n=10+10)
/100-4           207 ± 0%       203 ± 0%      -1.93%  (p=0.000 n=10+10)
/1000-4        2.01k ± 0%     2.00k ± 0%      -0.50%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.0k ± 0%      -0.39%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      200k ± 0%      -0.38%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.00M ± 0%      -0.38%  (p=0.000 n=10+10)
```

stack vs cookiejar - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseStackv1.0.1.txt testdata/BenchmarkSlowDecreaseCookiejar.txt
name        old time/op    new time/op    delta
/1-4          30.9ns ± 2%    33.8ns ± 2%   +9.62%  (p=0.000 n=10+10)
/10-4          314ns ± 1%     349ns ± 4%  +11.34%  (p=0.000 n=9+10)
/100-4        3.05µs ± 1%    3.36µs ± 7%  +10.11%  (p=0.000 n=10+10)
/1000-4       30.4µs ± 2%    33.1µs ± 2%   +8.88%  (p=0.000 n=10+10)
/10000-4       303µs ± 1%     328µs ± 1%   +8.44%  (p=0.000 n=9+10)
/100000-4     3.03ms ± 1%    3.28ms ± 1%   +8.22%  (p=0.000 n=9+8)
/1000000-4    30.4ms ± 1%    32.8ms ± 1%   +7.92%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/100-4           100 ± 0%       100 ± 0%     ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (all equal)
```

stack vs cookiejar - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableStackv1.0.1.txt testdata/BenchmarkStableCookiejar.txt
name        old time/op    new time/op    delta
/1-4          28.9ns ± 2%    29.9ns ± 3%   +3.43%  (p=0.000 n=10+9)
/10-4          294ns ± 3%     333ns ±11%  +13.13%  (p=0.000 n=9+10)
/100-4        2.88µs ± 2%    3.12µs ± 6%   +8.54%  (p=0.000 n=10+10)
/1000-4       28.5µs ± 3%    31.5µs ± 8%  +10.71%  (p=0.000 n=10+10)
/10000-4       285µs ± 1%     318µs ± 8%  +11.62%  (p=0.000 n=10+10)
/100000-4     2.85ms ± 2%    3.08ms ± 7%   +8.28%  (p=0.000 n=9+10)
/1000000-4    28.4ms ± 1%    30.1ms ± 8%   +5.75%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/100-4           100 ± 0%       100 ± 0%     ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (all equal)
```

#### stack vs [deque](https://github.com/ef-ds/deque)
stack vs Deque - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillStackv1.0.1.txt testdata/BenchmarkFillDeque.txt
name        old time/op    new time/op    delta
/0-4          27.6ns ± 1%    37.4ns ± 0%   +35.73%  (p=0.000 n=10+9)
/1-4           156ns ± 1%     171ns ± 1%   +10.04%  (p=0.000 n=10+9)
/10-4          539ns ± 0%     583ns ± 0%    +8.27%  (p=0.000 n=9+9)
/100-4        4.13µs ± 0%    4.73µs ± 0%   +14.73%  (p=0.000 n=8+8)
/1000-4       34.2µs ± 0%    36.6µs ± 0%    +6.89%  (p=0.000 n=9+9)
/10000-4       321µs ± 0%     363µs ± 0%   +13.10%  (p=0.000 n=9+9)
/100000-4     3.50ms ± 1%    3.85ms ± 2%   +10.14%  (p=0.000 n=10+10)
/1000000-4    40.6ms ± 3%    44.0ms ± 3%    +8.40%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           16.0B ± 0%     64.0B ± 0%  +300.00%  (p=0.000 n=10+10)
/1-4            192B ± 0%      192B ± 0%      ~     (all equal)
/10-4           592B ± 0%      592B ± 0%      ~     (all equal)
/100-4        5.62kB ± 0%    7.20kB ± 0%   +28.21%  (p=0.000 n=10+10)
/1000-4       40.5kB ± 0%    34.0kB ± 0%   -16.03%  (p=0.000 n=10+10)
/10000-4       333kB ± 0%     323kB ± 0%    -2.85%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    3.22MB ± 0%    +0.06%  (p=0.000 n=10+10)
/1000000-4    32.1MB ± 0%    32.2MB ± 0%    +0.34%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            4.00 ± 0%      4.00 ± 0%      ~     (all equal)
/10-4           14.0 ± 0%      14.0 ± 0%      ~     (all equal)
/100-4           107 ± 0%       107 ± 0%      ~     (all equal)
/1000-4        1.01k ± 0%     1.01k ± 0%    +0.20%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     10.1k ± 0%    +0.36%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      101k ± 0%    +0.39%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     1.01M ± 0%    +0.39%  (p=0.000 n=10+10)
```

stack vs Deque - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillStackv1.0.1.txt testdata/BenchmarkRefillDeque.txt
name       old time/op    new time/op    delta
/1-4         3.21µs ± 1%    3.65µs ± 2%  +13.68%  (p=0.000 n=10+10)
/10-4        30.8µs ± 2%    35.9µs ± 1%  +16.37%  (p=0.000 n=10+10)
/100-4        296µs ± 2%     345µs ± 1%  +16.38%  (p=0.000 n=10+9)
/1000-4      3.09ms ± 1%    3.39ms ± 2%   +9.87%  (p=0.000 n=10+10)
/10000-4     33.0ms ± 0%    36.5ms ± 1%  +10.32%  (p=0.000 n=9+10)
/100000-4     350ms ± 1%     395ms ± 2%  +12.75%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%   +0.00%  (p=0.000 n=10+10)
/1000-4      2.42MB ± 0%    1.60MB ± 0%  -33.95%  (p=0.000 n=9+10)
/10000-4     31.6MB ± 0%    30.1MB ± 0%   -4.86%  (p=0.000 n=10+9)
/100000-4     320MB ± 0%     320MB ± 0%   -0.13%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%   -0.20%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     1.01M ± 0%   +0.30%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     10.1M ± 0%   +0.38%  (p=0.000 n=10+10)
```

stack vs Deque - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullStackv1.0.1.txt testdata/BenchmarkRefillFullDeque.txt
name       old time/op    new time/op    delta
/1-4         3.06µs ± 2%    3.42µs ± 1%  +11.89%  (p=0.000 n=10+10)
/10-4        30.8µs ± 2%    34.4µs ± 1%  +11.71%  (p=0.000 n=10+9)
/100-4        298µs ± 1%     340µs ± 2%  +13.98%  (p=0.000 n=10+10)
/1000-4      3.34ms ± 1%    3.35ms ± 0%     ~     (p=0.052 n=10+10)
/10000-4     33.5ms ± 0%    37.2ms ± 0%  +11.06%  (p=0.000 n=9+10)
/100000-4     346ms ± 0%     393ms ± 3%  +13.67%  (p=0.000 n=9+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      3.24MB ± 0%    1.60MB ± 0%  -50.69%  (p=0.000 n=10+10)
/10000-4     32.4MB ± 0%    30.1MB ± 0%   -7.27%  (p=0.000 n=10+10)
/100000-4     320MB ± 0%     320MB ± 0%   -0.13%  (p=0.000 n=10+9)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%   -0.40%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     1.01M ± 0%   +0.28%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     10.1M ± 0%   +0.38%  (p=0.000 n=10+10)
```

stack vs Deque - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseStackv1.0.1.txt testdata/BenchmarkSlowIncreaseDeque.txt
name        old time/op    new time/op    delta
/1-4           187ns ± 1%     205ns ± 1%   +9.73%  (p=0.000 n=10+10)
/10-4          824ns ± 1%     915ns ± 0%  +11.05%  (p=0.000 n=9+9)
/100-4        6.92µs ± 0%    7.98µs ± 0%  +15.31%  (p=0.000 n=8+10)
/1000-4       63.2µs ± 1%    68.8µs ± 0%   +8.80%  (p=0.000 n=10+10)
/10000-4       633µs ± 1%     688µs ± 0%   +8.60%  (p=0.000 n=9+10)
/100000-4     7.30ms ± 0%    7.90ms ± 1%   +8.25%  (p=0.000 n=8+10)
/1000000-4    83.1ms ± 2%    83.2ms ± 1%     ~     (p=0.579 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%      208B ± 0%     ~     (all equal)
/10-4           752B ± 0%      752B ± 0%     ~     (all equal)
/100-4        7.22kB ± 0%    8.80kB ± 0%  +21.95%  (p=0.000 n=10+10)
/1000-4       64.8kB ± 0%    50.0kB ± 0%  -22.73%  (p=0.000 n=10+10)
/10000-4       649kB ± 0%     483kB ± 0%  -25.55%  (p=0.000 n=10+10)
/100000-4     6.42MB ± 0%    4.82MB ± 0%  -24.94%  (p=0.000 n=10+10)
/1000000-4    64.1MB ± 0%    48.2MB ± 0%  -24.86%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      5.00 ± 0%     ~     (all equal)
/10-4           24.0 ± 0%      24.0 ± 0%     ~     (all equal)
/100-4           207 ± 0%       207 ± 0%     ~     (all equal)
/1000-4        2.01k ± 0%     2.01k ± 0%     ~     (all equal)
/10000-4       20.1k ± 0%     20.1k ± 0%   -0.01%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      201k ± 0%   -0.00%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.01M ± 0%   -0.00%  (p=0.000 n=10+10)
```

stack vs Deque - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseStackv1.0.1.txt testdata/BenchmarkSlowDecreaseDeque.txt
name        old time/op    new time/op    delta
/1-4          30.9ns ± 2%    35.8ns ± 2%  +15.84%  (p=0.000 n=10+10)
/10-4          314ns ± 1%     359ns ± 1%  +14.40%  (p=0.000 n=9+10)
/100-4        3.05µs ± 1%    3.51µs ± 1%  +14.90%  (p=0.000 n=10+10)
/1000-4       30.4µs ± 2%    34.9µs ± 1%  +14.82%  (p=0.000 n=10+9)
/10000-4       303µs ± 1%     350µs ± 1%  +15.63%  (p=0.000 n=9+9)
/100000-4     3.03ms ± 1%    3.49ms ± 0%  +15.10%  (p=0.000 n=9+8)
/1000000-4    30.4ms ± 1%    34.8ms ± 2%  +14.61%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/100-4           100 ± 0%       100 ± 0%     ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (all equal)
```

stack vs Deque - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableStackv1.0.1.txt testdata/BenchmarkStableDeque.txt
name        old time/op    new time/op    delta
/1-4          28.9ns ± 2%    33.6ns ± 1%  +16.18%  (p=0.000 n=10+9)
/10-4          294ns ± 3%     342ns ± 2%  +16.43%  (p=0.000 n=9+10)
/100-4        2.88µs ± 2%    3.36µs ± 2%  +16.79%  (p=0.000 n=10+10)
/1000-4       28.5µs ± 3%    33.2µs ± 1%  +16.62%  (p=0.000 n=10+9)
/10000-4       285µs ± 1%     333µs ± 4%  +17.09%  (p=0.000 n=10+10)
/100000-4     2.85ms ± 2%    3.32ms ± 2%  +16.71%  (p=0.000 n=9+10)
/1000000-4    28.4ms ± 1%    33.3ms ± 2%  +16.95%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/100-4           100 ± 0%       100 ± 0%     ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (all equal)
```
