// please excuse me for the horrors i've written to make this dll
#include <iostream>
#include <fstream>
#include <vector>

namespace blindpaint {
    std::vector<std::vector<int>> makeCanvas(int rows, int cols) {
        std::vector<std::vector<int>> canvas(rows, std::vector<int>(cols, 0));
        return canvas;
    }

    void paintPixel(std::vector<std::vector<int>> &canvas, int x, int y, int val) {
        if (val >= 2 || val < 0) {
            std::cerr << "\x1b[1;31merror:\x1b[39m unknown pixel value of " << val << " detected \x1b[0m\n";
        }
        canvas[x][y] = val;
    }

    void exportCanvas(const std::vector<std::vector<int>> &canvas, const std::string &filename) {
        std::ofstream file(filename);
        if (!file.is_open()) {
            std::cerr << "\x1b[1;31merror:\x1b[39m exporting canvas to file " << filename << " failed \x1b[0m\n";
        }

        file << "(tip: use a font that has letters with the same width!)\n\n";
        for (const auto& row : canvas) {
            for (int pixel : row) {
               file << (pixel == 1 ? "■ " : "□ ");
            }
            file << "\n";
        }
        file << "time created: not yet implemented!\nmade by: a very awesome person\n";
    }
    
    void fillRegion(std::vector<std::vector<int>> &canvas, int x1, int y1, int x2, int y2, int val) {
        // yes i used ChatGPT to make this (don't blame me it's 2 hours before midnight as of writing (september 21st 2024 10:23 PM))
        if (val >= 2 || val < 0) {
            std::cerr << "\x1b[1;31merror:\x1b[39m unknown pixel value of " << val << " detected \x1b[0m\n";
        }
        int startX = std::min(x1, x2);
        int endX = std::max(x1, x2);
        int startY = std::min(y1, y2);
        int endY = std::max(y1, y2);

        int rows = canvas.size();
        int cols = canvas[0].size();
        startX = std::max(0, startX);
        endX = std::min(rows - 1, endX);
        startY = std::max(0, startY);
        endY = std::min(cols - 1, endY);

        for (int i = startX; i <= endX; ++i) {
            for (int j = startY; j <= endY; ++j) {
                canvas[i][j] = val;
            }
        }
    }
}

/* 
* !
*
* hi
*
* if you want to export this file yourself, run this command:
* g++ -shared -fPIC -o blindpaint.so blindpaint_so.cpp
* or if on windows, change "so" in "blindpaint.so" to "dll"
*/