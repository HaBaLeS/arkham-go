module arkham-go

go 1.18


require github.com/HaBaLeS/arkham-go/modules/gpbge v0.0.1
require github.com/HaBaLeS/arkham-go/modules/arkham-game v0.0.1

replace (
	github.com/HaBaLeS/arkham-go/modules/gpbge v0.0.1 => ../arkham-go/modules/gpbge
	github.com/HaBaLeS/arkham-go/modules/arkham-game v0.0.1 => ../arkham-go/modules/arkham-game
)


