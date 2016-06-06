# Amara

Starts and manages an Accumulo instance in Marathon

## Installation

Clone Amara on cluster where you have HDFS and Zookeeper running.

```
git clone https://github.com/mikewalch/amara.git
```

## Configuration

Create a config file and edit it for your environment. 

```
cd amara/conf
cp config.json.example config.json
vim config.json
```

If you have `JAVA_HOME`, `HADOOP_PREFIX`, and `ZOOKEEPER_HOME` set as environment variables
in your shell, these values do not need to be set in `config.json` as Amara will use what is
set in your environment.

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
