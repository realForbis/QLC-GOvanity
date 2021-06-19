# QLC-GOvanity
Vanity address generator for QLC written in go

!notice: qlc address will always start with "qlc_1" or "qlc_3"

## usage
```shell
$ go run main.go --prefix=abc --n=2

"""
Output:

Estimated number of iterations needed: 16384

Tried 8256 (~50.39%)
Found matching address!
seed =  020d468bb0b75af7acf1ee5aad1fac95db92039ad5c337fc68eced266b036a3b
address =  qlc_1abcz1bk7wjstt361xrtw4xto8k8xhk444xn3otymu37betjhmp1kume9f5g

Tried 39732 (~242.50%)
Found matching address!
seed =  91bfb01a6382ede2074a6f26313c8d1e21d915876e558ba7a76b3ae961a3bbb1
address =  qlc_3abcmqx1ea3dauofymne75b4r4n79qoj9j5kefbys58epofs7phufh7qt5aj
"""
```

### on Windows
```shell
$ ./QLC-Govanity --prefix=xyz --n=2

"""
Output:

Estimated number of iterations needed: 16384

Tried 13932 (~85.03%)
Found matching address!
seed =  60399bd19e096897b97b73311ad840e931dc6aecdfe550ab05d68607fda2bf13
address =  qlc_3xyzopno78hfydh3t91dhcs46ffo7meeg7bct6szpuscz3arg3b4ny6dy3a5

Tried 7224 (~44.09%)
Found matching address!
seed =  6c4ef4c3c420803e8ea7bd6fe0b915e85595f360e4f86e6381e46689910c0ed1
address =  qlc_3xyzdrirdqxfr9ngtk7g8ddp6tnd8zzfbyxjzxbn4886f34bunbz99qqzhib
"""
```