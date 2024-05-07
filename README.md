# tmpmailgo

Small package that uses [1sec Mail](https://www.1secmail.com/) to generate temporary emails to receive messages

## Install

```
go get github.com/snopan/tmpmailgo
```

## Usage

Simple use case to generate email and receive certain email

```go
// Create a new email
email, err := tmpmailgo.NewEmail()
if err != nil {
  panic(err)
}

// Continously check if message have arrived
var messageID int
tk := time.NewTicker(2 * time.Second)
for range tk.C {

  // Read inbox that contains list of message summary
  inbox, err := email.GetInbox()

  // Check each message whether we have recieved the specifc email
  foundMessage := false
  for _, message := range inbox {
    if message.From == "bob@mail.com" {
      messageID = message.ID
      foundMessage = true
      break
    }
  }

  // Stop wait loop when message is found
  if foundMessage {
  	break
  }
}

// Get more detail with message id
fullMessage, err := email.ReadMessage(messageID)
if err != nil {
	panic(err)
}

```
