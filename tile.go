package main

/*
    Documentation: http://docs.gl/
    https://youtu.be/x0H--CL2tUI
*/

// use OpenGL 4.6
import (
    "fmt"
    "encoding/hex"
    "github.com/go-gl/gl/v4.6-core/gl"
	//"github.com/go-gl/glfw/v3.3/glfw"
)


// specify array size and quantities
const (
    nbRows        = 8 * 8
    nbVoxs        = 8 * 8 * 8
    nbFaces4Vox   = 6 // 1 voxel = 6 faces
    nbVerts4Face  = 6 // 1 quad  = 2 tris = 6 verts
    nbCoords4Vert = 4 // x,y,z,i
    sizeOfCoord   = 1 // 1 coord = 1 byte
)


// 3D tile made of 4 colors
type Tile struct {
    rows [nbRows]uint16
    VBO  uint32
}


// load voxels from a HEX string
func (tile Tile) LoadHEX (data string) error {
    if len(data) != nbRows * 4 { // 4 HEX characters per row of the tile
        return fmt.Errorf(
            "Cannot construct tile: expecting %d given %d", nbRows * 4, len(data))
    }

    array, err := hex.DecodeString(data)
    if err != nil {return err}

    // two bytes are read for every rows
    for i := 0; i < nbRows; i += 1 {
        high := uint16(array[i * 2    ])
        low  := uint16(array[i * 2 + 1])
        tile.rows[i] = (high << 8) | low
    }
    return nil
}


// build mesh from pixel data
func (tile Tile) MakeMesh () uint32 {
    // delete the previously assigned buffers
    gl.DeleteBuffers(2, &tile.VBO)

    // struct helper
    type faceDef struct {
        ix, iy, iz int
        face []uint
    }

    // constant helpers
    var defs = [...]faceDef {
        { 0, 0,-1, faceFront [:]},
        { 0,-1, 0, faceTop   [:]},
        {-1, 0, 0, faceLeft  [:]},
        { 0, 0, 1, faceBack  [:]},
        { 0, 1, 0, faceBottom[:]},
        { 1, 0, 0, faceRight [:]}}

    const (
        nbComp    = nbVerts4Face * nbCoords4Vert
        arraySize = nbVoxs * nbFaces4Vox * nbVerts4Face * nbCoords4Vert
    )

    // buffers to fill in
    var (
        vertices [arraySize]uint8 // buffer of vertices
        countVerts = 0 // count the number of vertices in the buffer
    )

    // iterate over each pixel
    for z := 0; z < 8; z += 1 {
    for y := 0; y < 8; y += 1 {
    for x := 0; x < 8; x += 1 {

        // nothing to do if the current pixel is empty
        color := tile.GetVoxel(x, y, z)
        if color != 0 {

            // for each face of the pixel
            for _, def := range defs {

                // if the neighboring pixel is empty, create a new face
                if tile.GetVoxel(x + def.ix, y + def.iy, z + def.iz) == 0 {

                    // copy vertices with an offset
                    s := countVerts * nbCoords4Vert
                    for i := 0; i < nbComp; i += nbCoords4Vert {
                        vertices[s + i    ] = uint8(def.face[i    ] + uint(x))
                        vertices[s + i + 1] = uint8(def.face[i + 1] + uint(y))
                        vertices[s + i + 2] = uint8(def.face[i + 2] + uint(z))
                        vertices[s + i + 3] = uint8(color)
                        countVerts += 1
                    }
                }
            }
        }
    }}}

    const sizeOfVert = nbCoords4Vert * sizeOfCoord
    bufferVerts := vertices[:(countVerts * nbCoords4Vert)]

    // once the arrays have been filled we can make VBOs and a VAO
    gl.GenBuffers(1, &tile.VBO)

    gl.BindBuffer(gl.ARRAY_BUFFER, tile.VBO)
    gl.BufferData(gl.ARRAY_BUFFER, countVerts * sizeOfVert, gl.Ptr(bufferVerts), gl.STATIC_DRAW)
    gl.EnableVertexAttribArray(0)
    gl.VertexAttribPointer(0, 3, gl.UNSIGNED_BYTE, false, sizeOfVert, nil)
    gl.VertexAttribPointer(1, 1, gl.UNSIGNED_BYTE, false, sizeOfVert, nil)

    gl.BindBuffer(gl.ARRAY_BUFFER, 0)

    return tile.VBO
}


// get the pixel at specified location
func (tile Tile) GetVoxel (x, y, z int) uint {
    // if the pixel is out of bounds, return 0
    if x < 0 || 8 <= x || y < 0 || 8 <= y || z < 0 || 8 <= z {return uint(0)}
    row := tile.rows[y + z * 8] // planes along z
    shf := 14 - x * 2           // reverse indexing 7 -> 0
    return uint((row >> shf) & 0b11) // shift and mask relevant bits
}


// declare arrays for placing faces
var (
    //                     /* triangle 1    */     /* triangle 2    */
    faceFront  = [...]uint {0,0,0, 0,1,0, 1,0,0,    1,1,0, 1,0,0, 0,1,0}
    faceTop    = [...]uint {0,0,0, 1,0,0, 0,0,1,    1,0,1, 0,0,1, 1,0,0}
    faceLeft   = [...]uint {0,0,0, 0,0,1, 0,1,0,    0,1,1, 0,1,0, 0,0,1}
    faceBack   = [...]uint {1,1,1, 0,1,1, 1,0,1,    0,0,1, 1,0,1, 0,1,1}
    faceBottom = [...]uint {1,1,1, 1,1,0, 0,1,1,    0,1,0, 0,1,1, 1,1,0}
    faceRight  = [...]uint {1,1,1, 1,0,1, 1,1,0,    1,0,0, 1,1,0, 1,0,1}
)
