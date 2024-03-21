# Zippy 

## Why

To make the lives of window devs easy

when I was looking for a CLI for zipping files in windows this is what I found:

``C:\Program Files\7-Zip\7z.exe" --help``

or powershell from Win CMD:

``
powershell "Compress-Archive input.txt output.zip"
``

I wanted something that was intuitive with minimal options. So I decided to do it myself. 

## How to install

// TODO: add github link


### Windows
git clone [add github link here]

go build -o zippy.exe .

move the binary to the bin folder 

Add the bin folder to the PATH

## How to use it

``zippy [option] inputName outputName``

### To get help write:

``zippy help``

### To zip files:

``zippy -f filename``

or
 
``zippy -f filename .``

or 

``zippy -f filename someOutputNanme``

**Note**: if not given the output file will be output.zip

### To zip directories:

``zippy -d``

or

``zippy -d .``

or

``zippy -d dirName``

or

``zippy -d . output``

or

``zippy -d . .``

or

``zippy -d dirname output``


``zippy -d . output``

**Note**: 

-  input directory name if left blank or with a "." will be treated as the working directory

- output file name if left bank or with a "." will be default to output.zip

### To unzip

in the works

### test the zip

## Contributions

Contributions are welcome. There are 2 things that I would like the focus on:

1. unzip
2. testing the validity of the zip

## Author:

Aarav Sibbal

## Licence

# Released under MIT License

Copyright (c) 2023 Aarav Sibbal.

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.