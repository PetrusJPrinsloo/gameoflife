# Game of life

This is a project I did from following a tutorial. The link to that repo can me found here: https://github.com/KyleBanks/conways-gol

I have made some changes mainly to the structure because I want to expand on the OpenGL parts a lot in another project.

## Install

You can download and build directly from the source: 

```sh
$ go get github.com/PetrusJPrinsloo/gameoflife
```

## Configure

To change how the application runs just edit the default.json file in the root of the project.

```json
{
  "width": 1000,      // width of the window
  "height": 1000,     // height of the window
  "rows": 80,         // rows of cells
  "columns": 80,      // columns of cells
  "threshold": 0.15,  // chance of cell being alive on start, so default is 15%
  "fps": 10           // Frames Per Second, each frame is a new generation
}
``` 