# Mediatr.Go

Is a library with a small mediator pattern implementation for golang mainly inspired by Mediatr library for .Net

## Installation



```bash
go get github.com/casmelad/Mediatr.Go
```

## Usage

### The coleague

```golang
//Creates the coleague and implements the coleague interface

type MessageHandler struct{}

func (h MessageHandler) IsColleagueFor(r mediatr.RequestMessage) (bool, error) {

	_, ok := r.(MediatrRequestWrapper) //Forces to validate if the coleague can handle the message

	return ok, nil //Quick implementation for the example
}

//Handles the message
func (h MessageHandler) Receive(ctx context.Context, r mediatr.RequestMessage) error {

	data := r.(MediatrRequestWrapper)

	fmt.Println(fmt.Printf("message uuid %s and event data %s", data.GetUUID(), data.EventInfo))

	return nil
}
```
### The coleague interface
```golang
//The coleague interface
type Coleague interface {
	IsColleagueFor(RequestMessage) (bool, error)
	Receive(context.Context, RequestMessage) error
}
```

### The message to handle
```golang
//Then we define the message to handle and the Mediatr wrapper
type MessageToHandle struct {
	EventInfo string
}

type MediatrRequestWrapper struct {
	mediatr.BaseRequestMessage //It is necessary to compose our message with the mediatr message
	MessageToHandle //This is where the real data is going to live, just to keep it clean
}
```
### All the pieces working together
```golang
func main() {

	mediator := mediatr.NewMediator() //Creates the mediatr instance

	mediator.RegisterColeague(MessageHandler{}) //Register our handler as coleague

	msgWithUUID := mediatr.NewRequestWithUUID() //Optional method to add a uuid to the message

	msg := MediatrRequestWrapper{
		BaseRequestMessage: *msgWithUUID,
	}

	//msg.UUID=uuid.New()
	ctx := context.Background()

	mediator.Send(ctx, msg)
	mediator.Send(ctx, msg)

}

```


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change. There are several gaps to cover, let's see how it goes.


## License
[MIT](https://choosealicense.com/licenses/mit/)
