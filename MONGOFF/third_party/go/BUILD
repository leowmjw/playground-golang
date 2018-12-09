# This says that all targets in this package that don't specify otherwise
# will be visible to the whole repo.
# This is often useful for third-party packages like this where we expect
# anyone to use them.
package(default_visibility = ['PUBLIC'])

# go_get invokes, not unsurprisingly, 'go get' to fetch a third-party package.
# Note that we can specify Git revisions to pin them to a particular version.
go_get(
    name = 'mgo',
    get = 'github.com/globalsign/mgo',
    deps = [
    ]
)

# go_get(
#     name = 'bson',
#     get = 'github.com/globalsign/mgo/bson',
# )