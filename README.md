## telemetry-perfbench

Performance benchmarks comparing various telemetry projects.

This repository exists to provide benchmarks of popular logging libraries to compare against logging using the API in `x/tools/internal/telemetry`.

## Running the benchmarks

1. Follow the instructions [here](https://go-review.googlesource.com/c/tools/+/212078) to get results for no logging, std logging, and `internal/telemetry` logging.
2. Run the following command to get results for third party loggers:
    > go test -run=^$ -bench=. -count=10

## Benchstat Results
Below is a comparison of the results for each logger, obtained using benchstat. These results were obtained on a machine running Ubuntu 18.04 with an i5-6300HQ and 16GB DDR4-2133 memory.
> benchstat no_log.txt stdlog.txt tellog.txt zerolog.txt



| test   | time/op     | alloc/op    | allocs/op      |
|--------|-------------|-------------|----------------|
|no_log  | 376ns ± 4%  | 80.0B ± 0%  | 5.00 ± 0%      |
|stdlog  | 6772ns ± 1% | 568.0B ± 0% | 28.00 ± 0%     |
|tellog  | 3792ns ± 0% | 728.0B ± 0% | 32.00 ± 0%     |
|zerolog | 3802ns ±12% | 488.0B ± 0% | 23.00 ± 0%     |

## Comments

From these results, telemetry logging's runtime outperforms std logging significantly and performs similarly to zerolog. However, telemetry logging has a noticeably higher amount of allocations and bytes allocated.
