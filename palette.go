package main

/*

*/

// use OpenGL 4.6
import (
    "fmt"
    "encoding/hex"
    //"github.com/go-gl/gl/v4.6-core/gl"
	//"github.com/go-gl/glfw/v3.3/glfw"
)


// specify array size and quantities
const (
    nbColors      = 64
    nbColors4Pal  =  3
    nbComps4Color =  3
    nbComps       = nbColors * nbComps4Color
)


// 3D tile made of 4 colors
type Palette struct {
    indices [nbColors4Pal                ]uint
    colors  [nbColors4Pal * nbComps4Color]float32
}


// set the colors of this palette
func (palette Palette) SetColor (index, color uint) {
    palette.indices[index] = color

    const nb = uint(nbComps4Color)
    for i := uint(0); i < nb; i += 1 {
        palette.colors[index * nb + i] = Colors[color * nb + i]
    }
}


// set uniforms to use this palette
func (palette Palette) Use () {
    // TODO
}


// list all usable colors here
var Colors = [nbComps]float32 {}


// Load usable colors from an HEX string
//( colors can be defined from a config file without recompiling the program )
func LoadColors (data string) error {
    if len(data) != nbComps * 2 { // 6 HEX characters per colors
        return fmt.Errorf(
            "Cannot construct colors: expecting %d given %d", nbComps * 2, len(data))
    }

    array, err := hex.DecodeString(data)
    if err != nil {return err}

    // one bytes is read for every color component
    for i := 0; i < nbComps; i += 1 {
        Colors[i] = float32(array[i]) / 255.0
    }
    return nil
}
