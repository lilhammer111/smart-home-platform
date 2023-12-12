# *** Project

## introduce

- Use the [Hertz](https://github.com/cloudwego/hertz/) framework
- Integration of pprof, cors, recovery, access_log, gzip and other extensions of Hertz.
- Generating the base code for unit tests.
- Provides basic profile functions.
- Provides the most basic MVC code hierarchy.

## Directory structure

|  catalog   | introduce  |
|  ----  | ----  |
| conf  | Configuration files |
| main.go  | Startup file |
| hertz_gen  | Hertz generated model |
| api  | Used for request processing, validation and return of response. |
| biz  | The actual business logic. |
| router  | Routing and middleware registration |
| utils  | Wrapped some common methods |

## How to run

```shell
sh build.sh
sh output/bootstrap.sh
```