# CLI

## Dice Roller

Single roll:

```shell
$ hxgm30 roll d20
```
```
12
```

Multiple rolls of the same die:

```shell
$ hxgm30 roll d6 10
```
```
d6: [1 5 2 2 4 2 2 2 6 1]
```

Various rolls combined:

```shell
$ hxgm30 roll d6 10 d4 4 d20 2
```
```
d6:
        [5 4 1 3 1 2 1 1 3 1]
d4:
        [2 4 4 3]
d20:
        [12 12]
```

Roll with metadata:

```shell
hxgm30 roll meta d6 10
```
```
d6:
        [1 6 3 3 4 3 5 5 3 6]
        average:3.9 count:10 high:6 low:1 sum:39
```

Various rolls with metadata:

```shell
$ hxgm30 roll meta d6 10 d4 4 d20 2 d100 1
```
```
d6:
        [6 1 4 6 6 2 6 2 6 1]
        average:4 count:10 high:6 low:1 sum:40
d4:
        [2 2 2 2]
        average:2 count:4 high:2 low:2 sum:8
d20:
        [11 12]
        average:11.5 count:2 high:12 low:11 sum:23
d100:
        98
```
