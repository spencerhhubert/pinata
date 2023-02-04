# pinata
a super simple message bus and event management system for robotics

## architecture
- the graph
    - the network of information passing through the system
- messages
    - the packets of information being passed around
    - must define a specific message type
    - YAML syntax (like python dictionaries)
    - timers
        - can pass an optional timer to a message that will pause the full delivery of the message once it has arrived at the destination node
    - acknowledgements
        - optional flag to request a return code like 200, 404, etc.
- nodes
    - all things, both real and abstract, in the system such as motors, sensors, computers, buses, etc., are nodes
    - although in the graph they will appear different, multiple nodes might be the same physical device. for example, if a microcontroller on the graph both recieves and publishes messages under certain circumstances
- pub and sub nodes
    - pub nodes can publish information to buses, like computers or sensors
    - sub nodes can subscribe to buses for information, like motors
    - sub nodes don't actually do the work of retrieving the message, it's the bus that delivers it.
        - event-driven. publishers decide how often information gets published. does it happen on a regular basis, like say 100Hz? does it happen according to an event? let the publisher decide
- buses
    - the abstract connective tissue between all the things that need to talk to other things
    - must define what message type they work on
        - if it's not that type, do not accept. maybe send back error 400 but that sounds hard
    - arbitrary fan in/out. you could have a bus with 100 nodes publishing to it and one subbing or vice versa
    - pub nodes can pub to buses
    - sub nodes can sub to buses
- queue
    - all messages are queued
    - messages are passed to a stack and removed from said stack (a list)
- services
    - optionally run on nodes and execute a function on a regular basis, like 1Hz

### thoughts/questions
- encrypt?
- should publisher nodes be able to specify which subscriber node they want to send information to? or should it remain limited to "you can publish to a set of buses and subscribers can subscribe to a set of buses"
- return errors worth the complexity?
