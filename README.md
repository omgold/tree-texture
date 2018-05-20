# Tree procedural texture

Proof of principle implementation in Go for tree-like/crack-like procedural texture algorithm

## Getting Started

### Downloading pre-built Linux x86_64 static executable

https://github.com/omgold/tree-texture/raw/pre-built/tree-texture

### Building

### Prerequisites

- Go compiler
- GNU make

### Build steps

- checkout master branch
- type `make` in top-level directory to build just the executable
- type `make all` to also build the sample images

### Usage

Running the executable without arguments produces a 512x512 PNG image of a simple example.

Check `Makefile` for parameters of the sample images.

Command line arguments:

```
Usage of ./tree-texture:
  -bd int
        recursion depth (number of nested branches). (default 3)
  -bj float
        Branch jitter (random displacement of branches along parent) (default 0.1)
  -cx float
        iteration cutoff distance in direction perpendicular to branch (this is a performance optimization parameter, higher values make computation take longer, but if too low, branches will be truncated, leading to ugly artifacts). (default 2)
  -cy float
        iteration cutoff distance in direction along branch (this is a performance optimization parameter, higher values make computation take longer, but if too low, branches will be truncated, leading to ugly artifacts). (default 0.5)
  -di float
        initial relative distance of branches on same parent (near root of parent). Use -sl to make branches denser near tip. (default 0.55)
  -dm float
        maximum distance of branches from root of parent (in multiples of parent length) (default 0.9)
  -dol float
        distance of first left branch from root of parent (in multiples of parent length) (default 0.4)
  -dor float
        distance of first right branch from root of parent (in multiples of parent length) (default 0.6)
  -f string
        output filename (use empty string or '-' for stdout) (default "tree.png")
  -height int
        height of output image in pixels (default 512)
  -is float
        intensity scale factor of branch (higher values make more deeply nested branches finer) (default 1.1)
  -na float
        amount of noise (higher values will result make more noisy, zero will produce completely straight branches) (default 0.1)
  -nd int
        dimension of fractal noise (higher values will produce larger spectrum) (default 3)
  -ng float
        gain of noise (higher values will make high-frequency components more prominent) (default 0.5)
  -nl float
        lacunarity of fractal noise (higher values will stretch spectrum to higher frequencies) (default 1.7)
  -re float
        rotation of branch (near tip of parent) in multiples of 180 degrees (default 0.2)
  -rmb float
        relative scale of branches starting near tip of parent. Note: if significantly different from -rml, result may show visible discontinuities. (default 0.1)
  -rml float
        relative width of branch at tip (default 0.1)
  -rs float
        rotation of branch (near root of parent) in multiples of 180 degrees (default 0.3)
  -sb float
        scale (relative size) of branch (default 0.6)
  -seed float
        random seed (changing this will produce different random values)
  -shapeb float
        same as -shapel but for -rmb. Note: if significantly different from -rml, result may show visible discontinuities.
  -shapel float
        shape of branch (positive values make branches more pointy, negative ones more blunted)
  -sl float
        distance scale of branches at same parent (lower values make branches denser near tip of parent). Note: if too low, number of branches goes to infinity, resuning in an infinite loop (sl must be > dm*/(1-dl)) (default 0.7)
  -width int
        width of output image in pixels (default 512)
```

## Samples

![Sample 1](https://github.com/omgold/tree-texture/blob/pre-built/e1.png)
![Sample 2](https://github.com/omgold/tree-texture/blob/pre-built/e2.png)
![Sample 3](https://github.com/omgold/tree-texture/blob/pre-built/e3.png)
![Sample 4](https://github.com/omgold/tree-texture/blob/pre-built/e4.png)
![Sample 5](https://github.com/omgold/tree-texture/blob/pre-built/e5.png)
![Sample 6](https://github.com/omgold/tree-texture/blob/pre-built/e6.png)
![Sample 7](https://github.com/omgold/tree-texture/blob/pre-built/e7.png)
![Sample 8](https://github.com/omgold/tree-texture/blob/pre-built/e8.png)
