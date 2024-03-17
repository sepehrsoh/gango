[![codebeat badge](https://codebeat.co/badges/da2c9829-f570-4555-ab2a-9bcb7bc9e1e5)](https://codebeat.co/projects/github-com-sepehrsoh-gango-main)

# Gango: A Go-Based Backend Service Framework

Gango is an open-source project written in Go that provides a robust and flexible foundation for developing backend
services. Whether you're a seasoned developer or just getting started, Gango simplifies the process of creating backend
systems by automating many of the essential tasks.

## Features

- **Service Generation**: With Gango, you can easily generate a new backend service project. This includes creating a
  new directory with all the basic components you need for your backend system.

- **Docker Integration**: Gango streamlines the process of containerizing your service. A Dockerfile is provided, and
  you can build a Docker image with a simple command.

## Getting Started

### Prerequisites

Before you begin, ensure you have the following prerequisites:

- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/)

### Installation

1. Clone the Gango repository to your local machine:

   ```bash
   git clone https://github.com/sepehrsoh/gango.git
   cd gango
   ```

2. Run the following command to set up your Gango environment:

   ```bash
   make all
   ```

3. Create a new project using Gango:

   ```bash
   ./gango generate -n project-name -p redis -p elastic -p postgres
   ```
    * **_Note: supported providers listed below:_**
   
        * `redis`

        * `elastic`

        * `postgres`

   This command will create a new directory with your chosen service name, preconfigured with the basics you need for a
   backend system.

4. If you want to dockerized your service, navigate to the new service directory:

   ```bash
   cd [service-name]
   ```

5. Build a Docker image for your service:

   ```bash
   make build-docker
   ```

## Usage

Now that you have your Gango-based backend service up and running, you can start building and customizing your project.
Here are some commands and conventions you should be aware of:

- Modify your service code, endpoints, and database configurations in the `[service-name]` directory.
- Create routes, add middleware, and define models and controllers.
- Leverage the power of Go to build your backend system according to your requirements.

For more detailed information on how to work with Gango, please check
the [documentation](https://github.com/sepehrsoh/gango/wiki).

## Contributing

We welcome contributions from the community to make Gango even better. If you'd like to contribute, please follow
our [Contribution Guidelines](CONTRIBUTING.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

We'd like to express our gratitude to the open-source community for their valuable contributions and support.

---

Thank you for choosing Gango for your backend development needs. We hope it simplifies the process of building robust
and scalable backend services. If you have any questions or encounter issues, please feel free to open an issue on the
GitHub repository.

[GitHub Repository](https://github.com/sepehrsoh/gango)

[Documentation](https://github.com/sepehrsoh/gango/wiki)

[Report Issues](https://github.com/sepehrsoh/gango/issues)

Happy coding with Gango! 🚀
