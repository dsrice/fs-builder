package fsb_test

import (
	"fsb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"reflect"
	"testing"
	"unsafe"
)

type ExpressionSuite struct {
	suite.Suite
}

// Test_EqString is a unit test for the EqString method.
// It tests the functionality of the EqString method in the ExpressionSuite type.
// The EqString method creates an Expression object using the Eq function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_EqString() {
	ex := fsb.Eq("test", "user1")

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test = 'user1'", *ps)
}

// Test_EqInt is a unit test for the EqInt method.
// It tests the functionality of the EqInt method in the ExpressionSuite type.
// The EqInt method creates an Expression object using the Eq function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_EqInt() {
	ex := fsb.Eq("test", 1)

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test = 1", *ps)
}

// Test_NeqString is a unit test for the NeqString method.
// It tests the functionality of the NeqString method in the ExpressionSuite type.
// The NeqString method creates an Expression object using the Neq function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_NeqString() {
	ex := fsb.Neq("test", "user1")

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test != 'user1'", *ps)
}

// Test_NeqInt is a unit test for the NeqInt method.
// It tests the functionality of the NeqInt method in the ExpressionSuite type.
// The NeqInt method creates an Expression object using the Neq function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_NeqInt() {
	ex := fsb.Neq("test", 1)

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test != 1", *ps)
}

// Test_GtString is a unit test for the GtString method.
// It tests the functionality of the GtString method in the ExpressionSuite type.
// The GtString method creates an Expression object using the Gt function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_GtString() {
	ex := fsb.Gt("test", "user1")

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test > 'user1'", *ps)
}

// Test_GtInt is a unit test for the GtInt method.
// It tests the functionality of the GtInt method in the ExpressionSuite type.
// The GtInt method creates an Expression object using the Gt function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_GtInt() {
	ex := fsb.Gt("test", 1)

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test > 1", *ps)
}

// Test_GteString is a unit test for the GteString method.
// It tests the functionality of the GteString method in the ExpressionSuite type.
// The GteString method creates an Expression object using the Gte function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_GteString() {
	ex := fsb.Gte("test", "user1")

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test >= 'user1'", *ps)
}

// Test_GteInt is a unit test for the GteInt method.
// It tests the functionality of the GteInt method in the ExpressionSuite type.
// The GteInt method creates an Expression object using the Gte function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_GteInt() {
	ex := fsb.Gte("test", 1)

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test >= 1", *ps)
}

// Test_LtString is a unit test for the LtString method.
// It tests the functionality of the LtString method in the ExpressionSuite type.
// The LtString method creates an Expression object using the Lt function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_LtString() {
	ex := fsb.Lt("test", "user1")

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test < 'user1'", *ps)
}

// Test_LtInt is a unit test for the LtInt method.
// It tests the functionality of the LtInt method in the ExpressionSuite type.
// The LtInt method creates an Expression object using the Lt function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_LtInt() {
	ex := fsb.Lt("test", 1)

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test < 1", *ps)
}

// Test_LteString is a unit test for the LteString method.
// It tests the functionality of the LteString method in the ExpressionSuite type.
// The LteString method creates an Expression object using the Lte function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_LteString() {
	ex := fsb.Lte("test", "user1")

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test <= 'user1'", *ps)
}

// Test_LteInt is a unit test for the LteInt method.
// It tests the functionality of the LteInt method in the ExpressionSuite type.
// The LteInt method creates an Expression object using the Lte function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_LteInt() {
	ex := fsb.Lte("test", 1)

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test <= 1", *ps)
}

// Test_Like is a unit test for the Like method.
// It tests the functionality of the Like method in the ExpressionSuite type.
// The Like method creates an Expression object using the Like function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_Like() {
	ex := fsb.Like("test", "user1%")

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test LIKE 'user1%'", *ps)
}

// Test_Nlike is a unit test for the Nlike method.
// It tests the functionality of the Nlike method in the ExpressionSuite type.
// The Nlike method creates an Expression object using the Nlike function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_Nlike() {
	ex := fsb.Nlike("test", "user1%")

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test NOT LIKE 'user1%'", *ps)
}

// Test_PmString is a unit test for the PmString method.
// It tests the functionality of the PmString method in the ExpressionSuite type.
// The PmString method creates an Expression object using the Pm function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_PmString() {
	ex := fsb.Pm("test", "user1")

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test LIKE 'user1%'", *ps)
}

