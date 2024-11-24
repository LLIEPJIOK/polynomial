package polynomial

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/LLIEPJIOK/polynomial/pkg/polynomial"
)

func readPolFile(polFileName string) (*polynomial.Polynomial, error) {
	fileContent, err := os.ReadFile(polFileName)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	polStr := strings.Trim(string(fileContent), " \n\r")

	pol, err := polynomial.FromStr(polStr)
	if err != nil {
		return nil, fmt.Errorf("create polynomial from string %q: %w", polStr, err)
	}

	return pol, nil
}

type Input struct {
	first  *polynomial.Polynomial
	second *polynomial.Polynomial
	deg    int
	op     string
}

func readInputFile(inputFileName string) (Input, error) {
	fileContent, err := os.ReadFile(inputFileName)
	if err != nil {
		return Input{}, fmt.Errorf("read file: %w", err)
	}

	content := strings.Trim(string(fileContent), " \n\r")

	operatorID := strings.IndexAny(content, "+*/^")
	if operatorID == -1 {
		pol, err := polynomial.FromStr(content)
		if err != nil {
			return Input{}, fmt.Errorf("polynomial: %w", err)
		}

		return Input{
			first: pol,
			op:    "reduce",
		}, nil
	}

	firstPol, err := polynomial.FromStr(content[:operatorID])
	if err != nil {
		return Input{}, fmt.Errorf("first operand: %w", err)
	}

	if content[operatorID] == '^' {
		if len(content[operatorID+1:]) == 2 && content[operatorID+1:] == "-1" {
			return Input{
				first: firstPol,
				op:    "^-1",
			}, nil
		}

		deg, err := strconv.Atoi(content[operatorID+1:])
		if err != nil {
			return Input{}, fmt.Errorf("parse degree: %w", err)
		}

		return Input{
			first: firstPol,
			deg:   deg,
			op:    "^",
		}, nil
	}

	secondPol, err := polynomial.FromStr(content[operatorID+1:])
	if err != nil {
		return Input{}, fmt.Errorf("second operand: %w", err)
	}

	return Input{
		first:  firstPol,
		second: secondPol,
		op:     string(content[operatorID]),
	}, nil
}

func writeFile(outputFileName string, pols ...*polynomial.Polynomial) error {
	polsStr := make([]string, len(pols))
	for i, pol := range pols {
		polsStr[i] = pol.String()
	}

	if err := os.WriteFile(outputFileName, []byte(strings.Join(polsStr, " ")), 0o600); err != nil {
		return fmt.Errorf("write file %q: %w", outputFileName, err)
	}

	return nil
}

func process(in Input, polMod *polynomial.Polynomial, outputFileName string) error {
	if in.op == "reduce" {
		pols := in.first.Reduce()
		if err := writeFile(outputFileName, pols...); err != nil {
			return fmt.Errorf("write file: %w", err)
		}

		return nil
	}

	in.first.ToMod(polMod)

	if in.second != nil {
		in.second.ToMod(polMod)
	}

	var ans *polynomial.Polynomial

	switch in.op {
	case "+":
		ans = polynomial.Add(in.first, in.second, polMod)

	case "*":
		ans = polynomial.Multiply(in.first, in.second, polMod)

	case "/":
		ans = polynomial.Del(in.first, in.second, polMod)

	case "^":
		ans = polynomial.Pow(in.first, in.deg, polMod)

	case "^-1":
		ans = polynomial.Inv(in.first, polMod)
	default:
		NewErrUnknownOp(in.op)
	}

	if err := writeFile(outputFileName, ans); err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}

func Start() error {
	var (
		polFileName    string
		inputFileName  string
		outputFileName string
	)

	flag.StringVar(&polFileName, "pol", "polynom.txt", "specify file with polynomial for module")
	flag.StringVar(&polFileName, "p", "polynom.txt", "specify file with polynomial for module")

	flag.StringVar(&inputFileName, "input", "input.txt", "specify input operation file")
	flag.StringVar(&inputFileName, "i", "input.txt", "specify input operation file")

	flag.StringVar(&outputFileName, "output", "output.txt", "specify output file for polynomial operation")
	flag.StringVar(&outputFileName, "o", "output.txt", "specify output file for polynomial operation")

	flag.Parse()

	in, err := readInputFile(inputFileName)
	if err != nil {
		return fmt.Errorf("readInputFile(%q): %w", inputFileName, err)
	}

	var polMod *polynomial.Polynomial

	if in.op != "reduce" {
		var err error

		polMod, err = readPolFile(polFileName)
		if err != nil {
			return fmt.Errorf("readPolFile(%q): %w", polFileName, err)
		}
	}

	if err := process(in, polMod, outputFileName); err != nil {
		return fmt.Errorf("calculating: %w", err)
	}

	return nil
}
