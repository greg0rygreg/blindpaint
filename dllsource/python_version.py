# first, import ctypes, comes preinstalled with python i think
from ctypes import cdll

# reference the blindpaint lib by using LoadLibrary
blindpaint = cdll.LoadLibrary("./blindpaint.so")

# define rows and columns
row, col = 3, 3

# make the canvas
canvas = blindpaint.makeCanvas(row, col)

# paint some pixels
blindpaint.paintPixel(canvas, 0, 0, 1)
blindpaint.paintPixel(canvas, 0, 2, 1)
blindpaint.paintPixel(canvas, 2, 0, 1)
blindpaint.paintPixel(canvas, 2, 1, 1)
blindpaint.paintPixel(canvas, 2, 2, 1)
# (these make a face just like in the C variant of this script)

# save the canvas (for some weird reason you CAN'T use more than 1 char for the filename)
blindpaint.exportCanvas(canvas, row, col, "cujh.txt")

# IMPORTANT!!! free the reserved memory for the canvas after everything
blindpaint.freeCanvas(canvas, col)

# now you can run this script using the following command:
# python3 python_version.py
# or if you have python 3.11 installed like me:
# python3.11 python_version.py