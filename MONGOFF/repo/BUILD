go_library(
    name = "repo",
    srcs = glob(
        ["*.go"],
        exclude = [
            "*_test.go",
        ],
    ),
    visibility = ["PUBLIC"],
    deps = [
        '//third_party/go:mgo',
        '//third_party/go:ulid',
    ],
)
