# import ctypes, comes preinstalled with python i think
import ctypes

# reference the blindpaint lib by using LoadLibrary
blindpaint = ctypes.cdll.LoadLibrary("./blindpaint.so")

# define arg and response types
blindpaint.makeCanvas.argtypes = [ctypes.c_int, ctypes.c_int]; blindpaint.makeCanvas.restypes = ctypes.c_int
blindpaint.paintPixel.argtypes = [ctypes.c_int, ctypes.c_int, ctypes.c_int, ctypes.c_int]; blindpaint.makeCanvas.restypes = ctypes.c_int
blindpaint.exportCanvas.argtypes = [ctypes.c_int, ctypes.c_int, ctypes.c_int, ctypes.c_wchar_p]; blindpaint.makeCanvas.restypes = ctypes.c_int
blindpaint.freeCanvas.argtypes = [ctypes.c_int, ctypes.c_int]; blindpaint.makeCanvas.restypes = None

# define rows and columns
row, col = 3, 3

# make the canvas
canvas = blindpaint.makeCanvas(row, col)

# if the canvas is None or NULL then exit
if canvas == None: exit(1)

# paint a few pixels
blindpaint.paintPixel(canvas, 0,0, 1)
blindpaint.paintPixel(canvas, 0,2, 1)
blindpaint.paintPixel(canvas, 2,0, 1)
blindpaint.paintPixel(canvas, 2,1, 1)
blindpaint.paintPixel(canvas, 2,2, 1)
# (these make a face just like in the C variant of this script)

# save the canvas (for some weird reason you CAN'T use more than 1 char for the filename)
blindpaint.exportCanvas(canvas, row, col, "cujh.txt")

# IMPORTANT!!! free the reserved memory for the canvas after everything
blindpaint.freeCanvas(canvas, col)

# now you can run this script using the following command:
# python3 python_version.py
# or if you have python 3.11 installed like me:
# python3.11 python_version.py
# DO NOT FORGET TO PUT THE BLINDPAINT LIB IN THE SAME DIRECTORY AS OF THIS FILE!!!