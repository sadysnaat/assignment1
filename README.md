#### Install
Run 
```
go get github.com/sadysnaat/assignment1
```

#### Running
Run 
```
cd <path_to_installed_location>
go build
```

then 
```
./assignment1
```

#### Output 
```
Doing some heavy work, please wait a while
results [{0} {1} {2} {3} {4} {5} {6} {7} {8} {9}]
Doing the same heavy work but async, expect result as soon as they are ready
got 0
got 1
got 2
got 3
got 4
got 5
got 6
got 7
got 8
got 9
completed
[{0} {1} {2} {3} {4} {5} {6} {7} {8} {9}]
```

#### Design
Both WaitTillComplete() and Subscribe() implement the same workload where they iterate over 10 times and pause for 100ms 
