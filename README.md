# Colored Log

[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/powered-by-black-magic.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/for-you.svg)](https://forthebadge.com)

Print message in terminal with colored text/message.

## How to install

```
go get github.com/mohammadanang/golog
```

## How to use

### without alias

```
...

import "github.com/mohammadanang/golog"

func main() {
  golog.Success("Get success log")
}

...
```

### using alias

```
...

import logger "github.com/mohammadanang/golog"

func main() {
  logger.Success("Get success log")
}

...
```