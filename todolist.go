package main

import "fmt"

type Todo struct{
	Id int
	Item string
	Done bool
}

func (i Todo) GetId() int{
	return i.Id
}

func (i Todo) GetItem() string{
	return i.Item
}

func (i Todo) GetDone() bool{
	return i.Done
}

func AddItem(Id int, Uname string, Udone bool) *Todo{
	i := Todo{}
	i.Id = Id
	i.Item = Uname
	i.Done = Udone
	return &i
}

func (i *Todo) UpdateItem(Uname string){
	if Uname != ""{
		i.Item = Uname
	}	
}

func (i *Todo) MarkComplete(){
	i.Done = true
}

func (i *Todo) MarkIncomplete(){
	i.Done = false
}

func(i *Todo) ToggleDone(){
	i.Done = !(i.Done)
}
 
func main(){

	item1:= Todo{Id:1,Item: "Milk", Done: true }
	fmt.Println(item1)
	fmt.Println(item1.GetDone(), item1.GetId(), item1.GetItem())
	item1.MarkComplete()
	fmt.Println(item1.Done)
	item1.ToggleDone()
	fmt.Println(item1.Done)
	item1.UpdateItem("Titan")
	fmt.Println(item1.GetDone(), item1.GetId(), item1.GetItem())
}