package cube

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSide(t *testing.T) {
	s := createSide(3, "GREEN")
	assert.Equal(t, s[0][0].color, "GREEN")
	if !s.isSolved() {
		t.Errorf("side Creation failed")
	}

}

func TestCreateCube(t *testing.T) {
	c := createCube(3)
	for sideLabel, side := range c {
		if !side.isSolved() {
			t.Errorf("Cube Creation failed for sideLabe: %v", sideLabel)
			return
		}
	}

}

func TestRotateU(t *testing.T) {
	c := createCube(3)
	c.rotate("U")
	c.rotate("XU")
	if !c.isSolved() {
		t.Errorf("Cube Rotation failed for sideLabe: U")
		return
	}

}

func TestRotateD(t *testing.T) {
	c := createCube(3)
	c.rotate("D")
	c.rotate("XD")
	if !c.isSolved() {
		t.Errorf("Cube Rotation failed for sideLabe: D")
		return
	}
}

func TestRotateR(t *testing.T) {
	c := createCube(3)
	c.rotate("R")
	c.rotate("XR")
	if !c.isSolved() {
		t.Errorf("Cube Rotation failed for sideLabe: R")
		return
	}
}

func TestRotateL(t *testing.T) {
	c := createCube(3)
	c.rotate("L")
	c.rotate("XL")
	if !c.isSolved() {
		t.Errorf("Cube Rotation failed for sideLabe: L")
		return
	}
}

func TestRotateB(t *testing.T) {
	c := createCube(3)
	c.rotate("B")
	c.rotate("XB")
	if !c.isSolved() {
		t.Errorf("Cube Rotation failed for sideLabe: B")
		return
	}
}

func TestRotateF(t *testing.T) {
	c := createCube(3)
	c.rotate("F")
	c.rotate("XF")
	if !c.isSolved() {
		t.Errorf("Cube Rotation failed for sideLabe: B")
		return
	}
}
