# router
A simple sample Go application for routing binary messages to different handlers based on their contents.

## Disclaimer
This is a work in progress, and will be updated whenever I find the time to refine this solution. 

## Flexibility 
Any kind of listener can be created, as long as it satisfies the `listener.Listener` interface. An example UDP server
has been implemented for demonstrative purposes.