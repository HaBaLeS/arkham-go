package engine

func startScenario(engine *Engine, args []string) {

}

func exitScenario(engine *Engine, args []string) {
	println("Terminating Game")
	engine.currentState = StateEnd
}
