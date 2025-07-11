load("@gazelle//:def.bzl", "gazelle")
load("@pypi//:requirements.bzl", "all_whl_requirements")
load("@rules_python//python:defs.bzl", "py_binary", "py_library")
load("@rules_python//python:pip.bzl", "compile_pip_requirements")
load("@rules_python_gazelle_plugin//manifest:defs.bzl", "gazelle_python_manifest")
load("@rules_python_gazelle_plugin//modules_mapping:def.bzl", "modules_mapping")

gazelle(
    name = "gazelle",
)

# This rule fetches the metadata for python packages we depend on. That data is
# required for the gazelle_python_manifest rule to update our manifest file.
modules_mapping(
    name = "modules_map",
    wheels = all_whl_requirements,
)

# Gazelle python extension needs a manifest file mapping from
# an import to the installed package that provides it.
# This macro produces two targets:
# - //:gazelle_python_manifest.update can be used with `bazel run`
#   to recalculate the manifest
# - //:gazelle_python_manifest.test is a test target ensuring that
#   the manifest doesn't need to be updated
gazelle_python_manifest(
    name = "gazelle_python_manifest",
    modules_mapping = ":modules_map",
    # This is what we called our `pip_parse` rule, where third-party
    # python libraries are loaded in BUILD files.
    pip_repository_name = "pypi",
    # This should point to wherever we declare our python dependencies
    # (the same as what we passed to the modules_mapping rule in WORKSPACE)
    # This argument is optional. If provided, the `.test` target is very
    # fast because it just has to check an integrity field. If not provided,
    # the integrity field is not added to the manifest which can help avoid
    # merge conflicts in large repos.
    requirements = "//:requirements_lock.txt",
)

compile_pip_requirements(
    name = "requirements",
    src = "requirements.in",
    requirements_txt = "requirements_lock.txt",
)

# gazelle:resolve py frontmatter python-frontmatter

py_binary(
    name = "cat2tag",
    srcs = ["blog/cat2tag.py"],
    visibility = ["//:__subpackages__"],
    deps = ["//:python-frontmatter"],
)

py_library(
    name = "wordsonsand",
    srcs = ["blog/cat2tag.py"],
    visibility = ["//:__subpackages__"],
    deps = ["//:python-frontmatter"],
)
