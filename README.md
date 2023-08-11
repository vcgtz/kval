# kval CLI Tool

The *kval CLI Tool* is a command-line tool that allows you to store and retrieve values locally on your computer. It provides a simple and convenient way to manage key-value pairs, making it ideal for storing and accessing configuration settings, temporary data, and other types of information that you need to persist locally.

**Features:**
- Easy-to-use: The *kval CLI Tool* offers a user-friendly interface, allowing you to store and retrieve values with straightforward commands.
- Key-Value Storage: It enables you to store data using key-value pairs, making it flexible for various use cases.
- Secure Storage: Your values are stored locally on your computer, ensuring data privacy and security.
Whether you need a convenient way to store configuration settings, API keys, or any other type of data locally, *kval CLI Tool* simplifies the process and provides a reliable solution for managing your stored values.

**NOTE**: This tool is a go-based version of this package: [config-storage-cli](https://github.com/vcgtz/config-storage-cli).

## :arrow_double_down: Installation
To install this tool, you need to have installed Go. Then, run the next command:
```bash
go install github.com/vcgtz/kval@latest
```

Now you can use it by typing:
```text
kval
```

## :arrow_forward: Available Commands
### ⚡ `ls`: list all your stored keys
```text
kval ls
```

### ⚡ `set`: set a new key-value pair
```text
kval set <KEY> <VALUE>
```

### ⚡ `get`: get a value from a key
```text
kval get <KEY>
```

You can use the flag `--copy` or `-c` to copy the value to the clipboard:
```text
kval get <KEY> -c
```

### ⚡ `del`: delete an existing key
```text
kval del <KEY>
```

### ⚡ `clean`: remove all the key-value pairs stored
```text
kval clean
```

## License
[MIT](https://github.com/vcgtz/kval/blob/main/LICENSE)
