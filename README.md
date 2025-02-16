# Number Sequence Table

## Description
Prints a table of numbers based on the user's preferences.

## Table of Contents
- [Installation](#installation)
- [Usage examples](#usage-examples)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Installation
You can install the app either by downloading the pre-built executable or by building it from the source code.

### Option 1: Install Pre-Built Executable

#### Download and Install
Using `wget`:
```sh
wget -O nst https://github.com/epeleh/number-sequence-table/releases/latest/download/nst
```
Using `curl`:
```sh
curl -L -o nst https://github.com/epeleh/number-sequence-table/releases/latest/download/nst
```

Make the file executable:
```sh
chmod +x nst
```

Run the application:
```sh
./nst
```

---
### Option 2: Install from Source

#### Prerequisites
- Go 1.24.0 installed (Download from [golang.org](https://golang.org/dl/))

#### Steps
```sh
# Clone the repository
git clone https://github.com/epeleh/number-sequence-table
cd number-sequence-table

# Install dependencies
go mod tidy

# Build the project
go build -o bin/nst
```

Run the application:
```sh
./bin/nst
```

## Usage examples

### General usage
```sh
./nst
=> Please give table dimensions (<width>x<height>)
-> 5x4
=> Should I use (P)rime numbers or (F)ibonacci numbers?
-> F
=> Multiplication (*) or Addition (+)
-> M

  1   1   2   3   5
  1   1   2   3   5
  2   2   4   6  10
  3   3   6   9  15


```

### Or you can answer the questions using command arguments
```sh
./nst 5x4 F M
=> Please give table dimensions (<width>x<height>)
-> 5x4
=> Should I use (P)rime numbers or (F)ibonacci numbers?
-> F
=> Multiplication (*) or Addition (+)
-> M

  1   1   2   3   5
  1   1   2   3   5
  2   2   4   6  10
  3   3   6   9  15


```

### Piped input is also supported
```sh
printf "5x4\nF\nM\n" | ./nst
=> Please give table dimensions (<width>x<height>)
-> 5x4
=> Should I use (P)rime numbers or (F)ibonacci numbers?
-> F
=> Multiplication (*) or Addition (+)
-> M

  1   1   2   3   5
  1   1   2   3   5
  2   2   4   6  10
  3   3   6   9  15


```

## Contributing
1. Fork the repository
2. Create a new branch (`git checkout -b feature-branch`)
3. Commit your changes (`git commit -m 'Add new feature'`)
4. Push to the branch (`git push origin feature-branch`)
5. Open a Pull Request

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact
For issues, suggestions, or contributions, contact:
- **epeleh** - [epeleh@evrone.team](mailto:epeleh@evrone.team)
- GitHub: [epeleh](https://github.com/epeleh)
