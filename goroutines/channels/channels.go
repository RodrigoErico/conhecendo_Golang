package main


func main() {
	func write(text string) {
		for i := 0; i < 5; i++ {
			fmt.Println(text)
			time.Sleep(time.Second)
		}
}