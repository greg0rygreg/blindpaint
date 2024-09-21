// please excuse me for the horrors i've written to make this dll
// and also that i switched to C++ it's just higher-end and easier to use
#include <iostream>
#include <fstream>
#include <vector>

extern "C" {
    std::vector<std::vector<int>> makeCanvas(int rows, int cols) {
        std::vector<std::vector<int>> canvas(rows, std::vector<int>(cols, 0));
        return canvas;
    }

    int paintPixel(std::vector<std::vector<int>>& canvas, int x, int y, int val) {
        if (val >= 2 || val < 0) {
            std::cerr << "\x1b[1;31merror:\x1b[39m incorrectly painted pixel detected (" << x << "," << y << " with " << val << ") \x1b[0m\n";
            return 1;
        }
        canvas[x][y] = val;
        return 0;
    }

    int exportCanvas(const std::vector<std::vector<int>>& canvas, const std::string& filename) {
        std::ofstream file(filename);
        if (!file.is_open()) {
            std::cerr << "\x1b[1;31merror:\x1b[39m exporting canvas to file " << filename << " failed \x1b[0m\n";
            return 1;
        }

        file << "(tip: use a font that has letters with the same width!)\n\n";
        for (const auto& row : canvas) {
            for (int pixel : row) {
                if (pixel == 1) {
                    file << "■ ";
                } else {
                    file << "□ ";
                }
            }
            file << "\n";
        }
        file << "time created: not yet implemented!\nmade by: a very awesome person\n";
        return 0;
    }
}

// !

// hi

// if you want to export this file yourself, run this command:
// g++ -shared -fPIC -o blindpaint.so blindpaint_so.cpp

// or if on windows, change "so" in "blindpaint.so" to "dll"