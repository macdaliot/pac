In the ssh dir of this service directory lies the ssh config necessary to run port forwarding to ElasticSearch and DocumentDB
Copy and paste it to your .ssh dir and fill in the missing data
you may then run 'ssh awstunnel' to enable the port forwarding

create or obtain a .env for this service and add the following information

MONGO_CONN_STRING=mongodb://<user>:<password>@<your ip>:27017/<projectname>
PROJECTNAME=<projectName>
ES_ENDPOINT=https://<yourIp>:9200


If you wish to test data on a separate 'test' db, just change the project name. This will create new indices for
both ES and DocDb.