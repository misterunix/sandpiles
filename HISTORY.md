# History

Code start, but not working right. Some kind of problem with my loop.

Removed all the graphics and just printed the raw numbers.

Found the issue. grid2 values were being overwritten and causing all sorts of wierd issues. 

Added graphics back in.

First set up tests. Slow as heck.

## Timing amd tuning tests

- Triple loop
  - 6000 seeds
    - Time 1m22.959370164s
- Double loop 
  - 6000 seeds 
    - Time 56.509960683s
  - 600000 
    -Time 2h21m6.07087493s
- Single loop 
  - failed 
  - took 15 seconds but it had the wrong results

## Rectangle bounding

- Double loop
  - 6000 seeds
    - Time 16.838954968s **WOW**
  - 60000 seeds
    - Time 3m36.471748631s
  - 600000 seeds

