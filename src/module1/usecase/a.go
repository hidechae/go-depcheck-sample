package usecase

import (
	_ "github.com/hidechae/go-depcheck-sample/src/module1/domain"
	_ "github.com/hidechae/go-depcheck-sample/src/module1/infrastructure" // invalid dependency
)
