package main

import (
	"context"
	knote "github.com/gwen0x4c3/easy_note/kitex_gen/knote"
)

// NoteServiceImpl implements the last service interface defined in the IDL.
type NoteServiceImpl struct{}

// CreateNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) CreateNote(ctx context.Context, req *knote.CreateNoteRequest) (resp *knote.CreateNoteResponse, err error) {
	// TODO: Your code here...
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
	// TODO: Your code here...
	return
}