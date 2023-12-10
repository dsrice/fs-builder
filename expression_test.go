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
// The method creates an Expression object using the Pm function from the fsb package, passing in the target attribute and the integer value.
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

func TestExpressionSuite(t *testing.T) {
	suite.Run(t, new(ExpressionSuite))
}
