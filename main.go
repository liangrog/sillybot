package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	dimension = 5

	faces = []string{"NORTH", "EAST", "SOUTH", "WEST"}
)

func main() {
	execute()
}

func execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("file", "f", "", "Input file")
}

var rootCmd = &cobra.Command{
	Use:   "bot",
	Short: "Silly bot",
	Run: func(cmd *cobra.Command, args []string) {
		scanner := bufio.NewScanner(os.Stdin)
		f := cmd.Flags().Lookup("file").Value.String()
		if len(f) > 0 {
			orders, err := readYml(f)
			if err != nil {
				//log.Fatal("Failed to read given file")
				fmt.Println("Failed to read given file")
			}

			for _, ins := range orders {
				dispatch(ins)
			}
		} else {
			for {
				scan(scanner)
			}
		}
	},
}

func readYml(f string) ([]string, error) {
	var c []string
	content, err := ioutil.ReadFile(f)
	if err != nil {
		return c, err
	}
	err = yaml.Unmarshal(content, &c)
	if err != nil {
		return c, err
	}

	return c, nil
}

func scan(scanner *bufio.Scanner) {
	fmt.Print("Your command: ")
	if scanner.Scan() {
		command := scanner.Text()
		if len(command) == 0 {
			fmt.Println("No command was given!")
		}

		dispatch(command)
	} else {
		log.Fatal("Failed to get command")
	}
}

func dispatch(cmd string) {
	c := strings.Split(cmd, " ")
	switch c[0] {
	case "PLACE":
		if err := place(c[1]); err != nil {
			fmt.Println(err)
		}
	case "MOVE":
		if bot.proc {
			_ = move()
		}
	case "LEFT":
		if bot.proc {
			turnFace("left")
		}
	case "RIGHT":
		if bot.proc {
			turnFace("right")
		}
	case "REPORT":
		report()
	default:
		fmt.Println("Only command PLACE, MOVE, LEFT, RIGHT, REPORT allowed")
	}
}

type sbot struct {
	x    int
	y    int
	f    string
	proc bool
}

var bot sbot

func ifFalling(x, y int) bool {
	if (x >= 0 && x < dimension) && (y >= 0 && y < dimension) {
		return false
	}

	return true
}

func place(i string) error {
	bot = sbot{}

	ins := strings.Split(i, ",")
	x, _ := strconv.Atoi(ins[0])
	y, _ := strconv.Atoi(ins[1])
	if len(ins) != 3 {
		return errors.New("Number of params for place are wrong")
	}

	// validate faces
	foundFace := false
	for _, v := range faces {
		if ins[2] == v {
			foundFace = true
			break
		}
	}

	if !foundFace {
		return errors.New("Invalid orientation")
	}

	if ifFalling(x, y) {
		return errors.New("Falling out")
	}

	bot.x = x
	bot.y = y
	bot.f = ins[2]
	bot.proc = true

	return nil
}

func move() error {
	x := bot.x
	y := bot.y

	switch bot.f {
	case "NORTH":
		y = y + 1
	case "SOUTH":
		y = y - 1
	case "EAST":
		x = x + 1
	case "WEST":
		x = x - 1
	}

	if ifFalling(x, y) {
		return errors.New("Falling out")
	}

	bot.x = x
	bot.y = y

	return nil
}

func turnFace(t string) {
	var index int
	for i, v := range faces {
		if bot.f == v {
			index = i
		}
	}

	switch t {
	case "left":
		ind := index - 1
		if ind < 0 {
			bot.f = faces[len(faces)-1]
		} else {
			bot.f = faces[ind]
		}
	case "right":
		bot.f = faces[(index+1)%len(faces)]
	}
}

func report() {
	fmt.Printf("%d, %d, %s\n", bot.x, bot.y, bot.f)
}
