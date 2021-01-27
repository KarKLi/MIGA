package param

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	// configMap defines the variable name in configuration file.
	configMap map[string]struct{} = map[string]struct{}{
		// Hyperparam.epoch
		"epoch": {},
		// Hyperparam.batch
		"batch": {},
		// Hyperparam.subdivisions
		"subdivisions": {},
		// Hyperparam.height
		"height": {},
		// Hyperparam.width
		"width": {},
		// Hyperparam.channels
		"channels": {},
		// Hyperparam.momentum
		"momentum": {},
		// Hyperparam.decay
		"decay": {},
		// Hyperparam.learningRate
		"lr": {},
		// Hyperparam.coreNum
		"cpus": {},
		// IPParam.grayScale
		"grayscale": {},
		// IPParam.vfilping
		"vfilp": {},
		// IPParam.hfilping
		"hfilp": {},
	}
)

// TextHpParser defines the hyperparameter parser of TEXT(.txt) file
type TextHpParser struct{}

// ParseHyperParam parses the input configuration file and returns the struct.
func (t *TextHpParser) ParseHyperParam(hyperparamFile string) (*Hyperparam, error) {
	var hp *Hyperparam
	file, err := os.Open(hyperparamFile)
	if err != nil {
		return nil, err
	}

	// Allocate memory for hp
	hp = new(Hyperparam)

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
		err = parseHyperParam(hp, key, value)
		if err != nil {
			return nil, err
		}
	}
	return hp, nil
}

func parseHyperParam(hp *Hyperparam, key string, value string) error {
	var err error
	if key == "epoch" {
		var epoch int
		epoch, err = strconv.Atoi(value)
		if err == nil {
			hp.epoch = epoch
		}
	} else if key == "batch" {
		var batch int
		batch, err = strconv.Atoi(value)
		if err == nil {
			hp.batch = batch
		}
	} else if key == "subdivisions" {
		var subdivisions int
		subdivisions, err = strconv.Atoi(value)
		if err == nil {
			hp.subdivisions = subdivisions
		}
	} else if key == "height" {
		var height int
		height, err = strconv.Atoi(value)
		if err == nil {
			hp.height = height
		}
	} else if key == "width" {
		var width int
		width, err = strconv.Atoi(value)
		if err == nil {
			hp.width = width
		}
	} else if key == "channels" {
		var channels int
		channels, err = strconv.Atoi(value)
		if err == nil {
			hp.channels = channels
		}
	} else if key == "momentum" {
		var momentum float64
		momentum, err = strconv.ParseFloat(value, 64)
		if err == nil {
			hp.momentum = momentum
		}
	} else if key == "decay" {
		var decay float64
		decay, err = strconv.ParseFloat(value, 64)
		if err == nil {
			hp.decay = decay
		}
	} else if key == "lr" {
		var lr float64
		lr, err = strconv.ParseFloat(value, 64)
		if err == nil {
			hp.learningRate = lr
		}
	} else if key == "cpus" {
		var cpus int
		cpus, err = strconv.Atoi(value)
		if err == nil {
			hp.coreNum = cpus
		}
	}
	if err != nil {
		return fmt.Errorf("parse %s param failed", key)
	}
	return nil
}
