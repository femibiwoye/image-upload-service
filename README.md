# Image Upload Service

This Repo consists of a GRPC service that upload images and a cli tool which is a client for the service.
The two binaries are located in the `/cmd` folder.

## Getting started

- Install Go
- Build the cli tool in the `/cmd/image-uploader-cli` directory
- Build the grpc server in `/cmd/image-uploader` directory
- Start the server
- Use the CLI to upload images.

The `run-server.sh` and `run-cli.sh` scripts contains commands to help run the server and upload the two images in the `/assets` directory. You open two terminals, make sure .config.yaml variables are set and run the `sh run-server.sh` command on one terminal and run `sh run-cli.sh` on the second terminal once the server is up and running.

### Server

To run the grpc server, create the config file `.config.yaml` check the `.config-sample.yaml` file.

### CLI Tool

The cli has only one command `upload` which can take an arbitary number of arguments and returns the location where the files are uploaded.

```bash
[USAGE]:
    $ image-uploader-cli upload assets/image-1.jpg assets/image-2.png
```

### Running Tests

running automated tests

```bash
make test
```