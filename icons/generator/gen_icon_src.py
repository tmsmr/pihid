import sys

from string import Template

from wand.image import Image
from wand.color import Color

fname = sys.argv[1]
name = sys.argv[2]

t = Template("""package icons

var $name = Icon{
	Alpha16x16: []uint8{$values_16},
	Alpha32x32: []uint8{$values_32},
	Alpha64x64: []uint8{$values_64},
}
""")

values = {16: None, 32: None, 64: None}

with Image(filename=fname, background=Color('transparent')) as svg:
    for res in values:
        with svg.clone() as img:
            img.resize(res, res)
            pixels = []
            for x in range(res):
                for y in range(res):
                    pixels.append(img[x, y].alpha_int8)
            values[res] = ", ".join(str(a) for a in pixels)

print(t.substitute(name=name, values_16=values[16], values_32=values[32], values_64=values[64]))
