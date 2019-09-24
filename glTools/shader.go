package glTools

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/go-gl/gl/v4.5-core/gl"
)

// Shader is a shader struct
type Shader struct {
	handle uint32
}

// Program is a structure of a program
type Program struct {
	handle  uint32
	shaders []*Shader
}

// Delete remove a shader
func (shader *Shader) Delete() {
	gl.DeleteShader(shader.handle)
}

// Delete remove a program
func (prog *Program) Delete() {
	for _, shader := range prog.shaders {
		shader.Delete()
	}
	gl.DeleteProgram(prog.handle)
}

// Attach attach a shader to a program
func (prog *Program) Attach(shaders ...*Shader) {
	for _, shader := range shaders {
		gl.AttachShader(prog.handle, shader.handle)
		prog.shaders = append(prog.shaders, shader)
	}
}

// Use enable a program
func (prog *Program) Use() {
	gl.UseProgram(prog.handle)
}

// Link create a link between OpenGL and the program
func (prog *Program) Link() error {
	gl.LinkProgram(prog.handle)
	return getGlError(prog.handle, gl.LINK_STATUS, gl.GetProgramiv, gl.GetProgramInfoLog,
		"PROGRAM::LINKING_FAILURE")
}

// GetUniformLocation get location of a program
func (prog *Program) GetUniformLocation(name string) int32 {
	return gl.GetUniformLocation(prog.handle, gl.Str(name+"\x00"))
}

// GetAttribLocation get location of a program
func (prog *Program) GetAttribLocation(name string) int32 {
	return gl.GetAttribLocation(prog.handle, gl.Str(name+"\x00"))
}

// BindFragDataLocation binds
func (prog *Program) BindFragDataLocation(n uint32, name string) {
	gl.BindFragDataLocation(prog.handle, n, gl.Str(name+"\x00"))
}

// NewProgram create a new program
func NewProgram(shaders ...*Shader) (*Program, error) {
	prog := &Program{handle: gl.CreateProgram()}
	prog.Attach(shaders...)

	if err := prog.Link(); err != nil {
		return nil, err
	}

	return prog, nil
}

// NewShader create a new shader
func NewShader(src string, sType uint32) (*Shader, error) {

	handle := gl.CreateShader(sType)
	glSrc := gl.Str(src + "\x00")
	gl.ShaderSource(handle, 1, &glSrc, nil)
	gl.CompileShader(handle)
	err := getGlError(handle, gl.COMPILE_STATUS, gl.GetShaderiv, gl.GetShaderInfoLog,
		"SHADER::COMPILE_FAILURE::")
	if err != nil {
		return nil, err
	}
	return &Shader{handle: handle}, nil
}

// NewShaderFromFile create a shader from a file
func NewShaderFromFile(file string, sType uint32) (*Shader, error) {
	src, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	handle := gl.CreateShader(sType)
	csources, free := gl.Strs(string(src) + "\x00")
	gl.ShaderSource(handle, 1, csources, nil)

	free()

	gl.CompileShader(handle)
	err = getGlError(handle, gl.COMPILE_STATUS, gl.GetShaderiv, gl.GetShaderInfoLog,
		"SHADER::COMPILE_FAILURE::"+file)
	if err != nil {
		return nil, err
	}
	return &Shader{handle: handle}, nil
}

type getObjIv func(uint32, uint32, *int32)
type getObjInfoLog func(uint32, int32, *int32, *uint8)

func getGlError(glHandle uint32, checkTrueParam uint32, getObjIvFn getObjIv,
	getObjInfoLogFn getObjInfoLog, failMsg string) error {

	var success int32
	getObjIvFn(glHandle, checkTrueParam, &success)

	if success != gl.TRUE {
		var logLength int32
		getObjIvFn(glHandle, gl.INFO_LOG_LENGTH, &logLength)

		log := gl.Str(strings.Repeat("\x00", int(logLength+1)))
		getObjInfoLogFn(glHandle, logLength, nil, log)

		return fmt.Errorf("%s: %s", failMsg, gl.GoStr(log))
	}

	return nil
}
