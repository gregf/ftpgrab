package model

import "time"

// App holds application details
type App struct {
	ID      string
	Name    string
	Desc    string
	URL     string
	Author  string
	Version string
}

// Ftp holds data necessary for FTP configuration
type Ftp struct {
	Host               string   `yaml:"host,omitempty"`
	Port               int      `yaml:"port,omitempty"`
	Username           string   `yaml:"username,omitempty"`
	Password           string   `yaml:"password,omitempty"`
	ConnectionsPerHost int      `yaml:"connections_per_host,omitempty"`
	Timeout            int      `yaml:"timeout,omitempty"`
	DisableEPSV        bool     `yaml:"disable_epsv,omitempty"`
	TLS                TLS      `yaml:"tls,omitempty"`
	Sources            []string `yaml:"sources,omitempty"`
}

// TLS holds data necessary for TLS FTP configuration
type TLS struct {
	Enable             bool `yaml:"enable,omitempty"`
	Implicit           bool `yaml:"implicit,omitempty"`
	InsecureSkipVerify bool `yaml:"insecure_skip_verify,omitempty"`
}

// Db holds data necessary for Database configuration
type Db struct {
	Enable bool   `yaml:"enable,omitempty"`
	Path   string `yaml:"path,omitempty"`
}

// Download holds download configuration details
type Download struct {
	Output        string    `yaml:"output,omitempty"`
	UID           int       `yaml:"uid,omitempty"`
	GID           int       `yaml:"gid,omitempty"`
	ChmodFile     int       `yaml:"chmod_file,omitempty"`
	ChmodDir      int       `yaml:"chmod_dir,omitempty"`
	Include       []string  `yaml:"include,omitempty"`
	Exclude       []string  `yaml:"exclude,omitempty"`
	Since         time.Time `yaml:"since,omitempty"`
	Retry         int       `yaml:"retry,omitempty"`
	HideSkipped   bool      `yaml:"hide_skipped,omitempty"`
	CreateBasedir bool      `yaml:"create_basedir,omitempty"`
}

// Mail holds mail notification configuration details
type Mail struct {
	Enable             bool   `yaml:"enable,omitempty"`
	Host               string `yaml:"host,omitempty"`
	Port               int    `yaml:"port,omitempty"`
	SSL                bool   `yaml:"ssl,omitempty"`
	InsecureSkipVerify bool   `yaml:"insecure_skip_verify,omitempty"`
	Username           string `yaml:"username,omitempty"`
	Password           string `yaml:"password,omitempty"`
	From               string `yaml:"from,omitempty"`
	To                 string `yaml:"to,omitempty"`
}
