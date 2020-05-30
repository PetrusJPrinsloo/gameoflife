package cell

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"gitlab.com/PetrusJPrinsloo/gameoflife/config"
	"gitlab.com/PetrusJPrinsloo/gameoflife/graphics"
	"gitlab.com/PetrusJPrinsloo/gameoflife/shape"
)

type Cell struct {
	Drawable uint32

	Alive     bool
	AliveNext bool

	X int
	Y int
}

func NewCell(x, y int, cnf *config.Config) *Cell {
	points := make([]float32, len(shape.Square), len(shape.Square))

	copy(points, shape.Square)
	for i := 0; i < len(points); i++ {
		var position float32
		var size float32
		switch i % 3 {
		case 0:
			size = 1.0 / float32(cnf.Columns)
			position = float32(x) * size
		case 1:
			size = 1.0 / float32(cnf.Rows)
			position = float32(y) * size
		default:
			continue
		}

		if points[i] < 0 {
			points[i] = (position * 2) - 1
		} else {
			points[i] = ((position + size) * 2) - 1
		}
	}

	return &Cell{
		Drawable: graphics.MakeVao(points),

		X: x,
		Y: y,
	}
}

func (c *Cell) CheckState(cells [][]*Cell) {
	c.Alive = c.AliveNext
	c.AliveNext = c.Alive

	liveCount := c.liveNeighbors(cells)

	if c.Alive {
		// 1. Any live Cell with fewer than two live neighbours dies, as if caused by underpopulation.
		if liveCount < 2 {
			c.AliveNext = false
		}

		// 2. Any live Cell with two or three live neighbours lives on to the next generation.
		if liveCount == 2 || liveCount == 3 {
			c.AliveNext = true
		}

		// 3. Any live Cell with more than three live neighbours dies, as if by overpopulation.
		if liveCount > 3 {
			c.AliveNext = false
		}
	} else {
		// 4. Any dead Cell with exactly three live neighbours becomes a live Cell, as if by reproduction.
		if liveCount == 3 {
			c.AliveNext = true
		}
	}
}

// liveNeighbors returns the number of live neighbors for a Cell.
func (c *Cell) liveNeighbors(cells [][]*Cell) int {
	var liveCount int
	add := func(x, y int) {
		// If we're at an edge, check the other side of the board.
		if x == len(cells) {
			x = 0
		} else if x == -1 {
			x = len(cells) - 1
		}
		if y == len(cells[x]) {
			y = 0
		} else if y == -1 {
			y = len(cells[x]) - 1
		}

		if cells[x][y].Alive {
			liveCount++
		}
	}

	add(c.X-1, c.Y)   // To the left
	add(c.X+1, c.Y)   // To the right
	add(c.X, c.Y+1)   // up
	add(c.X, c.Y-1)   // down
	add(c.X-1, c.Y+1) // top-left
	add(c.X+1, c.Y+1) // top-right
	add(c.X-1, c.Y-1) // bottom-left
	add(c.X+1, c.Y-1) // bottom-right

	return liveCount
}

func (c *Cell) Draw() {
	if !c.Alive {
		return
	}

	gl.BindVertexArray(c.Drawable)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(shape.Square)/3))
}
