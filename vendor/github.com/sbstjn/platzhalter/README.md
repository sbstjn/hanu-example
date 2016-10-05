# platzhalter

**Platzhalter** is German for *placeholder* and a `golang` library to define strings with placeholders and check if others strings are matching the defined string. This can be useful when using a command line tool, writing a slack bot or some other crazy ideas …

### Usage

Check if a string matches you defined command:

```go
cmd := platzhalter.NewCommand("command <paramOne> <paramTwo>")

if (cmd.Matches("command lorem ipsum")) {
  fmt.Println(`\o/`)
}
```

Get paramters from a string matching your command:

```go
cmd := platzhalter.NewCommand("command <paramOne> <paramTwo>")

fmt.Println(cmd.Param("command lorem ipsum", "paramOne")) // lorem
fmt.Println(cmd.Param("command lorem ipsum", "paramTwo")) // ipsum
```
