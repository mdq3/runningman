package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// Frame definitions converted from the frames file
var frame1 = []string{
	"                                         .                     ",
	"                                        `-::-``                 ",
	"                                        sdNNmo+                 ",
	"                                       `dNMMMyy                 ",
	"                                        hNMMMyy`                ",
	"                                     .yhNMMmh:-                 ",
	"                                   `:mMMMMmy                    ",
	"                                   -yMMMMMNd.                   ",
	"                                  `oNMMMMMMNo`                  ",
	"                                  .sMMMMMMMNs`                  ",
	"                                  .sMMMMMMMN/                   ",
	"                                  /hMMMMMMMNy:-                 ",
	"                                 `hNMMMMMMMMMmm.                ",
	"                                 +mMMMMMMMMMMNm.`               ",
	"                                 yNMMMMMMMmyoyy.`               ",
	"                                 hNMMMMMMMd+                    ",
	"                                 yNMMMMMMMmo                    ",
	"                                 /dMMMMMMMh/                    ",
	"                                  -sMMMMMMms                    ",
	"                        -+o:-......oNMMMMMMm-                   ",
	"                       `dNMMNmmmmdhmMMMMMMMNo`                  ",
	"                      .oMMMMMMMMMMMMMMMMMMMNy`                  ",
	"                     .+h/:-/oyyhhhddmMMMMMMm:                   ",
	"                     `.`         ```:mMMMMd+                    ",
	"                                   `:NMMMNs.                    ",
	"                                   :hMMMNo.                     ",
	"                                   +NMMMh-                      ",
	"                                  `oNMMMy-                      ",
	"                                   /dNNNNho:                    ",
	"                                      ``..``                    ",
}

var frame2 = []string{
	"                                          .                     ",
	"                                        `:++/:.                 ",
	"                                       `+mMMNNo.                ",
	"                                       `oNMMMMs-                ",
	"                                       .oNMMMMo.                ",
	"                                     -odNMMho/.                 ",
	"                                    :yMMMMM+.                   ",
	"                                   .sNMMMMMy-                   ",
	"                                  `dNMMMMMMm/`                  ",
	"                                 `+NMMMMMMMN+``   ``            ",
	"                                 `/NMMMMMMMMdyso+oyy`           ",
	"                                  -NMMMMMMMMMMMMMNys.           ",
	"                                 .hMMMMMMMNo:---:-`             ",
	"                                `:MMMMMMMMh.                    ",
	"                                `/MMMMMMMMh`                    ",
	"                                `:MMMMMMMMMo.                   ",
	"                                `-NMMMMMMMMs-                   ",
	"                                 `oNMMMMMMMo-                   ",
	"                                   yNMMMMMMN+.`                 ",
	"                                   /dMMMMMMMdy-                 ",
	"                            `..--:/ymMMMMMMMMM+.                ",
	"                            +hNNNMMMMMMMMMMMMMs-                ",
	"                            +dMNNmMMMMMNmhyyo/.                 ",
	"                            sdy:/+MMMMMs-`                      ",
	"                            oo: :hMMMmy-                        ",
	"                              ./yMMMd+                          ",
	"                              :dmMMm/`                          ",
	"                             `/mNMMd/`                          ",
	"                              -ooydNds/`                        ",
	"                                  `.-:-`                        ",
}

var frame3 = []string{
	"                                          `.```                 ",
	"                                        `smMMNm-`               ",
	"                                        -dMMMMM/`               ",
	"                                        -hNMMMM/`               ",
	"                                      `/yNMMNmy.                ",
	"                                  `-+dmNMMMN+/`                 ",
	"                                 -smMMMMMMMMho                  ",
	"                                `hNMMMMMMMMMNh`    ..`          ",
	"                               `-MMMMMMMMMMMMNo-/sdmm+          ",
	"                                `smMMMMMMMMMMMNNNMMms`          ",
	"                                  odMMMMMMMMMMMMNh+-            ",
	"                                  /dMMMMMMMm+++/.`              ",
	"                                 .hMMMMMMMNy                    ",
	"                                 sNMMMMMMMMm/:                  ",
	"                                `mMMMMMMMMMMmh-                 ",
	"                               `:MMMMMMMMMMMMNs`                ",
	"                               `/MMMMMMMMMMMMMNs-               ",
	"                               .sMMMMNd+odNMMMMNs:              ",
	"                              `:NMMMNo/  .+mMMMMMms/            ",
	"                             `/sMMMNd..`:sdNMMMMMNNs.           ",
	"                             -dmMMMNh:/yMMMMMMMNmdy/            ",
	"                            :yMMMMMMNNNMMMMNhyo:-`              ",
	"                           :hMMMMmmNMMNmho:`                    ",
	"                          .hNMMMN-:yNhs-`                       ",
	"                         `oNMMMmo `om+:                         ",
	"                        -:NMMNds`  `--.                         ",
	"                       `ddMMNh..                                ",
	"                       `osmMdo                                  ",
	"                          .shh:.                                ",
	"                            ````                                ",
}

