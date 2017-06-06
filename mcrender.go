package mcrender

import (
	"bufio"
	"io"
	"regexp"
	"strconv"

	fgl "github.com/fogleman/fauxgl"
)

var (
	fillRxp = regexp.MustCompile(`fill ~(-*\d*) ~(-*\d*) ~(-*\d*) ~(-*\d*) ~(-*\d*) ~(-*\d*).*`)
)

type xyz struct {
	x int
	y int
	z int
}

func fillXVector(mesh *fgl.Mesh, a, b, y, z int) {
	for i := a; i <= b; i++ {
		c := fgl.NewCube()
		c.Transform(fgl.Translate(fgl.Vector{X: float64(i), Y: float64(y), Z: float64(z)}))
		mesh.Add(c)
	}
}

func fillBox(mesh *fgl.Mesh, c1, c2 xyz) {
	fromX, toX := c1.x, c2.x
	if c1.x > c2.x {
		fromX, toX = c2.x, c1.x
	}
	fromY, toY := c1.y, c2.y
	if c1.y > c2.y {
		fromY, toY = c2.y, c1.y
	}
	fromZ, toZ := c1.z, c2.z
	if c1.z > c2.z {
		fromZ, toZ = c2.z, c1.z
	}

	for i := fromY; i <= toY; i++ {
		for j := fromZ; j <= toZ; j++ {
			fillXVector(mesh, fromX, toX, i, j)
		}
	}
}

func CreateSTLFromInput(r io.Reader, stlname string) error {
	mesh := fgl.NewEmptyMesh()

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		match := fillRxp.FindStringSubmatch(scanner.Text())
		if match != nil {
			x, err := strconv.Atoi(match[1])
			if err != nil {
				return err
			}
			y, err := strconv.Atoi(match[2])
			if err != nil {
				return err
			}
			z, err := strconv.Atoi(match[3])
			if err != nil {
				return err
			}
			c1 := xyz{x: x, y: y, z: z}
			x, err = strconv.Atoi(match[4])
			if err != nil {
				return err
			}
			y, err = strconv.Atoi(match[5])
			if err != nil {
				return err
			}
			z, err = strconv.Atoi(match[6])
			if err != nil {
				return err
			}
			c2 := xyz{x: x, y: y, z: z}

			fillBox(mesh, c1, c2)
		}
	}

	fgl.SaveSTL(stlname, mesh)

	return nil
}
