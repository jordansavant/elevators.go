Lets see if we can use Golang to make a threaded elevator simulator.

Would be cool to be able to hook it into a websocket that adds people to the queue.

```
$ go run sim.go
```

```
running
@ Bank is starting elevators
- Bob requests a lift
@ Bank assigns elevator to 1
. EL01 button requested for 1
. EL01 at goal opening doors at ready
- Bob elevator arrived, getting on and pressing 3
. EL01 button requested for 3
. EL01 closing doors and moving to idle
. EL01 moving towards 3
. EL01 moving up
. EL01 moving up
. EL01 moving up
. EL01 moving up
. EL01 moving up
. EL01 moving up
. EL01 moving up
. EL01 moving up
. EL01 arrived at 3
. EL01 moving to ready and opening doors
- Bob elevator ready at level, going to work
- Bob is working for 10 seconds
. EL01 closing doors and moving to idle
- Bob done working, going idle
- Bob requests a lift
@ Bank assigns elevator to 3
. EL01 button requested for 3
. EL01 at goal opening doors at ready
- Bob elevator arrived, getting on and pressing 2
. EL01 button requested for 2
. EL01 closing doors and moving to idle
. EL01 moving towards 2
. EL01 moving down
. EL01 moving down
. EL01 moving down
. EL01 moving down
. EL01 arrived at 2
. EL01 moving to ready and opening doors
- Bob elevator ready at level, going to work
- Bob is working for 5 seconds
. EL01 closing doors and moving to idle
- Bob done working, going idle
- Bob requests a lift
@ Bank assigns elevator to 2
. EL01 button requested for 2
. EL01 at goal opening doors at ready
- Bob elevator arrived, getting on and pressing 1
. EL01 button requested for 1
. EL01 closing doors and moving to idle
. EL01 moving towards 1
. EL01 moving down
. EL01 moving down
. EL01 moving down
. EL01 moving down
. EL01 arrived at 1
. EL01 moving to ready and opening doors
- Bob elevator ready at level, going to work
- Bob is working for 0 seconds
- Bob done working, going idle
- Bob leaving
ending
```
