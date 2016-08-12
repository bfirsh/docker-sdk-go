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

container, err := client.Containers.Run(&docker.RunOptions{
    Image: "alpine",
    Cmd: []string{"echo", "hello world"},
})

if err != nil {
    return err
}

out, err := container.Logs(nil)
if err != nil {
    return err
}

fmt.Printf("%s", out)
```

Running a detached container:

```go
client, err := docker.FromEnv()
if err != nil {
    return err
}

container, err := client.Containers.Run(&docker.RunOptions{
    Image:  "bfirsh/reticulate-splines",
    Detach: true,
})
if err != nil {
    return err
}

out, err := container.Logs(nil)
if err != nil {
    return err
}

fmt.Printf("%s", out)

if err := container.Stop(nil); err != nil {
    return err
}
```
