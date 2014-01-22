package go_pack

import (
  "io/ioutil"
)

/**
 *
 */
type PackFile struct {
  Path string
  Packages []Package `json:"packages"`
}

type Package struct {
  Source string `json:"source"`
  Target string `json:"target"`
  BasePath string `json:"base_path"`
}

/**
 * Creates a new PackFile instance
 */
func NewPackFile(path string) *PackFile {
  packFile := new(PackFile)
  packFile.Path = path

  return packFile
}

/**
 * Reads a packfile
 */
func (self *PackFile) Read() error {
  contents, error := ioutil.ReadFile(self.Path)

  if error != nil {
    return error
  }

  json.Unmarshal(rawConfig, &self)

  return nil
}

/**
 * Writes a packfile
 *
 * @TODO implement me
 */
func (self *PackFile) Write() error {
  return nil
}
