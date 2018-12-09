go_binary(
  name = 'mongoff',
  srcs = ['main.go'], 
  out = 'mongoff',
  deps =  [
    '//third_party/go:mgo', 
    '//:ff', '//:ctxlog', '//:ulid',
    '//:run', '//:mergo', '//:hclog', '//repo'
  ],
)

# All the deps ..
# go_get(name = "mgo", get = "github.com/globalsign/mgo")

go_get(name = 'ff', get = 'github.com/peterbourgon/ff', deps = ['//:pkgerr'])
go_get(name = 'pkgerr', get = 'github.com/pkg/errors')

go_get(name = 'ctxlog', get = 'github.com/peterbourgon/ctxlog', deps = ['//:gokitlog'])
go_get(name = 'gokitlog', get = "github.com/go-kit/kit/log", deps = ['//:logfmt'])
go_get(name = 'logfmt', get = "github.com/go-logfmt/logfmt")

go_get(name = 'ulid', get = 'github.com/oklog/ulid')

go_get(name = 'run', get = 'github.com/oklog/run')

go_get(name = 'mergo', get = 'github.com/imdario/mergo')

go_get(name = 'hclog', get = 'github.com/hashicorp/go-hclog')