package testdata

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"

	"fmt"

	"github.com/buger/jsonparser"
	tokens2 "github.com/rackspace/gophercloud/openstack/identity/v2/tokens"
	tokens3 "github.com/rackspace/gophercloud/openstack/identity/v3/tokens"
)

type (
	TestData struct {
		version  string
		data     []byte
		filename string
	}
)

func New(version string) *TestData {
	root, err := os.Getwd()
	if err != nil {
		panic("os.Getwd(): " + err.Error())
	}

	return NewWithFilename(root, version)
}

func NewWithFilename(root, version string) *TestData {
	// try testdata.development.json
	filename := path.Join(root, "testdata.development.json")
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		filename = path.Join(root, "testdata.json")
	}

	filename = path.Clean(filename)

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("ioutil.ReadFile(" + filename + "): " + err.Error())
	}

	return &TestData{
		version:  version,
		data:     b,
		filename: filename,
	}
}

func (td *TestData) Version() string {
	return td.version
}

func (td *TestData) Data() []byte {
	return td.data
}

func (td *TestData) API(key string) ([]byte, error) {
	key = td.version + "." + key

	return td.Get(key)
}

func (td *TestData) APIString(key string) string {
	data, err := td.API(key)
	if err != nil {
		panic(err.Error())
	}

	return string(data)
}

func (td *TestData) Get(key string) ([]byte, error) {
	var (
		data = td.data

		err error
	)

	keys := strings.Split(key, ".")
	for _, yek := range keys {
		// is the yek a array subscript?
		n, e := strconv.ParseInt(yek, 10, 32)
		if e != nil {
			data, _, _, err = jsonparser.Get(data, yek)
		} else {
			var i int64 = 0
			_, err = jsonparser.ArrayEach(data, func(arrBuf []byte, arrType jsonparser.ValueType, arrOffset int, arrErr error) {
				if i == n {
					data = arrBuf
					err = arrErr
				}

				i += 1
			})
		}

		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

func (td *TestData) ApiSet(key, content string, force bool) (err error) {

	// not force to rewrite key result when key exist
	if !force {
		res, err := td.API(key)
		if err != nil {
			return err
		}

		if len(res) > 0 {
			return nil
		}
	}

	var data MapData
	err = json.Unmarshal(td.Data(), &data)
	if err != nil {
		return
	}

	versionData := data[td.Version()]
	mapdata, ok := versionData.(map[string]interface{})
	if !ok {
		return fmt.Errorf("cant assertion version data %s to map data", td.Version())

	}

	mapdata[key] = content

	// pretty json output
	td.data, err = json.MarshalIndent(data, "", "    ")
	if err != nil {
		return
	}

	return ioutil.WriteFile(td.filename, td.data, 0666)
}

func (td *TestData) Set(key, content string, force bool) (err error) {

	// not force to rewrite key result when key exist
	if !force {
		res, err := td.Get(key)
		if err != nil {
			return err
		}

		if len(res) > 0 {
			return nil
		}
	}

	var data MapData
	err = json.Unmarshal(td.Data(), &data)
	if err != nil {
		return
	}

	data[key] = content

	// pretty json output
	td.data, err = json.MarshalIndent(data, "", "    ")
	if err != nil {
		return
	}

	return ioutil.WriteFile(td.filename, td.data, 0666)
}

func (td *TestData) GetString(key string) string {
	data, err := td.Get(key)
	if err != nil {
		panic(err.Error())
	}

	return string(data)
}

func (td *TestData) MockURL(subpath string) string {
	endpoint := td.GetString("user.endpoint")
	endpoint = strings.TrimPrefix(endpoint, "mitm://")
	endpoint = strings.TrimSuffix(endpoint, "/")

	if subpath[0] != '/' {
		subpath = "/" + subpath
	}

	return "http://" + endpoint + subpath
}

func (td *TestData) MockURLWithSSL(subpath string) string {
	absurl := td.MockURL(subpath)

	return "https" + absurl[4:]
}

func (td *TestData) MockAdminURL(subpath string) string {
	endpoint := td.GetString("admin.endpoint")
	endpoint = strings.TrimPrefix(endpoint, "mitm://")
	endpoint = strings.TrimSuffix(endpoint, "/")

	if subpath[0] != '/' {
		subpath = "/" + subpath
	}

	return "http://" + endpoint + subpath
}

func (td *TestData) MockAdminURLWithSSL(subpath string) string {
	absurl := td.MockAdminURL(subpath)

	return "https" + absurl[4:]
}

func (td *TestData) MockResourceURL(subpath string) string {
	endpoint := td.GetString("endpoint")
	endpoint = strings.TrimPrefix(endpoint, "mitm://")
	endpoint = strings.TrimSuffix(endpoint, "/")

	if subpath[0] != '/' {
		subpath = "/" + subpath
	}

	return "http://" + endpoint + subpath
}

func (td *TestData) MockResourceURLWithPort(port, subpath string) string {
	endpoint := td.GetString("endpoint") + ":" + port
	endpoint = strings.TrimPrefix(endpoint, "mitm://")
	endpoint = strings.TrimSuffix(endpoint, "/")

	if subpath[0] != '/' {
		subpath = "/" + subpath
	}

	return "http://" + endpoint + subpath
}

func (td *TestData) MockResourceURLWithSSL(path string) string {
	absurl := td.MockResourceURL(path)

	return "https" + absurl[4:]
}

func (td *TestData) MockCatalog() interface{} {
	switch td.version {
	case "v2":
		data, err := td.API("scoped.access")
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(string(data))

		var result struct {
			Entries []tokens2.CatalogEntry `json:"serviceCatalog"`
		}

		err = json.Unmarshal(data, &result)
		if err != nil {
			panic(err.Error())
		}

		return &tokens2.ServiceCatalog{
			Entries: result.Entries,
		}

	case "v3":
		data, err := td.API("scoped.token")

		// TODO: convert v3 catalog entry to v2's
		var result struct {
			Entries []tokens3.CatalogEntry `json:"catalog"`
		}

		err = json.Unmarshal(data, &result)
		if err != nil {
			panic(err.Error())
		}

		return &tokens3.ServiceCatalog{
			Entries: result.Entries,
		}

	}

	panic("Unknown api version of " + td.version + "!")
}
