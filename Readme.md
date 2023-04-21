# Motivation:

When it comes to upgrading a blockchain, it is crucial to ensure that it can handle the expected load before rolling out the new version to production. Otherwise, scaling issues and potential downtime can occur, leading to significant losses for businesses and organizations. Therefore, it is essential to have a reliable tool that can benchmark a blockchain section and provide valuable insights into its performance under stress conditions.

Our benchmarking tool offers a comprehensive solution for running end-to-end tests on your blockchain, simulating a high volume of transactions and measuring its transaction latency and throughput. By doing so, you can identify potential bottlenecks and performance issues before deploying the new version to production. Additionally, you can use the tool to monitor the blockchain's performance during the upgrade process to detect any scaling problems.

With our benchmarking tool, you can feel confident that your blockchain can handle the expected load and that you are making an informed decision about upgrading to a new version. Don't leave your blockchain's performance to chance. Use our benchmarking tool to ensure that your blockchain can handle the stress of production and scale smoothly during upgrades.

# How it works?
The tool accept as parameters a list of one or more connections endpoints to the chain and a specified desired rate.
The tps stats are derived by having each connection send transactions at the specified rate (or as close as it can get) for the specified time. After the specified time, it iterates over all of the blocks that were created in that time. The average and standard per second are computed based off of that, by grouping the data by second.

To send transactions at the specified rate in each connection, we loop through the number of transactions. If its too slow, the loop stops at one second. If its too fast, we wait until the one second mark ends. The transactions per second stat is computed based off of what ends up in the block.

Note that there will be edge effects on the number of transactions in the first and last blocks. This is because transactions may start sending midway through when tendermint starts building the next block, so it only has half as much time to gather txs that tm-bench sends. Similarly the end of the duration will likely end mid-way through tendermint trying to build the next block.

Measuring Latency - When load is high, storing the start time and waiting for a commit response for each transaction can slow our client and sqew results. For this reason, we have two load generators. The first is the main load generator, which follows the specified rate load. The second request generator measures latency and has a much lower load; think of it as a single client in comparison to the rest of the system. Even if the system sends back replies to each and every request (as some systems do, such as a KV-store), we can easily drop all replies to the load generator and only measure the latency from the request generator.