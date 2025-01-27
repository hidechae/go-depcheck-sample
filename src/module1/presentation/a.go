package presentation

import (
	_ "github.com/hidechae/go-depcheck-sample/src/module1/domain"
	_ "github.com/hidechae/go-depcheck-sample/src/module1/infrastructure" // invalid dependency
	_ "github.com/hidechae/go-depcheck-sample/src/module1/usecase"
)
