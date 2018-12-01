## Neox

This is an early stage work in progress that extends the neo4j bolt driver with some useful utlities.

## Examples:

```go
// Instead of

for result.Next() {
    r := result.Record()

    value := r.GetByIndex(0).(float64)
    name := r.GetByIndex(1).(string)
    isActive := r.GetByIndex(2).(bool)

    user := User {
        Value: value,
        Name: name,
        isActive: isActive,
    }
}

// Pass the pointer to the struct
for result.Next() {
    var user User
    err := result.ToStruct(&user)
    if err != nil {
        log.Fatal("that didn't work out")
    }
}

```


## Docker Run
```
docker run \
    --name=neotest \
    --publish=7474:7474 --publish=7687:7687 \
    --volume=$HOME/neo4j/data:/data \
    --volume=$HOME/neo4j/logs:/logs \
    --volume=$HOME/neo4j/import:/var/lib/neo4j/import \
    --env=NEO4J_dbms_memory_pagecache_size=3G \
    --env=NEO4J_dbms_memory_heap_max__size=3G \
    neo4j:latest
```