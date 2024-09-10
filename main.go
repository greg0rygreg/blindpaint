package main

import (
	"fmt"
	"math/rand"
	"os"
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

func prtWar(name string) {
	color.New(color.FgYellow).Add(color.Bold).Printf("warning: ")
	color.New(color.Bold).Println(name)
}
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
	fmt.Println("---------------------------------------------------")
}

func main() {
	VERSION := "pre1"
	rand.New(rand.NewSource(time.Now().UnixNano()))
	byemsgs := []string{"bye!", "cya : )", ":3", "i hope you painted something. i have no idea if you did. i just hope you did.", "see you next time!", "thanks for using this console app!", "i'm surprised someone actually downloads from my (greg) repos."}
	randomIndex := rand.Intn(len(byemsgs))
	clear()
	fmt.Println("welcome to")
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
└(3) exit.
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
└(2) finish and leave to main menu.
>> `)
				fmt.Scan(&input_canvas)
				if input_canvas == 1 {
					clear()
					var input_row int
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
					}
					clear()
					var input_paint int
					fmt.Printf("input pixel color 1 or 0 (1:white; 0:black).\n>> ")
					fmt.Scan(&input_paint)
					if input_paint < 0 || input_paint > 1 {
						clear()
						prtErr("pixel color out of range", "now exiting to main menu")
						sep()
						break
					}
					canvas[input_row][input_col] = input_paint
					clear()
					fmt.Println("successfully painted pixel.")
					sep()
				} else if input_canvas == 2 {
					file, _ := os.Create("blindart.txt")
					defer file.Close()
					var content string
					for i := range canvas {
						for j := range canvas[i] {
							if canvas[i][j] == 1 {
								content += "■"
							} else {
								content += "□"
							}
						}
						content += "\n"
					}
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
			fmt.Println("blindpaint version " + VERSION + ".\ncrafted with much love by greg <3")
			sep()
		} else if input == 3 {
			clear()
			fmt.Println(byemsgs[randomIndex])
			os.Exit(0)
			sep()
		} else {
			clear()
			prtErr("input exception", "no option made for input "+strconv.Itoa(input))
			sep()
		}
	}
}
