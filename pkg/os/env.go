package os

import "os"

func GetEnv(env string) string {

    os_env := os.Getenv(env)

    if os_env == ""{
        return "localhost"
    }

    return os_env

}
