package n0489

type Robot struct{}

func (robot *Robot) Move() bool { panic("implement from leetcode") }
func (robot *Robot) TurnLeft()  {}
func (robot *Robot) TurnRight() {}
func (robot *Robot) Clean()     {}

type Pair struct{ x, y int }

func cleanRoom(robot *Robot) {
	visited := make(map[Pair]int)
	directs := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	var backtrack func(x, y, d int)

	backtrack = func(x, y, direct int) {
		visited[Pair{x, y}]++
		robot.Clean()
		for i := 0; i < 4; i++ {
			newD := (direct + i) % 4
			newX := x + directs[newD][0]
			newY := y + directs[newD][1]

			if _, ok := visited[Pair{newX, newY}]; !ok && robot.Move() {
				backtrack(newX, newY, newD)
				goBack(robot)
			}

			robot.TurnRight()
		}
	}

	backtrack(0, 0, 0)
}

func goBack(robot *Robot) {
	robot.TurnRight()
	robot.TurnRight()
	robot.Move()
	robot.TurnRight()
	robot.TurnRight()
}
