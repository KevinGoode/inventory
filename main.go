package main

import (
	"fmt"
	"os"

	impl "github.com/KevinGoode/inventory/service"
	"go.uber.org/dig"
)

func buildContainer() *dig.Container {
	container := dig.New()
	container.Provide(impl.NewDatabaseFactory)
	container.Provide(impl.NewCassandraDbCreater)
	container.Provide(impl.NewParameterReader)
	container.Provide(impl.NewParameterValidator)
	container.Provide(impl.NewCommandRunner)
	return container
}

func main() {
	//This code uses dig dependency injector for more details see
	//https://godoc.org/go.uber.org/dig
	//https://blog.drewolson.org/dependency-injection-in-go
	container := buildContainer()
	err := container.Invoke(func(runner impl.CommandRunnerAPI) {
		exitCode, OutputMessage := runner.Run()
		if exitCode == 0 {
			fmt.Println(OutputMessage)
		} else {
			os.Exit(exitCode)
		}
	})
	if err != nil {
		fmt.Println("General error executing command")
	}
}
