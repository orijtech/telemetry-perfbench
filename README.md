## telemetry-perfbench

Performance benchmarks comparing various telemetry projects.

This repository exists to provide benchmarks of popular logging libraries to compare against logging using the API in `x/tools/internal/telemetry`.

## Benchmarking Locally
This section is based on the instructions in [this](https://go-review.googlesource.com/c/tools/+/212078) CL.
### Running these benchmarks
1. Get this repository:
> go get -t -d -u github.com/kirbyquerby/go-logger-benchmarks
2. In the root directory of this repository, run the following command:
> go test -run=^$ -bench=. -count=10
3. Place each benchmark's output in its own file (e.g. `stdlog.txt`, `zerolog.txt`, etc).
4. Ensure all benchmark results have a common name (e.g. s/BenchmarkLoggingStdlib/BenchmarkIt/g).



### Benchmarking `internal/telemetry`
1. Clone the [tools](https://golang.org/x/tools) subrepository.
2. Inside the repository, `cd` to `internal/telemetry/log`.
3. Run the following command to obtain results:
> go test -run=^$ -bench=. -count=10
4. Retrieve results for the benchmarks you are interested in and store them according to the previous instructions.

### Comparing Results
All results can be compared using benchstat:
> benchstat no_log.txt stdlog.txt tellog.txt zerolog.txt zap_sugared.txt

Two results can also be compared to show percent change:
> benchstat stdlog.txt tellog.txt
## Benchstat Results
Below is a comparison of the results for each logger, obtained using benchstat. These results were obtained on a machine running Linux Mint 19 with an i7-8086K@5.0ghz and 32GB DDR4-3200 memory.




| test       | time/op     | alloc/op    | allocs/op      |
|------------|-------------|-------------|----------------|
|no_log      | 216ns ± 1%  | 80.0B ± 0%  | 5.00 ± 0%      |
|stdlog      | 3805ns ± 4% | 568.0B ± 0% | 28.00 ± 0%     |
|tellog      | 2215ns ± 3% | 728.0B ± 0% | 32.00 ± 0%     |
|zap_sugared | 6238ns ± 1% | 569.0B ± 0% | 28.00 ± 0%     |
|zerolog     | 2282ns ± 1% | 488.0B ± 0% | 23.00 ± 0%     |

## Comments

From these results, telemetry logging's runtime outperforms std logging significantly and performs similarly to zerolog. However, telemetry logging has a noticeably higher amount of allocations and bytes allocated.
From these results, telemetry logging's runtime outperforms std logging and sugared zap significantly and performs similarly to zerolog. However, telemetry logging has a noticeably higher amount of allocations and bytes allocated.
