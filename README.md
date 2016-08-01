# Docker SDK for Go

A Go library for controlling Docker and building apps on top of Docker.

## Usage

Import this library, and it will be available as `docker`:

```go
import "github.com/docker/docker-sdk-go"
```

Running a container:

```go
client, err := docker.FromEnv()
if err != nil {
    return err
}

opts := &docker.RunOptions{
    Image: "alpine",
    Cmd: []string{"echo", "hello world"},
}

container, err := client.Containers.Run(nil, opts)
if err != nil {
    return err
}

out, err := container.Logs()
if err != nil {
    return err
}

fmt.Println(out)
}
```

Running a container in the background:

```go
client, err := docker.FromEnv()
if err != nil {
    return err
}

opts := &docker.RunOptions{
    Image: "alpine",
    Cmd: []string{"echo", "hello world"},
}

container, err := client.Containers.Run(nil, "tasks/reticulate-splines")
if err != nil {
    return err
}
fmt.PrintLn(container.ID)

out, err := container.Logs()
if err != nil {
    return err
}
fmt.Println(out)

err := container.Stop()
}
```

