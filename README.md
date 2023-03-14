# File System Crawler

A file system crawler command-line tool which crawls into file system directories looking for specific files. It has the following features:

- Finds files based on criteria like size, modified date and file extension
- Lists files
- Compresses files (in .gz format) before delete. Files are archived relative to source directory path
- Deletes files
- Logs deleted files to a log file

## Attention!

Be careful when trying this tool on your system. The files will be deleted without any prompt or user confirmation. Never run this tool as a privileged user such as root or Administrator because it can cause irreversible damage to your system.

## Usage

### List all files in directory (e.g. /tmp/dir)

```go
-root /tmp/dir/ -list
```

### List only log files in directory

```go
-root /tmp/dir/ -ext .log -list
```

### List log and txt files in directory 

```go
-root /tmp/dir/ -ext .log|.txt -list
```

### List only files large than 20MB

```go
-root /tmp/dir/ -size 20 -list
```

### List only .gz or .pdf files modifed after 10 March 2023

```go
-root /tmp/dir/ -date 2023-Mar-10 -ext ".gz|.pdf" -list
```

### Delete all log files in directory

```go
-root /tmp/dir/ -ext .log -del
```

### Log file to deleted_files.log log file before delete

```go
-root /tmp/dir/ -ext .pdf -log deleted_files.log -del
```

### Archive txt file to another directory (e.g. dir_archive) before delete

```go
-root /tmp/dir/ -ext .txt -log deleted_files.log -archive /dir_archive -del
```