// Test_PmInt is a unit test for the PmInt method.
// It tests the functionality of the PmInt method in the ExpressionSuite type.
// The PmInt method creates an Expression object using the Pm function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
// Test_PmInt tests the logic that constructs a LIKE condition with an integer value in the ExpressionSuite.
// The method creates an Expression object using the Pm function from the fsb package,
// passing in the target attribute and the integer value.
// It retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Finally, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_PmInt() {
	ex := fsb.Pm("test", 1)

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test LIKE '1%'", *ps)
}

// Test_NpmString is a unit test for the NpmString method.
// It tests the functionality of the NpmString method in the ExpressionSuite type.
// The NpmString method creates an Expression object using the Nsm function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_NpmString() {
	ex := fsb.Npm("test", "user1")

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test NOT LIKE 'user1%'", *ps)
}

// Test_NpmInt is a unit test for the NpmInt method.
// It tests the functionality of the NpmInt method in the ExpressionSuite type.
// The NpmInt method creates an Expression object using the Nsm function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_NpmInt() {
	ex := fsb.Npm("test", 1)

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test NOT LIKE '1%'", *ps)
}

// Test_SmString is a unit test for the SmString method.
// It tests the functionality of the SmString method in the ExpressionSuite type.
// The SmString method creates an Expression object using the Sm function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_SmString() {
	ex := fsb.Sm("test", "user1")

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test LIKE '%user1'", *ps)
}

// Test_SmInt is a unit test for the SmInt method.
// It tests the functionality of the SmInt method in the ExpressionSuite type.
// The SmInt method creates an Expression object using the Sm function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_SmInt() {
	ex := fsb.Sm("test", 1)

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test LIKE '%1'", *ps)
}

// Test_NsmString is a unit test for the NsmString method.
// It tests the functionality of the NsmString method in the ExpressionSuite type.
// The NsmString method creates an Expression object using the Nsm function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_NsmString() {
	ex := fsb.Nsm("test", "user1")

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test NOT LIKE '%user1'", *ps)
}

// Test_NsmInt is a unit test for the NsmInt method.
// It tests the functionality of the NsmInt method in the ExpressionSuite type.
// The NsmInt method creates an Expression object using the Nsm function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
// Example: fsb.Nsm("test", 1) creates an Expression object with the condition "test NOT LIKE '%1'".
// This unit test checks that the condition field in the created Expression object matches the expected value.
func (s *ExpressionSuite) Test_NsmInt() {
	ex := fsb.Nsm("test", 1)

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test NOT LIKE '%1"+
		"'", *ps)
}

// Test_PsmString is a unit test for the PsmString method.
// It tests the functionality of the PsmString method in the ExpressionSuite type.
// The PsmString method creates an Expression object using the Psm function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_PsmString() {
	ex := fsb.Psm("test", "user1")

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test LIKE '%user1%'", *ps)
}

// Test_PsmInt is a unit test for the PsmInt method.
// It tests the functionality of the PsmInt method in the ExpressionSuite type.
// The PsmInt method creates an Expression object using the Psm function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
// Example: "test LIKE '%1%'"
func (s *ExpressionSuite) Test_PsmInt() {
	ex := fsb.Psm("test", 1)

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test LIKE '%1%'", *ps)
}

// Test_NpsmString is a unit test for the NpsmString method.
// It tests the functionality of the NpsmString method in the ExpressionSuite type.
// The NpsmString method creates an Expression object using the Npsm function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_NpsmString() {
	ex := fsb.Npsm("test", "user1")

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test NOT LIKE '%user1%'", *ps)
}

// Test_NpsmInt is a unit test for the NpsmInt method.
// It tests the functionality of the NpsmInt method in the ExpressionSuite type.
// The NpsmInt method creates an Expression object using the Npsm function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_NpsmInt() {
	ex := fsb.Npsm("test", 1)

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test NOT LIKE '%1%'", *ps)
}

// Test_BetweenString is a unit test for the BetweenString method.
// It tests the functionality of the BetweenString method in the ExpressionSuite type.
// The BetweenString method creates an Expression object using the Between function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_BetweenString() {
	ex := fsb.Between("test", "user1", "user2")

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test BETWEEN 'user1' TO 'user2'", *ps)
}