var frame4 = []string{
	"                                           .+++/-`              ",
	"                                          .dMMMMd/              ",
	"                                          .mMMMMd+              ",
	"                                         `:mMMMMd+              ",
	"                                     .:ohdNMMNms:`              ",
	"                                  `-smNMMMMMMNm/`  ``           ",
	"                                ./odMMMMMMMMMMMh/ `so/.         ",
	"                               `+NMMMMMMMMMMMMMMmhdMho.         ",
	"                                :ymMMMMMMMMMMMMMMMMNo.`         ",
	"                                `.omMMMMMMMMMNmdmNNo.           ",
	"                                  `/NMMMMMMMMy+.`-/`            ",
	"                                   -mMMMMMMMN+.                 ",
	"                                  :yMMMMMMMMNo-`                ",
	"                                `-sNMMMMMMMMMmh+-               ",
	"                               `/mNMMMMMMMMMMMMms.              ",
	"                               -sMMMMMMMMMMMMMMMNd/`            ",
	"                              `smMMMMNmdmmmmMMMMMMmo-`          ",
	"                              smMMMMN/`    `/ydNMMMMdy:`        ",
	"                            .oMMMMMm:        .-smMMMMN+`        ",
	"                          `:yMMMMMdo        `/smMMMMhs-         ",
	"                        `:sdMMMMds:`       `omMMMMNo-           ",
	"                       .ymMMMMd/.        `-hMMMMmy.             ",
	"                      .dNMMMNy.        `./dMMMMh/`              ",
	"                   `-/yMMMMmy.         .ohMMMdy:`               ",
	"                  `:NNMMNdy:           `.:yNm+`                 ",
	"                   .syMm+.                `/ms/`                ",
	"                    -+No.                  `---`                ",
	"                     .+:`                                       ",
}

var frame5 = []string{
	"                                            `-::-.`             ",
	"                                           .yNMMNdo             ",
	"                                           .mMMMMms             ",
	"                                          `-dMMMMms             ",
	"                                      `--oydNMMNy+-             ",
	"                                   `-omNMMMMMMMNo-  `           ",
	"                                  :/hMMMMMMMMMMMmo.:o:.         ",
	"                                `:NNMMMMMMMMMMMMMNdmNo-`        ",
	"                                `-NMMNmNMMMMMMMMMMMMm/`         ",
	"                                 `yhMMNMMMMMMMMMdmMM/.          ",
	"                                  -/MMMMMMMMMMdy-.:/`           ",
	"                                   `oNMMMMMMMMo:`               ",
	"                                   `:NMMMMMMMMo-`               ",
	"                                  `.hMMMMMMMMMmh+-              ",
	"                                 `+oMMMMMMMMMMMMNs.             ",
	"                                 `hdMMMMMMMMMMMMMNd/`           ",
	"                                `-NNMMMMNNmNNMMMMMMmo-          ",
	"                                -hMMMMNo:....+hmNMMMMmh:`       ",
	"                               .oMMMMM+.      `./smMMMMy:       ",
	"                           `.+hNMMMMMd`          .+MMMMy:       ",
	"                       .:ohNNMMMMMMNd-           -sMMMMo.       ",
	"                  ./yyydNMMMMMNmds+/.            :yMMMM+`       ",
	"                  :yMMMMMNmys/.`                 :yMMMM/`       ",
	"                  :yMNmyso/.`                    -yMMMN/`       ",
	"                  /hyo.`                         -sMMMN/`       ",
	"                  .:``                           .oMMMMh+-      ",
	"                                                 `:hhhdmds-     ",
	"                                                        ``      ",
}

var frame6 = []string{
	"                                            ``..``              ",
	"                                           `shmdhy`             ",
	"                                           /mMMMMN.`            ",
	"                                           /mMMMMM-`            ",
	"                                        `-/hNMMMNd`             ",
	"                                    `/shmMMMMMMh/               ",
	"                                  `-smMMMMMMMMMms               ",
	"                                 -sNMMMMMMMMMMMMm+`---`         ",
	"                                 sNMMmNMMMMMMMMMMNymMd/         ",
	"                                 /hMMmNMMMMMMMMMMMMNN/`         ",
	"                                 `+NMMMMMMMMMMMmhmm+:`          ",
	"                                  -yMMMMMMMMMmh/`..`            ",
	"                                   `hNMMMMMMMy-`                ",
	"                                   .mMMMMMMMMy:.                ",
	"                                  -sMMMMMMMMMNms-               ",
	"                                  /NMMMMMMMMMMMmo               ",
	"                                  /NMMMMMMMMMMMNy`              ",
	"                                 -sMMMMMdssNMMMMNd.             ",
	"                                 +dMMMNd/  /dMMMNNh:.           ",
	"                           ``-///hNMMMy/.   -sdNMMMdy.          ",
	"                  `.-//oshmmNMMMMMMMMd/      ./yNMMNN:`         ",
	"                  -ydNMMMMMMMMMNNNNNd:`        /dMMMMo.         ",
	"                  -hmMmhhyyso/:-...`           `/NMMMd/         ",
	"                  :ddh/-..``                    -hMMMNo.        ",
	"                  -s+-                          `-NMMMmo`       ",
	"                                                  +mMMMms+oo/.  ",
	"                                                  `sdMMMMNmm/.  ",
	"                                                   ``ohho-``    ",
	"                                                      ``        ",
}

