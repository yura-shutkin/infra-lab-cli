package utils

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func IsBinaryInPath(binary string) bool {
	_, err := exec.LookPath(binary)
	return err == nil
}

func BinaryNotFoundError(binary string) error {
	return fmt.Errorf("%s not found", binary)
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

func ByteCountIEC(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}
