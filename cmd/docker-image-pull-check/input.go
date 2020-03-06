package main

import (
	"os"
	"strings"
)

func parseImages() {
	imagesEnv = os.Getenv("IMAGES")
	images = strings.Split(",")
}
