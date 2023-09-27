// meatcounter/meatcounter_test.go
package meatcounter

import (
    "reflect"
    "testing"
)

func TestCountMeatNames(t *testing.T) {
    // Input text with meat names
    inputText := "Fatback t-bone t-bone, pastrami  ..   t-bone.  pork, meatloaf jowl enim.  Bresaola t-bone."

    // Expected result
    expectedResult := map[string]int{
        "t-bone":   4,
        "fatback":  1,
        "pastrami": 1,
        "pork":     1,
        "meatloaf": 1,
        "jowl":     1,
        "enim":     1,
        "bresaola": 1,
    }

    // Call the CountMeatNames function
    result := CountMeatNames(inputText)

    // Compare the result with the expected result
    if !reflect.DeepEqual(result, expectedResult) {
        t.Errorf("CountMeatNames(%s) = %v, expected %v", inputText, result, expectedResult)
    }
}

