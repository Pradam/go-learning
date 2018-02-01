package main

import ("fmt")

type Mapping interface{
    GetMap() map[string]int
}


//Method 1
type PradamMap map[string]int

func (pm *PradamMap) GetMap() map[string]int{
    var pram = *pm
    for key := range pram{
        pram[key] += 3
    }
    return pram
}

//Method 2
type PradamVertex struct{
    X map[string]int
}

func (pv PradamVertex) GetMap() map[string]int{
    for key := range pv.X{
        pv.X[key] += 2
    }
    return pv.X
}

func main(){
    var maps Mapping
    var pradam = map[string]int{"abhi":21,"lash":87,"pradam":98,"tiger":100}
    v := &PradamVertex{pradam}
    pk := PradamMap(pradam)
    maps = *v
    fmt.Println(maps.GetMap())
    maps = &pk
    fmt.Println(maps.GetMap())

}