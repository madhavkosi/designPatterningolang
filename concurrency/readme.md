### Concurrency in Go


#### Concurrency vs. Parallelism
- **Concurrency**: Executing multiple tasks at the same time within the same process.
- **Parallelism**: Executing multiple tasks on two or more processes at the same time.

#### Golang's Concurrency Management
- Go uses a scheduler strategy known as M:N, scheduling M goroutines on N OS threads.
- **Goroutine**: A logical unit of execution containing the program or function to run and its related information.
- **OS Thread**: Managed by the OS, used by the Go scheduler to run goroutines.
- **Context/Processor**: Manages the execution of goroutines, using an internal scheduler and runqueue.

#### Go Fork-Join Model
- Go controls goroutines using the fork-join model, allowing processes to split and rejoin.
- **Incomplete Fork-Join**: When the result of a goroutine is not needed for further execution.
  - Example: Saving a new client record and an audit record without needing to wait for the audit record's confirmation.
- **Complete Fork-Join**: When the process splits and then rejoins after all goroutines have completed.


### Goroutines in Go

#### Definition
- A goroutine is a logical unit of execution in Go that defines the program or functions to run.
- It contains important information such as stack memory, the machine/thread it is running on, and the stack trace.
- Essentially, a goroutine is a function executed concurrently with the main goroutine.

#### Creating a Goroutine
- In Go, a new "thread" (goroutine) is created using the keyword `go` before the function call.