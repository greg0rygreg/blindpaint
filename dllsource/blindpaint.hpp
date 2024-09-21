#ifndef BLINDPAINT_HPP
#define BLINDPAINT_HPP

#include <vector>
#include <string>

extern "C" {
    std::vector<std::vector<int>> makeCanvas(int rows, int cols);
    void paintPixel(std::vector<std::vector<int>>& canvas, int x, int y, int val);
    void exportCanvas(const std::vector<std::vector<int>>& canvas, const std::string& filename);
}

#endif