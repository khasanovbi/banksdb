package gen

import (
	"encoding/json"
	"gopkg.in/src-d/go-billy.v4"
	"gopkg.in/src-d/go-billy.v4/memfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"strings"
)

const (
	repoUrl = "https://github.com/ramoona/banks-db"
	dbPath  = "banks"
)

type Bank struct {
	Name       string `json:"name"`
	Country    string `json:"country"`
	LocalTitle string `json:"localTitle"`
	EngTitle   string `json:"engTitle"`
	URL        string `json:"url"`
	Color      string `json:"color"`
	Prefixes   []int  `json:"prefixes"`
}

func unmarshalBankFromFile(fs billy.Filesystem, path string) (*Bank, error) {
	jsonFile, err := fs.Open(path)
	if err != nil {
		return nil, err
	}
	//noinspection GoUnhandledErrorResult
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	bank := Bank{}
	err = json.Unmarshal(byteValue, &bank)
	if err != nil {
		return nil, err
	}
	return &bank, nil

}

func readBanks(fs billy.Filesystem, path string, banks *[]Bank) {
	contents, err := fs.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, content := range contents {
		contentPath := filepath.Join(path, content.Name())
		if content.IsDir() {
			log.Printf("enter dir: path='%s'", contentPath)
			readBanks(fs, contentPath, banks)
			log.Printf("exit dir: path='%s'", contentPath)
		} else {
			if !strings.HasSuffix(content.Name(), ".json") {
				log.Printf("skip not json file: path='%s'", contentPath)
				continue
			}

			log.Printf("handle file: path='%s'", contentPath)
			bank, err := unmarshalBankFromFile(fs, contentPath)
			if err != nil {
				log.Fatal(err)
			}
			*banks = append(*banks, *bank)
		}
	}
}

func ParseBanks() []Bank {
	log.Printf("clone repo: repoUrl='%s'", repoUrl)
	fs := memfs.New()
	repository, err := git.Clone(memory.NewStorage(), fs, &git.CloneOptions{
		URL: repoUrl,
	})
	if err != nil {
		log.Fatal(err)
	}
	ref, err := repository.Head()
	if err != nil {
		log.Fatal(err)
	}
	commit, err := repository.CommitObject(ref.Hash())
	log.Printf("repo cloned: lastCommitTime='%s'", commit.Committer.When.Format("02 Jan 2006 15:04:05"))
	banks := make([]Bank, 0)
	readBanks(fs, dbPath, &banks)
	sort.Slice(banks, func(i, j int) bool {
		if banks[i].Country == banks[j].Country {
			return banks[i].Name < banks[j].Name
		}
		return banks[i].Country < banks[j].Country
	})
	return banks
}
