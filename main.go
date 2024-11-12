package main

// https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797
// https://www.fileformat.info/info/unicode/char/1B/index.htm
// https://www.ascii-code.com/

func main() {
	//
	// fmt.Println("\u001b[31m Red")
	// fmt.Println("\x1b[31m Red")
	// fmt.Println("\033[31m Red")

	//
	// fmt.Println("\x1b[30m Black")
	// fmt.Println("\x1b[31m Red")
	// fmt.Println("\x1b[32m Green")
	// fmt.Println("\x1b[33m Yellow")
	// fmt.Println("\x1b[34m Blue")
	// fmt.Println("\x1b[35m Magenta")
	// fmt.Println("\x1b[36m Cyan")
	// fmt.Println("\x1b[37m White")
	// fmt.Println("\x1b[0m Reset")

	//
	// fmt.Println("\x1b[40m Black      \x1b[0m")
	// fmt.Println("\x1b[41m Red        \x1b[0m")
	// fmt.Println("\x1b[42m Green      \x1b[0m")
	// fmt.Println("\x1b[43m Yellow     \x1b[0m")
	// fmt.Println("\x1b[44m Blue       \x1b[0m")
	// fmt.Println("\x1b[45m Magenta    \x1b[0m")
	// fmt.Println("\x1b[46m Cyan       \x1b[0m")
	// fmt.Println("\x1b[47m White      \x1b[0m")

	//
	// fmt.Printf("\x1b[41m                            \x1b[0m\n")
	// fmt.Printf("\x1b[41m            Red             \x1b[0m\n")
	// fmt.Printf("\x1b[41m                            \x1b[0m\n")

	//
	// fmt.Println("\x1b[33mFake download...\x1b[0m")
	// for i := 0; i <= 20; i++ {
	// 	time.Sleep(100 * time.Millisecond)
	// 	fmt.Print("Progress: [")
	// 	for j := 0; j < i; j++ {
	// 		fmt.Print("#")
	// 	}
	// 	fmt.Printf("] %d%%\n", i*5)
	// }
	// fmt.Println("\x1b[32mDownload Completed!\x1b[0m")

	//
	// fmt.Printf("1\n")
	// fmt.Printf("2\n")
	// fmt.Printf("3\n")
	// fmt.Printf("4\n")
	// fmt.Printf("5\n")

	// fmt.Printf("\x1b[31m")
	// fmt.Printf("\r1")
	// time.Sleep(200 * time.Millisecond)
	// fmt.Printf("\r2")
	// time.Sleep(200 * time.Millisecond)
	// fmt.Printf("\r3")
	// time.Sleep(200 * time.Millisecond)
	// fmt.Printf("\r4")
	// time.Sleep(200 * time.Millisecond)
	// fmt.Printf("\r5")
	// time.Sleep(200 * time.Millisecond)
	// fmt.Printf("\r6")
	// time.Sleep(200 * time.Millisecond)
	// fmt.Printf("\r7")
	// time.Sleep(200 * time.Millisecond)
	// fmt.Printf("\r8")
	// time.Sleep(200 * time.Millisecond)
	// fmt.Printf("\r9")
	// time.Sleep(200 * time.Millisecond)
	// fmt.Printf("\r10")
	// fmt.Printf("\x1b[0m\n")

	//
	// fmt.Println("\x1b[33mFake download...\x1b[0m")
	// for i := 0; i <= 20; i++ {
	// 	time.Sleep(100 * time.Millisecond)
	// 	fmt.Print("\rProgress: [")
	// 	for j := 0; j < i; j++ {
	// 		fmt.Print("#")
	// 	}
	// 	fmt.Printf("] %d%%", i*5)
	// }
	// fmt.Println("\x1b[32m\nDownload Completed!\x1b[0m")

}
