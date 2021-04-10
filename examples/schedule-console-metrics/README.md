- Output with a job that run every 5s on a fixed schedule and run for a random amount of time no more than 5 S
- Some metrics are printed to console every 5s

## Output

```json
2021-04-10T13:13: 35.554+0200    INFO    sched   sched/schedule.go: 96    Job Schedule Started    {"id": "every5s"}
2021-04-10T13:13: 35.554+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "every5s", "After": "5s", "At": "2021-04-10T13:13:40+02:00"
}
2021-04-10T13: 13: 40.556+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "every5s", "After": "5s", "At": "2021-04-10T13:13:45+02:00"}
2021-04-10T13: 13:40.556+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "every5s", "Instance": "bd21a8a3-a9e7-4d8f-83cd-6412913198d8"}
2021/04/10 13: 13: 40 Doing some work for random time...
2021/04/10 13: 13: 41 Finished Work.
2021-04-10T13: 13: 41.657+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "every5s", "Instance": "bd21a8a3-a9e7-4d8f-83cd-6412913198d8", "Duration": "1.101s", "State": "FINISHED"}
2021-04-10T13: 13: 41.657+0200    INFO    sched.metrics   sched/metric.go: 48      timer sched.run_actual_elapsed_time     {"id": "every5s", "name": "sched.run_actual_elapsed_time", "interval": "1.101133682s", "tags": {"ID": "every5s"}}
2021-04-10T13: 13:41.657+0200    INFO    sched.metrics   sched/metric.go: 48      timer sched.run_total_elapsed_time      {"id": "every5s", "name": "sched.run_total_elapsed_time", "interval": "1.101159753s", "tags": {"ID": "every5s"}}
2021-04-10T13: 13: 45.559+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "every5s", "After": "5s", "At": "2021-04-10T13:13:50+02:00"}
2021-04-10T13: 13: 45.559+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "every5s", "Instance": "484b9a12-610c-43c4-9709-23356d7a434d"}
2021/04/10 13: 13: 45 Doing some work for random time...
2021/04/10 13: 13: 45 Finished Work.
2021-04-10T13: 13: 45.763+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "every5s", "Instance": "484b9a12-610c-43c4-9709-23356d7a434d", "Duration": "204ms", "State": "FINISHED"
}
2021-04-10T13: 13: 45.763+0200    INFO    sched.metrics   sched/metric.go: 48      timer sched.run_actual_elapsed_time     {"id": "every5s", "name": "sched.run_actual_elapsed_time", "interval": "203.728891ms", "tags": {"ID": "every5s"}}
2021-04-10T13: 13: 45.763+0200    INFO    sched.metrics   sched/metric.go: 48      timer sched.run_total_elapsed_time      {"id": "every5s", "name": "sched.run_total_elapsed_time", "interval": "203.793517ms", "tags": {"ID": "every5s"
}
}
2021-04-10T13: 13: 50.557+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "every5s", "After": "5s", "At": "2021-04-10T13:13:55+02:00"}
2021-04-10T13: 13: 50.557+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "every5s", "Instance": "905996a9-7fe8-46b3-8e32-938ca18f2d03"
}
2021/04/10 13: 13: 50 Doing some work for random time...
2021/04/10 13: 13: 52 Finished Work.
2021-04-10T13:13: 52.757+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "every5s", "Instance": "905996a9-7fe8-46b3-8e32-938ca18f2d03", "Duration": "2.2s", "State": "FINISHED"
}
2021-04-10T13: 13: 52.757+0200    INFO    sched.metrics   sched/metric.go:48      timer sched.run_actual_elapsed_time     {"id": "every5s", "name": "sched.run_actual_elapsed_time", "interval": "2.200216259s", "tags": {"ID": "every5s"
}
}
2021-04-10T13: 13: 52.757+0200    INFO    sched.metrics   sched/metric.go: 48      timer sched.run_total_elapsed_time      {"id": "every5s", "name": "sched.run_total_elapsed_time", "interval": "2.200301621s", "tags": {
"ID": "every5s"
}
}
2021-04-10T13: 13: 55.558+0200    INFO    sched   sched/schedule.go:168   Job Next Run Scheduled  {"id": "every5s", "After": "5s", "At": "2021-04-10T13:14:00+02:00"}
2021-04-10T13:13: 55.558+0200    INFO    sched.metrics   sched/metric.go: 40      counter sched.runs      {"id": "every5s", "name": "sched.runs", "value": 3, "tags": {"ID":"every5s"
}
}
2021-04-10T13: 13: 55.558+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "every5s", "Instance": "39c3295f-63be-4f4a-b81e-79753bbcc91a"
}
2021/04/10 13:13: 55 Doing some work for random time...
2021-04-10T13: 13: 55.558+0200    INFO    sched.metrics   sched/metric.go:44      gauge sched.up  {"id": "every5s", "name": "sched.up", "value": 1, "tags": {"ID": "every5s"
}
}
2021/04/10 13: 13: 55 Finished Work.
2021-04-10T13: 13: 55.762+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "every5s", "Instance": "39c3295f-63be-4f4a-b81e-79753bbcc91a", "Duration": "204ms", "State": "FINISHED"
}
2021-04-10T13: 13: 55.762+0200    INFO    sched.metrics   sched/metric.go: 48      timer sched.run_actual_elapsed_time     {"id": "every5s", "name": "sched.run_actual_elapsed_time", "interval": "203.88384ms", "tags": {"ID": "every5s"}}
2021-04-10T13:13: 55.762+0200    INFO    sched.metrics   sched/metric.go: 48      timer sched.run_total_elapsed_time      {"id": "every5s", "name": "sched.run_total_elapsed_time", "interval": "203.935885ms", "tags": {"ID":"every5s"
}
}
2021-04-10T13: 14: 00.555+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "every5s", "After": "5s", "At": "2021-04-10T13:14:05+02:00"}
2021-04-10T13: 14: 00.556+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "every5s", "Instance": "1cad0ba1-ac35-4c9d-8ffb-794cd2902c66"}
2021/04/10 13: 14: 00 Doing some work for random time...
2021/04/10 13: 14: 04 Finished Work.
2021-04-10T13: 14:04.359+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "every5s", "Instance": "1cad0ba1-ac35-4c9d-8ffb-794cd2902c66", "Duration": "3.804s", "State": "FINISHED"
}
2021-04-10T13: 14: 04.359+0200    INFO    sched.metrics   sched/metric.go: 48      timer sched.run_actual_elapsed_time     {"id": "every5s", "name": "sched.run_actual_elapsed_time", "interval": "3.803781236s", "tags": {"ID": "every5s"}
}
2021-04-10T13: 14: 04.359+0200    INFO    sched.metrics   sched/metric.go: 48      timer sched.run_total_elapsed_time      {"id": "every5s", "name": "sched.run_total_elapsed_time", "interval": "3.803831971s", "tags": {
"ID": "every5s"
}
}
2021-04-10T13: 14: 05.558+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "every5s", "After": "5s", "At": "2021-04-10T13:14:10+02:00"}
2021-04-10T13: 14:05.558+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "every5s", "Instance": "8e70617b-8245-4b24-aff1-bf4a823ca6ef"}
2021/04/10 13: 14: 05 Doing some work for random time...
2021/04/10 13: 14: 07 Finished Work.
2021-04-10T13: 14: 07.660+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "every5s", "Instance": "8e70617b-8245-4b24-aff1-bf4a823ca6ef", "Duration": "2.102s", "State": "FINISHED"}
2021-04-10T13: 14: 07.660+0200    INFO    sched.metrics   sched/metric.go: 48      timer sched.run_actual_elapsed_time     {"id": "every5s", "name": "sched.run_actual_elapsed_time", "interval": "2.10165422s", "tags": {"ID": "every5s"}}
2021-04-10T13: 14:07.660+0200    INFO    sched.metrics   sched/metric.go: 48      timer sched.run_total_elapsed_time      {"id": "every5s", "name": "sched.run_total_elapsed_time", "interval": "2.101723938s", "tags": {"ID": "every5s"}}
2021-04-10T13: 14: 10.557+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "every5s", "After": "5s", "At": "2021-04-10T13:14:15+02:00"}
2021-04-10T13: 14: 10.557+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "every5s", "Instance": "b4d3b29e-6d21-43d5-8550-6e462625f7ee"}
2021/04/10 13: 14: 10 Doing some work for random time...
2021/04/10 13: 14: 11 Finished Work.
2021-04-10T13: 14: 11.462+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "every5s", "Instance": "b4d3b29e-6d21-43d5-8550-6e462625f7ee", "Duration": "905ms", "State": "FINISHED"
}
2021-04-10T13: 14: 11.462+0200    INFO    sched.metrics   sched/metric.go: 48      timer sched.run_actual_elapsed_time     {"id": "every5s", "name": "sched.run_actual_elapsed_time", "interval": "904.904666ms", "tags": {"ID": "every5s"}}
2021-04-10T13: 14: 11.462+0200    INFO    sched.metrics   sched/metric.go: 48      timer sched.run_total_elapsed_time      {"id": "every5s", "name": "sched.run_total_elapsed_time", "interval": "905.002981ms", "tags": {"ID": "every5s"
}
}
2021-04-10T13: 14: 15.558+0200    INFO    sched.metrics   sched/metric.go: 40      counter sched.runs      {"id": "every5s", "name": "sched.runs", "value": 4, "tags": {"ID": "every5s"}}
2021-04-10T13: 14: 15.558+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "every5s", "After": "5s", "At": "2021-04-10T13:14:20+02:00"
}
2021-04-10T13: 14: 15.558+0200    INFO    sched   sched/schedule.go:193   Job Run Starting        {"id": "every5s", "Instance": "ac3ff586-9d77-410a-b493-8274347d8b72"
}
2021/04/10 13: 14: 15 Doing some work for random time...
2021/04/10 13: 14:20 Finished Work.
2021-04-10T13: 14: 20.459+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "every5s", "Instance": "ac3ff586-9d77-410a-b493-8274347d8b72", "Duration": "4.901s", "State": "FINISHED"}
2021-04-10T13: 14: 20.459+0200    INFO    sched.metrics   sched/metric.go: 48      timer sched.run_actual_elapsed_time     {"id": "every5s", "name": "sched.run_actual_elapsed_time", "interval": "4.900706131s", "tags": {
"ID": "every5s"
}
}
2021-04-10T13: 14: 20.459+0200    INFO    sched.metrics   sched/metric.go:48      timer sched.run_total_elapsed_time      {"id": "every5s", "name": "sched.run_total_elapsed_time", "interval": "4.900801476s", "tags": {"ID": "every5s"
}
}
2021-04-10T13: 14: 20.555+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "every5s", "After": "5s", "At": "2021-04-10T13:14:25+02:00"
}
2021-04-10T13: 14: 20.555+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "every5s", "Instance": "1d623dcb-e297-4987-baef-e840479abd06"
}
2021/04/10 13:14: 20 Doing some work for random time...
2021/04/10 13: 14: 22 Finished Work.
2021-04-10T13: 14: 22.257+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "every5s", "Instance": "1d623dcb-e297-4987-baef-e840479abd06", "Duration": "1.702s", "State": "FINISHED"}
2021-04-10T13:14: 22.257+0200    INFO    sched.metrics   sched/metric.go: 48      timer sched.run_actual_elapsed_time     {"id": "every5s", "name": "sched.run_actual_elapsed_time", "interval": "1.70196874s", "tags": {"ID":"every5s"
}
}
2021-04-10T13: 14: 22.257+0200    INFO    sched.metrics   sched/metric.go: 48      timer sched.run_total_elapsed_time      {"id": "every5s", "name": "sched.run_total_elapsed_time", "interval": "1.702028643s", "tags": {"ID": "every5s"}}
2021-04-10T13:14: 25.557+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "every5s", "After": "5s", "At": "2021-04-10T13:14:30+02:00"
}
2021-04-10T13: 14: 25.557+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "every5s", "Instance": "1ef7274b-c899-4156-a4f2-edb85a6db71a"}
2021/04/10 13: 14: 25 Doing some work for random time...
2021-04-10T13: 14: 30.556+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "every5s", "After": "5s", "At": "2021-04-10T13:14:35+02:00"}
2021-04-10T13: 14: 30.556+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "every5s", "Instance": "aface894-ff68-4e3c-8c72-0f0d8d1a4838"}
2021/04/10 13: 14: 30 Doing some work for random time...
2021/04/10 13: 14: 30 Finished Work.
2021-04-10T13: 14:30.557+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "every5s", "Instance": "1ef7274b-c899-4156-a4f2-edb85a6db71a", "Duration": "5s", "State": "FINISHED"
}
2021-04-10T13: 14: 30.557+0200    INFO    sched.metrics   sched/metric.go: 48      timer sched.run_actual_elapsed_time     {"id": "every5s", "name": "sched.run_actual_elapsed_time", "interval": "5.000234092s", "tags": {"ID": "every5s"}
}
2021-04-10T13: 14: 30.557+0200    INFO    sched.metrics   sched/metric.go: 48      timer sched.run_total_elapsed_time      {"id": "every5s", "name": "sched.run_total_elapsed_time", "interval": "5.00025765s", "tags": {
"ID": "every5s"
}
}
2021-04-10T13: 14: 32.869+0200    INFO    sched   sched/schedule.go: 125   Stopping Schedule...    {"id": "every5s"}
2021-04-10T13: 14:32.869+0200    INFO    sched   sched/schedule.go: 130   Waiting active jobs to finish...        {"id": "every5s"}
2021-04-10T13:14: 32.869+0200    INFO    sched   sched/schedule.go: 171   Job Next Run Canceled   {"id": "every5s", "At": "2021-04-10T13:14:35+02:00"
}
2021/04/10 13: 14: 34 Finished Work.
2021-04-10T13: 14: 34.057+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "every5s", "Instance": "aface894-ff68-4e3c-8c72-0f0d8d1a4838", "Duration": "3.501s", "State": "FINISHED"}
2021-04-10T13: 14: 34.057+0200    INFO    sched.metrics   sched/metric.go: 48      timer sched.run_actual_elapsed_time     {"id": "every5s", "name": "sched.run_actual_elapsed_time", "interval": "3.50128391s", "tags": {
"ID": "every5s"
}
}
2021-04-10T13: 14: 34.057+0200    INFO    sched.metrics   sched/metric.go: 48      timer sched.run_total_elapsed_time      {"id": "every5s", "name": "sched.run_total_elapsed_time", "interval": "3.501411708s", "tags": {"ID": "every5s"}
}
2021-04-10T13: 14: 34.057+0200    INFO    sched   sched/schedule.go: 133   Job Schedule Stopped    {"id": "every5s"
}
```
