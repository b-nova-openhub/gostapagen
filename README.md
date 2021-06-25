![gostapagen-header](https://github.com/b-nova-openhub/jams-vanilla-content/raw/main/gostapagen-header.png)

# gostapagen (go-static-page-generator)

<!--- These are examples. See https://shields.io for others or to customize this set of shields. You might want to include dependencies, project status and licence info here --->
![GoDoc](https://godoc.org/github.com/go-git/go-git/v5?status.svg)
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
$ ❯  go install -v ./...
```

## Running gostapagen

To run gostapagen, follow these steps:

```
$ ❯  stapagen --repo=<content-git-repository-url> --branch=<branch>
```

Flags that you can use:

* `--repo`: The url to the git repository containing the markdown files. This repository is going to be cloned onto the
  filesystem by gostapagen. Make sure to use fully qualified without the .git extension at the end.

* `--branch`: The git branch that you want to use for the aforementioned repository.

* `--clonePath`: The absolute path to which the aforementioned git repository is going to be cloned to. Default path
  is `/tmp`.

* `--contentDir`: The relative path to the markdown files from within the git repository. Default value for this
  directory is `/content`.

* `--port`: Desired port to expose the RESTful api. Default port is `/8080`.

## Reading gostapagen's API

Once gostapagen is running, it does clone the git repository, converts its markdown files to html pages and forwards
them over a RESTful api. That api can be accessed over locally if the executable was run on your local
machine (`http://localhost:8080`) or over a K8s service definition.

* `/generate`: The generate endpoint does make an attempt at cloning the given git repository and directly converts the
  markdown files to html pages.

**GET-Response:**

```json
{
  "success": true,
  "errors": []
}
```

* `/pages`: The pages endpoint exposes all available, read successfully converted pages in a json format.

**GET-Response:**

```json
[
  {
    "title": "<title>",
    "permalink": "<permalink>",
    "author": "<author>",
    "tags": "<tag1, tag2, tag3>",
    "categories": "<cat1, cat2, cat3>",
    "publishDate": "<yyyy-mm-dd>",
    "description": "<description>",
    "showComments": "<true|false>",
    "isPublished": "<true|false>",
    "body": "<markdown rendered to html>"
  },
  {
    <secondpage>
  },
  ...
]
```

* `/page?id=<permalink>`: The page endpoint takes in an id query parameter where the value has to be the fully qualified
  permalink of the desired static page. As the `/pages` endpoint, this endpoint renders the page in a json format
  containing front matter data.

**GET-Response:**

```json
 {
  "title": "<title>",
  "permalink": "<permalink>",
  "author": "<author>",
  "tags": "<tag1, tag2, tag3>",
  "categories": "<cat1, cat2, cat3>",
  "publishDate": "<yyyy-mm-dd>",
  "description": "<description>",
  "showComments": "<true|false>",
  "isPublished": "<true|false>",
  "body": "<markdown rendered to html>"
}
```

* `/status`: The status endpoint exposes a range of technical information about the generation feature.

**GET-Response:**

```json
{
  "totalPages": 42,
  "publishedPages": 38,
  "unpublishedPages": 4,
  "lastPublishedPage": "2021-06-20",
  "lastGeneratedAt": "2021-06-25 10:22:42.603814 +0200 CEST m=+178.298962869"
}

```

## Makefile

There is a Makefile which automates the building process of the stapagen application. The Makefile has 6 build targets:
test, vet, fmt, mod, build, run, install and all. It can simply be run as such:

```
$ ❯  make build
```

## Dockerfile

There is also a Dockerfile by which one can containerize the stapagen application. The port that is being exposed is the
default 8080.

## Deployment to K8s

{coming soon}

### Quick Deployment to DigitalOcean's Kubernetes

[![Deploy to DO](https://www.deploytodo.com/do-btn-blue.svg)](https://cloud.digitalocean.com/apps/new?repo=https://github.com/b-nova-openhub/static-page-generator/tree/main)

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
