# atari_pingpong_go

Atari Ping Pong Golang Implementation using a DDD approach

Following the learning path [Here](https://gameswithgo.org/topics.html)

## Setup SDL development

[Youtube video reference](https://www.youtube.com/watch?v=OXSMx45kayw&list=PLDZujg-VgQlZUy1iCqBbe5faZLMkA3g2x&index=7&ab_channel=JackMott)

After setup, a very basic test application can be used to validate instalation.

```sh
go run ./app/sdl2/main.go
```

## Execute actual game

```sh
go run ./app/pingpong/main.go
```

## Domain Layer

### Entities

A **paddle** is what the player controls through the keyboard

A **ball** is the moving part between the players

### Value Objects

A **Keyboard Event** represents a new state the player wants to set to paddle

### Infra Layer

A **Dispatcher** is responsible for reading *SDL* keyboard events and calling all KeyBoardListener subscribed to the dispatcher


