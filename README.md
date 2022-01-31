# Zet-go

This program creates a CLI for managing my zet repo. It implements the following subcommands:

* `search [string]`: Searches title, tags or entire files with flags
    * `--title`, `--tags` or `--all`
* `edit [string]`: Presents a numbered list of files containing the search string, and opens selected file in default editor
* `new`: Creates a folder in the zet repo with UTC time as name, creates new markdown file, and opens it in default editor
* `commit`: Commits changes using title of added file as commit message, pulls, and pushes to github
* `help`: Displays usage. This is the default if command is not recognized