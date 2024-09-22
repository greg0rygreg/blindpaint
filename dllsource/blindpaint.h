#ifndef BLINDPAINT_HPP
#define BLINDPAINT_HPP

#include <vector>
#include <string>


/// @brief the blindpaint namespace. where you can blindpaint from.
///
/// welcome to the blindpaint API! if you're new read how the functions work
/// in the `blindpaint.hpp` file (where this text originates from).
namespace blindpaint {
    
    /// @brief create a canvas with size `rows,cols` to be used by other functions
    ///
    /// this function will return a new canvas with `rows` rows and `cols` columns,
    /// can be used multiple times in order to make multiple canvases.
    /// 
    /// @param rows the amount of rows to generate
    /// @param cols the amount of columns to generate    
    std::vector<std::vector<int>> makeCanvas(int rows, int cols);
    
    /// @brief paint a pixel at position `x,y` on canvas `canvas` with pixel value `val`
    ///
    /// this function will return nothing, but will paint a pixel at position `x,y` of canvas `canvas` with value `val`.
    /// keep in mind that `val` only takes 0 or 1.
    /// 
    /// @param canvas the canvas to paint the pixel
    /// @param x the x position to paint at
    /// @param y the y position to paint at
    /// @param val the value to paint pixel with
    void paintPixel(std::vector<std::vector<int>>& canvas, int x, int y, int val);
    
    /// @brief export canvas `canvas` to file `filename`
    ///
    /// this function will return nothing, but will export canvas `canvas` to the file with name specified in `filename`.
    /// 
    /// @param canvas the canvas to export
    /// @param filename the name of the file to export to
    void exportCanvas(const std::vector<std::vector<int>>& canvas, const std::string& filename);
    
    /// @brief fill a region from position `x1,y1` to `x2,y2` on canvas `canvas` with pixel value `val`
    ///
    /// this function will return nothing, but will fill a specified area from position `x1,y1` to position `x2,y2` with value `val`.
    /// 
    /// keep in mind that `val` only takes 0 or 1,
    /// and `x#` and `y#` parameters will stay inbound relative to the specified canvas.
    /// 
    /// @param canvas the canvas to paint the area
    /// @param x1 starting x position
    /// @param y1 starting y position
    /// @param x2 ending x position
    /// @param y2 ending y position
    /// @param val the value to paint area with
    void fillRegion(std::vector<std::vector<int>>& canvas, int x1, int y1, int x2, int y2, int val);
}

#endif
