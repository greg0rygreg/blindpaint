#ifndef BLINDPAINT_HPP
#define BLINDPAINT_HPP

#include <vector>
#include <string>

namespace blindpaint {
    //create a canvas with size `rows,cols` to be used by other functions
    std::vector<std::vector<int>> makeCanvas(int rows, int cols);
    //paint a pixel at position `x,y` on canvas `canvas` with pixel value `val`
    void paintPixel(std::vector<std::vector<int>>& canvas, int x, int y, int val);
    //export canvas `canvas` to file `filename`
    void exportCanvas(const std::vector<std::vector<int>>& canvas, const std::string& filename);
    //fill a region from position `x1,y1` to `x2,y2` on canvas `canvas` with pixel value `val`
    void fillRegion(std::vector<std::vector<int>>& canvas, int x1, int y1, int x2, int y2, int val);
}

#endif