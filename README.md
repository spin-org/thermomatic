# Thermomatic

This code challenge may be completed on your own time. The challenge represents the kind of work we're doing at Spin. However, details are mostly fictional :)

Fork this repo. When it's ready, please invite `azazeal` for review.

## Preface

As part of Spin's upcoming entry to the microfarming industry, you've been assigned the task of creating a concurrent logging server for our next-gen, over-the-ether thermometers!

These groundbreaking devices have the singular purpose of broadcasting exact temperature (and other readings) from wherever they may be placed at.

The way they work is simple:

- Each device connects, over TCP/IP, to our servers.
- When a device connects to our servers it sends a _login_ message.
- After the login message has been sent then the device sends a singular message containing all of its readings, _typically_ every _25ms_.

## Task definition

Your task is to implement a TCP/IP server which:

- Accepts and maintains connections with any client capable of handshaking via the simple handshake mechanism described in the _Login Message_ section of this document.
- Gathers `Reading` messages (as defined in the _Payload message_ section of this document) from the devices.
- Outputs each received *valid* `Reading` message to the standard output of the program in the format defined in the _Output format example_ section.
- Drops client connections which fail to send at least a _Reading_ every _2 seconds_.

## Protocol specification

### Basics

- All of our servers listen on `tcp/1337`.
- Server output is written to `os.Stdout`
- Logging messages are written to `os.Stderr`.

### Login message

When devices connect to our servers they send a 15-byte long message containing their IMEI code in decimal format.

- Should the IMEI code be not be valid then the server will drop the client connection.
- Should the device fail to send the login message within _1 second_ the server will drop the client connection.

### Payload message

After the login message has been sent by the device the device starts sending, typically every 25ms, _Reading_ messages.

These fixed-size messages are formatted as follows:

| Field         | Start Index   | Size (in bytes) | Notes                                                                |
| ------------- | ------------- | --------------- | ---------------------------------------------------------------------|
| Temperature   | 0             | 8               | The temperature reading of the device. Celcius. Min/Max: [-300, 300] |
| Altitude      | 8             | 8               | The altitude reading of the device. Meters. Min/Max: [-20000, 20000] |
| Latitude      | 16            | 8               | The latitude reading of the device. Degrees. Min/Max: [-90, 90]      |
| Longitude     | 24            | 8               | The longitude reading of the device. Degrees. Min/Max: [-180, 180]   |
| BatteryLevel  | 32            | 8               | The battery level of the device. Percentage. Min/Max: (0, 100]       |

All these fields are `IEEE 754` binary representations of `float64` values encoded in Big-Endian.

## Output format example

Given a `Reading` message originating from the device with IMEI code `490154203237518`, received `1257894000000000000` nanoseconds since `January 1, 1970 UTC`, carrying the following values:

- Temperature: `67.77`
- Altitude: `2.63555`
- Latitude: `33.41`
- Longitude: `44.4`
- BatteryLevel: `0.25666`

the corresponding logging record, were it a Go string, would be:

```go
record = "1257894000000000000,490154203237518,67.77,2.63555,33.41,44.4,0.2566\n"
```

## Things we expect to see

- Meaningful (including _performance_) tests with reasonable coverage.
- Benchmarks.
- Elimination of allocations wherever possible.
- An effort to remain in the stack vs escaping to the heap.
- Bounds check eliminations wherever possible.
- Code Documentation!
- Detailed logging of any client connection's lifecycle.
- Detailed logging of any server-side noteworthy events.
- Zero dependencies to 3rd party libraries.
- A series of [well-formed](https://github.com/golang/go/wiki/CommitMessage) commits.

## Bonus objectives

If you feel like spending a bit more time on this challenge, you may also extend your implementation to support the following HTTP GET endpoints:

- `/stats`: returns a JSON document which contains runtime statistical information about the server (i.e. number of goroutines, bytes read per second, etc.).
- `/readings/:imei`: if the device is online returns a JSON representation of the last reading the device has sent (timestamped)
- `/status/:imei`: reports whether the device is online or not.

## Hints

- [`IEEE 754`](https://golang.org/pkg/math/#Float64bits)
- `io.Reader` is great but Go doesn't currently support full program escape analysis.

## Postface

- You may alter any of the existing code in order to perfect your deliverable.
- You may devise your own strategy against resource exhaustion attacks.
- You may devise your own strategy for what should happen when a device attempts to login twice.
