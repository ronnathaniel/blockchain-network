
package net

import (
    "fmt"
    "strings"
    "github.com/satori/go.uuid"
)


func genKey() string {

    uuid_v4 := uuid.Must(uuid.NewV4(), nil)
    id := fmt.Sprintf("%s", uuid_v4)
    key := strings.ReplaceAll(id, "-", "")

    // fmt.Printf("private key -- 0x%x", key)
    key = fmt.Sprintf("0x%s", key)
    return key
}
