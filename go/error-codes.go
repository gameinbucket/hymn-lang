package main

// Parsing error codes
const (
	ECodeUnexpectedToken                  = 0
	ECodeUnknownType                      = 1
	ECodeLineIndentation                  = 2
	ECodeClassMemberNotFound              = 3
	ECodeMixedParameters                  = 4
	ECodeClassLazyParameter               = 5
	ECodeClassParameter                   = 6
	ECodeClassMissingGeneric              = 7
	ECodeClassDoesNotExist                = 8
	ECodeClassImplementationMismatch      = 9
	ECodeEnumBracketPosition              = 10
	ECodeEnumMissingParenthesis           = 11
	ECodeEnumMemberNotFound               = 12
	ECodeEnumLazyParameter                = 13
	ECodeEnumParameter                    = 14
	ECodeEnumIncompleteDeclaration        = 15
	ECodeEnumMissingGeneric               = 16
	ECodeEnumDoesNotExist                 = 17
	ECodeEnumImplementationMismatch       = 18
	ECodeEnumDoesNotHaveType              = 19
	ECodeNoDefaultValue                   = 20
	ECodeDoublePlus                       = 21
	ECodeUnknownAssignOperation           = 22
	ECodeAssignOperationRequiresInteger   = 23
	ECodeAssignOperationRequiresNumber    = 24
	ECodeVariableNotMutable               = 25
	ECodeVariableTypeMismatch             = 26
	ECodeVariableDoesNotExist             = 27
	ECodeMemberVariableTypeMismatch       = 28
	ECodeBadAssignment                    = 29
	ECodeExpectedToAssignVariable         = 30
	ECodeUnknownCalcPrefix                = 31
	ECodeUnknownCalcInfix                 = 32
	ECodeFunctionNotFound                 = 33
	ECodeFunctionLazyParameter            = 34
	ECodeFunctionParameter                = 35
	ECodeFunctionMixedParameters          = 36
	ECodeFunctionTooManyParameters        = 37
	ECodeFunctionMissingGeneric           = 38
	ECodeMaybeTypeRequiresPointer         = 39
	ECodeNoneTypeRequiresPointer          = 40
	ECodeNameConflict                     = 41
	ECodeUnknownInterface                 = 42
	ECodeMemberNameConflict               = 43
	ECodeClassRecursiveDefinition         = 44
	ECodeClassRequiresInterface           = 45
	ECodeExpectedClassTypeForInterface    = 46
	ECodeArraySizeRequiresInteger         = 47
	ECodeStaticVariableNotFound           = 48
	ECodeVariableOutOfScope               = 49
	ECodeInterfaceNotFound                = 50
	ECodeMissingRootEnum                  = 51
	ECodeUnexpectedMaybeType              = 52
	ECodeVariableNotAnArray               = 53
	ECodeVariableNotIndexable             = 54
	ECodeExpectingExpression              = 55
	ECodeReturnTypeMismatch               = 56
	ECodeRedundantNoneDefinition          = 57
	ECodeBooleanRequired                  = 58
	ECodeImportPath                       = 59
	ECodeDoubleModuleImport               = 60
	ECodeDoubleClassImport                = 61
	ECodeDoubleInterfaceImport            = 62
	ECodeDoubleEnumImport                 = 63
	ECodeDoubleFunctionImport             = 64
	ECodeDoubleStaticVariableImport       = 65
	ECodeNameAlreadyDefined               = 66
	ECodeReservedName                     = 67
	ECodeNoAdditionalGenerics             = 68
	ECodeOnlyPrimitiveLitreralsAllowed    = 69
	ECodeFunctionAndSignatureMismatch     = 70
	ECodeFunctionMissingParenthesis       = 71
	ECodeFunctionMainSignature            = 72
	ECodeInterfaceDefinitionType          = 73
	ECodeMissingInterface                 = 74
	ECodeInterfaceNameConflict            = 75
	ECodeGenericNotImplemented            = 76
	ECodeTernaryVoid                      = 77
	ECodeTernaryTypeMismatch              = 78
	ECodeClassInterfaceSignatureMismatch  = 79
	ECodeClassMissingRequiredInterface    = 80
	ECodeSliceCapacityRequiresInteger     = 81
	ECodeArrayMemberMismatch              = 82
	ECodeArrayDefinedSizeLessThanImplied  = 83
	ECodeUnknownIdentifier                = 84
	ECodeInvalidCast                      = 85
	ECodeUnknownPrimitive                 = 86
	ECodeBadIsStatement                   = 87
	ECodeNegationProhibited               = 88
	ECodeNoneTypeValueProhibited          = 89
	ECodeIsStatementExpectedEnum          = 90
	ECodeCannotMatchOnPrimitive           = 91
	ECodeCannotMatchOnClass               = 92
	ECodeEnumMatchNotNeeded               = 93
	ECodeEnumMatchRequired                = 94
	ECodeLiteralMismatch                  = 95
	ECodeUnexpectedType                   = 96
	ECodeTypeMismatch                     = 97
	ECodeStringConcatenation              = 98
	ECodeOperationExpectedNumber          = 99
	ECodeNumberTypeMismatch               = 100
	ECodeOperationRequiresDiscreteNumber  = 101
	ECodeClassTypeExpected                = 102
	ECodeClassAndInterfaceMissingGenerics = 103
	ECodeUnusedVariable                   = 104
)
