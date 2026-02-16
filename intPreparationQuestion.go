package main


func main(){
go value()
}
var result int

func value(){
	result = 10
}()
fmt.Println(result)

using channels communication becomes easy no ded locks.constchannel
ch := make(chan int) --UnBuffered channel ch :=make (chan int,10) --Buffered channel
go func(){
	ch <-10
}()result:= <-ch
fmt.Println(result)

//bidirecly channels sender sends only reciever receives only
func Sender(ch chan<-int){
	ch <-100

}
func Reciever(ch <-chan int ){
        <-ch
}
func main (){
	ch :=make (chan int)
 go Sender(ch)
	go Reciever(ch)
}

//DEAFLOACK:
deadlock happens when go rountes waits infinitely for each other they cannor proceed further
func main(){
	cc1:=make(chan int)
	go func(){
		cc1 <-10
	}
}

//here dead lock happen because there is not reciver for cc1 channel, it will infiitely wauts untill recives the value
//race condition: 
race condition will aper when multiple go routines access the same variable at the same time
//test:
go supports inbuilt testing package testing package
 TestAdd(t *Testing.T){
	results:= Add(2,4)
	if results !=6 {
		t.Errorf("expected 6, but got %d",results)
	}
 }

 // go provieds the comma ok syntax to check if the map contain key
 val,ok:=myMap["id"]
 if ok {
 	fmt.Println(val)
 }