# RDX (README EXPLORER)

RDX is a tool for loading and previewing README files in the browser. Point rdx to a directory and it will walk through the directory and sub directories, finding and parsing README files into HTML contents that can be loaded in the browser.

## FEATURES:
- GO server
- Markdown parser

## USAGE

To load the README files in the current directory

```sh
   rdx
```
Load README files in a specify directory and sub directories

```sh
  rdx -path "./my-directory"
```
Load README with max recursive depth

```sh
  rdx -path "./my-directory" -d 3
```
To specify a different port for previewing the README files
```sh
   rdx -path "./my-directory" -p 8000
```