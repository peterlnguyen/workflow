package workflow

type Workflow struct {
	Name    string
	Version int
	Arguments
	State
	Steps
}

type Steps struct {
}

type Arguments struct {
}

type State struct {
}
