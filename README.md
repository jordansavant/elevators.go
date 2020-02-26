

## The Simulation

I created an application to learn some basic Go and use those lovely Go Routines.

To do so I desired to simulate an elevator bank at an office building.

The Elevator Bank has a collection of elevators. People make requests to the bank for an elevator. When it arrives they enter and press the button for the floor they desire. The elevators ascend or descend to fulfill those requests. Once peoples' schedules have been complete the simulation ends.

The bank, elevators and people all run within `go routines`.

I used structs to object-orient the design and a Finite State Machine within each struct `Run()` method to manage their AI.

Since each object's AI runs within a go routine I used Mutexes to ensure requests made from object to object are safe.

It runs in the status of a Server or a Client.

To start the server:
```
$ ./bin/elevators.go server [number of floors] [number of elevators]
```

To start the gui client:
```
$ ./bin/elevators.go client gui
```

Once in the GUI you can add a person with "Enter" or turn on automatic with "A"

<img src="https://raw.githubusercontent.com/jordansavant/elevators.go/master/gui.gif" />

<img src="https://raw.githubusercontent.com/jordansavant/elevators.go/master/elevators.gif" />
