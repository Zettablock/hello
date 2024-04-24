# Example of pipelines to run in zrunner, which is Zettablock FaaS platform to stream onchain data

Some explanations:

1, one yml file define one pipeline, which has source, destination, metadata, handlers information

2, Multiple go files are built into one golang plugin file.

So all handlers should have different names. The best practice of naming is Handler<Event name>, HandleBlock, etc

event handlers and block handlers can be in separate files.

3, All handlers which have the same source and destination can be in one pipeline.

4, dao folder contains DAO to access destination tables.



