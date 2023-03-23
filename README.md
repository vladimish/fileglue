# FileGlue

FileGlue is a simple command-line utility that recursively reads and prints the content of all text files in a 
specified directory, with optional custom headers and footers written in pure go. 

## Installation

To install FileGlue, first ensure that you have [Go](https://go.dev/) installed on your system, then run:

```bash
go install github.com/vladimish/fileglue/cmd/fileglue@latest
```

## Usage


```bash
Usage: fileglue [path] [flags]
Flags:
-path        Path to the directory (default "./")
-top         Top is a text that will be printed at the top of any file. Use '@PATH' expression to insert file path. e.g. === @PATH === (default "@PATH:")
-bottom      Footer is a text that will be printed at the bottom of any file. Use '@PATH' expression to insert file path. e.g. === @PATH === (default "===")
-ignore      By default, fileglue declines all the files which contains byte of value less than 40 except some escape sequences. Passing this flag will allow reading all files (default false).
-exclude     Specify a regex pattern to exclude file paths.
```

## Example

To print the contents of all text files in the current directory:


```bash
fileglue
```

To print the contents of all text files in a specific directory:

```bash
fileglue -path /path/to/directory/
```

To print the contents of all files, including binary files, in the current directory:

```bash
fileglue -ignore
```

To print the contents of all text files in the current directory, excluding files that match a specific regex pattern:

```bash
fileglue -exclude ".*\.log$"
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[GPL](https://github.com/vladimish/fileglue/blob/master/LICENSE)