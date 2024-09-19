package main

func main() {
	channel := multiplex(writer("Hello"), writer("World"), writer("from"), writer("the"), writer("other"), writer("side"))

	for i := 0; i < 10; i++ {
		println(<-channel)
	}
}

func multiplex(channels ...<-chan string) <-chan string {
	outputChannel := make(chan string)

	for _, channel := range channels {
		go func(ch <-chan string) {
			for {
				message := <-ch
				outputChannel <- message
			}
		}(channel)
	}
	return outputChannel
}

func writer(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- text
		}
	}()

	return channel
}
