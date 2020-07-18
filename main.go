package main

import (
	"fmt"
	"time"
	//"strings"
	//"runtime"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
    FPS    =  60
    width  = 512
    height = 512
)

// main function
func main () {
    window := InitGlfw()
    defer glfw.Terminate()
    InitOpenGL()

    for !window.ShouldClose() {
		t := time.Now()

        glfw.PollEvents()
		time.Sleep(time.Second/time.Duration(FPS) - time.Since(t))
    }
}


// initializes glfw and returns a Window to use.
func InitGlfw () *glfw.Window {
    if err := glfw.Init(); err != nil {
		panic(err)
    }

    glfw.WindowHint(glfw.Resizable, glfw.True)
	// use OpenGL 4.6-core
    glfw.WindowHint(glfw.ContextVersionMajor, 4)
    glfw.WindowHint(glfw.ContextVersionMinor, 6)
    glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
    glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

    window, err := glfw.CreateWindow(width, height, "Vox-Legacy", nil, nil)
    if err != nil {
		panic(err)
    }
    window.MakeContextCurrent()
    return window
}


// initializes OpenGL
func InitOpenGL () {
    if err := gl.Init(); err != nil {
		panic(err)
    }
    version := gl.GoStr(gl.GetString(gl.VERSION))
    fmt.Println("OpenGL version", version)
}
