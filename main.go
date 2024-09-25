package main

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/disintegration/imaging"
)

func prtErr(name string, desc string) {
	fmt.Println("\x1b[1;31merror:\x1b[39m " + name + "\x1b[0m\n" + desc)
}

func prtWar(name string) {
	fmt.Println("\x1b[1;33mwarning:\x1b[39m " + name + "\x1b[0m")
}

func prtDeb(name any) {
	name2 := fmt.Sprint(name)
	fmt.Println("\x1b[1;33mdebug:\x1b[0m " + name2)
}

func clear() {
	fmt.Printf("\x1b[H\x1b[2J")
}

func contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func sep() {
	fmt.Println("-------------------------------------------------------------------------")
}

func title(name any) {
	title2 := fmt.Sprint(name)
	fmt.Printf("\x1b]2;" + title2 + "\007")
}

func inargs(arg1 string, arg2 string) bool {
	if contains(os.Args, arg1) || contains(os.Args, arg2) {
		return true
	}
	return false
}

// Cuh...
func fillRegion(canvas [][]int, x1, y1, x2, y2, val int) {
	startX := int(math.Min(float64(x1), float64(x2)))
	endX := int(math.Max(float64(x1), float64(x2)))
	startY := int(math.Min(float64(y1), float64(y2)))
	endY := int(math.Max(float64(y1), float64(y2)))

	rows := len(canvas)
	cols := len(canvas[0])

	startX = int(math.Max(0, float64(startX)))
	endX = int(math.Min(float64(rows-1), float64(endX)))
	startY = int(math.Max(0, float64(startY)))
	endY = int(math.Min(float64(cols-1), float64(endY)))

	for i := startX; i <= endX; i++ {
		for j := startY; j <= endY; j++ {
			canvas[i][j] = val
		}
	}
}

