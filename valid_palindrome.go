package main

import (
    "unicode"
    "fmt"
)

func isAlphaNumeric(char byte) bool {
    rchar := rune(char)
    return unicode.IsLetter(rchar) || unicode.IsDigit(rchar)
}

func isPalindrome(s string) bool {
    leftPointer := 0
    rightPointer := len(s) - 1

    if s == "" {
        return true
    }

    for leftPointer < rightPointer {

        for leftPointer < len(s) - 1 && !isAlphaNumeric(s[leftPointer]) {
            leftPointer++;
        }

        for rightPointer > 0 && !isAlphaNumeric(s[rightPointer]) {
            rightPointer--;
        }

        if leftPointer >= rightPointer {
            return true
        }

        if unicode.ToLower(rune(s[leftPointer])) != 
            unicode.ToLower(rune(s[rightPointer])) {
            return false
        }

        leftPointer++
        rightPointer--
    }

    return true
}
