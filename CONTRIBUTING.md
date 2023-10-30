# Contributing to Gango

Thank you for your interest in contributing to Gango, the Go-Based Backend Service Framework. We welcome contributions from the community to help improve and enhance the project. By following the guidelines below, you can get started with contributing to Gango.

## Cloning the Repository

To contribute to Gango, you'll first need to clone the repository to your local machine. Please follow these steps:

### Prerequisites

1. **Git**: If you don't already have Git installed, you can download and install it from [here](https://git-scm.com/downloads).

### Cloning the Repository

1. Open your terminal or command prompt.

2. Navigate to the directory where you want to clone the Gango repository.

3. Run the following command to clone the Gango repository:

   ```bash
   git clone git@github.com:sepehrsoh/gango.git
   ```


### Setting Up Git Remotes

1. After cloning the repository, navigate into the Gango directory:

   ```bash
   cd gango
   ```

2. To keep your local repository in sync with the upstream (main Gango repository), add a remote named "upstream" using the following command:

   ```bash
   git remote add upstream git@github.com:sepehrsoh/gango.git
   ```

   This allows you to fetch the latest changes from the main repository.

## Making Contributions

Now that you have a local copy of the Gango repository, you can make your contributions. Here are some general steps to follow:

1. Create a new branch for your work to isolate it from the `main` branch. You can name your branch something meaningful to your changes:

   ```bash
   git checkout -b my-feature
   ```

2. Make your changes to the code, documentation, or any other contributions.

3. Commit your changes with a descriptive commit message:

   ```bash
   git commit -m "Add feature: description of changes"
   ```

4. Push your changes to your fork of the repository on GitHub:

   ```bash
   git push origin my-feature
   ```

5. Finally, create a pull request (PR) from your branch to the `main` branch of the main Gango repository.

## Keeping Your Fork Updated

To ensure your fork stays up-to-date with the main repository, periodically fetch changes from the upstream repository and rebase your local `main` branch:

```bash
git fetch upstream
git checkout main
git rebase upstream/main
```

This will keep your `main` branch current with the latest changes.

## Thank You

We appreciate your interest in contributing to Gango! Your contributions help make Gango better for everyone. If you have any questions or need assistance with your contributions, please don't hesitate to ask by opening an issue or reaching out to the maintainers.

Happy contributing!

