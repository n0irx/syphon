# Syphon

A Cli app to save shell commands for latter use

## Installation

### Binaries
Check the release page here: https://github.com/n0irx/syphon/releases

### From source

```bash
git clone https://github.com/n0irx/syphon
cd syphon && go install
```

## Requirements
- Go

## Usage 

### Syphon Add

Description: Add command shell to database  

Input:  

```bash
# syntax
syphon add <command_alias> "<shell_command>" [<category>]

# example
syphon add ssh-ubuntu-server "ssh -i \"my_aws_key.pem\" ubuntu@ec2-xx-xxx-xxx-xxx.compute-1.amazonaws.com" sshs
```

Output:  

```
Command added:

Command:    ssh -i "my_aws_key.pem" ubuntu@ec2-xx-xxx-xxx-xxx.compute-1.amazonaws.com
Alias:      ssh-ubuntu-server
Category:   ssh
```

### Syphon List

Description: List saved command shell from database  

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

### Syphon Exec

Description: Run your saved shell command  

Input:  

```bash
# syntax
syphon exec <alias>

# example: 
syphone exec ssh-ubuntu-server
```

### Syphon Delete

Description: Delete your saved command  

Input:  

```bash
# delete by id
syphon delete <id>

# example 
syphon delete 1

# delete by alias
syphon delete --alias <alias>

# example
syphon delete --alias ssh-ubuntu-server
```

Output:  

```
Command deleted:
alias:  ssh-ubuntu-server
```

### Syphon Help

```bash
syphon <command> --help
```

## Contributing

The main purpose of this repository is to continue evolving Syphon app, making it faster and easier to use. 

### License

Syphon is [MIT licensed](./LICENSE).