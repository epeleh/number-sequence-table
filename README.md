# Number Sequence Table

## Description
Prints a table of numbers based on the user's preferences.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Installation

### Prerequisites
- Go 1.24.0 installed (Download from [golang.org](https://golang.org/dl/))

### Steps
```sh
# Clone the repository
git clone https://github.com/epeleh/number-sequence-table
cd number-sequence-table

# Install dependencies
go mod tidy

# Build the project
go build -o bin/nst
```

## Usage
```sh
$ ./bin/nst
=> Please give matrix dimension (<width>x<height>)
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
