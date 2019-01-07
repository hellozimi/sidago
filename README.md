## Sida

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

## Todo

- [ ] Support for menus in config.toml
- [ ] FuncMap to transform asset paths and links to `base_url`-links.
- [ ] RSS
- [ ] Sitemap.xml
- [ ] Documentation for template variables
- [ ] Define robots.txt
- [ ] Describe how to host on [now.sh](https://zeit.co/now)
- [ ] Nginx config to serve static html files
- [ ] 404.html
- [ ] Node.js with koa/express script to serve html