var frame7 = []string{
	"                                            .                   ",
	"                                          `-//:.`               ",
	"                                        ``smMMMy/`              ",
	"                                        `.yNMMMy+.              ",
	"                                        ..yNMMMy+.              ",
	"                                      .ohdNMMmo-`               ",
	"                                   .:sNMMMMMMd.                 ",
	"                                 ./sdMMMMMMMMM/`                ",
	"                                -hNMMMMMMMMMMMs-                ",
	"                                sNMMMMMMMMMMMMy:````            ",
	"                                -hMMMMMMMMMMMMNmdddms`          ",
	"                                 -omNMMMMMMMMNNmmddy:           ",
	"                                  -sdMMMMMMMm+:-..`             ",
	"                                  /mNMMMMMMNh`                  ",
	"                                 `+NMMMMMMMNh`                  ",
	"                                 -yMMMMMMMMMNo`                 ",
	"                                 :yMMMMMMMMMNs`                 ",
	"                                 .sMMMMMMMMMNo`                 ",
	"                                 :yMMMMMMMMMMN:`                ",
	"                     ```         /hMMMMNmmNMMMd:                ",
	"                    /dmmdhdmmmmmmdNMMMMo:-dMMMMo-`              ",
	"                   -yMMMMMMMMMMMMMMMMMm.` +dMMMhs.              ",
	"                  .ymmhhhmmmmddddmmNmy-   .sNMMNm:              ",
	"                  :/-`   ```     ````      :yMMMMo.             ",
	"                  ``                       ./MMMMh/             ",
	"                                            .hMMMMy-            ",
	"                                             -mMMMmy            ",
	"                                              ymMMMm+-.`        ",
	"                                              `smNNddhs+.       ",
	"                                               ```.```          ",
}

func clearScreen() {
	// Clear screen and move cursor to top-left
	fmt.Printf("\033[H\033[2J")
}

func moveToTop() {
	// Move cursor to top-left corner without clearing screen
	fmt.Printf("\033[H")
}

func clearAnimationArea() {
	// Move to top-left position
	fmt.Printf("\033[H")
}

func hideCursor() {
	// Hide cursor to reduce visual distraction
	fmt.Printf("\033[?25l")
}

func showCursor() {
	// Show cursor when exiting
	fmt.Printf("\033[?25h")
}

func printFrame(frame []string, frameNum int) {
	blueColor := "\033[1;34m"
	resetColor := "\033[0m"
	const maxHeight = 30  // Tallest frame height
	const frameWidth = 50 // Approximate max width to ensure complete coverage

	// Build the entire frame as a single string to reduce flicker
	var frameBuilder strings.Builder
	frameBuilder.WriteString(blueColor)

	// Print frame lines, ensuring each line fills the full width
	for _, line := range frame {
		// Pad each line to full width to ensure complete coverage
		if len(line) < frameWidth {
			frameBuilder.WriteString(fmt.Sprintf("%-*s", frameWidth, line))
		} else {
			frameBuilder.WriteString(line)
		}
		frameBuilder.WriteString("\n")
	}

	// Pad shorter frames with empty lines to reach maxHeight
	for i := len(frame); i < maxHeight; i++ {
		frameBuilder.WriteString(fmt.Sprintf("%*s", frameWidth, ""))
		frameBuilder.WriteString("\n")
	}

	frameBuilder.WriteString(resetColor)

	// Print the entire frame at once
	fmt.Print(frameBuilder.String())
}

func main() {
	// All frames in order
	frames := [][]string{frame1, frame2, frame3, frame4, frame5, frame6, frame7}

	// Setup signal handling for graceful exit
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Hide cursor and clear screen initially
	hideCursor()
	clearScreen()

	// Animation loop
	go func() {
		for {
			for _, frame := range frames {
				select {
				case <-c:
					return
				default:
					clearAnimationArea()
					printFrame(frame, 0)
					time.Sleep(100 * time.Millisecond) // Fast animation
				}
			}
		}
	}()

	// Wait for interrupt signal
	<-c
	// Restore cursor and clear screen on exit
	showCursor()
	clearScreen()
	fmt.Printf("Animation stopped.\n")
}
