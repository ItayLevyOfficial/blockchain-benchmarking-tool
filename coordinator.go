/*
This file is responsible for starting the WebSocket server and handling the incoming workers. 
It will wait for the specified number of workers to connect and then begin the test simultaneously by
distributing the specified rate between them. When measuring latency, it will also start the latency sampler
and give him a small sample of the rate.
*/
