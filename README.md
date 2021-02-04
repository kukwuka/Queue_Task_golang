# QueueLine for test
First argument is Writers Count

Second argument is Array Size

Third argument is Iterations Count

for run 
```bash
go run  main.go 9 68 4
```

for build
```bash
go build -race main.go
main  9 68 4
```
Result be like 
```bash
RoutineId :0 time: 2021-02-04 17:16:12.3005471 +0300 MSK m=+0.004991201 min: 1 avarage :21 max:22
RoutineId :0 time: 2021-02-04 17:16:12.3005471 +0300 MSK m=+0.004991201 min: 6 avarage :1 max:40
RoutineId :0 time: 2021-02-04 17:16:12.3005471 +0300 MSK m=+0.004991201 min: 30 avarage :18 max:42
RoutineId :3 time: 2021-02-04 17:16:12.3005471 +0300 MSK m=+0.004991201 min: 27 avarage :18 max:32
RoutineId :6 time: 2021-02-04 17:16:12.3005471 +0300 MSK m=+0.004991201 min: 51 avarage :45 max:56
RoutineId :4 time: 2021-02-04 17:16:12.3005471 +0300 MSK m=+0.004991201 min: 19 avarage :42 max:32
RoutineId :2 time: 2021-02-04 17:16:12.3005471 +0300 MSK m=+0.004991201 min: 32 avarage :59 max:40
RoutineId :5 time: 2021-02-04 17:16:12.3005471 +0300 MSK m=+0.004991201 min: 34 avarage :17 max:48
RoutineId :7 time: 2021-02-04 17:16:12.3005471 +0300 MSK m=+0.004991201 min: 2 avarage :52 max:26
RoutineId :8 time: 2021-02-04 17:16:12.3005471 +0300 MSK m=+0.004991201 min: 14 avarage :19 max:67
RoutineId :1 time: 2021-02-04 17:16:12.3005471 +0300 MSK m=+0.004991201 min: 62 avarage :64 max:57
RoutineId :0 time: 2021-02-04 17:16:12.3005471 +0300 MSK m=+0.004991201 min: 20 avarage :28 max:40
RoutineId :3 time: 2021-02-04 17:16:12.3515467 +0300 MSK m=+0.055990801 min: 4 avarage :4 max:46
RoutineId :6 time: 2021-02-04 17:16:12.3515467 +0300 MSK m=+0.055990801 min: 51 avarage :30 max:41
RoutineId :4 time: 2021-02-04 17:16:12.3525651 +0300 MSK m=+0.057009201 min: 23 avarage :21 max:28
RoutineId :2 time: 2021-02-04 17:16:12.3525651 +0300 MSK m=+0.057009201 min: 16 avarage :26 max:67
RoutineId :5 time: 2021-02-04 17:16:12.3525651 +0300 MSK m=+0.057009201 min: 31 avarage :43 max:7
RoutineId :7 time: 2021-02-04 17:16:12.3525651 +0300 MSK m=+0.057009201 min: 48 avarage :64 max:3
RoutineId :8 time: 2021-02-04 17:16:12.3535456 +0300 MSK m=+0.057989701 min: 29 avarage :3 max:65
RoutineId :1 time: 2021-02-04 17:16:12.3535456 +0300 MSK m=+0.057989701 min: 57 avarage :39 max:16
RoutineId :3 time: 2021-02-04 17:16:12.3535456 +0300 MSK m=+0.057989701 min: 33 avarage :18 max:8
RoutineId :6 time: 2021-02-04 17:16:12.3545489 +0300 MSK m=+0.058993001 min: 29 avarage :42 max:61
RoutineId :4 time: 2021-02-04 17:16:12.3545489 +0300 MSK m=+0.058993001 min: 32 avarage :65 max:17
RoutineId :2 time: 2021-02-04 17:16:12.3545489 +0300 MSK m=+0.058993001 min: 28 avarage :42 max:16
RoutineId :5 time: 2021-02-04 17:16:12.3545489 +0300 MSK m=+0.058993001 min: 5 avarage :8 max:36
RoutineId :7 time: 2021-02-04 17:16:12.3545489 +0300 MSK m=+0.058993001 min: 25 avarage :16 max:51
RoutineId :8 time: 2021-02-04 17:16:12.3545489 +0300 MSK m=+0.058993001 min: 2 avarage :43 max:28
RoutineId :1 time: 2021-02-04 17:16:12.3545489 +0300 MSK m=+0.058993001 min: 13 avarage :26 max:21
RoutineId :3 time: 2021-02-04 17:16:12.3545489 +0300 MSK m=+0.058993001 min: 57 avarage :6 max:12
RoutineId :6 time: 2021-02-04 17:16:12.3555449 +0300 MSK m=+0.059989001 min: 18 avarage :60 max:29
RoutineId :4 time: 2021-02-04 17:16:12.3555449 +0300 MSK m=+0.059989001 min: 44 avarage :25 max:13
RoutineId :2 time: 2021-02-04 17:16:12.3555449 +0300 MSK m=+0.059989001 min: 2 avarage :36 max:31
RoutineId :5 time: 2021-02-04 17:16:12.3555449 +0300 MSK m=+0.059989001 min: 1 avarage :12 max:14
RoutineId :7 time: 2021-02-04 17:16:12.3555449 +0300 MSK m=+0.059989001 min: 13 avarage :57 max:32
RoutineId :8 time: 2021-02-04 17:16:12.3555449 +0300 MSK m=+0.059989001 min: 22 avarage :67 max:61
RoutineId :1 time: 2021-02-04 17:16:12.3555449 +0300 MSK m=+0.059989001 min: 26 avarage :54 max:22

```