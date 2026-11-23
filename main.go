package main

import(
	"system"
	"os"
)

func main(){
	if len(os.Args) < 2{
		panic("an argument must be provided")
	}

	args := os.Args[1:]

	var displayPerm, displayType, displaySize bool
	var dir string

	for _, arg := range args{
		switch arg{
			case "-l":
			   displayPerm = true
			case "-t":
			   displayType = true
			case "-s":
			   displaySize = true
		  default:
		     dir = arg
		}
	}

	system.Display(dir, displayPerm, displayType, displaySize)
}
