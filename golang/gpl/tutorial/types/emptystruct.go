package types

import (
	"fmt"
	"time"
)

// Set example for struct{} usage
func setExample() {
	fmt.Printf("\n\n== Set example for struct{} usage ==\n\n")

	set := map[string]struct{}{}

	set["hello"] = struct{}{}
	set["world"] = struct{}{}

	if _, ok := set["hello"]; ok {
		fmt.Println("hello is in the set")
	}

	if _, ok := set["name"]; !ok {
		fmt.Println("name is not in the set")
	}
}

// Async task example for struct{} usage
func asyncTask(done chan struct{}) {
	fmt.Println("Worker: Doing some work...")
	time.Sleep(2 * time.Second)
	fmt.Println("Worker: Work finished.")
	done <- struct{}{} // Signal that work is done
}

func executeAsyncTask() {
	fmt.Printf("\n\n== Async task example for struct{} usage ==\n\n")

	done := make(chan struct{})
	go asyncTask(done)

	<-done

	fmt.Println("Main: Worker finished, continuing...")
}

// Compile time check example
type Lifecycle interface {
	Start() error
	Stop() error
}

type lifecycleMethods struct{} // See the empty struct{} usage

func (lifecycleMethods) Start() error {
	fmt.Println("Plugin started")
	return nil
}

func (lifecycleMethods) Stop() error {
	fmt.Println("Plugin stopped")
	return nil
}

func embeddingStructExample() {
	fmt.Printf("\n\n== Embedding struct example ==\n\n")

	type MyPlugin struct {
		// Embedding an anonymous field of type lifecycleMethods
		lifecycleMethods // This is the key part for the compile-time check
		Name             string
	}

	plugin := MyPlugin{
		Name: "MyPlugin",
	}

	runLifeCycle := func(l Lifecycle) {
		l.Start()
		l.Stop()
	}

	runLifeCycle(plugin)
}
