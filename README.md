# sandpiles

Inspired by
[Fractal Zero](https://www.youtube.com/watch?v=1MtEUErz7Gg)


Now using a bounding box to speed up the topples. Looping through the entire image space is very slow.

Still using a double loop. First loop copies values from one grid to the temp grid. Second loop does the actual topple.

Starting the run with the number of grains in the starting spot is faster than adding one grain at a time, then running topple. 

## Timing and tuning tests. No optimazations.

- Triple loop
  - 6000 grains
    - Time 1m22.959370164s
- Double loop 
  - 6000 grains 
    - Time 56.509960683s
  - 600000 
    -Time 2h21m6.07087493s
- Single loop 
  - failed 
  - took 15 seconds but it had the wrong results

## Rectangle bounding

- Double loop
  - 6000 grains
    - Time 16.838954968s **WOW**
  - 60000 grains
    - Time 3m36.471748631s
  - 600000 grains
    - Time 52m54.697291662s
  - 1000000 grains
    - Time 1h40m8.293423533s

## Changing ints to fixed size.

## Base line before changes. Loading cell with total grains.
- 4095 1<<12 - 1
  - Time 10.626969828s
  - Frames: 2926
  - w:68 h:68
- 8191 1<<13 - 1
  - Time Time 25.623525736s
  - Frames: 6384
  - w:86 h:86
- 16383 1<<14 - 1
  - Time 52.905092416s
  - Frames: 14428
  - w:114 h:114
- 32767 1<<15 - 1
  - Time 1m50.013343568s
  - Frames: 29026
  - w:154 h:154

## Putting 1 grain at a time in a loop

