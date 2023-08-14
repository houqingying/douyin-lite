package repository

import (
	"testing"

	"github.com/houqingying/douyin-lite/pkg/config"
	"k8s.io/klog"
)

// init
func init() {
	// 1. Initialize configuration
	if err := config.Setup(); err != nil {
		klog.Fatalf("config.Setup() error: %s", err)
	}
	// 2. Initialize database
	if err := Setup(&config.Config.Database); err != nil {
		klog.Fatalf("repository.Setup() error: %s", err)
	}
}

func TestCommentDao_CreateComment(t *testing.T) {

	type args struct {
		Comment Comment
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "create comment",
			args: args{
				Comment: Comment{
					VideoId: 1,
					UserId:  1,
					Comment: "test",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			co := &CommentDao{}
			if err := co.CreateComment(&tt.args.Comment); (err != nil) != tt.wantErr {
				t.Errorf("CreateComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCommentDao_DeleteComment(t *testing.T) {
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "delete comment",
			args: args{
				id: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			co := &CommentDao{}
			if err := co.DeleteComment(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCommentDao_DeleteCommentByVideoId(t *testing.T) {
	type args struct {
		videoId int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "delete comment by video id",
			args: args{
				videoId: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			co := &CommentDao{}
			if err := co.DeleteCommentByVideoId(tt.args.videoId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteCommentByVideoId() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCommentDao_QueryCommentById(t *testing.T) {
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		args    args
		want    *Comment
		wantErr bool
	}{
		{
			name: "query comment by id",
			args: args{
				id: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			co := &CommentDao{}
			got, err := co.QueryCommentById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryCommentById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got == nil) != tt.wantErr {
				t.Errorf("QueryCommentById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommentDao_QueryCommentsByUserId(t *testing.T) {
	type args struct {
		userId int
	}
	tests := []struct {
		name    string
		args    args
		want    []Comment
		wantErr bool
	}{
		{
			name: "query comments by user id",
			args: args{
				userId: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			co := &CommentDao{}
			got, err := co.QueryCommentsByUserId(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryCommentsByUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got == nil) != tt.wantErr {
				t.Errorf("QueryCommentsByUserId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommentDao_QueryCommentsByVideoId(t *testing.T) {
	type args struct {
		videoId int
	}
	tests := []struct {
		name    string
		args    args
		want    []Comment
		wantErr bool
	}{
		{
			name: "query comments by video id",
			args: args{
				videoId: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			co := &CommentDao{}
			got, err := co.QueryCommentsByVideoId(tt.args.videoId)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryCommentsByVideoId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got == nil) != tt.wantErr {
				t.Errorf("QueryCommentsByVideoId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommentDao_UpdateComment(t *testing.T) {
	type args struct {
		id      uint
		comment string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "update comment",
			args: args{
				id:      2,
				comment: "update comment",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			co := &CommentDao{}
			if err := co.UpdateComment(tt.args.id, tt.args.comment); (err != nil) != tt.wantErr {
				t.Errorf("UpdateComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCommentDao_DeleteCommentByUserId(t *testing.T) {
	type args struct {
		userId int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "delete comment by user id",
			args: args{
				userId: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			co := &CommentDao{}
			if err := co.DeleteCommentByUserId(tt.args.userId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteCommentByUserId() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
