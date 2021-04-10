# Output with a job that run every 2s, takes 1s to run, and expected is 500ms.

Notice the WARN logs.

## Output

```json
2021-04-10T13:57:59.558+0200    INFO    sched   sched/schedule.go:101   Job Schedule Started    {"id": "fixed2s"}
2021-04-10T13:57:59.558+0200    INFO    sched   sched/schedule.go:176   Job Next Run Scheduled  {"id": "fixed2s", "After": "2s", "At": "2021-04-10T13:58:01+02:00"}
2021-04-10T13:58:01.558+0200    INFO    sched   sched/schedule.go:176   Job Next Run Scheduled  {"id": "fixed2s", "After": "2s", "At": "2021-04-10T13:58:03+02:00"}
2021-04-10T13:58:01.558+0200    INFO    sched   sched/schedule.go:203   Job Run Starting        {"id": "fixed2s", "Instance": "4cb6b448-5b93-4689-a129-5448297e727e"}
2021/04/10 13:58:01 Doing some work...
2021-04-10T13:58:02.060+0200    WARN    sched   sched/schedule.go:211   Job Run Exceeded Expected Time  {"id": "fixed2s", "Instance": "4cb6b448-5b93-4689-a129-5448297e727e", "Expected": "1s"}
github.com/sherifabdlnaby/sched.(*Schedule).runJobInstance.func1
/Users/sherifabdlnaby/code/projects/sched/schedule.go:211
2021/04/10 13:58:02 Finished Work.
2021-04-10T13:58:02.560+0200    INFO    sched   sched/schedule.go:229   Job Finished    {"id": "fixed2s", "Instance": "4cb6b448-5b93-4689-a129-5448297e727e", "Duration": "1.002s", "State": "FINISHED"}
2021-04-10T13:58:03.561+0200    INFO    sched   sched/schedule.go:176   Job Next Run Scheduled  {"id": "fixed2s", "After": "2s", "At": "2021-04-10T13:58:05+02:00"}
2021-04-10T13:58:03.561+0200    INFO    sched   sched/schedule.go:203   Job Run Starting        {"id": "fixed2s", "Instance": "3d47d01b-f2d3-4bf3-830c-1815308e3b1f"}
2021/04/10 13:58:03 Doing some work...
2021-04-10T13:58:04.062+0200    WARN    sched   sched/schedule.go:211   Job Run Exceeded Expected Time  {"id": "fixed2s", "Instance": "3d47d01b-f2d3-4bf3-830c-1815308e3b1f", "Expected": "1s"}
github.com/sherifabdlnaby/sched.(*Schedule).runJobInstance.func1
/Users/sherifabdlnaby/code/projects/sched/schedule.go:211
2021/04/10 13:58:04 Finished Work.
2021-04-10T13:58:04.564+0200    INFO    sched   sched/schedule.go:229   Job Finished    {"id": "fixed2s", "Instance": "3d47d01b-f2d3-4bf3-830c-1815308e3b1f", "Duration": "1.003s", "State": "FINISHED"}
2021-04-10T13:58:05.561+0200    INFO    sched   sched/schedule.go:176   Job Next Run Scheduled  {"id": "fixed2s", "After": "2s", "At": "2021-04-10T13:58:07+02:00"}
2021-04-10T13:58:05.561+0200    INFO    sched   sched/schedule.go:203   Job Run Starting        {"id": "fixed2s", "Instance": "78059970-4fc9-4e85-b310-15a3aef08602"}
2021/04/10 13:58:05 Doing some work...
2021-04-10T13:58:06.066+0200    WARN    sched   sched/schedule.go:211   Job Run Exceeded Expected Time  {"id": "fixed2s", "Instance": "78059970-4fc9-4e85-b310-15a3aef08602", "Expected": "1s"}
github.com/sherifabdlnaby/sched.(*Schedule).runJobInstance.func1
/Users/sherifabdlnaby/code/projects/sched/schedule.go:211
2021/04/10 13:58:06 Finished Work.
2021-04-10T13:58:06.563+0200    INFO    sched   sched/schedule.go:229   Job Finished    {"id": "fixed2s", "Instance": "78059970-4fc9-4e85-b310-15a3aef08602", "Duration": "1.002s", "State": "FINISHED"}
2021-04-10T13:58:07.561+0200    INFO    sched   sched/schedule.go:176   Job Next Run Scheduled  {"id": "fixed2s", "After": "2s", "At": "2021-04-10T13:58:09+02:00"}
2021-04-10T13:58:07.561+0200    INFO    sched   sched/schedule.go:203   Job Run Starting        {"id": "fixed2s", "Instance": "2e846c92-e0d6-4028-bca4-d930caded0ce"}
2021/04/10 13:58:07 Doing some work...
2021-04-10T13:58:08.066+0200    WARN    sched   sched/schedule.go:211   Job Run Exceeded Expected Time  {"id": "fixed2s", "Instance": "2e846c92-e0d6-4028-bca4-d930caded0ce", "Expected": "1s"}
github.com/sherifabdlnaby/sched.(*Schedule).runJobInstance.func1
/Users/sherifabdlnaby/code/projects/sched/schedule.go:211
2021/04/10 13:58:08 Finished Work.
2021-04-10T13:58:08.561+0200    INFO    sched   sched/schedule.go:229   Job Finished    {"id": "fixed2s", "Instance": "2e846c92-e0d6-4028-bca4-d930caded0ce", "Duration": "1s", "State": "FINISHED"}
2021-04-10T13:58:09.245+0200    INFO    sched   sched/schedule.go:130   Stopping Schedule...    {"id": "fixed2s"}
2021-04-10T13:58:09.245+0200    INFO    sched   sched/schedule.go:141   Job Schedule Stopped    {"id": "fixed2s"}

```
