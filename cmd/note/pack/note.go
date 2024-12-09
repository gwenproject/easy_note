package pack

import (
    "github.com/gwen0x4c3/easy_note/cmd/note/dal/mysql"
    "github.com/gwen0x4c3/easy_note/kitex_gen/knote"
)

func Notes(notes []*mysql.Note) []*knote.Note {
    res := make([]*knote.Note, 0, len(notes))
    for _, n := range notes {
        res = append(res, Note(n))
    }
    return res
}

func Note(note *mysql.Note) *knote.Note {
    return &knote.Note{
        NoteId:     int64(note.ID),
        UserId:     note.UserId,
        Title:      note.Title,
        Content:    note.Content,
        CreateTime: note.CreatedAt.Unix(),
    }
}
