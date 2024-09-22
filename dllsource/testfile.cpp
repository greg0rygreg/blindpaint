//this was made to test new features without interfering with the using_blindpaint_so.cpp file.
#include "blindpaint.hpp"
#include <iostream>

int main(int argc, char *argv[]) {
    int row = 10, col = 10;
    auto canvas = blindpaint::makeCanvas(row,col);

    blindpaint::fillRegion(canvas, 1,1, 8,8, 1);
    blindpaint::fillRegion(canvas, 2,2, 7,7, 0);

    blindpaint::exportCanvas(canvas, "example.txt");

    return 0;
}