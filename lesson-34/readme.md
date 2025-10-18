# Contexts 

* Context allows you to signal that the operation can be cancelled. 
* You can associate a deadline with a context, specifying the maximum time for an operation to complete. 
* Context can be used to carry arbitrary data along with the function calls. 

## Uses 
* HTTP Request context can carry values through middlewares and handlers. 
* Database operations can rely on contexts for cancellation. 
* Long running goroutines can be cancelled with contexts. 