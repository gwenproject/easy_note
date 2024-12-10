package service

import (
    "context"
    "github.com/gwen0x4c3/easy_note/cmd/note/dal/mysql"
    "github.com/gwen0x4c3/easy_note/cmd/note/pack"
    "github.com/gwen0x4c3/easy_note/cmd/note/rpc"
    "github.com/gwen0x4c3/easy_note/kitex_gen/knote"
    "github.com/gwen0x4c3/easy_note/kitex_gen/kuser"
)

type NoteService struct {
    ctx context.Context
}

func NewNoteService() *NoteService {
    return &NoteService{}
}

var NoteServiceImpl *NoteService

func init() {
    NoteServiceImpl = NewNoteService()
}

func (s *NoteService) CreateNote(ctx context.Context, req *knote.CreateNoteRequest) error {
    dbNote := &mysql.Note{
        UserId:  req.UserId,
        Title:   req.Title,
        Content: req.Content,
    }
    return mysql.CreateNote(ctx, []*mysql.Note{dbNote})
}

func (s *NoteService) UpdateNote(ctx context.Context, req *knote.UpdateNoteRequest) error {
    dbNote := &mysql.Note{}
    dbNote.ID = uint(req.NoteId)
    if req.Title != nil {
        dbNote.Title = *req.Title
    }
    if req.Content != nil {
        dbNote.Content = *req.Content
    }
    return mysql.UpdateNote(ctx, dbNote)
}

func (s *NoteService) DeleteNote(ctx context.Context, req *knote.DeleteNoteRequest) error {
    return mysql.DeleteNote(ctx, req.NoteId)
}

func (s *NoteService) QueryNote(ctx context.Context, req *knote.QueryNoteRequest) ([]*knote.Note, int64, error) {
    noteModels, total, err := mysql.QueryNote(ctx, req.UserId, req.Keyword, int(req.Current), int(req.PageSize))
    if err != nil {
        return nil, 0, err
    }
    userMap, err := rpc.MGetUser(ctx, &kuser.MGetUserRequest{UserIds: []int64{req.UserId}})
    if err != nil {
        return nil, 0, err
    }
    notes := pack.Notes(noteModels)
    for _, note := range notes {
        if u, ok := userMap[note.UserId]; ok {
            note.UserName = u.UserName
            note.UserAvatar = u.Avatar
        }
    }
    return notes, total, err
}
