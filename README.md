Well hello there.
This is a hacky hackity hack miniscule project to benchmark the performance of protobuf.
As a baseline i also measured a json version and a simple post version.
All versions contained the same data, and were tested the same way.

Results:
Protobuf was clearly the winner on my stack, using 21 bytes for each request against the 29(i think) bytes of the simple POST and the 60 (!) total bytes of the same JSON encoded request
As per the performance gain, the Protobuf also takes the cake, by being 1/3 faster than JSON and 2/3 faster than simple POSTS
