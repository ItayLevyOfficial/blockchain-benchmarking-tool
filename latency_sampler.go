/*
This file is responsible for sampling the latency of the transactions.
It will function as an additional worker that waits for a commit response for every
transaction it sends and stores each transaction latency.
*/
package main