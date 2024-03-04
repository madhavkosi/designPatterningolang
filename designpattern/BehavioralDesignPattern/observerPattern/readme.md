<!-- 
Observer is a behavioral design pattern that allows some objects to notify 
other objects about changes in their state.
 -->


<!-- In the e-commerce website, items go out of stock from time to time. There can be customers who are interested in a particular item that went out of stock. There are three solutions to this problem:

The customer keeps checking the availability of the item at some frequency.
E-commerce bombards customers with all new items available, which are in stock.
The customer subscribes only to the particular item he is interested in and gets notified if the item is available. Also, multiple customers can subscribe to the same product.
Option 3 is most viable, and this is what the Observer pattern is all about. The major components of the observer pattern are:

Subject, the instance which publishes an event when anything happens.
Observer, which subscribes to the subject events and gets notified when they happen. -->