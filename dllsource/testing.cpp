//this was made to test new features without interfering with the example_use.cpp file.
#include <iostream>
#include "blindpaint.hpp"

int main(int argc, char *argv[]) {
    int row = 10, col = 10;
    blindpaint::Canvas canvas(row,col);

    canvas.fillRegion(1,1, 8,8, 1);
    canvas.fillRegion(2,2, 7,7, 0);

    canvas.save("testing.txt");

    return 0;
}