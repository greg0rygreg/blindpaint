// excuse me for the horrors i've written to make this dll
#include <stdio.h>
#include <stdlib.h>

int **makeCanvas(int row, int col) {
    int **canvas = malloc(row * sizeof(int*));

    if (canvas == NULL) {
        printf("\x1b[1;31merror:\x1b[39m allocation for canvas failed \x1b[0m\n");
        return NULL;
    }

    for (int i = 0; i < col; i++) {
        canvas[i] = malloc(col * sizeof(int));
        if (canvas[i] == NULL) {
            printf("\x1b[1;31merror:\x1b[39m allocation for column %d failed \x1b[0m\n", i);
            for (int j = 0; j < i; j++) {
                free(canvas[j]);
            }
            free(canvas);
            return NULL;
        }
    }
    return canvas;
}

int paintPixel(int** canvas, int x, int y, int val) {
    if (val >= 2 || val < 0) {
        printf("\x1b[1;31merror:\x1b[39m incorrectly painted pixel detected (%d,%d with %d) \x1b[0m\n", x, y, val);
        return 1;
    }
    canvas[x][y] = val;
    return 0;
}

int exportCanvas(int **canvas, int rows, int cols, char *file) {
    FILE *file2 = fopen(file, "w");
    if (file == NULL) {
        printf("\x1b[1;31merror:\x1b[39m exporting canvas to file %s failed \x1b[0m\n", file);
        return 1;
    }
    fprintf(file2, "(tip: use a font that has letters with the same width!)\n\n");
    for (int x = 0; x < rows; x++) {
        for (int y = 0; y < cols; y++) {
            if (canvas[x][y] == 1) {
                fprintf(file2, "■ ");
            } else {
                fprintf(file2, "□ ");
            }
        }
        fprintf(file2, "\n");
    }
    fprintf(file2, "time created: not yet implemented!\nmade by: a very awesome person\n");
    return 0;
}

void freeCanvas(int **canvas, int cols) {
    for (int i = 0; i < cols; i++) {
        free(canvas[i]);
    }
    free(canvas);
}

// oh hey you scrolled to here
// if you want to export this file yourself, run this command:
// gcc -shared -o blindpaint.so blindpaint_so.c