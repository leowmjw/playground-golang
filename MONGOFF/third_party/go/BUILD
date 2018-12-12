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
    deps = []
)

go_get(name = 'ff', get = 'github.com/peterbourgon/ff', deps = ['//third_party/go:pkgerr'])
go_get(name = 'pkgerr', get = 'github.com/pkg/errors')

go_get(name = 'ulid', get = 'github.com/oklog/ulid')

go_get(name = 'run', get = 'github.com/oklog/run')

go_get(name = 'mergo', get = 'github.com/imdario/mergo')

go_get(name = 'hclog', get = 'github.com/hashicorp/go-hclog')

go_get(name = 'ctxlog', get = 'github.com/peterbourgon/ctxlog', deps = ['//third_party/go:gokitlog'])
go_get(name = 'gokitlog', get = "github.com/go-kit/kit/log", deps = ['//third_party/go:logfmt'])
go_get(name = 'logfmt', get = "github.com/go-logfmt/logfmt")
