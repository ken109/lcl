package util

type Config struct {
	Mysql struct {
		User     string
		Password string
	}
	Option struct {
		Start struct {
			EmptyOnly bool `mapstructure:"empty-only"`
			Share     bool
		}
		Stop struct {
			DropDb bool `mapstructure:"drop-db"`
		}
		Staging struct {
			Ssh    string
			Domain string
		}
	}
}
