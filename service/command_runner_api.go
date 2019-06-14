package service

//CommandRunnerAPI defines api to process command
type CommandRunnerAPI interface {
	Run() (int, string)
}
