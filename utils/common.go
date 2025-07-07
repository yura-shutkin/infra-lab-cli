package utils

import (
	"os/exec"
	"strconv"
	"strings"
)

func IsBinaryInPath(binary string) bool {
	_, err := exec.LookPath(binary)
	if err != nil {
		return false
	}
	return true
}

func ConvertToMiB(size string) (convertedSize string, err error) {
	// Could be 2048, 2048M, 2048m, 2G, 2.5G, 2g
	result, err := strconv.Atoi(size)
	if err != nil {
		coefficient := 1.0
		if strings.Contains(strings.ToLower(size), "g") {
			coefficient = 1024.0
			size = strings.Replace(strings.ToLower(size), "g", "", 1)
		}
		if strings.Contains(strings.ToLower(size), "m") {
			size = strings.Replace(strings.ToLower(size), "m", "", 1)
		}
		var sizeFloat float64
		sizeFloat, err = strconv.ParseFloat(size, 32)
		if err != nil {
			return "0", err
		}
		return strconv.Itoa(int(sizeFloat * coefficient)), nil
	}
	return strconv.Itoa(result), nil
}

func ConvertMiBToGiB(size int) (convertedSize float64) {
	// Expect only integer
	return float64(size) / 1024.0
}
