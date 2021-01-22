package dataset

import (
	"encoding/json"
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type basicAnnotation struct {
	X string `json:"x"`
	Y string `json:"y"`
	W string `json:"w"`
	H string `json:"h"`
	C string `json:"c"`
}

type eachAnnotation struct {
	Annotations []basicAnnotation `json:"annotations"`
	ImgName     string            `json:"img_name"`
}

type totalAnnotation struct {
	Annotations map[string][]basicAnnotation `json:"annotations"`
	Class       string                       `json:"class"`
}

// JSONChecker a struct for checking the json annotation.
type JSONChecker struct{}

// CheckEachAnno the path is pointed to an individual JSON annotation file.
func (checker *JSONChecker) CheckEachAnno(path string, classLength int) error {
	var data eachAnnotation

	// Read the json file
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return err
	}

	// Now check the image exist or not
	var reader *os.File
	currentPath := path[:strings.LastIndex(path, "/")+1]
	reader, err = os.Open(currentPath + data.ImgName)
	if err != nil {
		return err
	}
	if reader == nil {
		// Image doesn't exist, but we don't treat it as an error.
		//
		// That can reduce the delete useless annotation cost.
		return nil
	}

	// Image exists, recording the width and height, and release the image ASAP.
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		return err
	}

	var width, height int
	width = img.Width
	height = img.Height
	reader.Close()

	return checkEachAnno(data.Annotations, width, height, classLength)
}

// CheckTotalAnno the path is pointed to an total JSON annotation file.
func (checker *JSONChecker) CheckTotalAnno(path string) error {
	var data totalAnnotation

	// Read the json file
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return err
	}

	classLength := len(data.Class)
	for imgName, annos := range data.Annotations {
		// Check the image exists or not
		currentPath := path[:strings.LastIndex(path, "/")+1]
		reader, err := os.Open(currentPath + imgName)
		if err != nil {
			return err
		}
		if reader == nil {
			// Image doesn't exist, but we don't treat it as an error.
			//
			// That can reduce the delete useless annotation cost.
			continue
		}

		// Image exists, recording the width and height, and release the image ASAP.
		img, _, err := image.DecodeConfig(reader)
		if err != nil {
			return err
		}
		width := img.Width
		height := img.Height
		reader.Close()

		err = checkEachAnno(annos, width, height, classLength)
		if err != nil {
			return err
		}
	}
	return nil
}

func checkEachAnno(annos []basicAnnotation, width int, height int, classLength int) error {
	// Now read the annotation orderly.
	for _, anno := range annos {
		// Convert x,y,w,h,c into float
		x, err := strconv.ParseFloat(anno.X, 64)
		if err != nil {
			return err
		}
		y, err := strconv.ParseFloat(anno.Y, 64)
		if err != nil {
			return err
		}
		w, err := strconv.ParseFloat(anno.W, 64)
		if err != nil {
			return err
		}
		h, err := strconv.ParseFloat(anno.H, 64)
		if err != nil {
			return err
		}
		c, err := strconv.Atoi(anno.C)
		if err != nil {
			return err
		}
		if x < 0 || x+w > float64(width) || y < 0 || y+h > float64(height) || c < 0 || c > classLength-1 {
			return fmt.Errorf("Annotation cross over the boundary or class invaild")
		}
	}
	return nil
}
