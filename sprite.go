package main

/*
    Sprites composed of a tile and a palette
    Can be freely place in the world
*/

// use OpenGL 4.6
import (
    //"github.com/go-gl/gl/v4.6-core/gl"
	//"github.com/go-gl/glfw/v3.3/glfw"
)


// Sprite to place on the screen
type Sprite struct {
    // indices of tile and palette
    id_tile uint
    id_pal  uint

    // transforms to apply to the mesh
    pos Vector3 // position of the mesh
    rot Byte3   // rotate the mesh
    mir Bool3   // flip the mesh

    // OpenGL components to draw the associated mesh
    vbo  uint32   // Vertex Buffer Object
    pal *Palette  // pointer to the palette used
}


// draw the sprite in the scene with provided parameters
func (sprite Sprite) Draw () {
    sprite.pal.Use()

    // TODO
}


// set a new tile to this sprite
func (sprite Sprite) SetTile (id uint, tile *Tile) {
    sprite.id_tile = id
    sprite.vbo = tile.VBO
}


// set a new palette to this sprite
func (sprite Sprite) SetPalette (id uint, palette *Palette) {
    sprite.id_pal = id
    sprite.pal    = palette
}
