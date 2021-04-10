# Output with a job that run every 30s on a fixed schedule

## Output

```json
2021-04-10T12:58: 48.724+0200    INFO    sched   sched/schedule.go: 96    Job Schedule Started    {"id": "every30s"}
2021-04-10T12:58: 48.725+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "every30s", "After": "30s", "At": "2021-04-10T12:59:18+02:00"}
2021-04-10T12: 59: 18.729+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "every30s", "After": "30s", "At": "2021-04-10T12:59:48+02:00"}
2021-04-10T12: 59:18.729+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "every30s", "Instance": "e05aa702-f11c-46ba-8d7c-6ae4049c382d"}
2021/04/10 12: 59: 18 Doing some work...
2021/04/10 12: 59: 19 Finished Work.
2021-04-10T12: 59: 19.733+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "every30s", "Instance": "e05aa702-f11c-46ba-8d7c-6ae4049c382d", "Duration": "1.004s", "State": "FINISHED"}
2021-04-10T12:59: 48.724+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "every30s", "After": "30s", "At": "2021-04-10T13:00:18+02:00"}
2021-04-10T12: 59: 48.724+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "every30s", "Instance": "80ed574e-e4fc-4a9f-9741-6eda96ebd470"}
2021/04/10 12: 59: 48 Doing some work...
2021/04/10 12: 59: 49 Finished Work.
2021-04-10T12:59: 49.727+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "every30s", "Instance": "80ed574e-e4fc-4a9f-9741-6eda96ebd470", "Duration": "1.003s", "State": "FINISHED"
}
2021-04-10T13: 00: 18.726+0200    INFO    sched   sched/schedule.go:168   Job Next Run Scheduled  {"id": "every30s", "After": "30s", "At": "2021-04-10T13:00:48+02:00"}
2021-04-10T13:00: 18.726+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "every30s", "Instance": "6bc1402c-a5e4-4634-9e49-acaaa6dcb5d0"}
2021/04/10 13: 00: 18 Doing some work...
2021/04/10 13: 00: 19 Finished Work.
2021-04-10T13: 00: 19.729+0200    INFO    sched   sched/schedule.go:208   Job Finished    {"id": "every30s", "Instance": "6bc1402c-a5e4-4634-9e49-acaaa6dcb5d0", "Duration": "1.003s", "State": "FINISHED"}
2021-04-10T13: 00: 48.725+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "every30s", "After": "30s", "At": "2021-04-10T13:01:18+02:00"}
2021-04-10T13: 00: 48.725+0200    INFO    sched   sched/schedule.go:193   Job Run Starting        {"id": "every30s", "Instance": "ed82fdd5-63a0-4fb4-b6a3-7fb74a973d1a"}
2021/04/10 13: 00: 48 Doing some work...
2021/04/10 13: 00: 49 Finished Work.
2021-04-10T13: 00: 49.730+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "every30s", "Instance": "ed82fdd5-63a0-4fb4-b6a3-7fb74a973d1a", "Duration": "1.005s", "State": "FINISHED"}
2021-04-10T13: 01: 18.724+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "every30s", "After": "30s", "At": "2021-04-10T13:01:48+02:00"}
2021-04-10T13: 01: 18.724+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "every30s", "Instance": "a48eb5fa-2f62-4e9d-ae4e-550fbedc0cd6"}
2021/04/10 13: 01: 18 Doing some work...
2021/04/10 13: 01: 19 Finished Work.
2021-04-10T13: 01: 19.728+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "every30s", "Instance": "a48eb5fa-2f62-4e9d-ae4e-550fbedc0cd6", "Duration": "1.003s", "State": "FINISHED"}
2021-04-10T13: 01: 48.725+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "every30s", "After": "30s", "At": "2021-04-10T13:02:18+02:00"}
2021-04-10T13: 01: 48.725+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "every30s", "Instance": "2347a38f-82d8-45aa-abb9-1f5192f53a09"}
2021/04/10 13: 01:48 Doing some work...
2021/04/10 13: 01: 49 Finished Work.
2021-04-10T13: 01: 49.725+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "every30s", "Instance": "2347a38f-82d8-45aa-abb9-1f5192f53a09", "Duration": "1s", "State": "FINISHED"
}
2021-04-10T13: 02: 18.721+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "every30s", "After": "30s", "At": "2021-04-10T13:02:48+02:00"}
2021-04-10T13: 02: 18.721+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "every30s", "Instance": "b4f52aa8-999a-4d46-8cd0-f49c91b22ca3"}
2021/04/10 13: 02: 18 Doing some work...
2021/04/10 13: 02:19 Finished Work.
2021-04-10T13: 02: 19.722+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "every30s", "Instance": "b4f52aa8-999a-4d46-8cd0-f49c91b22ca3", "Duration": "1.001s", "State": "FINISHED"}
2021-04-10T13: 02: 48.721+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "every30s", "Instance": "e2c32b53-556c-48ca-98b0-9d39aa0fbd61"}
2021-04-10T13:02: 48.721+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "every30s", "After": "30s", "At": "2021-04-10T13:03:18+02:00"}
2021/04/10 13:02: 48 Doing some work...
2021/04/10 13: 02: 49 Finished Work.
2021-04-10T13: 02: 49.722+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "every30s", "Instance": "e2c32b53-556c-48ca-98b0-9d39aa0fbd61", "Duration": "1s", "State": "FINISHED"}
2021-04-10T13: 03: 18.719+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "every30s", "After": "30s", "At": "2021-04-10T13:03:48+02:00"}
2021-04-10T13: 03: 18.719+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "every30s", "Instance": "6a8aa794-3c9c-4ffe-a2a6-c4b78f31dd57"}
2021/04/10 13: 03: 18 Doing some work...
2021/04/10 13:03: 19 Finished Work.
...
...
...

```
