# import ctypes, comes preinstalled with python
import ctypes

# import the blindpaint lib
blindpaint = ctypes.CDLL("./blindpaint.so")

# no args and res types, too lazy

# define the row and col vars
row, col = 3, 3

# make a canvas
canvas = blindpaint.makeCanvas(row, col)

# draw some pixels
blindpaint.paintPixel(canvas, 0,0, 1)
blindpaint.paintPixel(canvas, 0,2, 1)
blindpaint.fillRegion(canvas, 2,0, 2,2, 1)
# (makes a face like in the C++ variant)

# export the canvas
blindpaint.saveCanvas(canvas, b"python.txt")

# important...? free the memory
blindpaint.freeCanvas(canvas)