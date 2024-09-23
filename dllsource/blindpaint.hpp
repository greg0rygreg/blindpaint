#ifndef BLINDPAINT_HPP
#define BLINDPAINT_HPP

#include <vector>
#include <string>


/// @brief the blindpaint namespace. where you can blindpaint from.
///
/// welcome to the blindpaint API! if you're new read how the functions work
/// in the `blindpaint.hpp` file (where this text originates from).
namespace blindpaint {
    /// @brief the main part of the blindpaint API
    ///
    /// this class called `Canvas` defines a canvas that soon will have some art on it (or not).
    class Canvas {
        /// array of pixels for the `Canvas` class.
        std::vector<std::vector<int>> pixels;
        public:
            Canvas(int rows, int cols);
            /// @brief paint a pixel at position `x,y` on current canvas with pixel value `val`
            ///
            /// this function paints a pixel at position `x,y` of the current canvas with value `val`.
            /// keep in mind that `val` only takes 0 or 1.
            /// 
            /// @param x the x position to paint at
            /// @param y the y position to paint at
            /// @param val the value to paint pixel with
            void paintPixel(int x, int y, int val);
            
            /// @brief export current canvas to file `filename`
            ///
            /// this function exports the current canvas to the file with name specified in `filename`.
            ///
            /// @param filename the name of the file to export to
            void save(const std::string &filename);
            
            /// @brief fill a region from position `x1,y1` to `x2,y2` on current canvas with pixel value `val`
            ///
            /// this function fills a specified area from position `x1,y1` to position `x2,y2` with value `val`.
            /// 
            /// keep in mind that `val` only takes 0 or 1,
            /// and `x#` and `y#` parameters will stay inbound relative to the current canvas.
            /// 
            /// @param x1 starting x position
            /// @param y1 starting y position
            /// @param x2 ending x position
            /// @param y2 ending y position
            /// @param val the value to paint area with
            void fillRegion(int x1, int y1, int x2, int y2, int val);
    };
    extern "C" {
        /// @brief same as the C++ variant
        /// @return `blindpaint::Canvas*` - there were no errors and a canvas was created successfully.
        Canvas* makeCanvas(int rows, int cols);
        /// @brief same as the C++ variant
        void paintPixel(Canvas* canvas, int x, int y, int val);
        /// @brief same as the C++ variant
        void saveCanvas(Canvas* canvas, const char* filename);
        /// @brief same as the C++ variant
        void fillRegion(Canvas* canvas, int x1, int y1, int x2, int y2, int val);
        /// @brief free memory from a canvas
        ///
        /// (C/PYTHON ONLY) clear the reserved memory for canvas `canvas`.
        /// useful if you don't want to deal with using `free()` (or something).
        /// @param canvas the canvas to free memory from
        void freeCanvas(Canvas* canvas);
    }
}

#endif