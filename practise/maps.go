package main 

import ("fmt")

type Vertex struct{
    age int
    gender string
}

var m map[string]Vertex


func main(){
   m = make(map[string]Vertex)
   m["Pradam"] = Vertex{25, "male"}
   m["Bhavya"] = Vertex{23, "female"}
   m["Tom"] = Vertex{21, "male"}
   m["Abhilash"] = Vertex{28, "male"}
   fmt.Println(m["Pradam"].age)
   var v = map[string]Vertex{"Google": Vertex{},
   "Yahoo":{23, "female"}}
   fmt.Println(v["Yahoo"])
   //Get Value from Map
   var value = m["Bhavya"].age
   m["Bhavya"] = Vertex{23, "male"}
   fmt.Println(value)
   delete(m, "Tom")
   fmt.Println(m)
   values, ok := m["Tom"]
   fmt.Println(values, ok)
}