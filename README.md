<img src="https://media.giphy.com/media/hu7xwqc6DMOgVo6sW9/giphy.gif" width="500" />

## What it be:

I created an application to learn some basic Go and use those lovely Go Routines.

To do so I desired to simulate an elevator bank at an office building.

The Elevator Bank has a collection of elevators. People make requests to the bank for an elevator. When it arrives they enter and press the button for the floor they desire. The elevators ascend or descend to fulfill those requests. Once peoples' schedules have been complete the simulation ends.

The bank, elevators and people all run within `go routines`.

I used structs to object-orient the design and a Finite State Machine within each struct `Run()` method to manage their AI.

Since each object's AI runs within a go routine I used Mutexes to ensure requests made from object to object are safe.

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
