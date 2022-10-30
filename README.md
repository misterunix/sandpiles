# sandpiles

Inspired by
[Fractal Zero](https://www.youtube.com/watch?v=1MtEUErz7Gg)


New code based on my 'c' version. Way faster. 

Maximum grid size is 12000 x 12000. This is much to large, but I have yet to find a way to overestimate out large a sandpile is going to be. The code is optimized to only work in the region of the active sandpile. Right now, its just a waste of memory.

The center of the grid starts with the listed starting grains. The maximum grains that can be placed is 2147483648 or half of uint32.


2^8 grains placed
grains 256
2^8 grains placed
Time:  0.000131559
32 32

2^9 grains placed
grains 512
Time:  0.000262601
38 38

2^10 grains placed
grains 1024
2^ 10 grains placed
Time:  0.000737656
44 44

2^11 grains placed
grains 2048
2^ 11 grains placed
Time:  0.002577955
54 54

2^12 grains placed
grains 4096
Time:  0.01035059
68 68

2^13 grains placed
grains 8192
Time:  0.037538124
86 86

2^14 grains placed
grains 16384
Time:  0.185819086
114 114

2^15 grains placed
grains 32768
Time:  0.620534637
154 154

2^16 grains placed
grains 65536
Time:  2.490630475
208 208

2^17 grains placed
grains 131072
Time:  9.817063459
286 286

2^18 grains placed
grains 262144
Time:  39.541612839
394 394

2^19 grains placed
grains 524288
Time:  160.914970642
550 550

2^20 grains placed
grains 1048576
Time:  636.310366064
766 766

2^21 grains placed
grains 2097152
Time:  2790.168505745
1072 1072

2^22 grains placed
grains 4194304
Time:  11797.657959528
1508 1508

2^23 grains placed
grains 8388608
Time:  48145.536773924
2120 2120

2^24 grains placed
grains 16777216
Time:  196337.852703953
2988 2988

