-- create keyspace
create keyspace if not exists dnslogs
    with replication = {'class': 'SimpleStrategy', 'replication_factor': 1};

-- create tabels
create table dnslogs.dnslogs (
    logtime timestamp primary key,
    sourceip varchar,
    query varchar
);

-- add data