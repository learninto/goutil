package db

// FieldTypeText 文本
const FieldTypeText = 0

// FieldTypeDate 日期
const FieldTypeDate = 100

// FieldTypeDateTime 日期时间
const FieldTypeDateTime = 101

// FieldTypePositiveInteger 正整数（包含0）
const FieldTypePositiveInteger = 200

// FieldTypeInteger 整数
const FieldTypeInteger = 201

// FieldTypePositiveRealNumber 正实数（包含0）
const FieldTypePositiveRealNumber = 200

// FieldTypeRealNumber 实数
const FieldTypeRealNumber = 301

// FieldTypeListOne 列表（只能单选）
const FieldTypeListOne = 400

// FieldTypeListMultiple 列表（可多选）
const FieldTypeListMultiple = 401

// FieldLogicLike 过滤逻辑 0：like ；
const FieldLogicLike = 0

// FieldLogicLeftLike 过滤逻辑 1：左 like；
const FieldLogicLeftLike = 1

// FieldLogicRightLike 过滤逻辑 2：右 like；
const FieldLogicRightLike = 2

// FieldLogicEq 过滤逻辑 100：= ；
const FieldLogicEq = 100

// FieldLogicGt 过滤逻辑 200：>;
const FieldLogicGt = 200

// FieldLogicEgt 过滤逻辑 201：>=
const FieldLogicEgt = 201

// FieldLogicLt 过滤逻辑 300 <；
const FieldLogicLt = 300

// FieldLogicElt 过滤逻辑 301: <=；
const FieldLogicElt = 301

// FieldLogicNeq 过滤逻辑 400： <>；
const FieldLogicNeq = 400

// FieldLogicIn 过滤逻辑 500： in
const FieldLogicIn = 500

// IsMultipleTrue 是否多重应用条件。是指在sql中，该过滤值在多个子查询中需要用到。 0：否；100：是
const IsMultipleTrue = 100

// IsMultipleFalse 是否多重应用条件。是指在sql中，该过滤值在多个子查询中需要用到。 0：否；100：是
const IsMultipleFalse = 0
