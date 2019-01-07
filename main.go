package main

import "strconv"
import "fmt"

type double float64

func (a double) IsEqual(b double) bool {
	r:=a-b
	if r==0.0{
		return true
	}else if r<0.0{
		return r>-0.0001
	}
	return r<0.0001
}

func IsEqual(a,b float64) bool{
	r:=a-b
	if r==0.0{
		return true
	}else if r<0.0{
		return r>-0.0001
	}
	return r<0.0001
}

type savelog func(msg string)

func stringtoint(s string,log savelog) int64{
	if value,err := strconv.ParseInt(s,0,0);err != nil{
		log(err.Error())
		return 0
	}else {
		return value
	}
}

func mylog(msg string){
	fmt.Println("find error:",msg)
}

func main(){
	var a double =1.999999
	var b double =1.9999998
	fmt.Println(a.IsEqual(b))
	fmt.Println(a.IsEqual(3))
	fmt.Println(IsEqual((float64)(a),(float64)(b)))

	fmt.Println("start1~")
	stringtoint("123",mylog)
	fmt.Println("start2~")
	stringtoint("asda",mylog)
	fmt.Println("start~3")
	println(stringtoint("123",mylog))
	fmt.Println("start~4")
	println(stringtoint("asda",mylog))


	println("note")
	println("note")
}