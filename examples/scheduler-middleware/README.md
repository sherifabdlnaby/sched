# 4 Schedules in a Scheduler Manager

1. Cron Every Minute
2. Cron Every 5 Minutes
3. Fixed Interval Every 30 Secs
4. *Once* after 10 Secs from schedule start.

Started and Stopped using `StartAll()` and `StopAll()`

## Output

```json
2021-04-10T13:26: 43.142+0200    INFO    sched   sched/schedule.go: 96    Job Schedule Started    {"id": "cronEveryMinute"}
2021-04-10T13:26: 43.142+0200    INFO    sched   sched/schedule.go: 96    Job Schedule Started    {"id": "cronEvery5Minute"}
2021-04-10T13:26: 43.142+0200    INFO    sched   sched/schedule.go: 96    Job Schedule Started    {"id": "fixedTimer30second"}
2021-04-10T13:26: 43.142+0200    INFO    sched   sched/schedule.go: 96    Job Schedule Started    {"id": "onceAfter10s"}
2021-04-10T13:26: 43.142+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "cronEveryMinute", "After": "17s", "At": "2021-04-10T13:27:00+02:00"
}
2021-04-10T13: 26: 43.143+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "onceAfter10s", "After": "10s", "At": "2021-04-10T13:26:53+02:00"}
2021-04-10T13: 26:43.142+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "cronEvery5Minute", "After": "3m17s", "At": "2021-04-10T13:30:00+02:00"
}
2021-04-10T13: 26: 43.143+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "fixedTimer30second", "After": "30s", "At": "2021-04-10T13:27:13+02:00"}
2021-04-10T13: 26: 53.143+0200    INFO    sched   sched/schedule.go: 162   No more Jobs will be scheduled  {"id": "onceAfter10s"}
2021-04-10T13: 26: 53.143+0200    INFO    sched   sched/schedule.go: 125   Stopping Schedule...    {"id": "onceAfter10s"}
2021-04-10T13: 26:53.143+0200    INFO    sched   sched/schedule.go: 130   Waiting active jobs to finish...        {"id": "onceAfter10s"}
2021-04-10T13:26: 53.143+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "onceAfter10s", "Instance": "47fe7d35-3494-43e0-8771-4282ecb80f3a"}
2021/04/10 13: 26: 53 onceAfter10s         Doing some work...
2021/04/10 13: 26: 54 onceAfter10s         Finished Work.
2021-04-10T13: 26: 54.148+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "onceAfter10s", "Instance": "47fe7d35-3494-43e0-8771-4282ecb80f3a", "Duration": "1.005s", "State": "FINISHED"}
2021-04-10T13: 26: 54.148+0200    INFO    sched.metrics   sched/metric.go: 48      timer sched.run_actual_elapsed_time     {"id": "onceAfter10s", "name": "sched.run_actual_elapsed_time", "interval": "1.004917899s", "tags": {
"ID": "onceAfter10s"
}
}
2021-04-10T13: 26: 54.148+0200    INFO    sched.metrics   sched/metric.go: 48      timer sched.run_total_elapsed_time      {"id": "onceAfter10s", "name": "sched.run_total_elapsed_time", "interval": "1.004966378s", "tags": {"ID": "onceAfter10s"}
}
2021-04-10T13: 26: 54.148+0200    INFO    sched   sched/schedule.go: 133   Job Schedule Stopped    {"id": "onceAfter10s"}
2021-04-10T13: 26: 54.148+0200    INFO    sched   sched/schedule.go: 153   Job Schedule Finished   {"id": "onceAfter10s"}
2021-04-10T13: 26: 57.662+0200    INFO    sched   sched/schedule.go: 125   Stopping Schedule...    {"id": "cronEveryMinute"}
2021-04-10T13: 26: 57.662+0200    INFO    sched   sched/schedule.go: 125   Stopping Schedule...    {"id": "fixedTimer30second"}
2021-04-10T13: 26: 57.662+0200    INFO    sched   sched/schedule.go: 130   Waiting active jobs to finish...        {"id": "cronEveryMinute"}
2021-04-10T13: 26: 57.662+0200    INFO    sched   sched/schedule.go: 125   Stopping Schedule...    {"id": "cronEvery5Minute"}
2021-04-10T13: 26: 57.662+0200    INFO    sched   sched/schedule.go: 133   Job Schedule Stopped    {"id": "cronEveryMinute"}
2021-04-10T13: 26: 57.662+0200    INFO    sched   sched/schedule.go: 130   Waiting active jobs to finish...        {"id": "cronEvery5Minute"
}
2021-04-10T13: 26: 57.662+0200    INFO    sched   sched/schedule.go: 171   Job Next Run Canceled   {"id": "fixedTimer30second", "At": "2021-04-10T13:27:13+02:00"
}
2021-04-10T13:26: 57.662+0200    INFO    sched   sched/schedule.go: 171   Job Next Run Canceled   {"id": "cronEvery5Minute", "At": "2021-04-10T13:30:00+02:00"
}
2021-04-10T13: 26: 57.662+0200    INFO    sched   sched/schedule.go: 133   Job Schedule Stopped    {"id": "cronEvery5Minute"
}
2021-04-10T13: 26: 57.662+0200    INFO    sched   sched/schedule.go: 130   Waiting active jobs to finish...        {"id": "fixedTimer30second"}
2021-04-10T13: 26: 57.662+0200    INFO    sched   sched/schedule.go: 133   Job Schedule Stopped    {"id": "fixedTimer30second"}
2021-04-10T13: 26: 57.662+0200    INFO    sched   sched/schedule.go: 171   Job Next Run Canceled   {"id": "cronEveryMinute", "At": "2021-04-10T13:27:00+02:00"}

```

