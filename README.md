![Example usage](https://github.com/nmaupu/gocube/blob/master/doc/title.png)

# What is *gocube*

*gocube* is a program to play with the Rubik's Cube ! It is able to:
- display a cube as ASCII, png (full view, top view, colors filtering, 3d)
- scramble the cube with a given algorithm
- generate PNG image given many parameters
- generate PDF with all your favorites algs

# Installation

Download a release, unzip and you are good to go !

# How to use it ?

## Command line help

```
$ ./gocube --help

Usage: gocube [-d] [-s] COMMAND [arg...]

Rubik's cube utilities written in Go

Options:
  -v, --version   Show the version and exit
  -s, --size      Size of the cube (default 3)
  -d, --debug     Enable debug mode

Commands:
  scramble        Scramble with the given algorithm
  reverse         Reverse the given algorithm
  generate        Generate algs
  exportImg       Generates a cube image
  exportPDF       Export as a PDF

Run 'gocube COMMAND --help' for more information on a command.
```

## Example usage

### Generating images

```
$ ./bin/gocube exportImg --help

Usage: gocube exportImg -o [-w] [-t] [-i...] [-p] [-a]

Generates a cube image

Options:
  -o, --output     Image file name (png)
  -w, --width      Generated image width (default 500)
  -t, --viewType   drawing style (3d, top, full, f2l) (default "full")
  -i, --include    Include color (colors: yellow, white, green, blue, red, orange) - default: all colors are included
  -a, --alg        Algorithm to execute
  -p, --preAlg     Algorithm to execute to setup up the cube (use x, y and z to position the needed cube colors on top and sides)
```

### Examples

Example 1: Generating an image with a given algorithm

```
./gocube exportImg -o /tmp/ex1.png -t full -a "R U R' U' B2 U2 F' L' F2 U2"

```
![example 1](https://github.com/nmaupu/gocube/blob/master/doc/ex1.png)

Example 2: Generating an OLL top view

```
$ ./gocube exportImg -w 200 -o /tmp/ex2.png -p "z2" -a "r U2 R' F R' F' R U' R U' r'" -t top -i yellow
```
![example 2](https://github.com/nmaupu/gocube/blob/master/doc/ex2.png)

Example 3: Generate a 3d representation of a F2L

```
$ ./gocube exportImg -o /tmp/ex3.png -p "z2" -a "R U R' U' R U R' U' R U R' U'" -t f2l -w 200
```
![example 3](https://github.com/nmaupu/gocube/blob/master/doc/ex3.png)

### Generating PDF

The command `exportPDF` can generate a **PDF** from a **config file**. The first alg of each entry will be reversed and used to generate an image to display what you should expect to see on the cube. You can also choose what kind of display you need (top view, full view, 3d view).

```
$ ./gocube exportPDF -f config-example -o /tmp/test.pdf
```

- See [the resulting PDF file](doc/oll.pdf) for an example usage.
- See [the example config file](config-example.yaml) for a complete config example.

# PDF config file description

Several formats are supported: JSON, TOML, YAML, HCL, and Java properties
I'm gonna describe YAML file as it is the simplest to use.

Three main sections:
- cube: cube description
- pdf: pdf description
- draw: what to draw on the PDF

## The `cube` section

The `cube` section only have one option: the size of the cube

Example:
```
cube:
  size: 3
```

## The `pdf` section

The `pdf` section contains two parameters:
- the pdf title (written on top of each page)
- a description (no used for now but can be used for documenting what the config file is all about)

Example:
```
pdf:
  title: My own algorithms
  description: My description
```

## The `draw` section

The `draw` section is an array. Each new entry creates a new page in the PDF.
An entry is as follow:
- view: the type of the view to display (top, full, 3d)
- colors: list of the colors to display
- preAlg: setup algorithm to prepare the cube before executing anything
- title: a title for the entry
- spec: contains a name and corresponding algs (see example below)

Example:
```
cube:
  size: 3

pdf:
  title: "CFOP: all 57 OLLs"
  description: All 57 OLLs used for CFOP method

draw:
  - view: top
    colors:
      - yellow
    preAlg: z2
    title: All edges oriented - All corners oriented
    spec:
      - name: OLL 21
        algs:
          - (R U2 R' U') (R U R' U') (R U' R')
      - name: OLL 22
        algs:
          - R U2 (R2 U' R2 U' R2) U2 R
- view: top
    colors:
      - yellow
    preAlg: z2
    title: T shapes - W shapes - Square shapes - P shapes
    spec:
      - name: OLL 33
        algs:
          - (R U R' U') (R' F R F')
      - name: OLL 45
        algs:
          - F (R U R' U') F'
```

# Title images' command

Here are the 4 commands used to generate the images seen on the top of the README file !
```
./gocube exportImg -w 150 -t 3d -p z2 -a "y2 U' L' U' F' R2 B' R F U B2 U B' L U' F U R F'" -o /tmp/title1.png
./gocube exportImg -w 150 -t top -p z2 -a "M U R U R' U' M2 U R U' r'" -o /tmp/title2.png -i yellow
./gocube exportImg -w 150 -t f2l -p z2 -a "R U R' U'" -o /tmp/title3.png
./gocube exportImg -w 150 -t 3d -p z2 -a "M2 E2 S2" -o /tmp/title4.png
```

# Building

```
make deps
make
```
