package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

func prtErr(name string, desc string) {
	color.New(color.FgRed).Add(color.Bold).Printf("error: ")
	color.New(color.Bold).Println(name)
	fmt.Println(desc)
}

/*func prtWar(name string) {
	color.New(color.FgYellow).Add(color.Bold).Printf("warning: ")
	color.New(color.Bold).Println(name)
}*/

func prtDeb(name any) {
	color.New(color.FgYellow).Add(color.Bold).Printf("debug: ")
	color.New(color.Reset).Println(name)
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

func main() {
	VERSION := "pre2"
	var HOME string
	if runtime.GOOS == "windows" {
		HOME = os.Getenv("APPDATA")
	} else {
		HOME = os.Getenv("HOME")
	}
	blindpaint_appdata_dir := filepath.Join(HOME, ".blindpaint")
	has_been_welcomed := filepath.Join(blindpaint_appdata_dir, "firsttime_welcomed")

	rand.New(rand.NewSource(time.Now().UnixNano()))
	byemsgs := []string{"bye!", "cya : )", ":3", "...",
		"i hope you painted something. i have no idea if you did. i just hope you did.", "see you next time!", "thanks for using this console app!",
		"i'm surprised someone actually downloads from my (greg) repos.", "keep blindpainting! i'm sure you'll get better : )",
		"don't use this app if you're tired! you might end up falling asleep."}
	welcomemsgs := []string{"welcome back to", "did you train your blindpainting skills? anyway, welcome back to", "...", ":3",
		"you learned how it works yet? welcome back to", "welcome back to blindpai- oh sorry i didn't notice the logo was below me... mistakes happen : )\nanyway, welcome back to",
	}
	randomBye := rand.Intn(len(byemsgs))
	randomWelcome := rand.Intn(len(welcomemsgs))

	clear()
	_, pluh := os.Stat(has_been_welcomed)
	if pluh == nil {
		fmt.Println(welcomemsgs[randomWelcome])
	} else {
		fmt.Println("welcome to")
	}

	os.MkdirAll(blindpaint_appdata_dir, os.ModePerm)
	os.Create(has_been_welcomed)
	for { // I WANT THE PAIN TO END
		fmt.Printf(`╭╮ ╭╮      ╭╮   a greg╭╮project.
││ ││      ││        ╭╯╰╮
│╰─┤│╭┬──┬─╯├──┬──┬┬─┴╮╭╯©️
│╭╮││├┤╭╮│╭╮│╭╮│╭╮├┤╭╮││
│╰╯│╰┤││││╰╯│╰╯│╭╮│││││╰╮╭╮
├──┴─┴┴╯╰┴──┤╭─┴╯╰┴┴╯╰┴─╯╰╯
│   blindly ││aint things.
│  made with╰╯Golang.
├(1) new canvas.
├(2) info.
├(3) tutorial.
└(0) exit.
>> `)
		var input int
		fmt.Scan(&input)

		if input == 1 {
			clear()

			var canvasSize string
			fmt.Printf("input canvas size (eg. 3,3).\n>> ")
			fmt.Scan(&canvasSize)

			canvasSize2 := strings.Split(canvasSize, ",")
			var canvasSize_int [2]int
			for i, num := range canvasSize2 {
				cuh, _ := strconv.Atoi(num)
				canvasSize_int[i] = cuh
			}

			canvas := make([][]int, canvasSize_int[0])
			for i := range canvas {
				canvas[i] = make([]int, canvasSize_int[1])
			}

			clear()

			for {
				var input_canvas int
				fmt.Printf(`┌actions.
├(1) paint a pixel.
└(0) finish and leave to main menu.
>> `)
				fmt.Scan(&input_canvas)

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
					fmt.Printf("input position (1-" + canvas_xlen + ",1-" + canvas_ylen + ").\n>> ")
					fmt.Scan(&input_pos)

					toPaintPos := strings.Split(input_pos, ",")
					var toPaintPos2 [2]int
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
					fmt.Printf("input pixel color 1 or 0 (1:white; 0:black).\n>> ")
					fmt.Scan(&input_paint)

					if input_paint < 0 || input_paint > 1 {
						clear()
						prtErr("pixel color out of range", "defaulting to 0")
						input_paint = 0
						sep()
					}

					canvas[toPaintPos2[0]][toPaintPos2[1]] = input_paint

					clear()
					fmt.Println("successfully painted pixel.")
					sep()

				} else if input_canvas == 0 {
					file, _ := os.Create("blindart.txt")
					defer file.Close()

					var content string

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

					cuh, _ := user.Current()
					content += "\ntime created: " + strings.Split(time.Now().String(), ".")[0] + "\nmade by: " + cuh.Username + "\n"
					_, _ = file.WriteString(content)

					clear()

					if contains(os.Args, "-d") || contains(os.Args, "--debug") {
						prtDeb(canvas)
					}

					fmt.Println("successfully saved your masterpiece to blindart.txt")
					sep()

					break
				} else {
					clear()
					prtErr("input exception", "no option made for input "+strconv.Itoa(input_canvas))
					sep()
				}
			}
		} else if input == 2 {
			clear()
			fmt.Println("blindpaint version " + VERSION + ".\ncrafted with much love and detail by greg <3")
			sep()
		} else if input == 3 {
			clear()
			fmt.Println(`HOW TO BLINDPAINT - a tutorial by greg.

1. on the main menu, press 1
 this is to start the main part blindpaint.

2. on the canvas selection, put a value like 3,3 or 5,5
 don't get greedy; saving might cause memory issues and a deformed canvas!

3. on the actions menu, select 1
 to choose a pixel to paint
 
4. choose a row and a column
 this is to choose the pixel to paint
 
5. choose 1 or 0 when it asks so
 this is to paint with white or black respectively (black to erase if you did a mistake!)
 
6. repeat steps 3, 4 and 5
 to fill your canvas with art, obviously!
 
7. after you've finished, save your masterpiece by pressing 2 on the actions menu
 i think you know what you need to do.
 
8. check the piece of art you've created!
 same here
 
9. (optional) repeat previous steps
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
