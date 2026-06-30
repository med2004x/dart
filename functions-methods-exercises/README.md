# Functions And Methods Exercises

This folder is a focused track for learning Go functions and methods.

The exercises start with simple function calls and move toward method receivers,
state changes, dependency boundaries, and a small capstone. Each project is a
separate `main` package.

## How To Use This Track

For each numbered folder:

1. Read the README.
2. Open `main.go`.
3. Add only the required functions or methods.
4. Run the program.
5. Prove the listed cases manually from output.

Do not jump to methods before functions are comfortable. Methods are functions
with a receiver; if normal parameters and returns are unclear, methods will feel
random.

## Exercise Map

| Exercise | Main focus |
|---|---|
| 01 Function Calls | call order and no-return functions |
| 02 Parameters | passing input into functions |
| 03 Return Values | computing and returning results |
| 04 Multiple Returns | returning value plus success |
| 05 Scope | local variables and shadowing |
| 06 Slice Parameters | functions that receive collections |
| 07 Return Updated State | changing state by returning values |
| 08 Function Composition | building larger behavior from smaller functions |
| 09 Value Receiver Methods | methods that read struct state |
| 10 Pointer Receiver Methods | methods that mutate struct state |
| 11 Functions Vs Methods | choosing the right form |
| 12 Error Returns | functions that report failure |
| 13 Method Sets | pointer/value receiver differences |
| 14 Higher-Order Functions | functions passed to functions |
| 15 Capstone | design a small library checkout model |

## Running

From a project folder:

```powershell
gofmt -w .\main.go
go run .
```

From this track root:

```powershell
go test ./...
```
