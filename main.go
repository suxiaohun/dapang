package main

import (
	"dapang/cgo"
	"dapang/grpc/face_worker"
	"dapang/lesson"
	"dapang/rsa_tools"
)

func main() {

	face_worker.Run()

	//testAll()
	//lesson.DNS()
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
	lesson.Reverse("123456789")
	lesson.WaitGroupWithUnsafeCounter()
	lesson.WaitGroupWithSafeCounter()

	cgo.HelloWorldWithCMethod()
	cgo.Say()
	cgo.Hey()

	rsa_tools.GenerateSignInfo()
	rsa_tools.GenerateRsa()
	rsa_tools.GenerateSignature()

	//lesson.PicGifOnBrowser()
	//lesson.Http()
	//aliyun.GetObjectToFile()

	//lesson.FetchUrl1()
	//lesson.FetchUrl2()
	//lesson.FetchUrl3()
	//lesson.FetchUrl4()
	//lesson.FetchUrl5()

	//cgo.HelloWorldWithCMethod()
	//cgo.Hello2()
}
