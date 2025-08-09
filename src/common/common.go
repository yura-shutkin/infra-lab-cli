package common

import (
	"bufio"
	"fmt"
	"os"
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

func ConvertToMiB(size string) (convertedSize int, err error) {
	// Could be 2048, 2048M, 2048m, 2G, 2.5G, 2g
	result, err := strconv.Atoi(size)
	if err != nil {
		// TODO: consider to switch to switch
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
			return 0, err
		}
		return int(sizeFloat * coefficient), nil
	}
	return result, nil
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

// TODO: is it wise to split this function to 2 different: exec and interactiveExec?

func ExecBinaryCommand(binaryName, args string, showOutput, inputRequired bool, envs []string) (stdout, stderr []string, err error) {
	cmd := exec.Command(binaryName, strings.Split(args, " ")...)
	cmd.Env = append(os.Environ(), envs...)

	if inputRequired {
		cmd.Stdin = os.Stdin
	}

	stdoutPipe, _ := cmd.StdoutPipe()
	stderrPipe, _ := cmd.StderrPipe()

	_ = cmd.Start()

	go func() {
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() {
			stdout = append(stdout, scanner.Text())
			if showOutput {
				fmt.Println(scanner.Text())
			}
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderrPipe)
		for scanner.Scan() {
			stderr = append(stderr, scanner.Text())
			if showOutput {
				fmt.Println(scanner.Text())
			}
		}
	}()

	return stdout, stderr, cmd.Wait()
}

func IfStringInSlice(str string, list []string) bool {
	for _, item := range list {
		if item == str {
			return true
		}
	}

	return false
}
