# Output with a cron expression that run every minute

## Output for 3 minutes

```json
2021-04-10T12:30: 13.132+0200    INFO    sched   sched/schedule.go: 96    Job Schedule Started    {"id": "cron"}
2021-04-10T12:30: 13.132+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "cron", "After": "47s", "At": "2021-04-10T12:31:00+02:00"
}
2021-04-10T12: 31: 00.000+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "cron", "After": "1m0s", "At": "2021-04-10T12:32:00+02:00"}
2021-04-10T12: 31:00.000+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "cron", "Instance": "8e1044ab-20b6-4acf-8a15-e06c0418522c"}
2021/04/10 12: 31: 00 Doing some work...
2021/04/10 12: 31: 01 Finished Work.
2021-04-10T12: 31: 01.001+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "cron", "Instance": "8e1044ab-20b6-4acf-8a15-e06c0418522c", "Duration": "1.001s", "State": "FINISHED"}
2021-04-10T12:32: 00.002+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "cron", "After": "1m0s", "At": "2021-04-10T12:33:00+02:00"
}
2021-04-10T12: 32: 00.002+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "cron", "Instance": "baae94eb-f818-4b34-a1f4-45b521a360a1"}
2021/04/10 12: 32: 00 Doing some work...
2021/04/10 12: 32: 01 Finished Work.
2021-04-10T12:32: 01.005+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "cron", "Instance": "baae94eb-f818-4b34-a1f4-45b521a360a1", "Duration": "1.003s", "State": "FINISHED"
}
2021-04-10T12: 33: 00.001+0200    INFO    sched   sched/schedule.go:168   Job Next Run Scheduled  {"id": "cron", "After": "1m0s", "At": "2021-04-10T12:34:00+02:00"}
2021-04-10T12:33: 00.001+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "cron", "Instance": "71c8f0bf-3624-4a92-909c-b4149f3c62a3"}
2021/04/10 12: 33: 00 Doing some work...
2021/04/10 12: 33: 01 Finished Work.
2021-04-10T12: 33: 01.004+0200    INFO    sched   sched/schedule.go:208   Job Finished    {"id": "cron", "Instance": "71c8f0bf-3624-4a92-909c-b4149f3c62a3", "Duration": "1.003s", "State": "FINISHED"}


```

## Output With CTRL+C

```json
2021-04-10T12:28: 45.591+0200    INFO    sched   sched/schedule.go: 96    Job Schedule Started    {"id": "cron"}
2021-04-10T12:28: 45.592+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "cron", "After": "14s", "At": "2021-04-10T12:29:00+02:00"
}
2021-04-10T12: 29: 00.000+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "cron", "After": "1m0s", "At": "2021-04-10T12:30:00+02:00"}
2021-04-10T12: 29:00.000+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "cron", "Instance": "786540f1-594b-44a0-9a66-7181619e38a6"}
2021/04/10 12: 29: 00 Doing some work...
CTRL+C
2021-04-10T12: 29: 00.567+0200    INFO    sched   sched/schedule.go: 125   Stopping Schedule...    {"id": "cron"}
2021-04-10T12: 29: 00.567+0200    INFO    sched   sched/schedule.go: 130   Waiting active jobs to finish...        {"id": "cron"
}
2021-04-10T12: 29: 00.567+0200    INFO    sched   sched/schedule.go: 171   Job Next Run Canceled   {"id": "cron", "At": "2021-04-10T12:30:00+02:00"
}
2021/04/10 12: 29: 01 Finished Work.
2021-04-10T12: 29:01.000+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "cron", "Instance": "786540f1-594b-44a0-9a66-7181619e38a6", "Duration": "1s", "State": "FINISHED"
}
2021-04-10T12: 29: 01.000+0200    INFO    sched   sched/schedule.go: 133   Job Schedule Stopped    {
"id": "cron"
}
```
