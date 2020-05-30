package main

import (
	"github.com/PetrusJPrinsloo/gameoflife/cell"
	"github.com/PetrusJPrinsloo/gameoflife/config"
	"github.com/PetrusJPrinsloo/gameoflife/graphics"
	"io/ioutil"
	"log"
	"math/rand"
	"runtime"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

var (
	cnf *config.Config
)

func main() {
	cnf = config.ReadFile("default.json")
	vertexShaderSource := getShaderFileContents("resources\\shaders\\vertex\\shader.glsl")
	fragmentShaderSource := getShaderFileContents("resources\\shaders\\fragment\\shader.glsl")

	runtime.LockOSThread()

	window := graphics.InitGlfw(cnf)
	defer glfw.Terminate()
	program := graphics.InitOpenGL(vertexShaderSource, fragmentShaderSource)

	cells := makeCells()

	for !window.ShouldClose() {
		t := time.Now()

		for x := range cells {
			for _, c := range cells[x] {
				c.CheckState(cells)
			}
		}

		draw(cells, window, program)
		time.Sleep(time.Second/time.Duration(cnf.Fps) - time.Since(t))
	}
}

func getShaderFileContents(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string
	text := string(content)
	return text
}

// loop over cells and tell them to draw
func draw(cells [][]*cell.Cell, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	for x := range cells {
		for _, c := range cells[x] {
			c.Draw()
		}
	}

	glfw.PollEvents()
	window.SwapBuffers()
}

// Set up the game board
func makeCells() [][]*cell.Cell {
	rand.Seed(time.Now().UnixNano())
	cells := make([][]*cell.Cell, cnf.Rows, cnf.Columns)

	for x := 0; x < cnf.Rows; x++ {
		for y := 0; y < cnf.Columns; y++ {
			c := cell.NewCell(x, y, cnf)

			c.Alive = rand.Float64() < cnf.Threshold
			c.AliveNext = c.Alive

			cells[x] = append(cells[x], c)
		}
	}

	return cells
}
