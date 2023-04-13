# Itemis Programming Challenge

I chose problem 3. 
You can find the full description in [Problem.md](PROBLEM.md).

Usually, I squash commits and delete the branch after merging from it.
However, I made an exception this time so you can follow along on the respective feature branches.


# Considerations

A parser is suitable for this kind of problem.
We can wrap the parser in a Command Line Interface (CLI) application to provide some quality of life (QoL) features.
These features may include hints, tab completion, and more.

For this project, I will use golang because I am keen to try out the language, and the scope of this project seems like a good opportunity to do so.
Additionally, golang has a nice and readable syntax so that should help reviewing purposes as well.

| Topic | Choice | Notes
| --- | --- | --- |
| Programming Language | [go](https://go.dev/) | `go version go1.20.3`
| CLI Framework |[urfave/cli](https://github.com/urfave/cli) | [Reference](https://github.com/shadawck/awesome-cli-frameworks)
| Testing Framework | `testing` | golang builtin testing framework
