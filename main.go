package witch

import (
	"encoding/json"

	"github.com/shibukawa/configdir"
)

const filename string = "config.json"

func Read(o interface{}, vendor string, name string) {
	configDirs := configdir.New(vendor, name)
	folders := configDirs.QueryFolderContainsFile(filename)
	if folders != nil {
		data, _ := folders.ReadFile(filename)
		_ = json.Unmarshal(data, o)
	}
}

func Write(o interface{}, vendor string, name string) {
	bytes, _ := json.Marshal(o)
	configDirs := configdir.New(vendor, name)
	folders := configDirs.QueryFolderContainsFile(filename)
	if folders != nil {
		folders.WriteFile(filename, bytes)
	} else {
		folder := configDirs.QueryFolders(configdir.Global)[0]
		folder.WriteFile(filename, bytes)
	}
}
