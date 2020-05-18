---
title: "Change Orientation of an Image"
date: 2020-05-17T17:30:23+05:30
tags:
- css
draft: true
---


A single-line CSS transform is all it takes to change the orientation of an
image (or any DOM element, really).

For instance, if I have an image that is _on its side_, I can use the following
[`rotate`
transform](https://developer.mozilla.org/en-US/docs/Web/CSS/transform-function/rotate)
to orient it correctly.

```css
img {
  transform: rotate(90deg);
}
```

Testing you babay

```python
import os
os.path.join()
```
---

Go hello world
```golang
import (
"fmt"
)

func main() {
    fmt.Println("hello world")
}
```



It takes an angle which can be specified in degrees. Here I use `90deg`. If I
was going for a different effect, I could do something like `45deg`.
