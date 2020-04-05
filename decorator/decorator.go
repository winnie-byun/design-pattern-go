package decorator

import "log"

// Decorator structural pattern allows extending the function
// of an existing object dynamically without altering its internals.

type Command func(string) string

func RunWithLog(fn Command, name string) Command {
	return func(n string) string {
		log.Println("Starting the execution with command : ", name, n)

		result := fn(n)

		log.Println("Execution is completed with command : ", result)

		return result
	}
}
