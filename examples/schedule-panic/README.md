# Output with a job that panics and is handled by schdeuler

## Output

```json
2021-04-10T13:06: 48.005+0200    INFO    sched   sched/schedule.go: 96    Job Schedule Started    {"id": "every10s"}
2021-04-10T13:06: 48.005+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "every10s", "After": "10s", "At": "2021-04-10T13:06:58+02:00"
}
2021-04-10T13: 06: 58.007+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "every10s", "After": "10s", "At": "2021-04-10T13:07:08+02:00"}
2021-04-10T13: 06:58.007+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "every10s", "Instance": "0f972951-37d0-44cf-bf9d-fa4a3e2ce0c4"}
2021/04/10 13: 06: 58 Doing some work...
2021-04-10T13: 06: 59.010+0200    ERROR   sched   sched/schedule.go:203   Job Error       {"id": "every10s", "Instance": "0f972951-37d0-44cf-bf9d-fa4a3e2ce0c4", "Duration": "1.002s", "State": "PANICKED", "error": "job panicked: Oops, I panicked, we all do, sorry."}
github.com/sherifabdlnaby/sched.(*Schedule).runJobInstance
/Users/sherifabdlnaby/code/projects/sched/schedule.go: 203
2021-04-10T13:06: 59.010+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "every10s", "Instance": "0f972951-37d0-44cf-bf9d-fa4a3e2ce0c4", "Duration": "1.002s", "State": "PANICKED"
}
2021-04-10T13: 07: 08.007+0200    INFO    sched   sched/schedule.go:168   Job Next Run Scheduled  {"id": "every10s", "After": "10s", "At": "2021-04-10T13:07:18+02:00"}
2021-04-10T13:07: 08.007+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "every10s", "Instance": "122536fa-074a-42c8-912c-7eb1dd21e0da"}
2021/04/10 13: 07: 08 Doing some work...
2021-04-10T13: 07: 09.009+0200    ERROR   sched   sched/schedule.go: 203   Job Error       {"id": "every10s", "Instance": "122536fa-074a-42c8-912c-7eb1dd21e0da", "Duration": "1.002s", "State": "PANICKED", "error": "job panicked: Oops, I panicked, we all do, sorry."}
github.com/sherifabdlnaby/sched.(*Schedule).runJobInstance
/Users/sherifabdlnaby/code/projects/sched/schedule.go: 203
2021-04-10T13: 07: 09.009+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "every10s", "Instance": "122536fa-074a-42c8-912c-7eb1dd21e0da", "Duration": "1.002s", "State": "PANICKED"
}
...
...
...

```
