# *** Project

## introduce

- Use the [Kitex](https://github.com/cloudwego/kitex/) framework
- Generating the base code for unit tests.
- Provides basic config functions
- Provides the most basic MVC code hierarchy.

## Directory structure

|  catalog   | introduce  |
|  ----  | ----  |
| conf  | Configuration files |
| main.go  | Startup file |
| biz/handler.go  | Used for request processing return of response. |
| kitex_gen  | kitex generated code |
| biz/service  | The actual business logic. |
| biz/model  | the data access layer and relative database operations |

## How to run

```shell
sh deploy/build.sh
sh deploy/output/bootstrap.sh
```