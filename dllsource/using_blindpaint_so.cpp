#include <stdio.h>
// to use the blindpaint lib you must reference the header file
#include "blindpaint.hpp"

//then make a main() func
int main() {
    // this is required for later functions, define the amount of rows and columns
    int row = 3, col = 3;
    // make a canvas with the previously defined variables
    auto canvas = makeCanvas(row,col);

    // paint a few pixels (this makes a face)
    paintPixel(canvas, 0,0, 1);
    paintPixel(canvas, 0,2, 1);
    paintPixel(canvas, 2,0, 1);
    paintPixel(canvas, 2,1, 1);
    paintPixel(canvas, 2,2, 1);

    // export the canvas
    exportCanvas(canvas, "example.txt");

    // i don't think you need to free the memory when using C++...

    // return exit code 0 if successful (you should've known how to do this)
    return 0;
}

// then, compile this using the following command:
// g++ -o example_binary using_blindpaint_so.cpp ./blindpaint.so
// DO NOT FORGET TO PUT THE BLINDPAINT LIB IN THE SAME DIRECTORY AS OF THE MAIN SOURCE FILE!!!
