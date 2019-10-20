

## The Simulation

I created an application to learn some basic Go and use those lovely Go Routines.

To do so I desired to simulate an elevator bank at an office building.

The Elevator Bank has a collection of elevators. People make requests to the bank for an elevator. When it arrives they enter and press the button for the floor they desire. The elevators ascend or descend to fulfill those requests. Once peoples' schedules have been complete the simulation ends.

The bank, elevators and people all run within `go routines`.

I used structs to object-orient the design and a Finite State Machine within each struct `Run()` method to manage their AI.

Since each object's AI runs within a go routine I used Mutexes to ensure requests made from object to object are safe.

```
$ go run sim.go
```

<img src="https://media.giphy.com/media/hu7xwqc6DMOgVo6sW9/giphy.gif" />
