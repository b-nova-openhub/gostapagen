# gostapagen (go-static-page-generator)

<!--- These are examples. See https://shields.io for others or to customize this set of shields. You might want to include dependencies, project status and licence info here --->
![GitHub repo size](https://img.shields.io/github/repo-size/b-nova-openhub/static-page-generator)
![GitHub contributors](https://img.shields.io/github/contributors/b-nova-openhub/static-page-generator)
![GitHub stars](https://img.shields.io/github/stars/b-nova-openhub/static-page-generator?style=social)
![GitHub forks](https://img.shields.io/github/forks/b-nova-openhub/static-page-generator?style=social)
![Twitter Follow](https://img.shields.io/twitter/follow/b__nova?style=social)

gostapagen is a `tool` that allows `markdown content creators` to do `generate and expose statically generated pages`.

More specifically, gostapagen is microservice written in Go that converts markdown files in a given git repository,
converts them to html and exposes those statically rendered pages over a RESTful interface. It is operates on the
principles of KISS and is best integrated in a JAMstack.

## Prerequisites

Before you begin, ensure you have met the following requirements:
<!--- These are just example requirements. Add, duplicate or remove as required --->

* You have installed the latest version of `go1.16.5`
* You have a `Linux/Mac` machine with working knowledge of the underlying filesystem and Go build process.

## Installing gostapagen

To install gostapagen, follow these steps:

Linux and macOS:

```
go build cmd/stapagen/main.go
```

## Using gostapagen

To use gostapagen, follow these steps:

```
./main --repo=<content-git-repository-url> --branch=<branch>
```

Flags that you can use:

* `--repo`: The url to the git repository containing the markdown files. This repository is going to be cloned onto the
  filesystem by gostapagen. Make sure to use fully qualified without the .git extension at the end.

* `--branch`: The git branch that you want to use for the aforementioned repository.

* `--clonePath`: The absolute path to which the aforementioned git repository is going to be cloned to. Default path
  is `/tmp`.

* `--contentDir`: The relative path to the markdown files from within the git repository. Default value for this
  directory is `/content`.

## Contributing to gostapagen

<!--- If your README is long or you have some specific process or steps you want contributors to follow, consider creating a separate CONTRIBUTING.md file--->
To contribute to gostapagen, follow these steps:

1. Fork this repository.
2. Create a branch: `git checkout -b <branch_name>`.
3. Make your changes and commit them: `git commit -m '<commit_message>'`
4. Push to the original branch: `git push origin <project_name>/<location>`
5. Create the pull request.

Alternatively see the GitHub documentation
on [creating a pull request](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request)
.

## Contributors

Thanks to the following people who have contributed to this project:

* [@raffaelschneider](https://github.com/raffaelschneider) 💻📖🐛
* [@stefanwelsch](https://github.com/bnova-stefan) 💻🧑‍🏫
* [@tomtrapp](https://github.com/tomtrapp) 🤔👀
* [@rickyelfner](https://github.com/ricky-bnova) 💬🐛

You might want to consider using something like
the [All Contributors](https://github.com/all-contributors/all-contributors) specification and
its [emoji key](https://allcontributors.org/docs/en/emoji-key).

## Contact

If you want to contact me you can reach me at [info@b-nova.com](info@b-nova.com).

## License

<!--- If you're not sure which open license to use see https://choosealicense.com/--->

This project uses the following license: [MIT License](https://opensource.org/licenses/MIT)
. https://opensource.org/licenses/MIT
