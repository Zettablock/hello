# Example of pipelines to run in zrunner, which is Zettablock FaaS platform to stream onchain data

## Some explanations:

1, one yml file define one pipeline, which has source, destination, metadata, handlers information

2, in the repo base directory, we can have multiple project directories to hold go/yml files separately. Some go files can be in the repo base directory, so they can be shared by all project directories.

3, Multiple go files (in one directory) are built into one golang plugin file. During building customer repo into golang plugins, we copy the go files in the repo base directory into all project directories.

So all handlers (which are built into one plugin) should have different names. The best practice of naming is Handler<Event name>, HandleBlock, etc

4, event handlers and block handlers can be in separate files.

5, All handlers which have the same source and destination can be in one pipeline.

6, dao folder contains DAO to access destination tables.

7, pipeline names should be unique.




