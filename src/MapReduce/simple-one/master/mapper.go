package master

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"

	log "github.com/segmentio/go-log"
)

type MapInput struct {
	Filename string
	Nr       int
}

func mapper(nr int, fileDir string) {
	for {
		val, ok := <-MapChanIn
		if !ok { // channel close
			break
		}
		inputFilename := val.Filename
		nr := val.Nr
		file, err := os.Open(inputFilename)
		if err != nil {
			errMsg := fmt.Sprintf("Read file(%s) error in mapper(%d)", inputFilename, nr)
			logger.Error(errMsg)
			MapChanOut <- ""
			continue
		}

		mp := make(map[string]int)
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			str := scanner.Text()
			//logger.Info(str)
			mp[str]++
		}
		outputFilename := path.Join(fileDir, "mapper-output-"+strconv.Itoa(nr))
		outputFileHandler, err := os.Create(outputFilename)
		if err != nil {
			errMsg := fmt.Sprintf("Write file(%s) error in mapper(%d)", outputFilename, nr)
			log.Error(errMsg)
		} else {
			for k, v := range mp {
				str := fmt.Sprintf("%s %d\n", k, v)
				outputFileHandler.WriteString(str)
			}
			outputFileHandler.Close()
		}

		MapChanOut <- outputFilename

	}
}
