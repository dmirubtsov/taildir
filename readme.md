taildir
==============================

[![Go Report Card](https://goreportcard.com/badge/github.com/dmirubtsov/taildir)](https://goreportcard.com/report/github.com/dmirubtsov/taildir)
[![Docker Build Status](https://img.shields.io/docker/cloud/build/mazy/taildir.svg)](https://cloud.docker.com/repository/docker/mazy/taildir)

Tailing logs from directories recursively.

Installation

```
go get github.com/dmirubtsov/taildir
```

Usage
------------------------------

Useful as sidekiq kubernetes container that watch emptyDir with logs and write it to container stdout.

```
$ taildir dir1 [dir2 ...]
```

Kubernetes example:

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp-deployment
  labels:
    app: myapp
spec:
  replicas: 3
  selector:
    matchLabels:
      app: myapp
  template:
    metadata:
      labels:
        app: myapp
    spec:
      containers:
      - name: myapp
        image: myapp:1.3.37
        volumeMounts:
          - name: logs
            mountPath: /app/storage/logs
      - name: logger
        image: mazy/taildir:latest
        args: ["/app/storage/logs"]
        volumeMounts:
          - name: logs
            mountPath: /app/storage/logs
      volumes:
        - name: logs
          emptyDir: {}

```

