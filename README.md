# event-management-system
Cloud-Native-programming with Golang . SaaS application for event management
```
Microservice architecture :
There are two main microservices that represent the logic of MyEvents:
Event Service: This is the service that handles the events, their locations, and changes that happen to them
Booking Service: This service handles bookings made by users
All our services will be integrated using a publish/subscribe architecture based on message queues. Since we aim to provide you with practical knowledge in the world of microservices and cloud computing, we will support multiple types of message queues. We will support Kafka, RabbitMQ, and SQS from AWS.

The persistence layer will support multiple database technologies as well, in order to expose you to various practical database engines that empower your projects. We will support MongoDB, and DynamoDB</span>.

All of our services will support metrics APIs, which will allow us to monitor the statistics of our services via Prometheus.

The MyEvents platform is designed in a way that will build strong foundations of knowledge and exposure to the powerful world of microservices and cloud computing
Event collaboration describes an architectural principle that works well together with an event-driven publish/subscribe architecture.

Consider the following example that uses the regular request/reply communication pattern—a user requests the booking service to book a ticket for a certain event. Since the events are managed by another microservice (the EventService), the BookingService will need to request information on both the event and its location from the EventService. Only then can the BookingService check whether there are still seats available and save the user's booking in its own database. The requests and responses required for this transaction are illustrated in the following diagram:
```
![image](https://github.com/haitham911/event-management-/blob/main/img/1.png)

```
requests and responses

Now, consider the same scenario in a publish/subscribe architecture, in which the BookingService and EventService are integrated using events: every time data changes in the EventService, it emits an event (for example, a new location was created, a new event was created, an event was updated, and so on).

Now, the BookingService can listen to these events. It can build its own database of all currently existing locations and events. Now, if a user requests a new booking for a given event, the BookingService can simply use the data from its own local database, without having to request this data from another service. Refer to the following diagram for another illustration of this principle:
```
![image](https://github.com/haitham911/event-management-/blob/main/img/2.png)
