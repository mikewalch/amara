# Amara

Start and manages Accumulo instance in Marathon

## Installation

Clone Amara on cluster where you have HDFS and Zookeeper running.

```
git clone https://github.com/mikewalch/amara.git
```

## Configuration

Create a config file and edit for you environment

```
cd amara/conf
cp config.json.example config.json
vim config.json
```

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
