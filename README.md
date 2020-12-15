# Syphon

A Cli app to save shell commands for latter use

## Installation

## From source

```bash
git clone https://github.com/n0irx/syphon
cd syphon && go install
```

## Go Get

TODO

## Requirements
- Go

## Usage 
``

### Syphon Add

Input:  

```bash
syphon add ssh-ubuntu-server "ssh -i \"my_aws_key.pem\" ubuntu@ec2-xx-xxx-xxx-xxx.compute-1.amazonaws.com" sshs
```

Output:  

```
Command added:

Command:    ubuntu@ec2-xx-xxx-xxx-xxx.compute-1.amazonaws.com
Alias:      ssh-ubuntu-server
Category:   ssh
```

### Syphon List

Input:  

```bash
syphon list
```

Output:  

```
+----+---------------------+---------------------------------------------------+----------+
| ID |        ALIAS        |                      COMMAND                      | CATEGORY |
+----+---------------------+---------------------------------------------------+----------+
|  1 | ssh-ubuntu-server   | ssh -i "my_aws_key.pem"                           | ssh      |
|    |                     | ubuntu@ec2-xx-xxx-xxx-xxx.compute-1.amazonaws.com |          |
+----+---------------------+---------------------------------------------------+----------+
```

### Syphon Delete

Input:  

```bash
# delete by id
syphon delete id

# delete by alias
syphon delete --alias ssh-ubuntu-server
```

Output:  

```
Command deleted:
alias:  ssh-ubuntu-server-2
```

### Syphon Help

```bash
syphon <command> --help
```