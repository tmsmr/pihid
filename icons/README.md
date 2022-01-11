# PiHID Icons
This package contains generated/transformed (Font Awesome-) icons expressed in `Icon` structs.

I figured this is better than embedding the whole TTF-File in a Go binary...

## `Icon` struct
Each instance contains the icon's alpha values for the resolutions 16x16, 32x32, 64x64. For example, a 16x16 icon is expressed as 256 (8-bit) alpha values.
```
type Icon struct {
	Alpha16x16 []uint8
	Alpha32x32 []uint8
	Alpha64x64 []uint8
}
```
Some graphic libraries and display drivers allow to write the pixel values for a rectangular sections en bloc, that's why the icon arrays are 1-dimensional:
```
| x0 | x1 | ... | xn |
|----|----|-----|----|
| →  | →  | →   | ↙  |
| →  | →  | →   | ↙  |
| →  | →  | →   | ↙  |
```
The order of the alpha values is row by row, for example, this icon:
```
| x0 | x1 |
|----|----|
| 4  | 7  |
| 1  | 1  |
```
would be `[]uint8{4, 7, 1, 1}`

## (Re-) Generating the icons
- `./generate_icons.sh`
- Requires Docker
- Check the source for details, it's only a couple LOC...

## Attribution
- The icons are generated from the SVG versions of https://fontawesome.com, thanks for that awesome icon collection!
