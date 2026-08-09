package constants

type CourseItemType int

const (
	COURSE_ITEM_ASSIGNMENT CourseItemType = iota + 1
	COURSE_ITEM_ACTIVITY
	COURSE_ITEM_RESOURCE
)

var courseItemStrArr = []string{
	"undefined",
	"assignment",
	"activity",
	"resource",
}

func (r CourseItemType) ToString() string {
	typeInt := int(r)
	typeStr := courseItemStrArr[0]
	if typeInt >= 1 && typeInt < len(courseItemStrArr) {
		typeStr = courseItemStrArr[typeInt]
	}

	return typeStr
}

func CourseItemTypeFromString(itemTypeStr string) CourseItemType {
	itemTypeNum := 0
	for i := 1; i < len(courseItemStrArr) && itemTypeNum == 0; i++ {
		if itemTypeStr == courseRoleStrArr[i] {
			itemTypeNum = i
		}
	}

	return CourseItemType(itemTypeNum)
}
