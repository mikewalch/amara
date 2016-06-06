# Amara

Starts and manages an Accumulo instance in Marathon

## Installation

Clone Amara on cluster where you have HDFS and Zookeeper running.

```
git clone https://github.com/mikewalch/amara.git
```

## Configuration

1. Create a config file and edit it for your environment. If you have `JAVA_HOME`, `HADOOP_PREFIX`,
and `ZOOKEEPER_HOME` set as environment variables in your shell, these values do not need to be set
in `config.json` as Amara will use what is set in your environment.

        cd amara/conf
        cp config.json.example config.json
        vim config.json

2. Download an Accumulo release tarball that matches the `AccumuloVersion` set in `config.json` and 
copy it to the `upload/` directory.

        cp accumulo-1.7.1-bin.tar.gz /path/to/amara/upload

## Start and manage Accumulo

Initialize Accumulo
```
cd amara/
./bin/amara init
```

Start Accumulo in Marathon
```
./bin/amara start
```

Destroy Accumulo in Marathon
```
./bin/amara destroy
```
