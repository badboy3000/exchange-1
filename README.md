# crypto currency exchange

## Features

* [ ] SMS and Google Two-Factor authenticaton
* [ ] KYC Verification
* [ ] High performance matching-engine
* [ ] 使用redis的list来作为时序queue
* [ ] 使用nats queue group来作为消息系统，保证消息只处理，但可以部署多个subscriber

## Libraries

* [emirpasic/gods](https://github.com/emirpasic/gods) Implementation of various data structures and algorithms in Go

## Why not?

#### [nsq](https://nsq.io/overview/features_and_guarantees.html)

* messages are delivered at least once, messages can be delivered multiple times
* messages received are un-ordered
