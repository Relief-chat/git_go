package testing

import (
    "bytes"
    "fmt"
    "reflect"
)

func (mf mapFmt) Format(s fmt.State, c rune, userInput string) {
    refVal := mf.m
    key := keys[i]
    val := refVal.MapIndex(key)

    // bad
    meth := key.MethodByName(userInput)
    meth.Call(nil)[0]

    return
}

func Test1(job interface{}, userInput string) {
    jobData := make(map[string]interface{})

    valueJ := reflect.ValueOf(job).Elem()

    // bad
    jobData["color"] = valueJ.FieldByName(userInput).String()

    return jobData
}

func OkTest(job interface{}, userInput string) {
    jobData := make(map[string]interface{})

    valueJ := reflect.ValueOf(job).Elem()

    // ok 
    meth := valueJ.MethodByName("Name")
    // ok 
    jobData["color"] = valueJ.FieldByName("color").String()

    return jobData
}
