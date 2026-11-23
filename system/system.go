package system

import(
	"os"
	"github.com/fatih/color"
	"fmt"
)

func displayFilePermission(entry os.DirEntry) string{
	file, err := entry.Info()

	if err != nil{
		panic(err)
	}

	perm := ("permission: " + file.Mode().String())
	return perm
}

func displayFileType(entry os.DirEntry) string {
	file , err := entry.Info()

	if err != nil{
		panic(err)
	}

	switch t := file.Mode();{
	case t.IsRegular():
	   return "regular"
	case t.IsDir():
	   return "directory"
	}

	return "unknown"
}

func displayFileSize(entry os.DirEntry) string{
	file, err := entry.Info()
	if err != nil{
		panic(err)
	}

	size:= file.Size()
	
	const(
		gb int64 = 1024 * 1024 * 1024
		mb int64 = 1024 * 1024
		kb int64 = 1024
	)

	switch { 
	case size >= gb:
	   return fmt.Sprintf("%.2f Gb", float64(size/gb))
	case size >= mb:
	   return fmt.Sprintf("%.2f Mb", float64(size/mb))
	case size > kb:
	   return fmt.Sprintf("%.2f Kb", float64(size/kb))
	default:
	   return fmt.Sprintf("%d B", size)
	}

	return "unknown size"
}

func prompt(what string, showPerm bool, showTypes bool, showSize bool){

	files, err := os.ReadDir(what)
	if err != nil{
		panic(err)
	}

	for _, file := range files{

		str := ""

		if showPerm{
			str += displayFilePermission(file)
		}

		if showTypes{
			str += " " + displayFileType(file)
		}

		str += " " + file.Name()

		if showSize{
			str += " " + displayFileSize(file)
		}

		color.RGB(255, 255, 255).Println(str)
	}

}

func Display(what string, showPerms bool, showTypes bool, showSize bool){
	prompt(what, showPerms, showTypes, showSize)
}


