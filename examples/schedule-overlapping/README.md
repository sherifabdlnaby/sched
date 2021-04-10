# Output with a job that run every 30s on a fixed schedule

## Output

```json
2021-04-10T13:35: 42.078+0200    INFO    sched   sched/schedule.go: 97    Job Schedule Started    {"id": "every1Sec"}
2021/04/10 13: 35: 42 Stopping schedule Automatically after 5 seconds
2021-04-10T13: 35: 42.078+0200    INFO    sched   sched/schedule.go: 172   Job Next Run Scheduled  {"id": "every1Sec", "After": "1s", "At": "2021-04-10T13:35:43+02:00"
}
2021-04-10T13: 35: 43.081+0200    INFO    sched   sched/schedule.go:172   Job Next Run Scheduled  {"id": "every1Sec", "After": "1s", "At": "2021-04-10T13:35:44+02:00"}
2021-04-10T13:35: 43.081+0200    INFO    sched   sched/schedule.go: 197   Job Run Starting        {"id": "every1Sec", "Instance": "12e1c9f3-69a7-4a31-8082-7dd40416f00b"}
2021/04/10 13: 35: 43 Doing some work for 5 seconds...
2021-04-10T13: 35: 44.082+0200    INFO    sched   sched/schedule.go: 172   Job Next Run Scheduled  {"id": "every1Sec", "After": "1s", "At": "2021-04-10T13:35:45+02:00"
}
2021-04-10T13: 35: 44.082+0200    INFO    sched   sched/schedule.go: 197   Job Run Starting        {"id": "every1Sec", "Instance": "46ccdc9c-1792-4fc6-8e7b-aa269ea8b5a0"
}
2021/04/10 13:35: 44 Doing some work for 5 seconds...
2021-04-10T13: 35: 45.078+0200    INFO    sched   sched/schedule.go:172   Job Next Run Scheduled  {"id": "every1Sec", "After": "1s", "At": "2021-04-10T13:35:46+02:00"}
2021-04-10T13:35: 45.078+0200    INFO    sched   sched/schedule.go: 197   Job Run Starting        {"id": "every1Sec", "Instance": "4c663f7a-40a8-47a1-8eaa-4ed52cc80a94"}
2021/04/10 13: 35: 45 Doing some work for 5 seconds...
2021-04-10T13: 35: 46.079+0200    INFO    sched   sched/schedule.go: 172   Job Next Run Scheduled  {"id": "every1Sec", "After": "1s", "At": "2021-04-10T13:35:47+02:00"
}
2021-04-10T13: 35: 46.079+0200    INFO    sched   sched/schedule.go: 197   Job Run Starting        {"id": "every1Sec", "Instance": "22a192be-2d42-4840-bbfe-05e28bee1716"
}
2021/04/10 13:35: 46 Doing some work for 5 seconds...
2021-04-10T13: 35: 47.080+0200    INFO    sched   sched/schedule.go:172   Job Next Run Scheduled  {"id": "every1Sec", "After": "1s", "At": "2021-04-10T13:35:48+02:00"}
2021-04-10T13:35: 47.080+0200    INFO    sched   sched/schedule.go: 126   Stopping Schedule...    {"id": "every1Sec"}
2021-04-10T13: 35: 47.080+0200    INFO    sched   sched/schedule.go: 132   Waiting for '5' active jobs still running...    {"id": "every1Sec"}
2021-04-10T13: 35: 47.080+0200    INFO    sched   sched/schedule.go: 197   Job Run Starting        {"id": "every1Sec", "Instance": "08cb1143-c0cc-436c-a7e6-097b0ef4cbf0"
}
2021-04-10T13: 35: 47.080+0200    INFO    sched   sched/schedule.go: 175   Job Next Run Canceled   {"id": "every1Sec", "At": "2021-04-10T13:35:48+02:00"}
2021/04/10 13: 35: 47 Doing some work for 5 seconds...
2021/04/10 13: 35: 48 Finished Work.
2021-04-10T13: 35: 48.082+0200    INFO    sched   sched/schedule.go: 212   Job Finished    {"id": "every1Sec", "Instance": "12e1c9f3-69a7-4a31-8082-7dd40416f00b", "Duration": "5.001s", "State": "FINISHED"}
2021/04/10 13: 35: 49 Finished Work.
2021-04-10T13: 35: 49.086+0200    INFO    sched   sched/schedule.go:212   Job Finished    {"id": "every1Sec", "Instance": "46ccdc9c-1792-4fc6-8e7b-aa269ea8b5a0", "Duration": "5.004s", "State": "FINISHED"}
2021/04/10 13: 35: 50 Finished Work.
2021-04-10T13:35: 50.081+0200    INFO    sched   sched/schedule.go: 212   Job Finished    {"id": "every1Sec", "Instance": "4c663f7a-40a8-47a1-8eaa-4ed52cc80a94", "Duration": "5.003s", "State": "FINISHED"
}
2021/04/10 13: 35: 51 Finished Work.
2021-04-10T13: 35: 51.084+0200    INFO    sched   sched/schedule.go: 212   Job Finished    {"id": "every1Sec", "Instance": "22a192be-2d42-4840-bbfe-05e28bee1716", "Duration": "5.005s", "State": "FINISHED"}
2021/04/10 13: 35: 52 Finished Work.
2021-04-10T13: 35:52.080+0200    INFO    sched   sched/schedule.go: 212   Job Finished    {"id": "every1Sec", "Instance": "08cb1143-c0cc-436c-a7e6-097b0ef4cbf0", "Duration": "5s", "State": "FINISHED"
}
2021-04-10T13: 35: 52.080+0200    INFO    sched   sched/schedule.go: 137   Job Schedule Stopped    {
"id": "every1Sec"
}

```
