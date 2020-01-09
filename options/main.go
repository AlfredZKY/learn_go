package main

func main(){
	srv := NewService(
		Name("golang"),
		Age(10),
	)
	
	srv.Output()
}

