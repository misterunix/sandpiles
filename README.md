# sandpiles

Abelian sandpile models

Inspired by
[Fractal Zero](https://www.youtube.com/watch?v=1MtEUErz7Gg)


New code based on my 'c' version. Way faster. 

The maximum grid size is 10000 x 10000. This is much too large, but I have yet to find a way to overestimate how large a sandpile is going to be. The code automatically scales down and up the number of cells it has to process each round so the maximum grid size isn't iterated over each time. The code is optimized to only work in the region of the active sandpile. Right now, it is just a waste of memory.

The center of the grid starts with the listed starting grains. The maximum number of grains that can be placed is 2147483648 or half of uint32.

Times are relative to the hardware they are running on. I have several i3,i5,i7, and AMD laptops that the code has run on at one time or another. Just for fun, I let it run for a month on a raspberry pi 2b and it was **SLOW**! But what do you expect from such a limited system?

If you have hardware/time to spare consider running one of the larger sandpiles that have not been completed yet.

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

25
grains 33554432
Time:  824001 (9.53 days)
4212 4212

26
grains 67108864
Time: 3319832 (38.42 days)
5944 5944

Interesting but useless

|shift|seconds|minutes|hours|days|
|-:|-:|-:|-:|-:|
|17|9|0.2|0.003|0.0001|
|18|39|0.7|0.011|0.0005|
|19|160|2.7|0.044|0.0019|
|20|636|10.6|0.177|0.0074|
|21|2790|46.5|0.775|0.0323|
|22|11797|196.6|3.277|0.1365|
|23|48145|802.4|13.374|0.5572|
|24|196337|3272.3|54.538|2.2724|

