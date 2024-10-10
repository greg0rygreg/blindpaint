// please excuse me for the horrors i've written to make this dll
#include <iostream>
#include <fstream>
#include <vector>
#include "blindpaint.hpp"
#include <ctime>
// oh yeah also you need to download stb image write header file to both compile and use the so
#define STB_IMAGE_WRITE_IMPLEMENTATION
#include "stb_image_write.h"

namespace blindpaint
{
    Canvas::Canvas(int rows, int cols) : pixels(rows, std::vector<int>(cols, 0)) {}

    void Canvas::paintPixel(int x, int y, int val)
    {
        if (val >= 2 || val < 0)
        {
            std::cerr << "\x1b[1;31merror:\x1b[39m unknown pixel value of " << val << " detected \x1b[0m(defaulting to 1)\n";
            val = 1;
        }
        pixels[x][y] = val;
    }
    void Canvas::save(const std::string &filename)
    {
        time_t currtime = time(0);
        std::ofstream file(filename);
        if (!file.is_open())
        {
            std::cerr << "\x1b[1;31merror:\x1b[39m exporting canvas to file " << filename << " failed \x1b[0m\n";
        }
        file << "(tip: use a font that has letters with the same width!)\n";
        for (const auto &row : pixels)
        {
            for (int pixel : row)
            {
                file << (pixel == 1 ? "#" : " ");
            }
            file << "\n";
        }
        char *dt = ctime(&currtime);
        file << "\ntime created: " << dt << "made by: a very awesome person\n";
    }

    void Canvas::fillRegion(int x1, int y1, int x2, int y2, int val)
    {
        // yes i used ChatGPT to make this (don't blame me it's 2 hours before midnight as of writing (september 21st 2024 10:23 PM))
        if (val >= 2 || val < 0)
        {
            std::cerr << "\x1b[1;31merror:\x1b[39m unknown pixel value of " << val << " detected \x1b[0m(defaulting to 1)\n";
            val = 1;
        }
        int startX = std::min(x1, x2);
        int endX = std::max(x1, x2);
        int startY = std::min(y1, y2);
        int endY = std::max(y1, y2);

        int rows = pixels.size();
        int cols = pixels[0].size();
        startX = std::max(0, startX);
        endX = std::min(rows - 1, endX);
        startY = std::max(0, startY);
        endY = std::min(cols - 1, endY);

        for (int i = startX; i <= endX; ++i)
        {
            for (int j = startY; j <= endY; ++j)
            {
                pixels[i][j] = val;
            }
        }
    }

    void Canvas::savePng(const std::string &filename)
    {
        // i too used chatgpt to make this
        int h = pixels.size();
        int w = pixels[0].size();
        std::vector<unsigned char> image(w * h * 3);

        for (int y = 0; y < h; y++)
        {
            for (int x = 0; x < w; x++)
            {
                image[(y * w + x) * 3 + 0] = pixels[y][x] * 255;
                image[(y * w + x) * 3 + 1] = pixels[y][x] * 255;
                image[(y * w + x) * 3 + 2] = pixels[y][x] * 255;
            }
        }

        stbi_write_png(filename.c_str(), w, h, 3, image.data(), w * 3);
    }

    ////////////////////////////////////////////////// C IMPLEMENTATION STARTS HERE //////////////////////////////////////////////////

    Canvas *makeCanvas(int rows, int cols)
    {
        return new Canvas(rows, cols);
    }

    void paintPixel(Canvas *canvas, int x, int y, int val)
    {
        canvas->paintPixel(x, y, val);
    }

    void fillRegion(Canvas *canvas, int x1, int y1, int x2, int y2, int val)
    {
        canvas->fillRegion(x1, y1, x2, y2, val);
    }

    void saveCanvas(Canvas *canvas, const char *filename)
    {
        canvas->save(filename);
    }

    void freeCanvas(Canvas *canvas)
    {
        delete canvas;
    }

    void saveCanvasPng(Canvas *canvas, const char *filename)
    {
        canvas->savePng(filename);
    }
}

/*
 * !
 *
 * hi
 *
 * if you want to export this file yourself, run this command:
 * g++ -shared -fPIC -o blindpaint.so blindpaint_so.cpp
 */
