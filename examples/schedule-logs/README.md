# 4 Logging with Other logrus

1. Cron Every Minute
2. Cron Every 5 Minutes
3. Fixed Interval Every 30 Secs
4. *Once* after 10 Secs from schedule start.

Started and Stopped using `StartAll()` and `StopAll()`

## Output

```json
INFO[0000] Job Schedule Started                          From=sched id=cronEvery5Minute
INFO[0000] Job Schedule Started                          From=sched id=fixedTimer30second
INFO[0000] Job Schedule Started                          From=sched id=onceAfter10s
INFO[0000] Job Schedule Started                          From=sched id=cronEveryMinute
INFO[0000] Job Next Run Scheduled                        After=10s At="2021-04-19T13:09:08+08:00" From=sched id=onceAfter10s
INFO[0000] Job Next Run Scheduled                        After=30s At="2021-04-19T13:09:28+08:00" From=sched id=fixedTimer30second
INFO[0000] Job Next Run Scheduled                        After=2s At="2021-04-19T13:09:00+08:00" From=sched id=cronEveryMinute
INFO[0000] Job Next Run Scheduled                        After=1m2s At="2021-04-19T13:10:00+08:00" From=sched id=cronEvery5Minute
INFO[0001] Job Next Run Scheduled                        After=1m0s At="2021-04-19T13:10:00+08:00" From=sched id=cronEveryMinute
INFO[0001] Job Run Starting                              From=sched Instance=2c114df6-e795-419a-8bf3-e4aada560a37 id=cronEveryMinute
2021/04/19 13:09:00 every-minute-cron    Doing some work...
2021/04/19 13:09:01 every-minute-cron    Finished Work.
INFO[0002] Job Finished                                  Duration=1s From=sched Instance=2c114df6-e795-419a-8bf3-e4aada560a37 State=FINISHED id=cronEveryMinute
INFO[0002] timer sched.run_actual_elapsed_time           From=metrics id=cronEveryMinute interval=1.0004944s name=sched.run_actual_elapsed_time tags="map[id:cronEveryMinute]"
INFO[0002] timer sched.run_total_elapsed_time            From=metrics id=cronEveryMinute interval=1.0006126s name=sched.run_total_elapsed_time tags="map[id:cronEveryMinute]"
INFO[0010] No more Jobs will be scheduled                From=sched id=onceAfter10s
INFO[0010] Stopping Schedule...                          From=sched id=onceAfter10s
INFO[0010] Waiting for '1' active jobs still running...  From=sched id=onceAfter10s
INFO[0010] Job Run Starting                              From=sched Instance=80691c48-8468-4029-9bc8-0264698a1cc2 id=onceAfter10s
2021/04/19 13:09:08 onceAfter10s         Doing some work...
2021/04/19 13:09:09 onceAfter10s         Finished Work.
INFO[0011] Job Finished                                  Duration=1s From=sched Instance=80691c48-8468-4029-9bc8-0264698a1cc2 State=FINISHED id=onceAfter10s
INFO[0011] timer sched.run_actual_elapsed_time           From=metrics id=onceAfter10s interval=1.0004959s name=sched.run_actual_elapsed_time tags="map[id:onceAfter10s]"
INFO[0011] timer sched.run_total_elapsed_time            From=metrics id=onceAfter10s interval=1.0006474s name=sched.run_total_elapsed_time tags="map[id:onceAfter10s]"
INFO[0011] Job Schedule Stopped                          From=sched id=onceAfter10s
INFO[0011] Job Schedule Finished                         From=sched id=onceAfter10s
INFO[0030] Job Next Run Scheduled                        After=30s At="2021-04-19T13:09:58+08:00" From=sched id=fixedTimer30second
INFO[0030] Job Run Starting                              From=sched Instance=eded3660-2b7c-46b0-8dcc-8a393102771b id=fixedTimer30second
2021/04/19 13:09:28 fixedEvery30Second   Doing some work...
2021/04/19 13:09:29 fixedEvery30Second   Finished Work.
INFO[0031] Job Finished                                  Duration=1.001s From=sched Instance=eded3660-2b7c-46b0-8dcc-8a393102771b State=FINISHED id=fixedTimer30second
INFO[0031] timer sched.run_actual_elapsed_time           From=metrics id=fixedTimer30second interval=1.0006731s name=sched.run_actual_elapsed_time tags="map[id:fixedTimer30second]"
INFO[0031] timer sched.run_total_elapsed_time            From=metrics id=fixedTimer30second interval=1.0008061s name=sched.run_total_elapsed_time tags="map[id:fixedTimer30second]"
^CINFO[0052] Stopping Schedule...                          From=sched id=cronEveryMinute
INFO[0052] Job Schedule Stopped                          From=sched id=cronEveryMinute
INFO[0052] Job Next Run Canceled                         At="2021-04-19T13:10:00+08:00" From=sched id=cronEveryMinute
INFO[0052] Stopping Schedule...                          From=sched id=cronEvery5Minute
INFO[0052] Job Schedule Stopped                          From=sched id=cronEvery5Minute
INFO[0052] Job Next Run Canceled                         At="2021-04-19T13:10:00+08:00" From=sched id=cronEvery5Minute
INFO[0052] Stopping Schedule...                          From=sched id=fixedTimer30second
INFO[0052] Job Schedule Stopped                          From=sched id=fixedTimer30second
```