// Test_BetweenInt is a unit test for the BetweenInt method.
// It tests the functionality of the BetweenInt method in the ExpressionSuite type.
// The BetweenInt method creates an Expression object using the Between function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_BetweenInt() {
	ex := fsb.Between("test", 1, 5)

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test BETWEEN 1 TO 5", *ps)
}

// Test_NbetweenString is a unit test for the NbetweenString method.
// It tests the functionality of the NbetweenString method in the ExpressionSuite type.
// The NbetweenString method creates an Expression object using the Nbetween function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_NbetweenString() {
	ex := fsb.Nbetween("test", "user1", "user2")

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test NOT BETWEEN 'user1' TO 'user2'", *ps)
}

// Test_NbetweenInt is a unit test for the NbetweenInt method.
// It tests the functionality of the NbetweenInt method in the ExpressionSuite type.
// The NbetweenInt method creates an Expression object using the Nbetween function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_NbetweenInt() {
	ex := fsb.Nbetween("test", 1, 5)

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test NOT BETWEEN 1 TO 5", *ps)
}

// Test_InString is a unit test for the InString method.
// It tests the functionality of the InString method in the ExpressionSuite type.
// The InString method creates an Expression object using the In function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_InString() {
	ex := fsb.In("test", "user1")

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test IN ('user1')", *ps)
}

// Test_InStringMulti is a unit test for the InStringMulti method.
// It tests the functionality of the InStringMulti method in the ExpressionSuite type.
// The InStringMulti method creates an Expression object using the In function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_InStringMulti() {
	ex := fsb.In("test", "user1", "user2")

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test IN ('user1', 'user2')", *ps)
}

// Test_InStringSlice is a unit test for the InStringSlice method.
// It tests the functionality of the InStringSlice method in the ExpressionSuite type.
// The InStringSlice method creates an Expression object using the In function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_InStringSlice() {
	ex := fsb.In("test", []string{"user1", "user2"})

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test IN ('user1', 'user2')", *ps)
}

// Test_InStringSliceMulti is a unit test for the InStringSliceMulti method.
// It tests the functionality of the InStringSliceMulti method in the ExpressionSuite type.
// The InStringSliceMulti method creates an Expression object using the In function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_InStringSliceMulti() {
	ex := fsb.In("test", []string{"user1", "user2"}, "user3")

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test IN ('user1', 'user2', 'user3')", *ps)
}

// Test_InInt is a unit test for the InInt method.
// It tests the functionality of the InInt method in the ExpressionSuite type.
// The InInt method creates an Expression object using the In function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_InInt() {
	ex := fsb.In("test", 1)

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test IN (1)", *ps)
}

// Test_InIntMulti is a unit test for the InIntMulti method.
// It tests the functionality of the InIntMulti method in the ExpressionSuite type.
// The InIntMulti method creates an Expression object using the In function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_InIntMulti() {
	ex := fsb.In("test", 1, 3)

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test IN (1, 3)", *ps)
}

// Test_InIntSlice is a unit test for the InIntSlice method.
// It tests the functionality of the InIntSlice method in the ExpressionSuite type.
// The InIntSlice method creates an Expression object using the In function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_InIntSlice() {
	ex := fsb.In("test", []int{1, 5})

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test IN (1, 5)", *ps)
}

// Test_InIntSliceMutli is a unit test for the InIntSliceMutli method.
// It tests the functionality of the InIntSliceMutli method in the ExpressionSuite type.
// The InIntSliceMutli method creates an Expression object using the In function from the fsb package.
// It then retrieves the condition field of the Expression object using reflection and unsafe pointer conversion.
// Lastly, it asserts that the condition field is equal to the expected value.
func (s *ExpressionSuite) Test_InIntSliceMutli() {
	ex := fsb.In("test", []int{1, 5}, 3)

	v := reflect.ValueOf(ex).Elem()
	r := v.FieldByName("condition")
	ps := (*string)(unsafe.Pointer(r.UnsafeAddr()))

	assert.Equal(s.T(), "test IN (1, 5, 3)", *ps)
}

func TestExpressionSuite(t *testing.T) {
	suite.Run(t, new(ExpressionSuite))
}
