update_user_kitex_gen:
	kitex -module github.com/gwen0x4c3/easy_note -I idl/ -type protobuf idl/user.proto

update_note_kitex_gen:
	kitex -module github.com/gwen0x4c3/easy_note idl/note.thrift