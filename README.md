# event-management-system
Cloud-Native-programming with Golang . SaaS application for event management
```
Microservice architecture :

We will use a ReactJS frontend to interface with the users of our applications. The ReactJS UI will use an API gateway (AWS  or local) to communicate with the different microservices that form the body of our application. There are two main microservices that represent the logic of MyEvents:
Event Service: This is the service that handles the events, their locations, and changes that happen to them
Booking Service: This service handles bookings made by users
All our services will be integrated using a publish/subscribe architecture based on message queues. Since we aim to provide you with practical knowledge in the world of microservices and cloud computing, we will support multiple types of message queues. We will support Kafka, RabbitMQ, and SQS from AWS.

The persistence layer will support multiple database technologies as well, in order to expose you to various practical database engines that empower your projects. We will support MongoDB, and DynamoDB</span>.

All of our services will support metrics APIs, which will allow us to monitor the statistics of our services via Prometheus.

The MyEvents platform is designed in a way that will build strong foundations of knowledge and exposure to the powerful world of microservices and cloud computing