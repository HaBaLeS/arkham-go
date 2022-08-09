package arkham_game

import "log"

func CreateEnemyPhase() *ArkhamPhase {
	return &ArkhamPhase{
		name:     "Enemy",
		execfunc: DoEnemyPhase,
	}
}

func DoEnemyPhase() {
	log.Printf("\t Enemys Move")
	log.Printf("\t Enemys Attack")
}
