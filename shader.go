package main

/*
    Sprites composed of a tile and a palette
    Can be freely place in the world
*/

// use OpenGL 4.6
import (
    "fmt"
    //"log"
    "strings"
    "io/ioutil"
	"github.com/go-gl/gl/v4.6-core/gl"
	//"github.com/go-gl/glfw/v3.3/glfw"
)


func CreateProgram () uint32 {
	vertShader, err := CompileShader("shaders/vertex.shader", gl.VERTEX_SHADER)
	if err != nil { panic(err) }
	fragShader, err := CompileShader("shaders/fragment.shader", gl.FRAGMENT_SHADER)
	if err != nil { panic(err) }

    program := gl.CreateProgram()
	gl.AttachShader(program, vertShader)
	gl.AttachShader(program, fragShader)
    gl.LinkProgram(program)
    return program
}

// compile the shader and return a pointer to it
func CompileShader (filepath string, shaderType uint32) (uint32, error) {
    source, err := ioutil.ReadFile(filepath)
    if err != nil { panic(err) }

	shader := gl.CreateShader(shaderType)
	compSrc, free := gl.Strs(string(source))
	gl.ShaderSource(shader, 1, compSrc, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength + 1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}
	return shader, nil
}
