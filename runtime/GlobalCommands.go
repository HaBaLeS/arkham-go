package runtime

func startScenario(ps *PlaySession, args []string) {

}

func exitScenario(ps *PlaySession, args []string) {
	println("Terminating Game")
	ps.currentState = StateEnd
}
