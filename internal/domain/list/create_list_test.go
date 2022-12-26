package list_test

import (
	"testing"

	"github.com/narcisobenigno/none-list/internal/domain/list"
	"github.com/narcisobenigno/none-list/pkg/es"
	"github.com/narcisobenigno/none-list/pkg/results"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ProductSuite struct {
	*require.Assertions
	suite.Suite
}

func TestProductSuite(t *testing.T) {
	s := &ProductSuite{}
	s.Assertions = require.New(t)
	suite.Run(t, s)
}

func (s *ProductSuite) Shows_interest_in_a_product() {
	store := es.NewInMemoryEventStore()
	subject := list.NewBus(store)

	result := subject.Execute(&list.AddItem{})

	s.Equal(results.Success(), result)
}
