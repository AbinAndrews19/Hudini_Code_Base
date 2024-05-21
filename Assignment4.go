// Objective:
// Learn how to send and receive values using channels.
package main

import "fmt"

//"time"
//"os"
// "sync"

// type Waitgroup struct {
// 	count int
// }

// func (c *Waitgroup) Add(num int) {
// 	c.count += num
// }
// func (c *Waitgroup) Done() {
// 	c.count -= 1
// }
// func (c *Waitgroup) Wait() {
// 	for {
// 		if c.count == 0 {
// 			os.Exit(0)
// 		}
// 	}
// }
// func loop(name string, wg *Waitgroup) {
// 	for i := 0; i <= 5; i++ {
// 		time.Sleep(time.Second)
// 		fmt.Println(name, i)
// 	}
// 	wg.Done()
// }
// func main() {
// 	wg := Waitgroup{}
// 	for i := 0; i < 3; i++ {
// 		wg.Add(1)
// 		go loop("Loop", &wg)
// 	}
// 	wg.Wait()
// }

// Task:

// Write a program that creates a goroutine to send a message "Hello, World!" to a channel.
// The main function should receive the message from the channel and print it.
// Hints:
// func msgToChannel(c chan string){
//      c <- "Hello, World!"
// }
// func main(){
//      Message := make(chan string)
// 	go msgToChannel(Message)
//      fmt.Println(<-Message)

// }

// Use an unbuffered channel for simplicity.

// -------------------------------------------------------------------------------------

// Objective:
// Learn how to create and use goroutines.

// Task:

// Write a program that launches three goroutines. Each goroutine should print numbers from 1 to 5 with a delay of 1 second between each number.
// Ensure that the main function waits for all goroutines to finish before exiting.
// Hints:
// Use time.Sleep for delays.
// Use a sync.WaitGroup to wait for all goroutines to complete.
// func printWithDelay( name string,wg *sync.WaitGroup){
//      // create a waitgroup
//      // call three goroutines
//      // wait for all three go rountines
//      for i:=0; i<=5; i++{
//         time.Sleep(time.Second)
//         fmt.Println(name,i)
//      }
//      wg.Done()
// }

// func main(){
//      // Create a waitgroup
//      var wg sync.WaitGroup
//      // call the go routines three times
//      for i:=0;i<3;i++{
//           wg.Add(1)
//           go printWithDelay("Routine",&wg)
//      }
//      wg.Wait()
// }
// // -------------------------------------------------------------------------------------

// Objective:
// Understand how to handle multiple senders with a single receiver using channels.

// Task:

// Write a program where two goroutines send different messages ("Hello" and "World") to the same channel.
// The main function should receive both messages from the channel and print them.
// func execute(name string , ch chan string){
//       ch <- name
//       time.Sleep(time.Second)
// }
// func main(){
//      ch:= make( chan string)
//      go execute("Hello",ch)
//      go execute("World",ch)
//      fmt.Println(<-ch)
//      fmt.Println(<-ch)
// }

// -------------------------------------------------------------------------------------

// Objective:
// Understand how to use channels for communication between goroutines.

// Task:

// Write a program where the main function creates a goroutine that generates numbers from 1 to 10 and sends them to a channel.
// The main function should receive these numbers from the channel and print them.
// Hints:

// Use an unbuffered channel for simplicity.
// Remember to close the channel when done sending values.
// func main(){
//     ch:= make(chan int)
//     go oneToTen(ch)
//     for i:=1;i<=10;i++{
//         fmt.Println(<-ch)
//    } 
// }
// func oneToTen(ch chan int){
//     for i:=1;i<=10;i++{
//       ch <-i
//     }

// close(ch)
// }
// -------------------------------------------------------------------------------------

// Objective:
// Learn how to use a buffered channel.

// Task:

// Write a program that creates a buffered channel with a capacity of 3.
// The main function should send three integers (1, 2, 3) to the buffered channel without blocking.
// Then, receive and print the integers from the channel.
// func main(){
//      message:= make(chan int,3)
//      message <- 1
//      message <- 2
//      message <- 3

//      fmt.Println(<-message)
//      fmt.Println(<-message)
//      fmt.Println(<-message)
// }
// -------------------------------------------------------------------------------------

// Objective:
// Learn how to use the select statement to handle multiple channels.

// Task:

// Write a program that launches two goroutines. Each goroutine sends a series of messages to a channel.
// The main function should use a select statement to receive messages from both channels and print them.
// Hints:

// Use two separate channels.
// Use the select statement inside a loop to receive from the channels.
func main(){
     chan1:= make(chan string)
     chan2:= make(chan string)
     ch1 := true
     ch2 := true

go func(){
     chan1 <- "Sending a series of message for the 1st time"
     chan1 <- "Sending a series of message for the 2nd time"
     chan1 <- "Sending a series of message for the 3rd time"
     chan1 <- "Sending a series of message for the 4th time"
     close(chan1)
}()
go func(){
     chan2 <- "Sending a series of message for the 1st time"
     chan2 <- "Sending a series of message for the 2nd time"
     chan2 <- "Sending a series of message for the 3rd time"
     chan2 <- "Sending a series of message for the 4th time"
     close(chan2)
}()
for {
     select {
     case msg1,ok := <- chan1:
     if !ok{
          ch1 = false
     
     }else{
          fmt.Println(msg1)
     }
case msg2,ok := <- chan2:
     if !ok{
          ch2 = false
     }else{
          fmt.Println(msg2)
     }
     }
   if !ch1 && !ch2{
     break
   }
}
}





// -------------------------------------------------------------------------------------

// Objective:
// Use sync.WaitGroup to wait for multiple goroutines to complete.

// Task:

// Write a program that launches three goroutines, each printing a different message.
// Use a sync.WaitGroup to ensure the main function waits for all goroutines to finish before exiting.
func main(){

}

func printMessage()


// -------------------------------------------------------------------------------------

// Objective:
// Learn how to use the select statement to handle multiple channels.

// Task:

// Write a program that creates two channels and two goroutines. Each goroutine sends a different message to its respective channel.
// Use a select statement in the main function to receive messages from both channels and print them.
