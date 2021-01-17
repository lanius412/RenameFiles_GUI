package rename

import(
	"path/filepath"
	"io/ioutil"

	"strings"
	"strconv"

	"os"

	"log"
)

func Rename(dir string) {

	basePath := filepath.Dir(dir) + "/" + filepath.Base(dir) + "/"

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	fileNum := len(files)

	log.Println(strconv.Itoa(fileNum-1) + "Files")

	if fileNum >= 1000 {
		log.Println("Sry, Support up to 3digits")
		os.Exit(0)
	}

	var fileCount int
	for _, file := range files {
		if file.Name()[:1] == "." { //Remove Hidden File
			continue
		} else {
			filePath := basePath + file.Name()
			renamedPath := basePath

			extPos := strings.LastIndex(file.Name(), ".")
			ext := file.Name()[extPos:] //ex).txt

			fileCount++
			if fileNum >= 100 && fileNum < 1000{ //Hundreds of Files
				renamedPath = ChangeNameandPath(fileCount, renamedPath, ext)
			} else if fileNum >= 10 { //Dozens of Files
				renamedPath = ChangeNameandPath(fileCount, renamedPath, ext)
			} else if fileNum < 10 { //Several Files
				renamedPath += "0" + strconv.Itoa(fileCount) + ext
			}

			log.Println(file.Name() + " -> " + strconv.Itoa(fileCount) + ext)

			os.Rename(filePath, renamedPath)
		}
	}

	log.Println(strconv.Itoa(fileNum-1) + " Files Rename Completed")

}

func ChangeNameandPath(fileCount int, renamedPath string, ext string) string{
	switch {
		case fileCount >= 100:
			renamedPath += strconv.Itoa(fileCount) + ext //3digits
		case fileCount >= 10:
			renamedPath += "0" + strconv.Itoa(fileCount) + ext //2digits
		case fileCount < 10:
			renamedPath += "00" + strconv.Itoa(fileCount) + ext //1digit
	}
	return renamedPath
}