- Output with a job that run every 5s on a fixed schedule and run for a random amount of time no more than 5 S
- Metrics at localhost:8080/metrics

## Output

```
# HELP sched_run_actual_elapsed_time sched_run_actual_elapsed_time summary
# TYPE sched_run_actual_elapsed_time summary
sched_run_actual_elapsed_time{id="every5s",quantile="0.5"} 0.203843151
sched_run_actual_elapsed_time{id="every5s",quantile="0.75"} 1.104031623
sched_run_actual_elapsed_time{id="every5s",quantile="0.95"} 1.104031623
sched_run_actual_elapsed_time{id="every5s",quantile="0.99"} 1.104031623
sched_run_actual_elapsed_time{id="every5s",quantile="0.999"} 1.104031623
sched_run_actual_elapsed_time_sum{id="every5s"} 1.307874774
sched_run_actual_elapsed_time_count{id="every5s"} 2
# HELP sched_run_errors sched_run_errors counter
# TYPE sched_run_errors counter
sched_run_errors{id="every5s"} 0
# HELP sched_run_exceed_expected_time sched_run_exceed_expected_time counter
# TYPE sched_run_exceed_expected_time counter
sched_run_exceed_expected_time{id="every5s"} 0
# HELP sched_run_total_elapsed_time sched_run_total_elapsed_time summary
# TYPE sched_run_total_elapsed_time summary
sched_run_total_elapsed_time{id="every5s",quantile="0.5"} 0.203880714
sched_run_total_elapsed_time{id="every5s",quantile="0.75"} 1.104065614
sched_run_total_elapsed_time{id="every5s",quantile="0.95"} 1.104065614
sched_run_total_elapsed_time{id="every5s",quantile="0.99"} 1.104065614
sched_run_total_elapsed_time{id="every5s",quantile="0.999"} 1.104065614
sched_run_total_elapsed_time_sum{id="every5s"} 1.307946328
sched_run_total_elapsed_time_count{id="every5s"} 2
# HELP sched_runs sched_runs counter
# TYPE sched_runs counter
sched_runs{id="every5s"} 2
# HELP sched_runs_overlapping sched_runs_overlapping counter
# TYPE sched_runs_overlapping counter
sched_runs_overlapping{id="every5s"} 0
# HELP sched_up sched_up gauge
# TYPE sched_up gauge
sched_up{id="every5s"} 1
```
