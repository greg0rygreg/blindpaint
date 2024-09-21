#include <stdio.h>
// to use the blindpaint dll you must reference these guys
// this one for creating a canvas
int **makeCanvas(int row, int col);
// this one for painting a pixel
int paintPixel(int **canvas, int col, int row, int val);
// this one for exporting the canvas
int exportCanvas(int **canvas, int rows, int cols, char *file);
// and this one for freeing the reserved memory for the canvas after everything
void freeCanvas(int **canvas, int cols);

//then make a main() func
int main() {
    // this is required for later functions, define the amount of rows and columns
    int row = 3, col = 3;
    // make a canvas with the previously defined variables
    int **canvas = makeCanvas(row,col);
    // check if it's null and return exit code 1 if true
    if (canvas == NULL) {return 1;}

    // paint a few pixels (this makes a face)
    paintPixel(canvas, 0,0, 1);
    paintPixel(canvas, 0,2, 1);
    paintPixel(canvas, 2,0, 1);
    paintPixel(canvas, 2,1, 1);
    paintPixel(canvas, 2,2, 1);

    // export the canvas
    exportCanvas(canvas, 3, 3, "example.txt");

    // IMPORTANT!!! free the allocated memory for the canvas (so your app doesn't cause a memory leak when used continuously)
    freeCanvas(canvas, col);

    // return exit code 0 if successful (you should've known how to do this)
    return 0;
}

// then, compile this using the following command:
// gcc -o example_binary using_blindpaint_dll.c ./blindpaint.so
// DO NOT FORGET TO PUT THE BLINDPAINT LIB IN THE SAME DIRECTORY AS OF THE MAIN SOURCE FILE!!!