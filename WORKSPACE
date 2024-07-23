load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# ------ gazelle stuff
http_archive(
    name = "io_bazel_rules_go",
    integrity = "sha256-M6zErg9wUC20uJPJ/B3Xqb+ZjCPn/yxFF3QdQEmpdvg=",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.48.0/rules_go-v0.48.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.48.0/rules_go-v0.48.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    integrity = "sha256-12v3pg/YsFBEQJDfooN6Tq+YKeEWVhjuNdzspcvfWNU=",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.37.0/bazel-gazelle-v0.37.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.37.0/bazel-gazelle-v0.37.0.tar.gz",
    ],
)


load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")


go_rules_dependencies()

go_register_toolchains(
    version = "1.22.4",
)

gazelle_dependencies()


# ----- csharp rules stuff

http_archive(
    name = "rules_dotnet",
    integrity = "sha256-KUH7wrIVcXczJjB1Mq7vA2ZqhDKURxlT4O3MJVbGwUg=",
    urls = [
        "https://github.com/bazelbuild/rules_dotnet/releases/download/v0.15.1/rules_dotnet-v0.15.1.tar.gz",
    ],
)

load(
    "@rules_dotnet//dotnet:repositories.bzl",
    "dotnet_register_toolchains",
)

dotnet_register_toolchains("dotnet", "8.0.100")

local_repository(
    name = "sharpazelle",
    path = "./",
)


