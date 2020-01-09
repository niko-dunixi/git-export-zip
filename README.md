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