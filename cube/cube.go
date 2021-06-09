package cube

import "fmt"

var sideToColorMap = map[string]string{
	"WHITE":  "U",
	"YELLOW": "D",
	"GREEN":  "R",
	"BLUE":   "L",
	"RED":    "B",
	"ORANGE": "F",
}

// square represent each square in the rubix cube
//  	id represent the exact location on an initial/solved rubix cube
//		color represent the color of the square
type square struct {
	color string
	idx   int
}

// side is a two dementional array
type side [][]square

func createSide(n int, color string) side {
	// TODO: OPTIMIZE value IDX_
	var s = make(side, n) // make a n row matrix
	idx_ := 1
	for row := 0; row < n; row++ {
		var r = make([]square, n) //make a single 1 * n col array
		for col := 0; col < n; col++ {
			r[col] = square{
				color: color,
				idx:   idx_,
			}
			idx_ = idx_ + 1
		}
		s[row] = r
	}
	return s
}

func (s side) isSolved() bool {
	/**
	*Return true if all square is of same color, else false
	**/

	color := s[0][0].color

	for _, squares := range s {
		for _, sq := range squares {
			if sq.color != color {
				return false
			}
		}
	}
	return true
}

// cube will be a map of side with key string representing the side (U,D,L,R,F,B)
type cube map[string]side

func createCube(n int) cube {
	/**
	*Return n x n cube
	**/
	c := make(cube, n)
	for colorLabel, sideLabe := range sideToColorMap {
		c[sideLabe] = createSide(n, colorLabel)
	}

	return c
}

func createCubeFromFile(filename string) cube {
	//TODO: Implement this
	return nil
}

func (c cube) printToScreen() {
	fmt.Println(c)
}

//TODO: Create function that compare two cube and see if they are equal

func (c *cube) isSolved() bool {

	for _, s := range *c {
		if !s.isSolved() {
			return false
		}
	}

	return true
}

func (c *cube) scramble() {
	//TODO: Implement this
}

func (c cube) rotate(opp string) {
	switch opp {
	case "U":
		rotateU(c)
	case "XU":
		rotateXU(c)
	case "D":
		rotateD(c)
	case "XD":
		rotateXD(c)
	case "R":
		rotateR(c)
	case "XR":
		rotateXR(c)
	case "L":
		rotateL(c)
	case "XL":
		rotateXL(c)
	case "B":
		rotateB(c)
	case "XB":
		rotateXB(c)
	case "F":
		rotateF(c)
	case "XF":
		rotateXF(c)
	default:

	}
}

func rotateU(c cube) {
	c["F"][0], c["R"][0], c["B"][0], c["L"][0] = c["R"][0], c["B"][0], c["L"][0], c["F"][0]
}

func rotateXU(c cube) {
	c["F"][0], c["L"][0], c["B"][0], c["R"][0] = c["L"][0], c["B"][0], c["R"][0], c["F"][0]
}

func rotateD(c cube) {
	n := len(c["F"][0]) - 1
	c["F"][n], c["R"][n], c["B"][n], c["L"][n] = c["R"][n], c["B"][n], c["L"][n], c["F"][n]
}

func rotateXD(c cube) {
	n := len(c["F"][0]) - 1
	c["F"][n], c["L"][n], c["B"][n], c["R"][n] = c["L"][n], c["B"][n], c["R"][n], c["F"][n]
}

func rotateR(c cube) {
	n := len(c["F"][0]) - 1
	for i := 0; i <= n; i++ {
		c["F"][i][n], c["D"][i][n], c["B"][n-i][0], c["U"][i][n] = c["D"][i][n], c["B"][n-i][0], c["U"][i][n], c["F"][i][n]
	}
}

func rotateXR(c cube) {
	n := len(c["F"][0]) - 1
	for i := 0; i <= n; i++ {
		c["F"][i][n], c["U"][i][n], c["B"][n-i][0], c["D"][i][n] = c["U"][i][n], c["B"][n-i][0], c["D"][i][n], c["F"][i][n]
	}
}

func rotateL(c cube) {
	n := len(c["F"][0]) - 1
	for i := 0; i <= n; i++ {
		c["F"][i][0], c["U"][i][0], c["B"][n-i][n], c["D"][i][0] = c["U"][i][0], c["B"][n-i][n], c["D"][i][0], c["F"][i][0]
	}
}

func rotateXL(c cube) {
	n := len(c["F"][0]) - 1
	for i := 0; i <= n; i++ {
		c["F"][i][0], c["D"][i][0], c["B"][n-i][n], c["U"][i][0] = c["D"][i][0], c["B"][n-i][n], c["U"][i][0], c["F"][i][0]
	}
}
func rotateB(c cube) {
	n := len(c["F"][0]) - 1
	for i := 0; i <= n; i++ {
		c["U"][0][i], c["R"][i][n], c["D"][n][n-i], c["L"][n-i][0] = c["R"][i][n], c["D"][n][n-i], c["L"][n-i][0], c["U"][0][i]

	}
}

func rotateXB(c cube) {
	n := len(c["F"][0]) - 1
	for i := 0; i <= n; i++ {
		c["U"][0][i], c["L"][n-i][0], c["D"][n][n-i], c["R"][i][n] = c["L"][n-i][0], c["D"][n][n-i], c["R"][i][n], c["U"][0][i]
	}
}

func rotateF(c cube) {
	n := len(c["F"][0]) - 1
	for i := 0; i <= n; i++ {
		c["U"][n][i], c["L"][n-i][n], c["D"][0][n-i], c["R"][i][0] = c["L"][n-i][n], c["D"][0][n-i], c["R"][i][0], c["U"][n][i]
	}
}

func rotateXF(c cube) {
	n := len(c["F"][0]) - 1
	for i := 0; i <= n; i++ {
		c["U"][n][i], c["R"][i][0], c["D"][0][n-i], c["L"][n-i][n] = c["R"][i][0], c["D"][0][n-i], c["L"][n-i][n], c["U"][n][i]
	}
}
