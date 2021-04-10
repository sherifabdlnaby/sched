- Output with a job that run every 5s on a fixed schedule and run for a random amount of time no more than 5 S
- Metrics at localhost:8080/metrics

## Output

```json
# HELP sched_run_actual_elapsed_time sched_run_actual_elapsed_time summary
# TYPE sched_run_actual_elapsed_time summary
sched_run_actual_elapsed_time{ID="every5s", quantile="0.5"
} 0.205075002
sched_run_actual_elapsed_time{
ID="every5s", quantile="0.75"
} 1.103087326
sched_run_actual_elapsed_time{
ID="every5s", quantile="0.95"
} 2.204325295
sched_run_actual_elapsed_time{
ID="every5s",quantile="0.99"
} 2.204325295
sched_run_actual_elapsed_time{
ID="every5s", quantile="0.999"
} 2.204325295
sched_run_actual_elapsed_time_sum{
ID="every5s"
} 3.713109351
sched_run_actual_elapsed_time_count{
ID="every5s"
} 4
# HELP sched_run_errors sched_run_errors counter
# TYPE sched_run_errors counter
sched_run_errors{ID="every5s"} 0
# HELP sched_run_total_elapsed_time sched_run_total_elapsed_time summary
# TYPE sched_run_total_elapsed_time summary
sched_run_total_elapsed_time{
ID="every5s", quantile="0.5"
} 0.205178562
sched_run_total_elapsed_time{ID="every5s", quantile="0.75"
} 1.103123691
sched_run_total_elapsed_time{
ID="every5s", quantile="0.95"
} 2.204367762
sched_run_total_elapsed_time{
ID="every5s", quantile="0.99"
} 2.204367762
sched_run_total_elapsed_time{
ID="every5s",quantile="0.999"
} 2.204367762
sched_run_total_elapsed_time_sum{
ID="every5s"
} 3.7133283589999997
sched_run_total_elapsed_time_count{ID="every5s"} 4
# HELP sched_runs sched_runs counter
# TYPE sched_runs counter
sched_runs{ID="every5s"} 3
# HELP sched_runs_overlapping sched_runs_overlapping counter
# TYPE sched_runs_overlapping counter
sched_runs_overlapping{ID="every5s"} 0
# HELP sched_up sched_up gauge
# TYPE sched_up gauge
sched_up{ID="every5s"} 1
```
