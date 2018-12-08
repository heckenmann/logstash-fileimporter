# LOGSTASH-FILEIMPORTER
(Written in GO)

Reads files from a directory, sends them to a host:port (logstash) and DELETES the files after success.

---

***INFO***

For logstash please use the new logstash-feature to delete files after import: https://www.elastic.co/guide/en/logstash/current/plugins-inputs-file.html#plugins-inputs-file-file_completed_action

---


## Docker Hub
https://hub.docker.com/r/heckenmann/logstash-fileimporter/

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
