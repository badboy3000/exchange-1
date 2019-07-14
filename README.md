# crypto currency exchange

## Generate doc

`swag init`

## Features

* [ ] SMS and Google Two-Factor authenticaton
* [ ] KYC Verification
* [ ] High performance matching-engine
* [ ] 使用redis的list来作为时序queue
* [ ] 更新的订单重新入matching-engine
* [ ] 使用[nats queue group](https://nats-io.github.io/docs/developer/concepts/queue.html)来作为消息系统，保证消息只处理一次，但可以部署多个subscriber
* [ ] 合约支持

## Libraries

* [emirpasic/gods](https://github.com/emirpasic/gods) Implementation of various data structures and algorithms in Go

## Why not?

#### [nsq](https://nsq.io/overview/features_and_guarantees.html)

* messages are delivered at least once, messages can be delivered multiple times
* messages received are un-ordered
