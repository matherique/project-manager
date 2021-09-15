# Project manager

A project manager that setup all my environment

## Motivation

I like to use _vim_ as my editor and _tmux_ as my "window manager", and
every time I have to create manually all panes and windows, so to
automatize that I create this project. I could make in `bash`
because is a small project, but I decide to use `golang` to practice
creating small utilities that helps my workflow

## Usage

```bash
project create <name>
project edit <name>
project open <name>
project config
project config <key>
project config <key> <value>
project config edit
```

## `project create <name>`

Open the editor configured in `project config` to create the tmux
script to open setup the project.

* `<name>` is the name of the project

## `project edit <name>`

Open the bash script that contains all the tmux code to setup the
project.

* `<name>` is the name of the project

## `project open`

Execute the bash script to run all tmux code

## `project config`

List all the config setup

## `project config <key>`

List the value in the config

* `<key>` is the name of the config that you looking for

## `project config <key> <value>`

Set the value to the config key

* `<key>` is the name of the config that you looking for
* `<value>` is the value to set

## `project config edit`

Open the config file in to the editor to edit
