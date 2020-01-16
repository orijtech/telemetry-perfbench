## telemetry-perfbench

Performance benchmarks comparing various telemetry projects.

This repository exists to provide benchmarks of popular logging libraries to compare against logging using the API in `x/tools/internal/telemetry`.

## Running the benchmarks

1. Follow the instructions [here](https://go-review.googlesource.com/c/tools/+/212078) to get results for no logging, std logging, and `internal/telemetry` logging.
2. Run the following command to get results for third party loggers:
    > go test -run=^$ -bench=. -count=10

## Results
