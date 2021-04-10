# 4 Schedules

1. Cron Every Minute
2. Cron Every 5 Minutes
3. Fixed Interval Every 30 Secs
4. *Once* after 10 Secs from schedule start.

## Output

```json
2021-04-10T12:46: 09.307+0200    INFO    sched   sched/schedule.go: 96    Job Schedule Started    {"id": "cron-every-minute"}
2021-04-10T12:46: 09.307+0200    INFO    sched   sched/schedule.go: 96    Job Schedule Started    {"id": "cron-every-5minute"}
2021-04-10T12:46: 09.307+0200    INFO    sched   sched/schedule.go: 96    Job Schedule Started    {"id": "fixed-every-30seconds"}
2021-04-10T12:46: 09.307+0200    INFO    sched   sched/schedule.go: 96    Job Schedule Started    {"id": "once-after-10seconds"}
2021-04-10T12:46: 09.307+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "once-after-10seconds", "After": "10s", "At": "2021-04-10T12:46:19+02:00"
}
2021-04-10T12: 46: 09.307+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "cron-every-minute", "After": "51s", "At": "2021-04-10T12:47:00+02:00"}
2021-04-10T12: 46:09.307+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "fixed-every-30seconds", "After": "30s", "At": "2021-04-10T12:46:39+02:00"
}
2021-04-10T12: 46: 09.307+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "cron-every-5minute", "After": "3m51s", "At": "2021-04-10T12:50:00+02:00"}
2021-04-10T12: 46: 19.311+0200    INFO    sched   sched/schedule.go: 162   No more Jobs will be scheduled  {"id": "once-after-10seconds"}
2021-04-10T12: 46: 19.311+0200    INFO    sched   sched/schedule.go: 125   Stopping Schedule...    {"id": "once-after-10seconds"}
2021-04-10T12: 46:19.311+0200    INFO    sched   sched/schedule.go: 130   Waiting active jobs to finish...        {"id": "once-after-10seconds"}
2021-04-10T12:46: 19.312+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "once-after-10seconds", "Instance": "6f971652-8394-47d6-86ff-3173cde2b6a1"}
2021/04/10 12: 46: 19 job once after 10 seconds    Doing some work...
2021/04/10 12: 46: 20 job once after 10 seconds    Finished Work.
2021-04-10T12: 46: 20.314+0200    INFO    sched   sched/schedule.go:208   Job Finished    {"id": "once-after-10seconds", "Instance": "6f971652-8394-47d6-86ff-3173cde2b6a1", "Duration": "1.002s", "State": "FINISHED"}
2021-04-10T12: 46: 20.314+0200    INFO    sched   sched/schedule.go: 133   Job Schedule Stopped    {"id": "once-after-10seconds"}
2021-04-10T12: 46: 20.314+0200    INFO    sched   sched/schedule.go: 153   Job Schedule Finished   {"id": "once-after-10seconds"}
2021-04-10T12: 46: 39.309+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "fixed-every-30seconds", "After": "30s", "At": "2021-04-10T12:47:09+02:00"
}
2021-04-10T12: 46: 39.309+0200    INFO    sched   sched/schedule.go:193   Job Run Starting        {"id": "fixed-every-30seconds", "Instance": "bf92e3f5-ea61-4044-b6a4-12d1ca654eec"
}
2021/04/10 12: 46: 39 job every 30 seconds         Doing some work...
2021/04/10 12: 46: 40 job every 30 seconds         Finished Work.
2021-04-10T12:46: 40.311+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "fixed-every-30seconds", "Instance": "bf92e3f5-ea61-4044-b6a4-12d1ca654eec", "Duration": "1.001s", "State": "FINISHED"
}
2021-04-10T12: 46: 59.999+0200    INFO    sched   sched/schedule.go:168   Job Next Run Scheduled  {"id": "cron-every-minute", "After": "0s", "At": "2021-04-10T12:47:00+02:00"}
2021-04-10T12:46: 59.999+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "cron-every-minute", "Instance": "fc39385f-5c52-4402-aec0-3ee0ec08f3d2"}
2021/04/10 12: 46: 59 job cron every minute        Doing some work...
2021-04-10T12: 47: 00.000+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "cron-every-minute", "After": "1m0s", "At": "2021-04-10T12:48:00+02:00"}
2021-04-10T12: 47: 00.000+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "cron-every-minute", "Instance": "b4911daa-5412-4816-9636-039af5f897ff"
}
2021/04/10 12: 47: 00 job cron every minute        Doing some work...
2021/04/10 12: 47: 01 job cron every minute        Finished Work.
2021/04/10 12: 47: 01 job cron every minute        Finished Work.
2021-04-10T12: 47: 01.002+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "cron-every-minute", "Instance": "b4911daa-5412-4816-9636-039af5f897ff", "Duration": "1.002s", "State": "FINISHED"}
2021-04-10T12: 47:01.002+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "cron-every-minute", "Instance": "fc39385f-5c52-4402-aec0-3ee0ec08f3d2", "Duration": "1.003s", "State": "FINISHED"
}
2021-04-10T12: 47: 09.309+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "fixed-every-30seconds", "After": "30s", "At": "2021-04-10T12:47:39+02:00"}
2021-04-10T12: 47:09.309+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "fixed-every-30seconds", "Instance": "bd50d728-807e-4d0e-aed3-45044e4b925a"}
2021/04/10 12: 47: 09 job every 30 seconds         Doing some work...
2021/04/10 12: 47: 10 job every 30 seconds         Finished Work.
2021-04-10T12: 47: 10.309+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "fixed-every-30seconds", "Instance": "bd50d728-807e-4d0e-aed3-45044e4b925a", "Duration": "1s", "State": "FINISHED"}
2021-04-10T12: 47:39.306+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "fixed-every-30seconds", "After": "30s", "At": "2021-04-10T12:48:09+02:00"
}
2021-04-10T12: 47: 39.306+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "fixed-every-30seconds", "Instance": "282dea9f-b64a-48ab-9ce8-3dcf533d95e4"}
2021/04/10 12: 47: 39 job every 30 seconds         Doing some work...
2021/04/10 12: 47: 40 job every 30 seconds         Finished Work.
2021-04-10T12: 47: 40.307+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "fixed-every-30seconds", "Instance": "282dea9f-b64a-48ab-9ce8-3dcf533d95e4", "Duration": "1.001s", "State": "FINISHED"
}
2021-04-10T12: 48: 00.002+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "cron-every-minute", "After": "1m0s", "At": "2021-04-10T12:49:00+02:00"}
2021-04-10T12: 48: 00.002+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "cron-every-minute", "Instance": "9311f3ef-3a43-4fe1-8c8d-740616f42294"
}
2021/04/10 12: 48: 00 job cron every minute        Doing some work...
2021/04/10 12: 48: 01 job cron every minute        Finished Work.
2021-04-10T12: 48: 01.007+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "cron-every-minute", "Instance": "9311f3ef-3a43-4fe1-8c8d-740616f42294", "Duration": "1.004s", "State": "FINISHED"}
2021-04-10T12: 48: 09.304+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "fixed-every-30seconds", "After": "30s", "At": "2021-04-10T12:48:39+02:00"
}
2021-04-10T12: 48: 09.304+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "fixed-every-30seconds", "Instance": "083d7198-1981-4542-9677-f837509a346c"
}
2021/04/10 12:48: 09 job every 30 seconds         Doing some work...
2021/04/10 12: 48: 10 job every 30 seconds         Finished Work.
2021-04-10T12: 48: 10.308+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "fixed-every-30seconds", "Instance": "083d7198-1981-4542-9677-f837509a346c", "Duration": "1.004s", "State": "FINISHED"
}
2021-04-10T12: 48: 39.333+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "fixed-every-30seconds", "After": "30s", "At": "2021-04-10T12:49:09+02:00"}
2021-04-10T12: 48: 39.333+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "fixed-every-30seconds", "Instance": "8b8cffcb-b026-4383-813d-3f704d6d37a6"}
2021/04/10 12: 48: 39 job every 30 seconds         Doing some work...
2021/04/10 12: 48: 40 job every 30 seconds         Finished Work.
2021-04-10T12: 48: 40.338+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "fixed-every-30seconds", "Instance": "8b8cffcb-b026-4383-813d-3f704d6d37a6", "Duration": "1.004s", "State": "FINISHED"}
2021-04-10T12: 49: 00.035+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "cron-every-minute", "After": "1m0s", "At": "2021-04-10T12:50:00+02:00"}
2021-04-10T12: 49: 00.035+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "cron-every-minute", "Instance": "a4c6a479-2138-4037-b613-8adcae1e882d"
}
2021/04/10 12: 49:00 job cron every minute        Doing some work...
2021/04/10 12: 49: 01 job cron every minute        Finished Work.
2021-04-10T12: 49: 01.039+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "cron-every-minute", "Instance": "a4c6a479-2138-4037-b613-8adcae1e882d", "Duration": "1.004s", "State": "FINISHED"}
2021-04-10T12: 49: 09.344+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "fixed-every-30seconds", "After": "30s", "At": "2021-04-10T12:49:39+02:00"}
2021-04-10T12: 49: 09.344+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "fixed-every-30seconds", "Instance": "6f17958b-2e83-4d81-808d-9432900cf24e"}
2021/04/10 12: 49: 09 job every 30 seconds         Doing some work...
2021/04/10 12:49: 10 job every 30 seconds         Finished Work.
2021-04-10T12: 49: 10.349+0200    INFO    sched   sched/schedule.go:208   Job Finished    {"id": "fixed-every-30seconds", "Instance": "6f17958b-2e83-4d81-808d-9432900cf24e", "Duration": "1.004s", "State": "FINISHED"}
2021-04-10T12: 49: 39.346+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "fixed-every-30seconds", "After": "30s", "At": "2021-04-10T12:50:09+02:00"
}
2021-04-10T12: 49: 39.346+0200    INFO    sched   sched/schedule.go:193   Job Run Starting        {"id": "fixed-every-30seconds", "Instance": "11da6889-dee3-4e1c-8c52-af2e2618999a"
}
2021/04/10 12: 49: 39 job every 30 seconds         Doing some work...
2021/04/10 12: 49: 40 job every 30 seconds         Finished Work.
2021-04-10T12:49: 40.348+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "fixed-every-30seconds", "Instance": "11da6889-dee3-4e1c-8c52-af2e2618999a", "Duration": "1.003s", "State": "FINISHED"
}
2021-04-10T12: 50: 00.003+0200    INFO    sched   sched/schedule.go:168   Job Next Run Scheduled  {"id": "cron-every-minute", "After": "1m0s", "At": "2021-04-10T12:51:00+02:00"}
2021-04-10T12:50: 00.003+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "cron-every-minute", "Instance": "35df9a8c-7900-4f79-943f-a0f0efe86032"}
2021/04/10 12: 50: 00 job cron every minute        Doing some work...
2021-04-10T12: 50: 00.035+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "cron-every-5minute", "After": "5m0s", "At": "2021-04-10T12:55:00+02:00"}
2021-04-10T12: 50: 00.035+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "cron-every-5minute", "Instance": "0ba1a33f-e75f-4273-ace7-98a6c51113ff"
}
2021/04/10 12: 50: 00 job cron every 5 minute      Doing some work...
2021/04/10 12: 50: 01 job cron every minute        Finished Work.
2021-04-10T12: 50: 01.003+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "cron-every-minute", "Instance": "35df9a8c-7900-4f79-943f-a0f0efe86032", "Duration": "1.001s", "State": "FINISHED"
}
2021/04/10 12: 50:01 job cron every 5 minute      Finished Work.
2021-04-10T12:50: 01.037+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "cron-every-5minute", "Instance": "0ba1a33f-e75f-4273-ace7-98a6c51113ff", "Duration": "1.002s", "State": "FINISHED"
}
CTRL+C
2021-04-10T12: 50: 08.092+0200    INFO    sched   sched/schedule.go: 125   Stopping Schedule...    {"id": "cron-every-minute"}
2021-04-10T12: 50: 08.092+0200    INFO    sched   sched/schedule.go: 130   Waiting active jobs to finish...        {"id": "cron-every-minute"
}
2021-04-10T12: 50: 08.092+0200    INFO    sched   sched/schedule.go: 133   Job Schedule Stopped    {"id": "cron-every-minute"
}
2021-04-10T12: 50: 08.092+0200    INFO    sched   sched/schedule.go: 125   Stopping Schedule...    {"id": "cron-every-5minute"}
2021-04-10T12: 50: 08.092+0200    INFO    sched   sched/schedule.go: 130   Waiting active jobs to finish...        {"id": "cron-every-5minute"
}
2021-04-10T12: 50: 08.092+0200    INFO    sched   sched/schedule.go: 171   Job Next Run Canceled   {"id": "cron-every-minute", "At": "2021-04-10T12:51:00+02:00"}
2021-04-10T12: 50: 08.092+0200    INFO    sched   sched/schedule.go:133   Job Schedule Stopped    {
"id": "cron-every-5minute"
}
2021-04-10T12: 50: 08.092+0200    INFO    sched   sched/schedule.go:171   Job Next Run Canceled   {"id": "cron-every-5minute", "At": "2021-04-10T12:55:00+02:00"}
2021-04-10T12: 50: 08.092+0200    INFO    sched   sched/schedule.go: 125   Stopping Schedule...    {"id": "fixed-every-30seconds"
}
2021-04-10T12: 50: 08.093+0200    INFO    sched   sched/schedule.go: 130   Waiting active jobs to finish...        {"id": "fixed-every-30seconds"}
2021-04-10T12: 50: 08.093+0200    INFO    sched   sched/schedule.go: 133   Job Schedule Stopped    {"id": "fixed-every-30seconds"}
2021-04-10T12: 50: 08.093+0200    INFO    sched   sched/schedule.go: 171   Job Next Run Canceled   {"id": "fixed-every-30seconds", "At": "2021-04-10T12:50:09+02:00"}

```

