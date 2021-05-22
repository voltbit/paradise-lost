### Solution description

1. Generator

Optimised for disk access at the expense of memory - the tokens are written in a single disk write call - the problem size allows for this as the data will take less than 100MB, if the task was too large I would have chunked the writes as well to use a limited data buffer.

The generator is implemented with concurrent workers, this is a bad idea if the storage does not support concurrent writes, but I wanted to see the performance of new NVME SSDs relative to concurrent writes. New SSDs usually have multiple memory cells that can be accessed concurrently, but in the end the OS seems to keep a single I/O queue so there is no improvement (on the contrary, there will be more syncs() because of the threads). The performance was degraded by the use of multiple routines so I kept the worker count set on 1. The code **would** however run faster on a RAID setup or similar storage.

In a later refactoring I decided the newlines are not necessary at all since both the encoding and the size of the tokens is given, so the final version of the code optimizes for size as well and does not write any newlines.

The generator was not seeded ( to allow easier observation about performance improvement with different settings ).

2. Token reader

Loads the whole data in memory, then generates a frequency map using a hashmap.

The DB write queries are batched with a configurable batch size. Each batch executes on a separate routine using a worker pool with configurable size.

About the token processor tradeoffs. If the size of the problem is known or estimated ahead of time - say in the range of 10 to 100 mil. entries then trading more memory for less I/O would be best. Allocating a hash map large enough to keep the whole frequency table and then writing it to the database can be a reasonable choice. It also maps well on a map-reduce model where a reducer sends the workload divided as batches of 100 mil. lines, the processor loads it and creates the frequency table in memory, then sends the data to the DB.

Assuming the task is a one time job then it is enough to load the file in memory, make a frequency table with the help of a hash map, then send it in batched writes to the DB. However if this is a repetitive task or a map-reduce process described above, then the DB must also be aware of our intention to have a frequency table (or read the existing DB state which is worse). For this I made the DB use the tokens as keys and added a conflict rule on insert to simply increment the existing frequency value with the incoming value.

The precomputed frequency map might be unnecessary too because the set of possible tokens is much larger that the set in the task so the probability of duplicates is very small - the gain in network I/O in sending only unique tokes vs all tokens is probably not very big. The map is only useful to ensure that in a single batch write there are no two identical entries which would make the write fail.

If the task is too large to do it at once in memory then reading the file in chunks and applying the same principle would work just as well, with a performance hit for multiple disk access (which would likely be negligible relative to network I/O).

3. Database

Nothing worth mentioning in the schema, you can see the table definition in the `migrations` directory.

### How to use

Run a Postgres instance:
```
docker-compose up
```

Execute the code from the root of the repository
```
time go run .
```
Running the code with the above will:

- generate a `data_file` file
- load the file
- write the tokens in the Postgres DB available on localhost:9000

To check the state of the database (if you have psql in the system PATH):
```
PGPASSWORD=tokenpass psql -h localhost -p 9000 -U tokenuser tokendb
PGPASSWORD=tokenpass psql -h localhost -p 9000 -U tokenuser tokendb -c 'select count(*) from tokens'
```

`go run .  14.42s user 1.17s system 24% cpu 1:04.95 total`
