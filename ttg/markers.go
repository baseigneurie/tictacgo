package ttg

var XMarker, OMarker, NoMarker [5]string

func LoadMarkers() {
	x := &XMarker
	o := &OMarker

	x[0] = `X   X`
	x[1] = ` X X `
	x[2] = `  X  `
	x[3] = ` X X `
	x[4] = `X   X`

	o[0] = ` OOO `
	o[1] = `O   O`
	o[2] = `O   O`
	o[3] = `O   O`
	o[4] = ` OOO `
}
