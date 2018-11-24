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