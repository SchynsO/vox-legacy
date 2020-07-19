package main


const (
    mapSize = 16
    nbTiles = mapSize * mapSize * mapSize
)


type TileMap struct {
    tils [nbTiles]uint8 // [8bits] Tile to place
    rots [nbTiles]uint8 // [6bits] Is the tile rotated
    mirs [nbTiles]uint8 // [3bits] Is the tile mirrored
    pals [nbTiles]uint8 // [2bits] Palette to use
}


// Set the tile in the tile map from 3 bytes
func (tm TileMap) Set (index uint, til, rot, mir, pal uint8) {
    tm.tils[index] = til
    tm.rots[index] = rot
    tm.mirs[index] = mir
    tm.pals[index] = pal
}


// Get the tile from the tile map
func (tm TileMap) Get (index uint) (uint8, uint8, uint8, uint8) {
    return tm.tils[index], tm.rots[index], tm.mirs[index], tm.pals[index]
}


// render the tile maps in OpenGL (having 4096 draw calls per frame is acceptable)
func DrawTileMaps (map1, map2 *TileMap, scroll Vector3) {
    var brush Sprite // use a sprite as a brush

    // draw tiles from the tile map based on the scrolling
    var index Vector3
    for ix := uint(0); ix < mapSize; ix += 1 {
    for iy := uint(0); iy < mapSize; iy += 1 {
    for iz := uint(0); iz < mapSize; iz += 1 {

        // find the tile data to display
        index.Set(ix, iy, iz)
        i, m := findTileWithScroll(index, scroll)
        map0 := map1; if m {map0 = map2}
        til, rot, mir, pal := map0.Get(i)

        // if tile 0 there is nothing to do
        if til != 0 {
            brush.SetTile   (uint(til), nil) // TODO get tile from a list
            brush.SetPalette(uint(pal), nil) // TODO get palette from a list
            brush.rot.SetByte(rot)
            brush.mir.SetByte(mir)
            brush.pos = index.ShiftL(3).Add(scroll.Mask(0x7))

            brush.Draw()
        }
    }}}
}


func findTileWithScroll (index, scroll Vector3) (uint, bool) {
    pos := index.Sub(scroll.ShiftR(3)).Mod(mapSize * 2)

    // identify the map to use based on crossboard pattern
    msk := 0
    if pos.x >= mapSize {msk += 1}
    if pos.y >= mapSize {msk += 1}
    if pos.z >= mapSize {msk += 1}

    // find index in the map
    x, y, z := pos.Mod(mapSize).Get()
    return z << 8 | y << 4 | x, (msk & 0x1) != 0
}
