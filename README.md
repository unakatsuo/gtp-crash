### Linux gtp kernel module crash case

This is the test case to reproduce kernel panic at creating a GTP device while flooding GTP packets.

crash.sh causes kernel panic around 80% probability.

```
./build.sh
sudo ./crash.sh
```

