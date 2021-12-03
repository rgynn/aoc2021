package day

import (
	"strconv"
	"strings"
)

func ThreePartOne(input []string) (interface{}, error) {
	var gamma_str, epsilon_str string
	for x := 0; x < len(input[0]); x++ {
		if criteria := getCriteria(x, input); criteria {
			gamma_str += "1"
			epsilon_str += "0"
		} else {
			gamma_str += "0"
			epsilon_str += "1"
		}
	}
	gamma, err := strconv.ParseInt(gamma_str, 2, 32)
	if err != nil {
		return nil, err
	}
	epsilon, err := strconv.ParseInt(epsilon_str, 2, 32)
	if err != nil {
		return nil, err
	}
	return int(gamma) * int(epsilon), nil
}

func ThreePartTwo(input []string) (interface{}, error) {
	o2_readings := copyInput(input)
	co2_readings := copyInput(input)
	var o2_criteria, co2_criteria string
	for x := 0; x < len(input[0]); x++ {
		if criteria := getCriteria(x, o2_readings); criteria {
			o2_criteria += "1"
			co2_criteria += "0"
		} else {
			o2_criteria += "0"
			co2_criteria += "1"
		}
		if len(o2_readings) > 1 {
			o2_readings = filterReadings(o2_readings, o2_criteria)
		}
		if len(co2_readings) > 1 {
			co2_readings = filterReadings(co2_readings, co2_criteria)
		}
	}
	o2, err := strconv.ParseInt(o2_readings[0], 2, 64)
	if err != nil {
		return nil, err
	}
	co2, err := strconv.ParseInt(co2_readings[0], 2, 64)
	if err != nil {
		return nil, err
	}
	return int(o2) * int(co2), nil
}

func getCriteria(position int, input []string) bool {
	zeros, ones := countBitsAtPosition(position, input)
	return zeros <= ones
}

func filterReadings(input []string, criteria string) []string {
	var result []string
	for _, reading := range input {
		if strings.HasPrefix(reading, criteria) {
			result = append(result, reading)
		}
	}
	return result
}

func countBitsAtPosition(position int, input []string) (zeros int, ones int) {
	for _, cmd := range input {
		switch cmd[position] {
		case '0':
			zeros++
		case '1':
			ones++
		}
	}
	return
}
