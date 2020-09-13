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

A **Paddle** is what the player controls through the keyboard

A **Ball** is the moving part between the players

### Value Objects

A **Keyboard Event** represents a new state the player wants to set to paddle

### Service

A **G** structure inside *game* package is responsible for holding all game logic. It abstracts the calls to Update and Draw by utilizing the interfaces defined for entities

### Infra Layer

A **EventDispatcher** is responsible for reading *SDL* keyboard events and calling all KeyBoardListener subscribed to the dispatcher.

A **EventPooler** is responsible for pooling *SDL* events and calling the respective dispatcher

## TODOS:

* ✔ Frame Rate Independence
* Score
* Game Over State -> Win/Lose
* Multiplayer
* AI more imperfect
* Window resize