# HACKERNEWS-GO-CLI

![Installation and Usage](https://github.com/ozbekburak/hackernews-go-cli/blob/master/assets/usage-gif.gif)

### Description

**hackernews-go-cli** is a CLI application that helps you to fetch
top stories from [Hacker News](https://news.ycombinator.com/) like

* top stories until 500 story

* show stories until 200 story

* ask stories until 200 story

* job stories until 200 story

developed using the [hacker news api](https://github.com/HackerNews/API) for those who don't like to leave the command-line screen.

### Installation and Usage

**Assume that you have go already installed on your system**

```
    $ git clone https://github.com/ozbekburak/hackernews-go-cli.git
    $ cd hackernews-go-cli
    $ go build -o hackernews-go-cli
    $ ./hackernews-go-cli -top 5
```

> or

```
    $ ./hackerknews-go-cli -show 92
```

```
    $ ./hackerknews-go-cli -ask 125
```

```
    $ ./hackerknews-go-cli -job 75
```
