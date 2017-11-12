# LOGSTASH-FILEIMPORTER
(Written in GO)

Reads files from a directory, sends them to a host:port (logstash) an DELETES the files after success.


## HowTo use
For an example you can run the docker-compose.yml and have a look at logstash/logstash.conf.

### Command
```
docker run -d --name logstash-fileimporter -v $(pwd)/input:/input -e INPUT_DIR=/input -e SLEEP=10 -e LOGSTASH_HOST=127.0.0.1 -e LOGSTASH_PORT=9600 -e FILEENDINGS=csv heckenmann/logstash-fileimporter
```
### Variables
|Var        | Desc                                      |
|-----|-----|
|INPUT_DIR  | Directory where to search files       |
|SLEEP      | Seconds to sleep between searchings for new files |
|LOGSTASH_HOST | IP where logstash is started |
|LOGSTASH_PORT | Port where logstash listens |
|FILEENDINGS | Filetypes to search for (comma separated) |
