package grpc

import (
	"context"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/LizaMeytner/foru/forum-service/internal/model"
	"github.com/LizaMeytner/foru/forum-service/internal/usecase"
	pb "github.com/LizaMeytner/foru/forum-service/proto"
)

type ForumService struct {
	pb.UnimplementedForumServiceServer
	postUseCase    usecase.PostUseCase
	commentUseCase usecase.CommentUseCase
	chatUseCase    usecase.ChatUseCase
}

func NewForumService(
	postUseCase usecase.PostUseCase,
	commentUseCase usecase.CommentUseCase,
	chatUseCase usecase.ChatUseCase,
) *ForumService {
	return &ForumService{
		postUseCase:    postUseCase,
		commentUseCase: commentUseCase,
		chatUseCase:    chatUseCase,
	}
}

// Post methods
func (s *ForumService) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.Post, error) {
	post := &model.Post{
		ID:        uuid.New(),
		AuthorID:  uuid.MustParse(req.UserId),
		Title:     req.Title,
		Content:   req.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.postUseCase.Create(ctx, post); err != nil {
		return nil, err
	}

	return &pb.Post{
		Id:        post.ID.String(),
		UserId:    post.AuthorID.String(),
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: timestamppb.New(post.CreatedAt),
		UpdatedAt: timestamppb.New(post.UpdatedAt),
	}, nil
}

func (s *ForumService) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.Post, error) {
	post, err := s.postUseCase.GetByID(ctx, uuid.MustParse(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.Post{
		Id:        post.ID.String(),
		UserId:    post.AuthorID.String(),
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: timestamppb.New(post.CreatedAt),
		UpdatedAt: timestamppb.New(post.UpdatedAt),
	}, nil
}

func (s *ForumService) ListPosts(ctx context.Context, req *pb.ListPostsRequest) (*pb.ListPostsResponse, error) {
	posts, err := s.postUseCase.GetAll(ctx, int(req.Offset), int(req.Limit))
	if err != nil {
		return nil, err
	}

	pbPosts := make([]*pb.Post, len(posts))
	for i, post := range posts {
		pbPosts[i] = &pb.Post{
			Id:        post.ID.String(),
			UserId:    post.AuthorID.String(),
			Title:     post.Title,
			Content:   post.Content,
			CreatedAt: timestamppb.New(post.CreatedAt),
			UpdatedAt: timestamppb.New(post.UpdatedAt),
		}
	}

	return &pb.ListPostsResponse{
		Posts: pbPosts,
		Total: int32(len(posts)),
	}, nil
}

func (s *ForumService) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.Post, error) {
	post := &model.Post{
		ID:      uuid.MustParse(req.Id),
		Title:   req.Title,
		Content: req.Content,
	}

	if err := s.postUseCase.Update(ctx, post); err != nil {
		return nil, err
	}

	updatedPost, err := s.postUseCase.GetByID(ctx, post.ID)
	if err != nil {
		return nil, err
	}

	return &pb.Post{
		Id:        updatedPost.ID.String(),
		UserId:    updatedPost.AuthorID.String(),
		Title:     updatedPost.Title,
		Content:   updatedPost.Content,
		CreatedAt: timestamppb.New(updatedPost.CreatedAt),
		UpdatedAt: timestamppb.New(updatedPost.UpdatedAt),
	}, nil
}

func (s *ForumService) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.Empty, error) {
	if err := s.postUseCase.Delete(ctx, uuid.MustParse(req.Id)); err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

// Comment methods
func (s *ForumService) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.Comment, error) {
	comment := &model.Comment{
		ID:        uuid.New(),
		PostID:    uuid.MustParse(req.PostId),
		AuthorID:  uuid.MustParse(req.UserId),
		Content:   req.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.commentUseCase.Create(ctx, comment); err != nil {
		return nil, err
	}

	return &pb.Comment{
		Id:        comment.ID.String(),
		PostId:    comment.PostID.String(),
		UserId:    comment.AuthorID.String(),
		Content:   comment.Content,
		CreatedAt: timestamppb.New(comment.CreatedAt),
		UpdatedAt: timestamppb.New(comment.UpdatedAt),
	}, nil
}

func (s *ForumService) GetComment(ctx context.Context, req *pb.GetCommentRequest) (*pb.Comment, error) {
	comment, err := s.commentUseCase.GetByID(ctx, uuid.MustParse(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.Comment{
		Id:        comment.ID.String(),
		PostId:    comment.PostID.String(),
		UserId:    comment.AuthorID.String(),
		Content:   comment.Content,
		CreatedAt: timestamppb.New(comment.CreatedAt),
		UpdatedAt: timestamppb.New(comment.UpdatedAt),
	}, nil
}

func (s *ForumService) ListComments(ctx context.Context, req *pb.ListCommentsRequest) (*pb.ListCommentsResponse, error) {
	comments, err := s.commentUseCase.GetByPostID(ctx, uuid.MustParse(req.PostId), int(req.Offset), int(req.Limit))
	if err != nil {
		return nil, err
	}

	pbComments := make([]*pb.Comment, len(comments))
	for i, comment := range comments {
		pbComments[i] = &pb.Comment{
			Id:        comment.ID.String(),
			PostId:    comment.PostID.String(),
			UserId:    comment.AuthorID.String(),
			Content:   comment.Content,
			CreatedAt: timestamppb.New(comment.CreatedAt),
			UpdatedAt: timestamppb.New(comment.UpdatedAt),
		}
	}

	return &pb.ListCommentsResponse{
		Comments: pbComments,
		Total:    int32(len(comments)),
	}, nil
}

func (s *ForumService) UpdateComment(ctx context.Context, req *pb.UpdateCommentRequest) (*pb.Comment, error) {
	comment := &model.Comment{
		ID:      uuid.MustParse(req.Id),
		Content: req.Content,
	}

	if err := s.commentUseCase.Update(ctx, comment); err != nil {
		return nil, err
	}

	updatedComment, err := s.commentUseCase.GetByID(ctx, comment.ID)
	if err != nil {
		return nil, err
	}

	return &pb.Comment{
		Id:        updatedComment.ID.String(),
		PostId:    updatedComment.PostID.String(),
		UserId:    updatedComment.AuthorID.String(),
		Content:   updatedComment.Content,
		CreatedAt: timestamppb.New(updatedComment.CreatedAt),
		UpdatedAt: timestamppb.New(updatedComment.UpdatedAt),
	}, nil
}

func (s *ForumService) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.Empty, error) {
	if err := s.commentUseCase.Delete(ctx, uuid.MustParse(req.Id)); err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

// Chat methods
func (s *ForumService) CreateChatMessage(ctx context.Context, req *pb.CreateChatMessageRequest) (*pb.ChatMessage, error) {
	message := &model.ChatMessage{
		ID:        uuid.New(),
		AuthorID:  uuid.MustParse(req.UserId),
		Content:   req.Content,
		CreatedAt: time.Now(),
	}

	if err := s.chatUseCase.Create(ctx, message); err != nil {
		return nil, err
	}

	return &pb.ChatMessage{
		Id:        message.ID.String(),
		UserId:    message.AuthorID.String(),
		Content:   message.Content,
		CreatedAt: timestamppb.New(message.CreatedAt),
	}, nil
}

func (s *ForumService) GetChatMessage(ctx context.Context, req *pb.GetChatMessageRequest) (*pb.ChatMessage, error) {
	message, err := s.chatUseCase.GetByID(ctx, uuid.MustParse(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.ChatMessage{
		Id:        message.ID.String(),
		UserId:    message.AuthorID.String(),
		Content:   message.Content,
		CreatedAt: timestamppb.New(message.CreatedAt),
	}, nil
}

func (s *ForumService) ListChatMessages(ctx context.Context, req *pb.ListChatMessagesRequest) (*pb.ListChatMessagesResponse, error) {
	messages, err := s.chatUseCase.GetAll(ctx, int(req.Offset), int(req.Limit))
	if err != nil {
		return nil, err
	}

	pbMessages := make([]*pb.ChatMessage, len(messages))
	for i, message := range messages {
		pbMessages[i] = &pb.ChatMessage{
			Id:        message.ID.String(),
			UserId:    message.AuthorID.String(),
			Content:   message.Content,
			CreatedAt: timestamppb.New(message.CreatedAt),
		}
	}

	return &pb.ListChatMessagesResponse{
		Messages: pbMessages,
		Total:    int32(len(messages)),
	}, nil
}

func (s *ForumService) DeleteOldChatMessages(ctx context.Context, req *pb.DeleteOldChatMessagesRequest) (*pb.Empty, error) {
	if err := s.chatUseCase.DeleteOldMessages(ctx, req.OlderThan.AsTime()); err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
