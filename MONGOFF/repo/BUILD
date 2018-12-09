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
    ],
)

# All the deps ..
# go_get(name = "mgo", get = "github.com/globalsign/mgo", visibility = ['PUBLIC'],)