func main() {
	if inargs("-h", "--help") {
		fmt.Println(`this isn't a command line tool but ok
-h --help           shows this message
-nw --no-welcome    the app won't welcome you
-d --debug          debug mode, print the list of pixels on save`)
		os.Exit(0)
	}
	// god knows when 2.0 will release
	VERSION := "1.0"

	var HOME string
	if runtime.GOOS == "windows" {
		HOME = os.Getenv("APPDATA")
	} else {
		HOME = os.Getenv("HOME")
	}
	var blindpaint_appdata_dir string
	if runtime.GOOS == "windows" {
		blindpaint_appdata_dir = filepath.Join(HOME, "blindpaint")
	} else {
		blindpaint_appdata_dir = filepath.Join(HOME, ".blindpaint")
	}
	os.MkdirAll(blindpaint_appdata_dir, os.ModePerm)
	byemsgs := []string{"bye!", "cya : )", ":3", "...",
		"i hope you painted something. i have no idea if you did. i just hope you did.", "see you next time!", "thanks for using this console app!",
		"i'm surprised someone actually downloads from my (greg) repos", "keep blindpainting! i'm sure you'll get better : )",
		"don't forget to join our discord!"}
	welcomemsgs := []string{"welcome back to", "did you train your blindpainting skills? anyway, welcome back to", "...", ":3", "press 2 to get the invite for our discord!",
		"you learned how it works yet? welcome back to", "welcome back to blindpai- oh sorry i didn't notice the logo was below me... mistakes happen : )\nanyway, welcome back to",
	}
	randomBye := rand.Intn(len(byemsgs))
	clear()
	if runtime.GOOS != "windows" {
		xtermwarn := filepath.Join(blindpaint_appdata_dir, "disablexterm")
		_, pluh := os.Stat(xtermwarn)
		if pluh != nil {
			prtWar("use xterm for the best experience! (run \x1b[4mtouch ~/.blindpaint/disablexterm\x1b[24m to disable this warning)")
			sep()
		}
	}
	if !inargs("-nw", "--no-welcome") {
		has_been_welcomed := filepath.Join(blindpaint_appdata_dir, "firsttime_welcomed")
		randomWelcome := rand.Intn(len(welcomemsgs))
		_, pluh := os.Stat(has_been_welcomed)
		if pluh == nil {
			fmt.Println(welcomemsgs[randomWelcome])
		} else {
			fmt.Println("welcome to blindpaint! if this is your first time, type 3 for the tutorial")
		}
		os.Create(has_been_welcomed)
	}
	for {
		title("blindpaint. - main menu")
		// why.
		fmt.Printf(`╭╮ ╭╮      ╭╮   a greg╭╮project
││ ││      ││        ╭╯╰╮
│╰─┤│╭┬──┬─╯├──┬──┬┬─┴╮╭╯(c)
│╭╮││├┤╭╮│╭╮│╭╮│╭╮├┤╭╮││
│╰╯│╰┤││││╰╯│╰╯│╭╮│││││╰╮╭╮
├──┴─┴┴╯╰┴──┤╭─┴╯╰┴┴╯╰┴─╯╰╯
│   blindly ││aint things.
│  made with╰╯Golang
├(1) new canvas
├(2) info
├(3) tutorial
└(0) exit
>> `)
		var input int
		fmt.Scan(&input)
		if input == 1 {
			clear()
			title("blindpaint. - new canvas")

			var canvasSize string
			fmt.Printf("input canvas size (eg. 3,3)\n>> ")
			fmt.Scan(&canvasSize)

			canvasSize2 := strings.Split(canvasSize, ",")
			var canvasSize_int [2]int
			for i, num := range canvasSize2 {
				cuh, _ := strconv.Atoi(num)
				canvasSize_int[i] = cuh
			}
			if len(canvasSize_int) != 2 || (canvasSize_int[0] == 0 || canvasSize_int[1] == 0) {
				clear()
				prtErr("invalid canvas size", "defaulting to 3,3")
				canvasSize_int = [2]int{3, 3}
				sep()
			}

			canvas := make([][]int, canvasSize_int[0])
			for i := range canvas {
				canvas[i] = make([]int, canvasSize_int[1])
			}

			clear()

			for {
				title("blindpaint. - " + strings.Join(canvasSize2, "x") + " canvas")
				var input_canvas int
				//you can't be serious.
				fmt.Printf(`      ╭╮
     ╭╯╰╮
╭──┬─┴╮╭┼┬──┬──┬──╮
│╭╮│╭─┤│├┤╭╮│╭╮│──┤
│╭╮│╰─┤╰┤│╰╯│││├──│
├╯╰┴──┴─┴┴──┴╯╰┴──╯
├(1) paint a pixel
├(2) fill a region
└(0) finish and leave to main menu
>> `)
				fmt.Scan(&input_canvas)
				var toPaintPos2 [2]int
				var toPaintPos3 [2]int
				if input_canvas == 1 {
					clear()

					/*var input_row int
					canvas_row_len := strconv.Itoa(len(canvas) - 1)
					fmt.Printf("input row from 0-" + canvas_row_len + ".\n>> ")
					fmt.Scan(&input_row)

					if input_row < 0 || input_row > len(canvas)-1 {
						clear()
						prtErr("row out of range", "now exiting to main menu")
						sep()
						break
					}

					clear()

					var input_col int
					canvas_col_len := strconv.Itoa(len(canvas[input_row]) - 1)
					fmt.Printf("input column from 0-" + canvas_col_len + ".\n>> ")
					fmt.Scan(&input_col)

					if input_col < 0 || input_col > len(canvas)-1 {
						clear()
						prtErr("column out of range", "now exiting to main menu")
						sep()
						break
					}*/

					var input_pos string
					canvas_xlen := strconv.Itoa(len(canvas))
					canvas_ylen := strconv.Itoa(len(canvas[0]))
					fmt.Printf("input position (1-" + canvas_xlen + ",1-" + canvas_ylen + ")\n>> ")
					fmt.Scan(&input_pos)

					toPaintPos := strings.Split(input_pos, ",")
					for i, num := range toPaintPos {
						cuh, _ := strconv.Atoi(num)
						toPaintPos2[i] = cuh - 1
					}

					clear()

					if toPaintPos2[0]+1 < 1 || toPaintPos2[0]+1 > len(canvas) {
						prtErr("X position out of range", "defaulting to 1")
						toPaintPos2[0] = 0
						sep()
					}

					if toPaintPos2[1]+1 < 1 || toPaintPos2[1]+1 > len(canvas[0]) {
						prtErr("Y position out of range", "defaulting to 1")
						toPaintPos2[1] = 0
						sep()
					}

					var input_paint int
					fmt.Printf("input pixel color 1 or 0 (1:white; 0:black)\n>> ")
					fmt.Scan(&input_paint)

					if input_paint < 0 || input_paint > 1 {
						clear()
						prtErr("pixel color out of range", "defaulting to 0")
						input_paint = 0
						sep()
					}

					canvas[toPaintPos2[0]][toPaintPos2[1]] = input_paint

					clear()
					fmt.Println("successfully painted pixel")
					sep()
				} else if input_canvas == 2 {
					clear()
					var input_pos1 string
					canvas_xlen := strconv.Itoa(len(canvas))
					canvas_ylen := strconv.Itoa(len(canvas[0]))
					fmt.Printf("input position 1 (1-" + canvas_xlen + ",1-" + canvas_ylen + ")\n>> ")
					fmt.Scan(&input_pos1)

					toPaintPos := strings.Split(input_pos1, ",")
					for i, num := range toPaintPos {
						cuh, _ := strconv.Atoi(num)
						toPaintPos2[i] = cuh - 1
					}

					clear()

					if toPaintPos2[0]+1 < 1 || toPaintPos2[0]+1 > len(canvas) {
						prtErr("X position out of range", "defaulting to 1")
						toPaintPos2[0] = 0
						sep()
					}

					if toPaintPos2[1]+1 < 1 || toPaintPos2[1]+1 > len(canvas[0]) {
						prtErr("Y position out of range", "defaulting to 1")
						toPaintPos2[1] = 0
						sep()
					}

					var input_pos2 string
					fmt.Printf("input position 2 (1-" + canvas_xlen + ",1-" + canvas_ylen + ")\n>> ")
					fmt.Scan(&input_pos2)

					toPaintPos = strings.Split(input_pos2, ",")
					for i, num := range toPaintPos {
						cuh, _ := strconv.Atoi(num)
						toPaintPos3[i] = cuh - 1
					}

					clear()

					if toPaintPos3[0]+1 < 1 || toPaintPos3[0]+1 > len(canvas) {
						prtErr("X position out of range", "defaulting to 1")
						toPaintPos3[0] = 0
						sep()
					}

					if toPaintPos3[1]+1 < 1 || toPaintPos3[1]+1 > len(canvas[0]) {
						prtErr("Y position out of range", "defaulting to 1")
						toPaintPos3[1] = 0
						sep()
					}

					var input_paint int
					fmt.Printf("input pixel color 1 or 0 (1:white; 0:black).\n>> ")
					fmt.Scan(&input_paint)

					if input_paint < 0 || input_paint > 1 {
						clear()
						prtErr("pixel color out of range", "defaulting to 0")
						input_paint = 0
						sep()
					}

					fillRegion(canvas, toPaintPos2[0], toPaintPos2[1], toPaintPos3[0], toPaintPos3[1], input_paint)

					clear()
					fmt.Println("successfully filled region")
					sep()

				} else if input_canvas == 0 {
					clear()
					var input_fn string
					var input_fnt int

					fmt.Printf("choose type to save to\n├(1) png (\x1b[31mWARNING: MAY NOT WORK\x1b[0m)\n└(2) txt\n>> ")
					fmt.Scan(&input_fnt)
					clear()

					if input_fnt == 1 {
						img := image.NewRGBA(image.Rect(0, 0, canvasSize_int[0], canvasSize_int[1]))
						for i := 0; i < canvasSize_int[0]; i++ {
							for j := 0; j < canvasSize_int[1]; j++ {
								img.Set(j, i, color.RGBA{uint8(canvas[i][j]) * 255, uint8(canvas[i][j]) * 255, uint8(canvas[i][j]) * 255, 255})
							}
						}

						fmt.Printf("save art to (.png will be added automatically):\n>> ")
						fmt.Scan(&input_fn)
						sep()

						clear()
						var input_upscale string
						fmt.Printf("do you want to upscale image x10? (Y/N)\n>> ")
						fmt.Scan(&input_upscale)
						var scaledImg image.Image
						if strings.ToLower(input_upscale) == "y" {
							scaledImg = imaging.Resize(img, canvasSize_int[0]*10, canvasSize_int[1]*10, imaging.NearestNeighbor)
						}

						fmt.Println("saving your image, this might take a while on bad pcs with big canvas sizes")
						if strings.ToLower(input_upscale) == "y" {
							imaging.Save(scaledImg, input_fn+".png")
						} else {
							imaging.Save(img, input_fn+".png")
						}

						clear()
						fmt.Println("successfully saved your masterpiece to " + input_fn + ".png")
						sep()

						break
					} else if input_fnt == 2 {
						fmt.Printf("save art to (.txt will be added automatically):\n>> ")
						fmt.Scan(&input_fn)
						file, _ := os.Create(input_fn + ".txt")
						defer file.Close()

						var content string
						content += "(tip: use a font that has letters with the same width!)\n\n"

						for i := range canvas {
							for j := range canvas[i] {
								if canvas[i][j] == 1 {
									content += "■ "
								} else {
									content += "□ "
								}
							}
							content += "\n"
						}

						content += "time created: " + strings.Split(time.Now().String(), ".")[0] + "\nmade by: a very awesome person\n"
						_, _ = file.WriteString(content)

						clear()

						if contains(os.Args, "-d") || contains(os.Args, "--debug") {
							prtDeb(canvas)
						}

						fmt.Println("successfully saved your masterpiece to " + input_fn + ".txt")
						sep()
						break
					}
				} else {
					clear()
					prtErr("input exception", "no option made for input "+strconv.Itoa(input_canvas))
					sep()
				}
			}
		} else if input == 2 {
			clear()
			fmt.Println("blindpaint. version " + VERSION + `
blindpaint is open-source and licensed under GPL v3.0
you can redistribute this app as long as you maintain the same license and credit me (greg)
join our lovely discord built off of another server! discord.gg/c2KTVEgxBn`)
			/* 																																																do it i BEG of you. */
			sep()
		} else if input == 3 {
			clear()
			fmt.Println(`HOW TO BLINDPAINT - a tutorial by greg

1. on the main menu, press 1
 this is to start the main part of blindpaint

2. on the canvas selection, put a value like 3,3 or 5,5
 don't get greedy; saving as txt might cause memory issues and a deformed canvas!

3. on the actions menu, select 1
 to choose a pixel to paint
 
4. choose a position
 this is to choose the pixel to paint
 
5. choose 1 or 0 when it asks so
 this is to paint with white or black respectively (black to erase if you did a mistake!)
 
6. repeat steps 3, 4 and 5
 to fill your canvas with art, obviously!
 
7. after you've finished, save your masterpiece
 by pressing 0 on the actions menu!

8. choose a filetype
 tip: png saves faster, even when upscaled 10x!
 
9. check the piece of art you've created!
 it's not going to be good, but you tried
 
10. (optional) repeat previous steps
 until you get good!

thanks for using blindpaint, i (greg) appreciate it, and i mean it :,)`)
			sep()
		} else if input == 0 {
			clear()

			fmt.Println(byemsgs[randomBye])
			os.Exit(0)

			sep()
		} else {
			clear()
			prtErr("input exception", "no option made for input "+strconv.Itoa(input))
			sep()
		}
	}
}
