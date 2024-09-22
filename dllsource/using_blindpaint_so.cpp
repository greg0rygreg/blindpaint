#include <iostream>
// to use the blindpaint lib you must reference the header file
#include "blindpaint.h"

//then make a main() func
int main() {
    // define the amount of rows and columns
    int row = 3, col = 3;
    // make a canvas with the previously defined variables
    blindpaint::Canvas canvas(row,col);

    // paint a few pixels (this makes a face)
    canvas.paintPixel(0,0, 1);
    canvas.paintPixel(0,2, 1);
    canvas.fillRegion(2,0, 2,2, 1);

    // export the canvas
    canvas.save("example.txt");

    // i don't think you need to free the memory when using C++...

    // return exit code 0 if successful (you should've known how to do this)
    return 0;
}

// then, compile this using the following command:
// g++ -o example_binary using_blindpaint_so.cpp ./blindpaint.so
// DO NOT FORGET TO PUT THE BLINDPAINT LIB IN THE SAME DIRECTORY AS OF THE MAIN SOURCE FILE!!!