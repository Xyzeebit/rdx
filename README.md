# RDX (README EXPLORER)

RDX is a tool for loading and previewing README files in the browser. Point rdx to a directory and it will walk through the directory and sub directories, finding and parsing README files into HTML contents that can be loaded in the browser.

## FEATURES:
- GO server
- Markdown parser

## USAGE

To load the README files in the current directory

```console
   rdx
```
Load README files in a specify directory and sub directories

```console
  rdx -path ./my-directory
```
Load README with max recursive depth

```console
  rdx -path ./my-directory -d 3
```
To specify a different port for previewing the README files
```console
   rdx -path ./my-directory -p 8000
```
To open and preview any Markdown file run
```console
   rdx -open my-file.md
```

## LICENSE
The MIT License (MIT)

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.