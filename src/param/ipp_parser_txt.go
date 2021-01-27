package param

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// TextIPParser defines the image pre-process parameter parser of TEXT(.txt) file
type TextIPParser struct{}

// ParseIPParam parses the input configuration file and returns the struct.
func (t *TextHpParser) ParseIPParam(IPFile string) (*IPParam, error) {
	var ip *IPParam
	file, err := os.Open(IPFile)
	if err != nil {
		return nil, err
	}

	// Allocate memory for hp
	ip = new(IPParam)

	// Read the text file by line
	scanner := bufio.NewScanner(file)
	// Record the file current line for debugging information.
	var line int

	for scanner.Scan() {
		line++
		str := scanner.Text()
		// Find the '#' symbol and ignore it and its latter character.
		idx := strings.Index(str, "#")
		if idx >= 0 {
			str = str[:idx]
		}
		// ignore all Space
		str = strings.ReplaceAll(str, " ", "")
		// And Tab
		str = strings.ReplaceAll(str, "\t", "")
		// Parse the variable
		var key, value string
		fmt.Sscanf(str, "%s=%s", &key, &value)
		if key == "" {
			return nil, fmt.Errorf("Invaild syntax:%s at Line %d", str, line)
		}
		_, ok := configMap[key]
		if !ok {
			return nil, fmt.Errorf("Invaild syntax:%s at Line %d", str, line)
		}
		err = parseIPParam(ip, key, value)
		if err != nil {
			return nil, err
		}
	}
	return ip, nil
}

func parseIPParam(ip *IPParam, key string, value string) error {
	var err error
	if key == "grayscale" {
		var grayscale float64
		grayscale, err = strconv.ParseFloat(value, 64)
		if err == nil {
			ip.grayScale = grayscale
		}
	} else if key == "vflip" {
		var vfilp float64
		vfilp, err = strconv.ParseFloat(value, 64)
		if err == nil {
			ip.vfilping = vfilp
		}
	} else if key == "hfilp" {
		var hflip float64
		hflip, err = strconv.ParseFloat(value, 64)
		if err == nil {
			ip.hfliping = hflip
		}
	}
	if err != nil {
		return fmt.Errorf("parse %s param failed", key)
	}
	return nil
}
