package raindrops

import "strconv"

func isDivisibleByThree(number int) bool {
    if number % 3 == 0 {
        return true
    } else{
        return false
    }
}

func isDivisibleByFive(number int) bool {
    if number % 5 == 0 {
        return true
    } else{
        return false
    }
}

func isDivisibleBySeven (number int) bool {
    if number % 7 == 0 {
        return true
    } else {
        return false
    }
}

func Convert(number int) string {
	res := ""
    if isDivisibleByThree(number) {
        res = res + "Pling"
    }
    if isDivisibleByFive(number) {
        res = res + "Plang"
    } 
    if isDivisibleBySeven(number) {
        res = res + "Plong"
    }
    if !(isDivisibleByThree(number) || isDivisibleByFive(number) || isDivisibleBySeven(number)) {
        res = strconv.Itoa(number)
    }
    return res
}
