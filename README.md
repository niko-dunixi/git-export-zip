# git-export-zip

## Abstract

`git-export-zip` is a simple command-line utility to
create a zip file from the current HEAD commit.

Essentially it wraps the most common commands I execute
when I want to create a zip file quickly.

For example, it is common for WGU students to need to
export their projects into a form that they can turn in
to their graders. This utility creates a zip that is
easy to identify and upload.

I found that I was navigating and iterating upon my
projects, I didn't have a reliable and quick way to
create these archives, except for "right-click > compress".

Programmers don't like clicking through menus, we
like code and unix-style commands.

## Installation

Ensure you have GoLang installed and execute
```bash
go get github.com/paul-nelson-baker/git-export-zip
```

## Usage

In your terminal
```bash
$ cd ~/your/project/directory
$ git-export-zip
> Exported to: <zip-filename-here>
```
## Why not a shell script?
I've been in what's commonly known as "shellscript hell".
Basically I have a lot of shellscripts already, which
depend on one another without any type-safety or simple
way to do what is trivial in any other language.

I actually love shellscripts, but I've found that
writing utilities in Go is trivial, powerful, and
safer.

Bonus points, I can cross compile and send the
binary to another machine. Even if that machine
doesn't have bash, it will still work because it's
a fully compiled executable!

## Prefer Bash?
If you prefer the raw command-line the equivalent
shellscript is this:

```
#!/usr/bin/env bash
set -ex
current_date="$(date +%m-%d-%Y)"
current_hash="$(git rev-parse HEAD | cut -c1-8)"
project_name="$(basename "$(git rev-parse --show-toplevel)" | tr '[:upper:]' '[:lower:]' | tr ' ' '_')"
git archive -o "../${project_name}-${current_date}-${current_hash}.zip" HEAD
```
