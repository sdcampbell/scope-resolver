package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"regexp"
)

// Function to check if the address is an IP address
func isIP(addr string) bool {
	return net.ParseIP(addr) != nil
}

// Function to process each line of the input
func processLine(line string) {
	ipRegex := regexp.MustCompile(`\b(?:\d{1,3}\.){3}\d{1,3}\b`)
	if ipRegex.MatchString(line) && isIP(line) {
		// It's an IP address, attempt a reverse DNS lookup
		names, err := net.LookupAddr(line)
		if err == nil && len(names) > 0 {
			//trimmed := strings.TrimSuffix(names[0], ".")
			fmt.Printf("%s\t(%s)\n", line, names) // Print the resolved hostname
		} else {
			fmt.Println(line) // If reverse lookup fails, print the IP address
		}
	} else {
		// It's a hostname, print it
		fmt.Println(line)
	}
}

// Prints the usage of the program
func printUsage() {
	fmt.Printf("Usage of %s:\n", os.Args[0])
	fmt.Println("  (no arguments)         Read from stdin")
	fmt.Println("  <file_path>            Read from the specified file")
	fmt.Println("  -h or --help           Print this help message")
	fmt.Println("\nExamples:")
	fmt.Printf("  cat ips.txt | %s\n", os.Args[0])
	fmt.Printf("  %s ips.txt\n", os.Args[0])
}

func main() {
	helpFlag := flag.Bool("help", false, "Print help information")
	flag.BoolVar(helpFlag, "h", false, "Print help information (shorthand)")

	flag.Parse()

	// If help flag is present or no arguments are provided, print usage and exit
	if *helpFlag || (len(os.Args) == 1 && len(flag.Args()) == 0) { // No input provided
		fi, _ := os.Stdin.Stat() // Get the FileInfo struct describing the standard input.
		if (fi.Mode() & os.ModeCharDevice) != 0 {
			printUsage()
			return
		}
	}

	var scanner *bufio.Scanner

	if len(flag.Args()) > 0 {
		filePath := flag.Arg(0) // Get the file path from the command-line argument
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			return
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	} else {
		// Read from stdin
		scanner = bufio.NewScanner(os.Stdin)
	}

	for scanner.Scan() {
		processLine(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error while reading: %v\n", err)
	}
}
