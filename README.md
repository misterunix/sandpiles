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
    
