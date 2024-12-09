package main

import (
    "context"
    "github.com/gwen0x4c3/easy_note/cmd/note/pack"
    "github.com/gwen0x4c3/easy_note/cmd/note/service"
    knote "github.com/gwen0x4c3/easy_note/kitex_gen/knote"
    "github.com/gwen0x4c3/easy_note/pkg/constants"
    "github.com/gwen0x4c3/easy_note/pkg/errno"
)

// NoteServiceImpl implements the last service interface defined in the IDL.
type NoteServiceImpl struct{}

// CreateNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) CreateNote(ctx context.Context, req *knote.CreateNoteRequest) (resp *knote.CreateNoteResponse, err error) {
    resp = new(knote.CreateNoteResponse)
    err = service.NoteServiceImpl.CreateNote(ctx, req)
    if err != nil {
        resp.BaseResp = pack.BuildBaseResp(err)
        return
    }
    resp.BaseResp = pack.BuildBaseResp(nil)
    return
}

// DeleteNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) DeleteNote(ctx context.Context, req *knote.DeleteNoteRequest) (resp *knote.DeleteNoteResponse, err error) {
    // TODO: Your code here...
    return
}

// UpdateNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) UpdateNote(ctx context.Context, req *knote.UpdateNoteRequest) (resp *knote.UpdateNoteResponse, err error) {
    // TODO: Your code here...
    return
}

// MGetNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) MGetNote(ctx context.Context, req *knote.MGetNoteRequest) (resp *knote.MGetNoteResponse, err error) {
    // TODO: Your code here...
    return
}

// QueryNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) QueryNote(ctx context.Context, req *knote.QueryNoteRequest) (resp *knote.QueryNoteResponse, err error) {
    resp = new(knote.QueryNoteResponse)

    if req.UserId <= 0 || req.Current <= 0 || req.PageSize <= 0 {
        resp.BaseResp = pack.BuildBaseResp(errno.ParamErr.WithMessage("参数有误"))
        return
    }
    if req.PageSize >= 100 {
        req.PageSize = constants.DefaultLimit
    }

    notes, total, err := service.NoteServiceImpl.QueryNote(ctx, req)
    if err != nil {
        resp.BaseResp = pack.BuildBaseResp(err)
        return
    }
    resp.Notes = notes
    resp.Total = total
    return
}
