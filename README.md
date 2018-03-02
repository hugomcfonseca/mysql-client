# mysql-tasker
Go(lang) client library for MySQL

`mysql-tasker` is an application, developed in Go, that contains a TCP client providing some helpful use cases when you want to execute some tasks in a remote MySQL database.

Currently, I am working to provide this use cases:

* Users creation, modification and deletion
* Database creation and deletion
* Import database from an existing one

## Getting Started

Provided two methods to help you run `mysql-tasker`: **Standalone** is the usual way to run any Go application, you just need to setup your Go environment, get dependencies, build it and run as normal binary file; and **Docker** is the containerized version of this tasker, that will run in the top of a Docker container and you can use it as an ephemeral container, that execute tasks and remove container after exit, or run the container permanently and execute all your tasks into the container.

### Standalone

Here I will give the instructions about `git clone`, `go get` and `go build` this tasker and start using it.

### Docker

Here I will give the instructions about `docker pull`, `docker build` and `docker run` the container.

### Use cases

...

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details

## Supported versions

## Contributing

Contribution, in any kind of way, is highly welcome! 
It doesn't matter if you are not able to write code.
Creating issues or holding talks and help other people to use [mysql-tasker](https://github.com/hugomcfonseca/mysql-tasker) is contribution, too!
A few examples:

* Correct typos in the README / documentation
* Reporting bugs
* Implement a new feature or task

If you are new to pull requests, checkout [Collaborating on projects using issues and pull requests / Creating a pull request](https://help.github.com/articles/creating-a-pull-request/).
If you've found a bug, a typo, have a question or a want to request new feature, please [report it as a GitHub issue](https://github.com/hugomcfonseca/mysql-tasker/issues).
