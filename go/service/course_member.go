package service

import (
	"errors"
	"frascati/comp/txhandler"
	"frascati/exception"
	"frascati/lambda"
	"frascati/obj/entity"
	"frascati/repository"
	"frascati/typing"
)

type CourseMemberService interface {
	FindByCourseId(ctx typing.Context, courseId typing.ID) ([]entity.CourseMember, exception.Exception)
	AddMultiple(ctx typing.Context, courseId typing.ID, newMember []entity.CourseMember) ([]entity.CourseMember, exception.Exception)
	DeleteMultiple(ctx typing.Context, courseId typing.ID, memberIds []typing.ID) (int64, exception.Exception)
	Update(ctx typing.Context, courseId typing.ID, members []entity.CourseMember) ([]entity.CourseMember, exception.Exception)
	GetAccessData(ctx typing.Context, courseId typing.ID, userId typing.ID) (entity.CourseMember, exception.Exception)
}

type courseMemberServiceImpl struct {
	memberRepo repository.CourseMemberRepository
	courseRepo repository.CourseRepository
	userRepo   repository.UserRepository
	transactor txhandler.Transactor
}

func NewCourseMemberService(
	memberRepo repository.CourseMemberRepository,
	courseRepo repository.CourseRepository,
	userRepo repository.UserRepository,
	transactor txhandler.Transactor,
) CourseMemberService {
	return courseMemberServiceImpl{
		memberRepo: memberRepo,
		courseRepo: courseRepo,
		userRepo:   userRepo,
		transactor: transactor,
	}
}

func (s courseMemberServiceImpl) FindByCourseId(ctx typing.Context, courseId typing.ID) ([]entity.CourseMember, exception.Exception) {
	course := entity.Course{
		Base: entity.Base{
			ID: courseId,
		},
	}

	res, err := s.memberRepo.FindByCourse(ctx, course)
	return res, err
}

func (s courseMemberServiceImpl) AddMultiple(ctx typing.Context, courseId typing.ID, newMember []entity.CourseMember) ([]entity.CourseMember, exception.Exception) {
	var res []entity.CourseMember
	err := s.transactor.WithTransaction(
		ctx, txhandler.TxOptionDefault, false,
		func(ctx typing.Context) exception.Exception {
			insertResult, err := s.addMultiple(ctx, courseId, newMember)
			res = insertResult
			return err
		},
	)

	return res, err
}

func (s courseMemberServiceImpl) addMultiple(ctx typing.Context, courseId typing.ID, newMember []entity.CourseMember) ([]entity.CourseMember, exception.Exception) {
	isExistErr := s.checkCourse(ctx, courseId)
	if isExistErr != nil {
		return nil, isExistErr
	}

	validUsers, err := s.userRepo.FilterExistingId(
		ctx, lambda.MapList(newMember, func(member entity.CourseMember) (userID typing.ID) {
			return member.User.ID
		}),
	)

	if err != nil {
		return nil, err
	}

	validMember := lambda.MapList(validUsers, func(user entity.User) entity.CourseMember {
		return entity.CourseMember{
			Base:   user.Base,
			User:   user,
			Course: entity.Course{Base: entity.Base{ID: courseId}},
			Role:   user.AcademicRole.ToDefaultCourseRole(),
		}
	})

	res, err := s.memberRepo.AddMultiple(ctx, validMember)
	return res, err
}

func (s courseMemberServiceImpl) DeleteMultiple(ctx typing.Context, courseId typing.ID, memberIds []typing.ID) (int64, exception.Exception) {
	isExistErr := s.checkCourse(ctx, courseId)
	if isExistErr != nil {
		return -1, isExistErr
	}

	res, err := s.memberRepo.DeleteMultiple(ctx, courseId, memberIds)
	return res, err
}

func (s courseMemberServiceImpl) Update(ctx typing.Context, courseId typing.ID, members []entity.CourseMember) ([]entity.CourseMember, exception.Exception) {
	isExistErr := s.checkCourse(ctx, courseId)
	if isExistErr != nil {
		return nil, isExistErr
	}

	res, err := s.memberRepo.Update(ctx, courseId, members)
	return res, err
}

func (s courseMemberServiceImpl) GetAccessData(ctx typing.Context, courseId typing.ID, userId typing.ID) (entity.CourseMember, exception.Exception) {
	member := entity.CourseMember{}
	member.Course.ID = courseId
	member.User.ID = userId

	res, err := s.memberRepo.FindAccessData(ctx, member)
	return res, err
}

func (s courseMemberServiceImpl) checkCourse(ctx typing.Context, courseId typing.ID) exception.Exception {
	isExist, err := s.courseRepo.IsExistById(ctx, courseId)
	if err != nil {
		return err
	}

	if !isExist {
		return exception.NewBaseException(
			exception.CAUSE_NOT_FOUND,
			"course_member/service",
			"record not found",
			errors.New("course not found"),
		)
	}

	return nil
}
