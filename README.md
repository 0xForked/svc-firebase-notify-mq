## Notify Service

### What's this?
[Firebase Cloud Messaging](https://firebase.google.com) ([Firebase Admin](https://github.com/firebase/firebase-admin-go)) 
implementation for push notification with [golang](https://golang.org/) (Go Programming Language).

### How it's works?

This service subscribe and handle an incoming queue's event from [Message Broker](https://medium.com/@acep.abdurohman90/mengapa-menggunakan-message-broker-c17453cb225e) 
([RabbitMQ](https://www.rabbitmq.com/)) via [amqp](https://www.amqp.org/) connection protocol and sent it to users via 
firebase cloud messaging using firebase admin module for golang.
  


<hr>
<p align="center">
The part of microservices architecture!
</p>