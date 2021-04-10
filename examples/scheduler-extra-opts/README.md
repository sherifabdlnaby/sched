# 4 Schedules in a Scheduler Manager

1. Cron Every Minute
2. Cron Every 5 Minutes
3. Fixed Interval Every 30 Secs
4. *Once* after 10 Secs from schedule start.

Started and Stopped using `StartAll()` and `StopAll()`

## Extra Options

While all schedules inherit the same opts passed to their scheduler. The `fixedTimer30second` schedule has an extra
Option passed to it. Extra Options override inherited options.

## Output

```json
2021-04-10T14:02:12.912+0200    INFO    sched   sched/schedule.go:101   Job Schedule Started    {"id": "cronEveryMinute"}
2021-04-10T14:02:12.912+0200    INFO    sched   sched/schedule.go:101   Job Schedule Started    {"id": "cronEvery5Minute"}
2021-04-10T14:02:12.912+0200    INFO    sched   sched/schedule.go:101   Job Schedule Started    {"id": "fixedTimer30second"}
2021-04-10T14:02:12.912+0200    INFO    sched   sched/schedule.go:101   Job Schedule Started    {"id": "onceAfter10s"}
2021-04-10T14:02:12.912+0200    INFO    sched   sched/schedule.go:176   Job Next Run Scheduled  {"id": "cronEveryMinute", "After": "47s", "At": "2021-04-10T14:03:00+02:00"}
2021-04-10T14:02:12.912+0200    INFO    sched   sched/schedule.go:176   Job Next Run Scheduled  {"id": "fixedTimer30second", "After": "30s", "At": "2021-04-10T14:02:42+02:00"}
2021-04-10T14:02:12.912+0200    INFO    sched   sched/schedule.go:176   Job Next Run Scheduled  {"id": "cronEvery5Minute", "After": "2m47s", "At": "2021-04-10T14:05:00+02:00"}
2021-04-10T14:02:12.912+0200    INFO    sched   sched/schedule.go:176   Job Next Run Scheduled  {"id": "onceAfter10s", "After": "10s", "At": "2021-04-10T14:02:22+02:00"}
2021-04-10T14:02:22.917+0200    INFO    sched   sched/schedule.go:170   No more Jobs will be scheduled  {"id": "onceAfter10s"}
2021-04-10T14:02:22.917+0200    INFO    sched   sched/schedule.go:130   Stopping Schedule...    {"id": "onceAfter10s"}
2021-04-10T14:02:22.917+0200    INFO    sched   sched/schedule.go:203   Job Run Starting        {"id": "onceAfter10s", "Instance": "9d70dff1-a120-446c-937c-64fae1c5922e"}
2021/04/10 14:02:22 onceAfter10s         Doing some work...
2021/04/10 14:02:23 onceAfter10s         Finished Work.
2021-04-10T14:02:23.921+0200    INFO    sched   sched/schedule.go:229   Job Finished    {"id": "onceAfter10s", "Instance": "9d70dff1-a120-446c-937c-64fae1c5922e", "Duration": "1.003s", "State": "FINISHED"}
2021-04-10T14:02:23.921+0200    INFO    sched.metrics   sched/metric.go:48      timer sched.run_actual_elapsed_time     {"id": "onceAfter10s", "name": "sched.run_actual_elapsed_time", "interval": "1.00318182s", "tags": {"ID":"onceAfter10s"}}
2021-04-10T14:02:23.921+0200    INFO    sched.metrics   sched/metric.go:48      timer sched.run_total_elapsed_time      {"id": "onceAfter10s", "name": "sched.run_total_elapsed_time", "interval": "1.003226808s", "tags": {"ID":"onceAfter10s"}}
2021-04-10T14:02:23.921+0200    INFO    sched   sched/schedule.go:141   Job Schedule Stopped    {"id": "onceAfter10s"}
2021-04-10T14:02:23.921+0200    INFO    sched   sched/schedule.go:161   Job Schedule Finished   {"id": "onceAfter10s"}
2021-04-10T14:02:42.916+0200    INFO    sched   sched/schedule.go:176   Job Next Run Scheduled  {"id": "fixedTimer30second", "After": "30s", "At": "2021-04-10T14:03:12+02:00"}
2021-04-10T14:02:42.916+0200    INFO    sched   sched/schedule.go:203   Job Run Starting        {"id": "fixedTimer30second", "Instance": "2a02ec70-d141-48bd-885f-14bc89329ea6"}
2021/04/10 14:02:42 fixedEvery30Second   Doing some work...
2021-04-10T14:02:43.419+0200    WARN    sched   sched/schedule.go:211   Job Run Exceeded Expected Time  {"id": "fixedTimer30second", "Instance": "2a02ec70-d141-48bd-885f-14bc89329ea6", "Expected": "1s"}
github.com/sherifabdlnaby/sched.(*Schedule).runJobInstance.func1
/Users/sherifabdlnaby/code/projects/sched/schedule.go:211
2021/04/10 14:02:43 fixedEvery30Second   Finished Work.
2021-04-10T14:02:43.921+0200    INFO    sched   sched/schedule.go:229   Job Finished    {"id": "fixedTimer30second", "Instance": "2a02ec70-d141-48bd-885f-14bc89329ea6", "Duration": "1.005s", "State": "FINISHED"}
2021-04-10T14:02:43.921+0200    INFO    sched.metrics   sched/metric.go:48      timer sched.run_actual_elapsed_time     {"id": "fixedTimer30second", "name": "sched.run_actual_elapsed_time", "interval": "1.004632668s", "tags": {"ID":"fixedTimer30second"}}
2021-04-10T14:02:43.921+0200    INFO    sched.metrics   sched/metric.go:48      timer sched.run_total_elapsed_time      {"id": "fixedTimer30second", "name": "sched.run_total_elapsed_time", "interval": "1.004697199s", "tags": {"ID":"fixedTimer30second"}}

```

