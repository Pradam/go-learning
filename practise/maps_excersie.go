package main

import (
         "fmt"
         "strings"
        )


func main(){

var sentence = "Pradam is goodboy and awesome dude."

string_array := strings.Fields(sentence)

var counter_letters = map[string]int{}

for i:=int(0); i < len(sentence); i++{
        var val string = string(sentence[i])
        if (val != " ") {
            counter_letters[string(sentence[i])] += 1
       }
    }

counter_sentences := make(map[string]int)

for j:=int(0); j < len(string_array); j++{
    counter_sentences[string_array[j]] += 1
}

for i, v := range counter_letters{
   fmt.Println(i,v)
   counter_sentences[i] = v
}

fmt.Println(counter_letters, counter_sentences)

fmt.Println(strings.Split(sentence, " "))
}