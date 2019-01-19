package ttg

var dashboard string

func DrawDash() string {
	dashboard = "Total Moves: XX   "
	dashboard += "Current Player: X\n\n"
	dashboard = Yellow(dashboard)

	return dashboard
}
