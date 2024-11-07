# Pub/Sub System for Events

## What is it?

### Small Example

Taken from [*THIS*](https://medium.com/globant/pub-sub-in-golang-an-introduction-8be4c65eafd4) medium article by Vaibhav Chichmalkar

```go
// channel to publish messages to
msgChannel := make(chan string)

// function to publish messages to the channel
func publishingMessage(message string) {
  msgChannel <- message
}

// function to receive messages from the channel
func receivingMessage() {
  for {
    msg := <-msgChannel
    fmt.Println("Received message:", msg)
  }
}

// goroutine to publish messages
go publishingMessage("Hello from Globant")

// goroutine to receive messages
go receivingMessage()
```