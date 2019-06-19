package gen

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/src-d/go-billy.v4"
	"gopkg.in/src-d/go-billy.v4/memfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

const (
	repoURL = "https://github.com/ramoona/banks-db"
	dbPath  = "banks"
)

// Bank represent bank info.
type Bank struct {
	Name       string `json:"name"`
	Country    string `json:"country"`
	LocalTitle string `json:"localTitle"`
	EngTitle   string `json:"engTitle"`
	URL        string `json:"url"`
	Color      string `json:"color"`
	Prefixes   []int  `json:"prefixes"`
}

// CountryBanks represent country with related banks.
type CountryBanks struct {
	Country string
	Banks   []Bank
}

func unmarshalBankFromFile(fs billy.Filesystem, path string) (*Bank, error) {
	jsonFile, err := fs.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close() // nolint: errcheck
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

func readCountryDir(fs billy.Filesystem, path string) []Bank {
	contents, err := fs.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	banks := make([]Bank, 0)
	for _, content := range contents {
		contentPath := filepath.Join(path, content.Name())
		if content.IsDir() {
			log.Fatalf("unexpected directory in country directory: path='%s'", contentPath)
		}

		if !strings.HasSuffix(content.Name(), ".json") {
			log.Printf("skip not json file: path='%s'", contentPath)
			continue
		}

		log.Printf("handle file: path='%s'", contentPath)
		bank, err := unmarshalBankFromFile(fs, contentPath)
		if err != nil {
			log.Fatal(err)
		}
		banks = append(banks, *bank)
	}

	sort.Slice(banks, func(i, j int) bool {
		return banks[i].Name < banks[j].Name
	})
	return banks
}

func readBanks(fs billy.Filesystem, path string) []CountryBanks {
	countryBanksSlice := make([]CountryBanks, 0)
	contents, err := fs.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, content := range contents {
		contentPath := filepath.Join(path, content.Name())
		if content.IsDir() {
			log.Printf("enter country dir: path='%s'", contentPath)
			countryBanksSlice = append(
				countryBanksSlice,
				CountryBanks{
					Country: content.Name(),
					Banks:   readCountryDir(fs, contentPath),
				},
			)
			log.Printf("exit country dir: path='%s'", contentPath)
		} else {
			log.Printf("skip file in root dir: path='%s'", contentPath)
		}
	}
	return countryBanksSlice
}

// ParseBanks fetch latest commit from repository and parse banks info for each country.
func ParseBanks() []CountryBanks {
	log.Printf("clone repo: repoURL='%s'", repoURL)
	fs := memfs.New()
	repository, err := git.Clone(memory.NewStorage(), fs, &git.CloneOptions{
		URL: repoURL,
	})
	if err != nil {
		log.Fatal(err)
	}
	ref, err := repository.Head()
	if err != nil {
		log.Fatal(err)
	}
	commit, err := repository.CommitObject(ref.Hash())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("repo cloned: lastCommitTime='%s'", commit.Committer.When.Format("02 Jan 2006 15:04:05"))

	countryBanksSlice := readBanks(fs, dbPath)
	sort.Slice(countryBanksSlice, func(i, j int) bool {
		return countryBanksSlice[i].Country < countryBanksSlice[j].Country
	})
	return countryBanksSlice
}
