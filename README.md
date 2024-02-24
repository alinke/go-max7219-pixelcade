# Go-Max7219

Library written in Go to allow controlling of the MAX7219 LED module.  
Source forked from https://github.com/adrianh-za/go-max7219-rpi which was forked from https://github.com/d2r2/go-max7219

## Enhancements ##

* Works with pre-assembled 4 and 8 module MAX7219 LED module matrices.
* Set the rotational direction of the MAX7219 LED modules.
    * Some of the pre-assembled 4 and 8 LED modules are connected "upsidedown" to the circuit board.  Using the <i>inverted</i> rotational directions allows for supporting of these MAX7219 4 and 8 module LED module matrices.
* Sliding of text with blank padding before and after text.

## Usage ##

1) go get [github.com/adrianh-za/go-max7219](https://github.com/alinke/go-max7219-pixelcade)
2) browse to $/go/src/github.com/alinke/go-max7219-pixelcade/examples
3) sudo -E go run [filename].go
4) run to end (or ctrl-c to quit)

Examples filenames
* chars-4in1.go
* chars.go
* font-4in1.go
* font.go
* slide.go
* slide-4in1.go
* slide-4in1-invert.go
* slide-8in1-and-still-text.go


## Compatibility ##

Tested on Raspberry PI 3 B+, Orange Pi Zero 2 (use SPI 1), and Orange Pi Zero 3 (use SPI 1)

## Acknowledgements ##

Thanks to <a href="https://github.com/d2r2" target="blank"><b>Denis Dyakov</b></a> for his excellent libraries

## Gits ##

https://github.com/d2r2/go-max7219
