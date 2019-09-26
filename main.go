package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"./generators"
	"./glTools"
	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

var (
	triangle = []float32{
		-1, -1, 0,
		-1, 1, 0,
		1, 1, 0,
		1, -1, 0,
	}
	triangleData = []float32{
		-1., -1., 0., 0., 0.,
		-1., 1., 0., 0., 1.,
		1., 1., 0., 1., 1.,

		1., -1., 0., 1., 0.,
		-1., -1., 0., 0., 0.,
		1., 1., 0., 1., 1.,
	}
	window  *glfw.Window
	program *glTools.Program
	vao     uint32
	texture *glTools.Texture

	width           = 2048
	height          = 2048
	seed            = rand.Intn(math.MaxInt16)
	vizEnabled      = false
	outputFilename  = "out.png"
	backgroundColor = "black"
	colors          = "rainbow"
	generatorMethod = generators.BasicDiagonalLine
	space           = 8
)

func fillImage(img *image.RGBA64, color color.RGBA64) {
	bounds := img.Bounds()

	for y := bounds.Min.Y; y <= bounds.Max.Y; y++ {
		for x := bounds.Min.X; x <= bounds.Max.X; x++ {
			img.Set(x, y, color)
		}
	}

	generators.SetBaseColor(color)
}

func saveImage(img *image.RGBA64, name string) {
	f, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func setVisualizer(state bool) {
	vizEnabled = state
}

func setWidth(value int) {
	width = value
}

func setHeight(value int) {
	height = value
}

func setSeed(value int) {
	seed = value
}

func setBackgroundColor(value string) {
	backgroundColor = value
}

func setGenerator(value func()) {
	generatorMethod = value
}

func setSpace(value int) {
	space = value
}

func setColors(value string) {
	colors = value

	generators.SetParamColors(generators.ColorsPacks[colors])
}

func valueToInt(value string) int {
	v, err := strconv.Atoi(value)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	return v
}

func applyArgWithValue(arg string, value string) {
	switch arg {
	case "--width":
		setWidth(valueToInt(value))
	case "--height":
		setHeight(valueToInt(value))
	case "--seed":
		setSeed(valueToInt(value))
	case "--bg":
		setBackgroundColor(value)
	case "--colors":
		setColors(value)
	case "--p1":
		generators.SetParam1(valueToInt(value))
	case "--p2":
		generators.SetParam2(valueToInt(value))
	case "--cycles":
		generators.SetCycles(valueToInt(value))
	}
}

func applyArg(arg string) {
	switch arg {
	case "-v":
		setVisualizer(true)
	case "-bdl":
		setGenerator(generators.BasicDiagonalLine)
	case "-bdl2":
		setGenerator(generators.BiSymDiagonalLine)
	case "-bdl4":
		setGenerator(generators.QuadSymDiagonalLine)
	case "-bdg":
		setGenerator(generators.BasicDiamondGrid)
	case "-bppp":
		setGenerator(generators.BasicPointPerPoint)
	case "-ibdl":
		setGenerator(generators.InversedSurroundedDiagonalLine)
	case "-ibdl2":
		setGenerator(generators.BiSymInversedSurroundedDiagonalLine)
	case "-ibdl4":
		setGenerator(generators.QuadSymInversedSurroundedDiagonalLine)
	case "-ths":
		setGenerator(generators.TurningHazardousShape)
	}
}

func applyArgs(args []string) {
	for i := range args {
		if strings.HasPrefix(args[i], "--") && i+1 < len(args) {
			applyArgWithValue(args[i], args[i+1])
			i++
		} else if strings.HasPrefix(args[i], "-") {
			applyArg(args[i])
		} else if strings.HasSuffix(args[i], ".png") {
			outputFilename = args[i]
		}
	}
}

func main() {
	args := os.Args[1:]

	applyArgs(args)

	rand.Seed(time.Now().UTC().UnixNano())
	seed = rand.Intn(math.MaxInt16)

	log.Println("Seed :", seed)

	img := image.NewRGBA64(image.Rect(0, 0, width, height))

	if vizEnabled {
		runtime.LockOSThread()

		window = initGlfw()

		program = initOpenGL()
		if program == nil {
			return
		}

		vao = makeVao(triangle)
		texture, _ = glTools.NewTexture(img, gl.CLAMP_TO_EDGE, gl.CLAMP_TO_EDGE)

		generators.SetDrawer(draw)
	}

	if color, ok := generators.Colors[backgroundColor]; ok {
		fillImage(img, color)
	} else {
		log.Println("Error : background color not found. Back to black.")
		fillImage(img, generators.Colors["black"])
	}

	generators.SetImage(img)
	generators.SetNumberSeed(seed)

	generatorMethod()
	//generators.AutomoveCenterLines(64, 4, color.RGBA64{0x0120, 0x0060, 0x0200, 0xFFFF})
	//generators.QuadSymInversedSurroundedDiagonalLine(space, generators.ColorsPacks[colors])
	log.Println("> End")

	saveImage(img, outputFilename)

	for vizEnabled && !window.ShouldClose() {
		draw(img)
	}
}

// initGlfw initializes glfw and returns a Window to use.
func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4) // OR 2
	glfw.WindowHint(glfw.ContextVersionMinor, 5)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width/2, height/2, "Kirina", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

// initOpenGL initializes OpenGL and returns an intiialized program.
func initOpenGL() *glTools.Program {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	vertShader, err := glTools.NewShaderFromFile("gl.vert", gl.VERTEX_SHADER)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	fragShader, err := glTools.NewShaderFromFile("gl.frag", gl.FRAGMENT_SHADER)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	prog, err := glTools.NewProgram(vertShader, fragShader)

	return prog
}

func draw(img *image.RGBA64) {
	texture.Update(img)

	gl.Uniform1i(program.GetUniformLocation("tex"), 0)

	program.BindFragDataLocation(0, "outputColor")

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	program.Use()

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	gl.BindVertexArray(vao)

	texture.Bind(gl.TEXTURE0)

	gl.DrawArrays(gl.TRIANGLES, 0, 6)

	window.SwapBuffers()
	glfw.PollEvents()
}

// makeVao initializes and returns a vertex array from the points provided.
func makeVao(points []float32) uint32 {
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	gl.BufferData(gl.ARRAY_BUFFER, len(triangleData)*4, gl.Ptr(triangleData), gl.STATIC_DRAW)

	vertAttr := uint32(program.GetAttribLocation("vert"))
	gl.EnableVertexAttribArray(vertAttr)
	gl.VertexAttribPointer(vertAttr, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))

	texCoord := uint32(program.GetAttribLocation("vertTexCoord"))
	gl.EnableVertexAttribArray(texCoord)
	gl.VertexAttribPointer(texCoord, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

	return vao
}
