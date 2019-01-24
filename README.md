# Sida

[![Build Status](https://travis-ci.org/hellozimi/sidago.svg?branch=master)](https://travis-ci.org/hellozimi/sidago)


Sida is a minimalistic static site generator written in Go. This is still a very early version, everything is pretty much subject to change. Sida let's you generate static html from markdown files which you later can host where ever you want. Perfect for simple blogs and small sites with a few pages.

```console
$ sida --help

A static site generator

Usage:
  sida [command]

Available Commands:
  generate    Generates html files from markdown
  help        Help about any command
  init        Creates a new sida
  new         Creates a new post or page
  version     Displays the current version flags

Flags:
  -h, --help          help for sida
  -p, --path string   base path for your sida location (default "./")

Use "sida [command] --help" for more information about a command.
```

## Usage

Initialize a new blog

```console
$ sida init ~/Blog
🚀 Your new sida is now created in /Users/whoami/Blog
```

Download and unzip the default theme

```console
$ curl -sSL https://github.com/hellozimi/sida-default-theme/archive/master.tar.gz > ~/Blog/layout.tar.gz
$ tar -xpvf ~/Blog/layout.tar.gz -C ~/Blog/layout --strip-components=1
$ rm ~/Blog/layout.tar.gz
```

Create a new post

```console
$ cd ~/Blog
$ sida new post "My first post"
```

Edit the post

```console
$ vim posts/2019-01-07_my-first-post.md
```

Generate html output

```console
$ sida generate
🔨 Generating page 2/2
📁 Copying static assets

🚀 Build completed...

```

Now's your page put into the `./build` directory, ready to be deployed.

## Variables

### Global Variables

|Variable|Description|
|---|---|
|`.Global.Title`|The page title provided in config.toml|
|`.Global.Description`|The page description provided in config.toml|
|`.Global.Copyright`|The copyright information provided in config.toml|
|`.Global.BaseURL`|The base url provided in config.toml|
|`.Global.Posts`|Returns a list with all non draft posts sorted by date|
|`.Global.Menu.Items`|Returns a list with all menu items defined in config.toml|

### Page Variables

|Variable|Description|
|---|---|
|`.Title`|The title of the page from either parsed front matter or from markdown file|
|`.Date`|The date from either parsed front matter or from markdown file (Only usable on posts)|
|`.Draft`|Boolean to see if the page is a draft or not (specified in front matter - defaults to false)|
|`.Content`|Returns the parsed html content|
|`.Summary`|Returns a summary from the content without any html tags|
|`.Slug`|Returns the page slug|
|`.PageMeta.URL`|The url to the current page|
|`.PageMeta.DateComponents.Year`|The year as string parsed from the posts date|
|`.PageMeta.DateComponents.Month`|The month as string parsed from the posts date|
|`.PageMeta.DateComponents.Day`|The day as string parsed from the posts date|


## config.toml

An example config

```toml
title = "Sida"
description = "A blog about sida"
copyright = "© Copyright 2019 Author"
base_url = "https://sidago.blog/"

[[menu]]
title = "About"
link = "/about.html"

[[menu]]
title = "Contact"
link = "/contact.html"
```

## Todo

- [x] Support for menus in config.toml
- [ ] FuncMap to transform asset paths and links to `base_url`-links.
- [ ] RSS
- [x] Sitemap.xml
- [x] Documentation for template variables
- [ ] Define robots.txt
- [ ] Describe how to host on [now.sh](https://zeit.co/now)
- [ ] Nginx config to serve static html files
- [ ] 404.html
- [ ] Node.js with koa/express script to serve html