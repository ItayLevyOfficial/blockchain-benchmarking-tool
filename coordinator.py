"""
This file is responsible for starting the websocket server and handling the incoming 
workers. It will wait for the specified number of workers to connect and then start the
test simoltaneously, by distributing the specified rate between the workers.
"""