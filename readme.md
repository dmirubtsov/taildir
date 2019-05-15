taildir
==============================

Tailing logs from directories recursively.

Usage
------------------------------

```
$ taildir dir1 [dir2 ...]
```

### Demo

[![asciicast](https://asciinema.org/a/FPqA0gaBhJvX2mdJQVAk9uBEw.svg)](https://asciinema.org/a/FPqA0gaBhJvX2mdJQVAk9uBEw)

1. Preparation for demo

    ```sh
    $ while :; do echo "example of new log entry in logdir1" >> logdir1/log.log; sleep 1; done
    $ while :; do echo "example of new log entry in logdir2" >> logdir2/log.log; sleep 1; done
    ```

2. Run

    ```
    $ taildir logdir1 logdir2
    ```

3. Installation

   ```
   go get github.com/dmirubtsov/taildir
   ```