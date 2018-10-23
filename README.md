# unilist

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/04e2384682014f69b7d63350a8a8ff85)](https://app.codacy.com/app/SourceFoundry/unilist?utm_source=github.com&utm_medium=referral&utm_content=source-foundry/unilist&utm_campaign=Badge_Grade_Dashboard)
[![Build Status](https://semaphoreci.com/api/v1/sourcefoundry/unilist/branches/master/badge.svg)](https://semaphoreci.com/sourcefoundry/unilist)
[![Build status](https://ci.appveyor.com/api/projects/status/c79jc4vl9jbiki0g/branch/master?svg=true)](https://ci.appveyor.com/project/chrissimpkins/unilist/branch/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/source-foundry/unilist)](https://goreportcard.com/report/github.com/source-foundry/unilist)

unilist is a command line executable that writes lists of Unicode code points out to the standard output stream.  It supports space-, comma-, and newline-delimited lists of Unicode code points that are defined with inclusion and exclusion arguments.  The application was developed to support the creation of Unicode code point lists for use with font subsetting tools.

```
# Example of Basic Latin - Latin Extended A sets, excluding U+0080-009F, comma-delimited

$ unilist -c 0020-017F x0080-009F
```

**Output**:
U+0020,U+0021,U+0022,U+0023,U+0024,U+0025,U+0026,U+0027,U+0028,U+0029,U+002A,U+002B,U+002C,U+002D,U+002E,U+002F,U+0030,U+0031,U+0032,U+0033,U+0034,U+0035,U+0036,U+0037,U+0038,U+0039,U+003A,U+003B,U+003C,U+003D,U+003E,U+003F,U+0040,U+0041,U+0042,U+0043,U+0044,U+0045,U+0046,U+0047,U+0048,U+0049,U+004A,U+004B,U+004C,U+004D,U+004E,U+004F,U+0050,U+0051,U+0052,U+0053,U+0054,U+0055,U+0056,U+0057,U+0058,U+0059,U+005A,U+005B,U+005C,U+005D,U+005E,U+005F,U+0060,U+0061,U+0062,U+0063,U+0064,U+0065,U+0066,U+0067,U+0068,U+0069,U+006A,U+006B,U+006C,U+006D,U+006E,U+006F,U+0070,U+0071,U+0072,U+0073,U+0074,U+0075,U+0076,U+0077,U+0078,U+0079,U+007A,U+007B,U+007C,U+007D,U+007E,U+007F,U+00A0,U+00A1,U+00A2,U+00A3,U+00A4,U+00A5,U+00A6,U+00A7,U+00A8,U+00A9,U+00AA,U+00AB,U+00AC,U+00AD,U+00AE,U+00AF,U+00B0,U+00B1,U+00B2,U+00B3,U+00B4,U+00B5,U+00B6,U+00B7,U+00B8,U+00B9,U+00BA,U+00BB,U+00BC,U+00BD,U+00BE,U+00BF,U+00C0,U+00C1,U+00C2,U+00C3,U+00C4,U+00C5,U+00C6,U+00C7,U+00C8,U+00C9,U+00CA,U+00CB,U+00CC,U+00CD,U+00CE,U+00CF,U+00D0,U+00D1,U+00D2,U+00D3,U+00D4,U+00D5,U+00D6,U+00D7,U+00D8,U+00D9,U+00DA,U+00DB,U+00DC,U+00DD,U+00DE,U+00DF,U+00E0,U+00E1,U+00E2,U+00E3,U+00E4,U+00E5,U+00E6,U+00E7,U+00E8,U+00E9,U+00EA,U+00EB,U+00EC,U+00ED,U+00EE,U+00EF,U+00F0,U+00F1,U+00F2,U+00F3,U+00F4,U+00F5,U+00F6,U+00F7,U+00F8,U+00F9,U+00FA,U+00FB,U+00FC,U+00FD,U+00FE,U+00FF,U+0100,U+0101,U+0102,U+0103,U+0104,U+0105,U+0106,U+0107,U+0108,U+0109,U+010A,U+010B,U+010C,U+010D,U+010E,U+010F,U+0110,U+0111,U+0112,U+0113,U+0114,U+0115,U+0116,U+0117,U+0118,U+0119,U+011A,U+011B,U+011C,U+011D,U+011E,U+011F,U+0120,U+0121,U+0122,U+0123,U+0124,U+0125,U+0126,U+0127,U+0128,U+0129,U+012A,U+012B,U+012C,U+012D,U+012E,U+012F,U+0130,U+0131,U+0132,U+0133,U+0134,U+0135,U+0136,U+0137,U+0138,U+0139,U+013A,U+013B,U+013C,U+013D,U+013E,U+013F,U+0140,U+0141,U+0142,U+0143,U+0144,U+0145,U+0146,U+0147,U+0148,U+0149,U+014A,U+014B,U+014C,U+014D,U+014E,U+014F,U+0150,U+0151,U+0152,U+0153,U+0154,U+0155,U+0156,U+0157,U+0158,U+0159,U+015A,U+015B,U+015C,U+015D,U+015E,U+015F,U+0160,U+0161,U+0162,U+0163,U+0164,U+0165,U+0166,U+0167,U+0168,U+0169,U+016A,U+016B,U+016C,U+016D,U+016E,U+016F,U+0170,U+0171,U+0172,U+0173,U+0174,U+0175,U+0176,U+0177,U+0178,U+0179,U+017A,U+017B,U+017C,U+017D,U+017E,U+017F

## Contents

- [Installation](#installation)
- [Usage](#usage)
- [Issues](#issues)
- [Contributing](#contributing)
- [License](#license)

## Installation

unilist is developed in Go and compiled to the command line executable `unilist` (`unilist.exe` on Windows). A variety of cross-compiled binaries are available for use on Linux, macOS, and Windows systems, or you can download the source and compile the application yourself. Instructions for both approaches follow.

## Installation Approaches

### Approach 1: Install the pre-compiled binary executable file

Download the latest compiled release file for your operating system and architecture from [the Releases page](https://github.com/source-foundry/unilist/releases/latest).

#### Linux / macOS

Unpack the tar.gz archive and move the `unilist` executable file to a directory on your system PATH (e.g. `/usr/local/bin`).  Execute this move with the following command in the root of the unpacked archive:

```
$ mv unilist /usr/local/bin/unilist
```

There are no dependencies contained in the archive.  You can delete all downloaded archive files after the above step.

#### Windows

Unpack the zip archive and move the `unilist.exe` executable file to a directory on your system PATH. See [details here](https://stackoverflow.com/questions/4822400/register-an-exe-so-you-can-run-it-from-any-command-line-in-windows) for more information about how to do this.

There are no dependencies contained in the archive.  You can delete all downloaded archive files after the above step.

### Approach 2: Compile from the source code and install

You must install the Go programming language (which includes the `go` tool) to compile the project from source.  Follow the [instructions on the Go download page](https://golang.org/dl/) for your platform. 

Once you have installed Go and configured your settings so that Go executables are installed on your system PATH, use the following command to (1) pull the master branch of the repository; (2) compile the `unilist` executable from source for your platform/architecture configuration; (3) install the executable on your system:

```
$ go get -u github.com/source-foundry/unilist
```

## Uninstall

The installation includes a single executable binary file.  If you installed with `go get` or added one of the pre-compiled binaries on your system `$PATH` on *.nix systems, you can uninstall with:

```
$ rm $(which unilist)
```

## Usage

```
$ unilist [option] [arg]
```

where `arg` is a single Unicode code point hexadecimal value or a range of hexadecimal values as defined below.

#### How to define included Unicode code points

##### Individual Unicode code points

```
$ unilist 0020
```

The syntax for single Unicode code point values is the hexadecimal value without prefix characters (i.e., do not use `U+`, `u+` or any other commonly used prefix characters to indicate that these are Unicode code point values) and must be at least 4 digits in length.  An example of a valid definition of the space glyph is `0020`.

##### A range of Unicode code points

```
$ unilist 0020-0025
```

The syntax for Unicode code point ranges is the start hexadecimal value, a dash `-`, the end hexadecimal value.  For example, a request to include Unicode code points with values `0020`, `0021`, `0022`, `0023`, `0024`, and `0025` should appear like this on the command line: `0020-0025`.

#### How to define excluded Unicode code points

unilist supports individual code point and code point range exclusions from the list of values with a case-insensitive `x` glyph at the first position of the command line argument.  The hexadecimal value in the argument must use the format described in the section above.

##### Individual Unicode code points

```
$ unilist 0020-0025 x0021
```

Add the `x` at the first character position of your Unicode code point argument to define it for exclusion from the list.  

For example, use the argument `x0021` to exclude the Unicode code point `0021` from the list.

##### A range of Unicode code points

```
$ unilist 0020-0025 x0021-0022
```

Add the `x` at the first character position of your Unicode code point range argument to define the range of Unicode code point values for exclusion from the list.  

For example, use the argument `x0021-0022` to exclude the Unicode code points `0021` and `0022` from the list.  Note that this is a shortcut for behavior that can be specified with multiple individual code point definitions like this: `x0021 x0022`.

### Default Behavior

The unilist default is a space-delimited list of Unicode code point values with the format `U+[value]`.  

Modify the delimiter between Unicode code point values with the use of an option in your command as described below.

### Options

```
  -c, --comma          Use comma delimiter
  -h, --help           Application help
  -n, --newline        Use newline delimiter
      --usage          Application usage
  -v, --version        Application version
```

## Issues

Please [file an issue report](https://github.com/source-foundry/unilist/issues/new) on the repository for any problems that arise with use.

## Contributing

Contributions to this project are welcomed. Please submit your changes as a pull request on this repository.

## License

Licensed under the [MIT License](LICENSE)