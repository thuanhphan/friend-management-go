package service

import (
	"friend-management-go/internal/model"
	"friend-management-go/internal/repository"
)

type FriendshipService struct {
	friendshipRepository repository.IFriendshipRepository
}

func NewFriendshipService(mFriendshipRepoitory repository.IFriendshipRepository) *FriendshipService {
	return &FriendshipService{friendshipRepository: mFriendshipRepoitory}
}

func (service *FriendshipService) MakeFriend(friendship model.Friendship) error {
	return service.friendshipRepository.MakeFriend(friendship)
}

func (service *FriendshipService) GetFriends(email string) ([]string, error) {
	return service.friendshipRepository.GetFriends(email)
}

func (service *FriendshipService) GetCommonFriends(email1, email2 string) ([]string, error) {
	return service.friendshipRepository.GetCommonFriends(email1, email2)
}

func (service *FriendshipService) UpdateFriendshipStatus(friendship model.Friendship) error {
	return service.friendshipRepository.UpdateFriendshipStatus(friendship)
}
