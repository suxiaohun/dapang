package main

import (
	"dapang/lesson"
	"fmt"
)

func main() {

	testAll()
}

func testAll() {
	lesson.HelloWorld()
	lesson.ForStandard(5)
	lesson.ForWithCondition(0)
	lesson.ForInfinite()
	lesson.ForRange()
	lesson.Join()
	lesson.BinaryConversion()
	lesson.GetUUID()
	lesson.Map2JSON()
	lesson.Array1()
	fmt.Println(lesson.Reverse("123456789"))

	//lesson.Dup1()
	//lesson.Dup2()

	//lesson.PicGifOnBrowser()

	//lesson.FetchUrl1()
	//lesson.FetchUrl2()
	//lesson.FetchUrl3()
	//lesson.FetchUrl4()
	//lesson.FetchUrl5()

	//cgo.HelloWorldWithCMethod()
	//cgo.Hello2()
}
