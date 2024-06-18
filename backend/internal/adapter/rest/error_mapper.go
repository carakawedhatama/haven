package rest

import "haven/pkg/formatter"

var CodeMap = map[error]formatter.Status{
	// domain.ErrDataNotFound: formatter.DataNotFound,
}

var StatusMap = map[error]int{
	// domain.ErrDataNotFound: fiber.StatusNotFound,
}
