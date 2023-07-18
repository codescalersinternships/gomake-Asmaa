package internal

type App struct {
	graph *Graph
}

func NewApp() *App {
	return &App{newGrap()}
}

// Run runs all functions and check if there is an error
func (app *App) Run(filePath string, target string) error {
	if len(target) == 0 {
		return ErrorNoTarget
	}
	var err error
	app.graph, err = ParseMakefile(filePath)
	if err != nil {
		return err
	}

	err = app.graph.CheckNoCommands()
	if err != nil {
		return err
	}

	err = app.graph.CheckCircularDependencies()
	if err != nil {
		return err
	}

	err = app.graph.RunTarget(target)
	if err != nil {
		return err
	}
	return nil
}
