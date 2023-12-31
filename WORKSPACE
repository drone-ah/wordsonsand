workspace(name = "wordsonsand")
############################################################
############################################################
## Go rules

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "278b7ff5a826f3dc10f04feaf0b70d48b68748ccd512d7f98bf442077f043fe3",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.41.0/rules_go-v0.41.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.41.0/rules_go-v0.41.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "d3fa66a39028e97d76f9e2db8f1b0c11c099e8e01bf363a923074784e451f809",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.33.0/bazel-gazelle-v0.33.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.33.0/bazel-gazelle-v0.33.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

############################################################
# Define your own dependencies here using go_repository.
# Else, dependencies declared by rules_go/gazelle will be used.
# The first declaration of an external repository "wins".
############################################################

load("//:deps.bzl", "go_dependencies")

# gazelle:repository_macro deps.bzl%go_dependencies
go_dependencies()

go_rules_dependencies()

go_register_toolchains(version = "1.21.0")

gazelle_dependencies()

############################################################
############################################################
## Go protobuf def

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

git_repository(
    name = "com_google_protobuf",
    remote = "https://github.com/protocolbuffers/protobuf",
    tag = "v24.4",
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

############################################################
############################################################
## Hugo

# Update these to latest
RULES_HUGO_COMMIT = "294a8ec626a394011d35397108c930be631ab9fa"
RULES_HUGO_SHA256 = "8df370f374dc72701b65b7c8a8add8ccb8423a845e973993fa9c68f8b516c9be"

http_archive(
    name = "build_stack_rules_hugo",
    url = "https://github.com/stackb/rules_hugo/archive/%s.zip" % RULES_HUGO_COMMIT,
    sha256 = RULES_HUGO_SHA256,
    strip_prefix = "rules_hugo-%s" % RULES_HUGO_COMMIT
)

load("@build_stack_rules_hugo//hugo:rules.bzl", "hugo_repository", "github_hugo_theme")

#
# Load hugo binary itself
#
# Optionally, load a specific version of Hugo, with the 'version' argument
hugo_repository(
    name = "hugo",
)

#
# This makes a filegroup target "@com_github_yihui_hugo_xmin//:files"
# available to your build files
#
github_hugo_theme(
    name = "com_github_yihui_hugo_xmin",
    owner = "yihui",
    repo = "hugo-xmin",
    commit = "c14ca049d0dd60386264ea68c91d8495809cc4c6",
)

#
# This creates a filegroup target from a released archive from GitHub
# this is useful when a theme uses compiled / aggregated sources NOT found
# in a source root.
#
http_archive(
    name = "com_github_thegeeklab_hugo_geekdoc",
    url = "https://github.com/thegeeklab/hugo-geekdoc/releases/download/v0.34.2/hugo-geekdoc.tar.gz",
    sha256 = "7fdd57f7d4450325a778629021c0fff5531dc8475de6c4ec70ab07e9484d400e",
    build_file_content="""
filegroup(
    name = "files",
    srcs = glob(["**"]),
    visibility = ["//visibility:public"]
)
    """
)

############################################################
############################################################
## Python

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "rules_python",
    sha256 = "8c8fe44ef0a9afc256d1e75ad5f448bb59b81aba149b8958f02f7b3a98f5d9b4",
    strip_prefix = "rules_python-0.13.0",
    url = "https://github.com/bazelbuild/rules_python/archive/refs/tags/0.13.0.tar.gz",
)

load("@rules_python//python:pip.bzl", "pip_parse")

pip_parse(
    name = "python_deps",
    requirements_lock = "//third_party/python:requirements.txt",
)

# Load the starlark macro, which will define your dependencies.
load("@python_deps//:requirements.bzl", "install_deps")

# Call it to define repos for your requirements.
install_deps()
